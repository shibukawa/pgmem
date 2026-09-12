package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_sts_begin_parallel_scan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v3 != 0 {
		F_BufFileClose(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			v6 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v6
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v6
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v10
			return
		}
	} else {
		v6 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v6
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v6
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v10
		return
	}
}
func F_sts_end_parallel_scan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v3 != 0 {
		F_BufFileClose(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
			return
		}
	} else {
		return
	}
}
func F_sts_parallel_scan_next(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
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
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
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
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	v7 = m.G0
	v9 = v7 - int32(1072)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v11 < v12 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L7
	} else {
		goto L74
	}
L2:
	;
	m.G0 = v9 + int32(1072)
	return v257
L3:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+8))
	if v147 != 0 {
		goto L42
	} else {
		goto L43
	}
L4:
	;
	goto L5
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v24 = v20 + v21*int32(28)
	v26 = v24 + int32(76)
	v28 = F_LWLockAcquire(m, v26, int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L7
	} else {
		goto L38
	}
L7:
	;
	return int32(0)
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v24)+92))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v32) < base.Ui32(v33) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v33
	v36 = v33
	goto L11
L10:
	;
	v36 = v32
	goto L11
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	if base.Ui32(v36) < base.Ui32(v37) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L6
L13:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v123 <= v122 {
		goto L5
	} else {
		goto L37
	}
L14:
	;
	v40 = v36 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v40
	F_LWLockRelease(m, v26)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L7
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	F_LWLockRelease(m, v26)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L7
	} else {
		goto L29
	}
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v45 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v48 + int32(12)
	v60 = F_pg_snprintf(m, v9+int32(32), int32(1024), int32(466253), v9+int32(16))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L7
	} else {
		goto L21
	}
L19:
	;
	v77 = v45
	goto L20
L20:
	;
	v80 = F_BufFileSeekBlock(m, v77, base.I64_extend_i32_u(v36))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L7
	} else {
		goto L23
	}
L21:
	;
	v62 = int32(4515712)
	v63 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v70 = int32(0)
	v72 = F_BufFileOpenFileSet(m, v67, v9+int32(32), v70, v70)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v72
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v63
	v77 = v72
	goto L20
L23:
	;
	if v80 != 0 {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_BufFileReadExact(m, v82, v9+int32(1064), int32(8))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1068))
	if v88 <= int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1064))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v93
	goto L13
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v88<<(uint(int32(2))%32) + v36
	goto L13
L29:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v103 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	F_BufFileClose(m, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L7
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v113 = base.I32_rem_s(v108+int32(1), v112)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v113
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v115 == v113 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	goto L32
L34:
	;
	v257 = int32(0)
	goto L2
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(0)
	goto L13
L37:
	;
	goto L3
L38:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v36
	F_errmsg(m, int32(387510), v9)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(499185), int32(550), int32(63505))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L7
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
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_BufFileReadExact(m, v148, l1, v147)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L7
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_BufFileReadExact(m, v156, v9+int32(1064), int32(4))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L7
	} else {
		goto L46
	}
L45:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v151 + v153
	goto L44
L46:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v164 = v162 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v164
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1064))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui32(v168) < base.Ui32(v167) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	if v166 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v188 = v166
	v189 = v167
	v190 = v164
	goto L49
L49:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v192 = int32(4)
	v193 = v188 + v192
	v195 = v189 - v192
	v197 = int32(32768) - v190
	if base.Ui32(v195) < base.Ui32(v197) {
		goto L58
	} else {
		goto L59
	}
L50:
	;
	F_pfree(m, v166)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L7
	} else {
		goto L53
	}
L51:
	;
	v174 = v168
	v175 = v167
	goto L52
L52:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v178 = v174 << (uint(int32(1)) % 32)
	if base.Ui32(v178) < base.Ui32(v175) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1064))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v174 = v173
	v175 = v172
	goto L52
L54:
	;
	v180 = v175
	goto L56
L55:
	;
	v180 = v178
	goto L56
L56:
	;
	v181 = F_MemoryContextAlloc(m, v176, v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v181
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1064))
	v188 = v181
	v189 = v186
	v190 = v185
	goto L49
L58:
	;
	v199 = v195
	goto L60
L59:
	;
	v199 = v197
	goto L60
L60:
	;
	F_BufFileReadExact(m, v191, v193, v199)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v202 + v199
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v205 + int32(1)
	v209 = v195 - v199
	if v209 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v213 = v209
	v216 = v193 + v199
	goto L65
L63:
	;
	goto L64
L64:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1064))
	*(*int32)(unsafe.Add(mBase, uint32(v254))) = v255
	v257 = v254
	goto L2
L65:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_BufFileReadExact(m, v217, v9+int32(32), int32(8))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L7
	} else {
		goto L67
	}
L66:
	;
	goto L64
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(8)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	if v225 == int32(0) {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v228 + int32(4)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v233 = int32(32760)
	if base.Ui32(v233) <= base.Ui32(v213) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v236 = v233
	goto L71
L70:
	;
	v236 = v213
	goto L71
L71:
	;
	F_BufFileReadExact(m, v232, v216, v236)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L7
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v241 + v236
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v244
	v247 = v213 - v236
	if v247 != 0 {
		v213 = v247
		v216 = v236 + v216
		goto L65
	} else {
		goto L73
	}
L73:
	;
	goto L66
L74:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L7
	} else {
		goto L75
	}
L75:
	;
	F_errmsg(m, int32(387573), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L7
	} else {
		goto L76
	}
L76:
	;
	F_errdetail_internal(m, int32(623953), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L7
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(499185), int32(468), int32(384160))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L7
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
