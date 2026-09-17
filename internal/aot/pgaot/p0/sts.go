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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
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
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
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
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
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
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
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
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	v8 = m.G0
	v10 = v8 - int32(1072)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v12 < v13 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L7
	} else {
		goto L74
	}
L2:
	;
	m.G0 = v10 + int32(1072)
	return v261
L3:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+8))
	if v149 != 0 {
		goto L42
	} else {
		goto L43
	}
L4:
	;
	goto L5
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v26 = v22 + v23*int32(28)
	v28 = v26 + int32(76)
	v30 = F_LWLockAcquire(m, v28, int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L7
	} else {
		goto L38
	}
L7:
	;
	return int32(0)
L8:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v26)+92))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.Ui32(v34) < base.Ui32(v35) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v35
	v38 = v35
	goto L11
L10:
	;
	v38 = v34
	goto L11
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	if base.Ui32(v38) < base.Ui32(v39) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L6
L13:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v124 <= v123 {
		goto L5
	} else {
		goto L37
	}
L14:
	;
	v42 = v38 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v42
	F_LWLockRelease(m, v28)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	F_LWLockRelease(m, v28)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L7
	} else {
		goto L29
	}
L17:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v47 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v50 + int32(12)
	v57 = v10 + int32(32)
	v62 = F_pg_snprintf(m, v57, int32(1024), int32(_a_F_sts_parallel_scan_next_0), v10+int32(16))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L7
	} else {
		goto L21
	}
L19:
	;
	v77 = v47
	goto L20
L20:
	;
	v80 = F_BufFileSeekBlock(m, v77, base.I64_extend_i32_u(v38))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L7
	} else {
		goto L23
	}
L21:
	;
	v64 = int32(_a_F_sts_parallel_scan_next_1)
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_sts_parallel_scan_next[0]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_sts_parallel_scan_next[0])) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v70 = int32(0)
	v72 = F_BufFileOpenFileSet(m, v69, v57, v70, v70)
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
	*(*int32)(unsafe.Add(mBase, _c_F_sts_parallel_scan_next[0])) = v65
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
	F_BufFileReadExact(m, v82, v10+int32(1064), int32(8))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1068))
	if v88 <= int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1064))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v93
	goto L13
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v88<<(uint(int32(2))%32) + v38
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
	v261 = int32(0)
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
	v131 = m.ExcPending
	if v131 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v38
	F_errmsg(m, int32(_a_F_sts_parallel_scan_next_2), v10)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_sts_parallel_scan_next_3), int32(550), int32(_a_F_sts_parallel_scan_next_4))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
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
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_BufFileReadExact(m, v150, l1, v149)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L7
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_BufFileReadExact(m, v158, v10+int32(1064), int32(4))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L7
	} else {
		goto L46
	}
L45:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v153 + v155
	goto L44
L46:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v166 = v164 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v166
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1064))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui32(v170) < base.Ui32(v169) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	if v168 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v190 = v169
	v191 = v168
	v192 = v166
	goto L49
L49:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v194 = int32(4)
	v195 = v191 + v194
	v197 = v190 - v194
	v199 = int32(_a_F_sts_parallel_scan_next_5) - v192
	if base.Ui32(v197) < base.Ui32(v199) {
		goto L58
	} else {
		goto L59
	}
L50:
	;
	F_pfree(m, v168)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L7
	} else {
		goto L53
	}
L51:
	;
	v176 = v170
	v177 = v169
	goto L52
L52:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v180 = v176 << (uint(int32(1)) % 32)
	if base.Ui32(v180) < base.Ui32(v177) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1064))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v176 = v175
	v177 = v174
	goto L52
L54:
	;
	v182 = v177
	goto L56
L55:
	;
	v182 = v180
	goto L56
L56:
	;
	v183 = F_MemoryContextAlloc(m, v178, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v183
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1064))
	v190 = v188
	v191 = v183
	v192 = v187
	goto L49
L58:
	;
	v201 = v197
	goto L60
L59:
	;
	v201 = v199
	goto L60
L60:
	;
	F_BufFileReadExact(m, v193, v195, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v204 + v201
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v207 + int32(1)
	v211 = v197 - v201
	if v211 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v216 = v211
	v218 = v201 + v195
	goto L65
L63:
	;
	goto L64
L64:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1064))
	*(*int32)(unsafe.Add(mBase, uint32(v258))) = v259
	v261 = v258
	goto L2
L65:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_BufFileReadExact(m, v220, v10+int32(32), int32(8))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
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
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	if v228 == int32(0) {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v231 + int32(4)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v236 = int32(_a_F_sts_parallel_scan_next_6)
	if base.Ui32(v236) <= base.Ui32(v216) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v239 = v236
	goto L71
L70:
	;
	v239 = v216
	goto L71
L71:
	;
	F_BufFileReadExact(m, v235, v218, v239)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L7
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v244 + v239
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v247
	v250 = v216 - v239
	if v250 != 0 {
		v216 = v250
		v218 = v239 + v218
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
	v277 = m.ExcPending
	if v277 != 0 {
		goto L7
	} else {
		goto L75
	}
L75:
	;
	F_errmsg(m, int32(_a_F_sts_parallel_scan_next_7), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L7
	} else {
		goto L76
	}
L76:
	;
	F_errdetail_internal(m, int32(_a_F_sts_parallel_scan_next_8), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L7
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_sts_parallel_scan_next_3), int32(468), int32(_a_F_sts_parallel_scan_next_9))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
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
