package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgstat_archiver_init_shmem_cb(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(-1)
	return
}
func F_pgstat_assoc_relation(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	v8 = *(*int32)(unsafe.Add(mBase, _consts[226]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+117)))
	if v10 != 0 {
		v11 = int32(0)
	} else {
		v11 = v8
	}
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v15 = F_pgstat_prep_pending_entry(m, int32(2), v11, base.I64_extend_i32_u(v12), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
		*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)) = uint8(v10)
		*(*int32)(unsafe.Add(mBase, uint32(v17))) = v12
		*(*int32)(unsafe.Add(mBase, uint32(l0)+272)) = v17
		*(*int32)(unsafe.Add(mBase, uint32(v17)+128)) = l0
		return
	}
}
func F_pgstat_bestart_initial(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int64
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v88 int64
	_ = v88
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v123 int32
	_ = v123
	var v125 int64
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
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
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	v17 = m.G0
	v19 = v17 - int32(320)
	m.G0 = v19
	v22 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+188))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+196))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+179)) = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+193))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+176)) = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)+201))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+168)) = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v22)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+171)) = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)+212))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v22)+216))
	v38 = v22 + int32(228)
	goto L2
L1:
	;
	v43 = v22 + int32(392)
	v45 = v22 + int32(201)
	v47 = v22 + int32(193)
	v51 = v22 + int32(24)
	v53 = *(*int64)(unsafe.Add(mBase, _consts[682]))
	v55 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	v57 = *(*int32)(unsafe.Add(mBase, _consts[522]))
	v59 = *(*int32)(unsafe.Add(mBase, _consts[360]))
	if v59 != 0 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v40 = F__emscripten_memcpy_bulkmem(m, v19+int32(4), v38, int32(164))
	mBase = m.M
	goto L4
L4:
	;
	goto L1
L5:
	;
	v73 = int32(4481700)
	v75 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v76 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v75 + v76
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v79 + v76
	*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v57
	v88 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v51)+24)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v51)+16)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v51)+8)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v51))) = v88
	goto L15
L6:
	;
	goto L10
L7:
	;
	goto L8
L8:
	;
	v72 = F__emscripten_memset_bulkmem(m, v19+int32(184), base.I32_extend8_s(int32(0)), int32(132))
	mBase = m.M
	goto L13
L9:
	;
	goto L5
L10:
	;
	v65 = F__emscripten_memcpy_bulkmem(m, v19+int32(184), v59+int32(144), int32(132))
	mBase = m.M
	goto L12
L12:
	;
	goto L9
L13:
	;
	goto L5
L14:
	;
	v101 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+192)) = uint8(v101)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+188)) = v23
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v19)+179))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+3)) = v104
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v19)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v106
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+200)) = uint8(v101)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v19)+171))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+3)) = v110
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v19)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v112
	*(*int64)(unsafe.Add(mBase, uint32(v22)+220)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+216)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v22)+212)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v22)+208)) = int32(1)
	goto L19
L15:
	;
	v99 = F__emscripten_memcpy_bulkmem(m, v22+int32(56), v19+int32(184), int32(132))
	mBase = m.M
	goto L17
L17:
	;
	goto L14
L18:
	;
	v125 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v43)+8)) = v125
	*(*int64)(unsafe.Add(mBase, uint32(v43))) = v125
	v129 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v33))) = uint8(v129)
	v132 = *(*int32)(unsafe.Add(mBase, _consts[360]))
	if v132 == v129 {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	v123 = F__emscripten_memcpy_bulkmem(m, v38, v19+int32(4), int32(164))
	mBase = m.M
	goto L21
L21:
	;
	goto L18
L22:
	;
	v257 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v34))) = uint8(v257)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+63)) = uint8(v257)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+63)) = uint8(v257)
	v264 = *(*int32)(unsafe.Add(mBase, _consts[683]))
	v266 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v34+v264-v266))) = uint8(v257)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v270 + v266
	v274 = int32(4481700)
	v276 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v276 - v266
	m.G0 = v19 + int32(320)
	return
L23:
	;
	v254 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v254)
	goto L22
L24:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v132)+280))
	if v135 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	goto L29
L26:
	;
	goto L22
L27:
	;
	v250 = F_strlen(m, v239)
	mBase = m.M
	goto L26
L29:
	;
	goto L30
L30:
	;
	v144 = int32(63)
	if (v23^v135)&int32(3) != 0 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v243 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v240))) = uint8(v243)
	goto L27
L32:
	;
	v224 = v219
	v225 = v220
	v226 = v221
	goto L54
L33:
	;
	if v214 == int32(0) {
		v239 = v212
		v240 = v213
		goto L31
	} else {
		goto L53
	}
L34:
	;
	v212 = v135
	v213 = v23
	v214 = v144
	goto L33
L35:
	;
	goto L36
L36:
	;
	if v135&int32(3) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v181 == int32(0) {
		v239 = v178
		v240 = v179
		goto L31
	} else {
		goto L46
	}
L38:
	;
	v178 = v135
	v179 = v23
	v180 = v144
	v181 = int32(1)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v157 = v135
	v158 = v23
	v159 = v144
	goto L41
L41:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	*(*uint8)(unsafe.Add(mBase, uint32(v158))) = uint8(v161)
	if v161 == int32(0) {
		v219 = v157
		v220 = v158
		v221 = v159
		goto L32
	} else {
		goto L43
	}
L42:
	;
	v178 = v172
	v179 = v166
	v180 = v168
	v181 = v170
	goto L37
L43:
	;
	v165 = int32(1)
	v166 = v158 + v165
	v168 = v159 - v165
	v169 = int32(0)
	v170 = base.B2i32(v168 != v169)
	v172 = v157 + v165
	if v172&int32(3) == v169 {
		v178 = v172
		v179 = v166
		v180 = v168
		v181 = v170
		goto L37
	} else {
		goto L44
	}
L44:
	;
	if v168 != 0 {
		v157 = v172
		v158 = v166
		v159 = v168
		goto L41
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	if v184 == int32(0) {
		v212 = v178
		v213 = v179
		v214 = v180
		goto L33
	} else {
		goto L47
	}
L47:
	;
	if base.Ui32(v180) < base.Ui32(int32(4)) {
		v212 = v178
		v213 = v179
		v214 = v180
		goto L33
	} else {
		goto L48
	}
L48:
	;
	v190 = v178
	v191 = v179
	v192 = v180
	goto L49
L49:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	v198 = int32(-2139062144)
	if (int32(16843008)-v195|v195)&v198 != v198 {
		v219 = v190
		v220 = v191
		v221 = v192
		goto L32
	} else {
		goto L51
	}
L50:
	;
	v212 = v206
	v213 = v204
	v214 = v208
	goto L33
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v191))) = v195
	v203 = int32(4)
	v204 = v191 + v203
	v206 = v190 + v203
	v208 = v192 - v203
	if base.Ui32(int32(3)) < base.Ui32(v208) {
		v190 = v206
		v191 = v204
		v192 = v208
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v219 = v212
	v220 = v213
	v221 = v214
	goto L32
L54:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	*(*uint8)(unsafe.Add(mBase, uint32(v225))) = uint8(v228)
	if v228 == int32(0) {
		v239 = v224
		v240 = v225
		goto L31
	} else {
		goto L56
	}
L55:
	;
	v239 = v235
	v240 = v233
	goto L31
L56:
	;
	v232 = int32(1)
	v233 = v225 + v232
	v235 = v224 + v232
	v237 = v226 - v232
	if v237 != 0 {
		v224 = v235
		v225 = v233
		v226 = v237
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
}
func F_pgstat_clear_snapshot(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int64
	_ = v3
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v1 = int32(0)
	v3 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[685])) = v3
	*(*int64)(unsafe.Add(mBase, _consts[686])) = v3
	*(*int64)(unsafe.Add(mBase, _consts[687])) = v3
	*(*uint8)(unsafe.Add(mBase, _consts[688])) = uint8(v1)
	*(*int32)(unsafe.Add(mBase, _consts[689])) = v1
	*(*int32)(unsafe.Add(mBase, _consts[690])) = v1
	v21 = *(*int32)(unsafe.Add(mBase, _consts[691]))
	if v21 != 0 {
		F_MemoryContextDelete(m, v21)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[691])) = int32(0)
			F_pgstat_clear_backend_activity_snapshot(m)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				v30 = int32(0)
				*(*uint8)(unsafe.Add(mBase, _consts[692])) = uint8(v30)
				return
			}
		}
	} else {
		F_pgstat_clear_backend_activity_snapshot(m)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			v30 = int32(0)
			*(*uint8)(unsafe.Add(mBase, _consts[692])) = uint8(v30)
			return
		}
	}
}
func F_pgstat_count_io_op(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	var v22 int64
	_ = v22
	var v28 int32
	_ = v28
	var v44 int32
	_ = v44
	var v47 int64
	_ = v47
	var v53 int64
	_ = v53
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	v13 = l0*int32(320) + l1<<(uint(int32(6))%32) + l2<<(uint(int32(3))%32)
	v16 = *(*int64)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[694])))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[694]))) = v16 + int64(1)
	v22 = *(*int64)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[695])))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[695]))) = v22
	v28 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	if base.Ui32(int32(16)) < base.Ui32(v28) {
	} else {
		if int32(1)<<(uint(v28)%32)&int32(115186) == int32(0) {
		} else {
			v44 = l0*int32(320) + l1<<(uint(int32(6))%32) + l2<<(uint(int32(3))%32)
			v47 = *(*int64)(unsafe.Add(mBase, uint32(v44)+uint32(_consts[696])))
			*(*int64)(unsafe.Add(mBase, uint32(v44)+uint32(_consts[696]))) = v47 + base.I64_extend_i32_u(int32(1))
			v53 = *(*int64)(unsafe.Add(mBase, uint32(v44)+uint32(_consts[697])))
			*(*int64)(unsafe.Add(mBase, uint32(v44)+uint32(_consts[697]))) = v53 + int64(0)
			v57 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _consts[137])) = uint8(v57)
			*(*uint8)(unsafe.Add(mBase, _consts[174])) = uint8(v57)
		}
	}
	v65 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[137])) = uint8(v65)
	*(*uint8)(unsafe.Add(mBase, _consts[170])) = uint8(v65)
	return
}
func F_pgstat_database_reset_timestamp_cb(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, uint32(l0)+280)) = l1
	return
}
func F_pgstat_drop_entry_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
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
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
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
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v14 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v19 - v17
	if v19 == v17 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L12
	} else {
		goto L24
	}
L4:
	;
	m.G0 = v12 + int32(32)
	return base.B2i32(v19 == int32(1))
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if l1 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	if l1 != 0 {
		goto L4
	} else {
		goto L22
	}
L8:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _consts[700]))
	F_dsa_free(m, v90, v25)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L12
	} else {
		goto L21
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[699]))
	F_dshash_delete_entry(m, v29, l0)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+36))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v34)+40))
	v46 = v35 + int32(base.Ui32(v37)>>(uint(int32(32)-v39)%32))<<(uint(int32(2))%32)
	goto L15
L12:
	;
	return int32(0)
L13:
	;
	goto L8
L14:
	;
	goto L8
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v54 == int32(0) {
		goto L14
	} else {
		goto L17
	}
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	F_dsa_free(m, v62, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L12
	} else {
		goto L20
	}
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v58 = F_dsa_get_address(m, v57, v54)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	if v58 != v36 {
		v46 = v58
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v61
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v34)+32))
	v74 = v67 + int32(base.Ui32(v37)>>(uint(int32(25))%32))*int32(20) + int32(24)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v75 - int32(1)
	goto L14
L21:
	;
	goto L4
L22:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _consts[699]))
	F_dshash_release_lock(m, v94, l0)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	goto L4
L24:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v116-int32(1)) <= base.Ui32(int32(11)) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v146 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v145)+68))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v147
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v146
	F_errmsg_internal(m, int32(38451), v12)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L12
	} else {
		goto L32
	}
L26:
	;
	v145 = v116*int32(72) + int32(1634656)
	goto L25
L27:
	;
	goto L28
L28:
	;
	if base.Ui32(int32(8)) < base.Ui32(v116-int32(24)) {
		v143 = int32(0)
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v145 = v143
	goto L25
L30:
	;
	v131 = int32(0)
	v133 = *(*int32)(unsafe.Add(mBase, _consts[684]))
	if v133 == v131 {
		v143 = v131
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v133+v116<<(uint(int32(2))%32)-int32(96))))
	v143 = v141
	goto L29
L32:
	;
	F_errfinish(m, int32(494039), int32(908), int32(308886))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgstat_fetch_stat_backend_by_pid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	v3 = int32(0)
	v7 = F_BackendPidGetProc(m, l0)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l1 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	goto L5
L4:
	;
	goto L5
L5:
	;
	if v7 != 0 {
		v50 = v7
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[94]))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v56 = base.I32_div_s(v50-v53, int32(640))
	v57 = F_pgstat_get_beentry_by_proc_number(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L20
	}
L7:
	;
	v13 = int32(0)
	if l0 == v13 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v47 != 0 {
		v50 = v47
		goto L6
	} else {
		goto L18
	}
L9:
	;
	v47 = int32(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[693]))
	v22 = v13
	goto L13
L12:
	;
	v47 = v42
	goto L8
L13:
	;
	v27 = v20 + v22*int32(640)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+44))
	if v28 == l0 {
		v42 = v27
		goto L12
	} else {
		goto L15
	}
L14:
	;
	v47 = int32(0)
	goto L8
L15:
	;
	v34 = v20 + (v22|int32(1))*int32(640)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+44))
	if v35 == l0 {
		v42 = v34
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v38 = v22 + int32(2)
	if v38 != int32(38) {
		v22 = v38
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	return int32(0)
L19:
	;
	return v83
L20:
	;
	if v57 == int32(0) {
		v83 = v3
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	if base.Ui32(int32(16)) < base.Ui32(v61) {
		v83 = v3
		goto L19
	} else {
		goto L22
	}
L22:
	;
	if int32(1)<<(uint(v61)%32)&int32(115186) == int32(0) {
		v83 = v3
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v70 != l0 {
		v83 = v3
		goto L19
	} else {
		goto L24
	}
L24:
	;
	if l1 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v61
	goto L27
L26:
	;
	goto L27
L27:
	;
	v76 = F_pgstat_fetch_entry(m, int32(6), int32(0), base.I64_extend_i32_s(v56))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if l1 == int32(0) {
		v83 = v76
		goto L19
	} else {
		goto L29
	}
L29:
	;
	if v76 != 0 {
		v83 = v76
		goto L19
	} else {
		goto L30
	}
L30:
	;
	v80 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v80
	v83 = v80
	goto L19
}
func F_pgstat_get_kind_info(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
		return l0*int32(72) + int32(1634656)
	} else {
		if base.Ui32(int32(8)) < base.Ui32(l0-int32(24)) {
			v29 = int32(0)
		} else {
			v17 = int32(0)
			v19 = *(*int32)(unsafe.Add(mBase, _consts[684]))
			if v19 == v17 {
				v29 = v17
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v19+l0<<(uint(int32(2))%32)-int32(96))))
				v29 = v27
			}
		}
		return v29
	}
}
func F_pgstat_get_my_query_id(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int64
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	if v3 == int32(0) {
		return int64(0)
	} else {
		v8 = *(*int64)(unsafe.Add(mBase, uint32(v3)+392))
		return v8
	}
}
func F_pgstat_io_init_shmem_cb(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	v2 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(-1)
	v9 = l0 + int32(16)
	v10 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v9))) = uint16(v10)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(-1)
	v17 = l0 + int32(32)
	v18 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v17))) = uint16(v18)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = int64(-1)
	v25 = l0 + int32(48)
	v26 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v25))) = uint16(v26)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = int64(-1)
	v33 = l0 - int32(-64)
	v34 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v33))) = uint16(v34)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v33)+8)) = int64(-1)
	v41 = l0 + int32(80)
	v42 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v41))) = uint16(v42)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v41)+8)) = int64(-1)
	v49 = l0 + int32(96)
	v50 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v49))) = uint16(v50)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = int64(-1)
	v57 = l0 + int32(112)
	v58 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v57))) = uint16(v58)
	*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v57)+8)) = int64(-1)
	v65 = l0 + int32(128)
	v66 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v65))) = uint16(v66)
	*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = int64(-1)
	v73 = l0 + int32(144)
	v74 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v73))) = uint16(v74)
	*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v73)+8)) = int64(-1)
	v81 = l0 + int32(160)
	v82 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v81))) = uint16(v82)
	*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v81)+8)) = int64(-1)
	v89 = l0 + int32(176)
	v90 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v89))) = uint16(v90)
	*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v89)+8)) = int64(-1)
	v97 = l0 + int32(192)
	v98 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v97))) = uint16(v98)
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v97)+8)) = int64(-1)
	v105 = l0 + int32(208)
	v106 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v105))) = uint16(v106)
	*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v105)+8)) = int64(-1)
	v113 = l0 + int32(224)
	v114 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v113))) = uint16(v114)
	*(*int32)(unsafe.Add(mBase, uint32(v113)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v113)+8)) = int64(-1)
	v121 = l0 + int32(240)
	v122 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v121))) = uint16(v122)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v121)+8)) = int64(-1)
	v129 = l0 + int32(256)
	v130 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v129))) = uint16(v130)
	*(*int32)(unsafe.Add(mBase, uint32(v129)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v129)+8)) = int64(-1)
	v137 = l0 + int32(272)
	v138 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v137))) = uint16(v138)
	*(*int32)(unsafe.Add(mBase, uint32(v137)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v137)+8)) = int64(-1)
	return
}
func F_pgstat_progress_incr_param(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	v5 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	if v5 == int32(0) {
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, _consts[34])))
		if v9 != int32(1) {
		} else {
			v12 = int32(4481700)
			v14 = *(*int32)(unsafe.Add(mBase, _consts[7]))
			v15 = int32(1)
			*(*int32)(unsafe.Add(mBase, _consts[7])) = v14 + v15
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v18 + v15
			v26 = v5 + l0<<(uint(int32(3))%32) + int32(232)
			v27 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
			*(*int64)(unsafe.Add(mBase, uint32(v26))) = v27 + l1
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v30 + v15
			v36 = *(*int32)(unsafe.Add(mBase, _consts[7]))
			*(*int32)(unsafe.Add(mBase, _consts[7])) = v36 - v15
		}
	}
	return
}
func F_pgstat_release_entry_ref(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v98 int64
	_ = v98
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v108 int64
	_ = v108
	var v114 int64
	_ = v114
	var v120 int64
	_ = v120
	var v125 int64
	_ = v125
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var __phi226 int32
	_ = __phi226
	var v231 int32
	_ = v231
	var __phi231 int32
	_ = __phi231
	var v232 int32
	_ = v232
	var __phi232 int32
	_ = __phi232
	var v235 int32
	_ = v235
	var __phi235 int32
	_ = __phi235
	var v237 int32
	_ = v237
	var v238 int64
	_ = v238
	var v239 int64
	_ = v239
	var v242 int64
	_ = v242
	var v243 int64
	_ = v243
	var v244 int64
	_ = v244
	var v246 int64
	_ = v246
	var v251 int64
	_ = v251
	var v257 int64
	_ = v257
	var v262 int64
	_ = v262
	var v267 int64
	_ = v267
	var v277 int64
	_ = v277
	var v279 int64
	_ = v279
	var v281 int64
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l1 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L16
	} else {
		goto L73
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L16
	} else {
		goto L70
	}
L3:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _consts[698]))
	v94 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v94
	v96 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v96
	v98 = int64(23)
	v101 = int64(2388976653695081527)
	v102 = (v96 ^ int64(base.Ui64(v96)>>(uint(v98)%64))) * v101
	v103 = int64(47)
	v108 = int64(-8645972361240307355)
	v114 = (v94 ^ int64(base.Ui64(v94)>>(uint(v98)%64))) * v101
	v120 = ((v102^int64(base.Ui64(v102)>>(uint(v103)%64))^int64(-9208349263878056368))*v108 ^ int64(base.Ui64(v114)>>(uint(v103)%64)) ^ v114) * v108
	v125 = (int64(base.Ui64(v120)>>(uint(v98)%64)) ^ v120) * v101
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v135 = base.I32_wrap_i64(int64(base.Ui64(v125)>>(uint(v103)%64)) ^ v125 - int64(base.Ui64(v125)>>(uint(int64(32))%64)))
	goto L29
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v17 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if l2 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v54 == int32(0) {
		goto L3
	} else {
		goto L19
	}
L8:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if base.Ui32(v22-int32(1)) <= base.Ui32(int32(11)) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	if v40 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v39 = v22*int32(72) + int32(1634656)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[684]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32+v22<<(uint(int32(2))%32)-int32(96))))
	v39 = v38
	goto L9
L13:
	;
	m.T0[v40].(func(*base.Module, int32))(m, l1)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	F_pfree(m, v20)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L16
	} else {
		goto L18
	}
L16:
	;
	return
L17:
	;
	goto L15
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(0)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v50
	goto L7
L19:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	v59 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v58 - v59
	if v58 != v59 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _consts[699]))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v68 = F_dshash_find(m, v65, v66, int32(1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	if v68 == int32(0) {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+24))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v73 == v74 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v68)+28))
	v78 = *(*int32)(unsafe.Add(mBase, _consts[699]))
	F_dshash_delete_entry(m, v78, v68)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L16
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _consts[699]))
	F_dshash_release_lock(m, v86, v68)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L16
	} else {
		goto L28
	}
L26:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _consts[700]))
	F_dsa_free(m, v82, v76)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L16
	} else {
		goto L27
	}
L27:
	;
	goto L3
L28:
	;
	goto L3
L29:
	;
	v145 = v135 & v134
	v148 = v133 + v145*int32(24)
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+16)))
	switch v149 {
	case 0:
		goto L32
	case 1:
		goto L33
	default:
		goto L31
	}
L31:
	;
	v135 = v145 + int32(1)
	goto L29
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L16
	} else {
		goto L67
	}
L33:
	;
	v150 = int32(16)
	goto L37
L34:
	;
	if v212 != 0 {
		goto L31
	} else {
		goto L52
	}
L35:
	;
	v212 = int32(0)
	goto L34
L36:
	;
	v186 = v181
	v187 = v182
	v188 = v183
	goto L46
L37:
	;
	if (v148|v13)&int32(3) != 0 {
		v181 = v148
		v182 = v13
		v183 = v150
		goto L36
	} else {
		goto L40
	}
L39:
	;
	if v171 == int32(0) {
		goto L35
	} else {
		goto L45
	}
L40:
	;
	v158 = v148
	v159 = v13
	v160 = v150
	goto L41
L41:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	if v163 != v164 {
		v181 = v158
		v182 = v159
		v183 = v160
		goto L36
	} else {
		goto L43
	}
L42:
	;
	goto L39
L43:
	;
	v166 = int32(4)
	v167 = v159 + v166
	v169 = v158 + v166
	v171 = v160 - v166
	if base.Ui32(int32(3)) < base.Ui32(v171) {
		v158 = v169
		v159 = v167
		v160 = v171
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v181 = v169
	v182 = v167
	v183 = v171
	goto L36
L46:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	if v191 == v192 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v212 = v191 - v192
	goto L34
L48:
	;
	v194 = int32(1)
	v199 = v188 - v194
	if v199 != 0 {
		v186 = v186 + v194
		v187 = v187 + v194
		v188 = v199
		goto L46
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	goto L47
L51:
	;
	goto L35
L52:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	v214 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v213 - v214
	v219 = (v145 + v214) & v134
	v222 = v133 + v219*int32(24)
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+16)))
	if v223 != v214 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v304 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v296)+16)) = uint8(v304)
	if l1 != 0 {
		goto L63
	} else {
		goto L64
	}
L54:
	;
	v296 = v148
	goto L53
L55:
	;
	goto L56
L56:
	;
	__phi226 = v148
	__phi231 = v222
	__phi232 = v134
	__phi235 = v219
	v226 = __phi226
	v231 = __phi231
	v232 = __phi232
	v235 = __phi235
	goto L57
L57:
	;
	v237 = v231 + int32(8)
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v237)))
	v239 = int64(23)
	v242 = int64(2388976653695081527)
	v243 = (int64(base.Ui64(v238)>>(uint(v239)%64)) ^ v238) * v242
	v244 = int64(47)
	v246 = *(*int64)(unsafe.Add(mBase, uint32(v231)))
	v251 = (int64(base.Ui64(v246)>>(uint(v239)%64)) ^ v246) * v242
	v257 = int64(-8645972361240307355)
	v262 = (int64(base.Ui64(v243)>>(uint(v244)%64)) ^ (v251^int64(base.Ui64(v251)>>(uint(v244)%64))^int64(-9208349263878056368))*v257 ^ v243) * v257
	v267 = (int64(base.Ui64(v262)>>(uint(v239)%64)) ^ v262) * v242
	if v235 == v232&base.I32_wrap_i64(int64(base.Ui64(v267)>>(uint(v244)%64))^v267-int64(base.Ui64(v267)>>(uint(int64(32))%64))) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v296 = v231
	goto L53
L59:
	;
	v296 = v226
	goto L53
L60:
	;
	goto L61
L61:
	;
	v277 = *(*int64)(unsafe.Add(mBase, uint32(v231)))
	*(*int64)(unsafe.Add(mBase, uint32(v226))) = v277
	v279 = *(*int64)(unsafe.Add(mBase, uint32(v231)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v226)+16)) = v279
	v281 = *(*int64)(unsafe.Add(mBase, uint32(v237)))
	*(*int64)(unsafe.Add(mBase, uint32(v226)+8)) = v281
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v285 = int32(1)
	v287 = v284 & (v235 + v285)
	v290 = v283 + v287*int32(24)
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290)+16)))
	if v291 == v285 {
		__phi226 = v231
		__phi231 = v290
		__phi232 = v284
		__phi235 = v287
		v226 = __phi226
		v231 = __phi231
		v232 = __phi232
		v235 = __phi235
		goto L57
	} else {
		goto L62
	}
L62:
	;
	goto L58
L63:
	;
	F_pfree(m, l1)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L16
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	m.G0 = v13 + int32(16)
	return
L66:
	;
	goto L65
L67:
	;
	F_errmsg_internal(m, int32(249565), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L16
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(494039), int32(667), int32(337245))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L16
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	F_errmsg_internal(m, int32(12285), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L16
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(494039), int32(640), int32(337245))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L16
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
	F_errmsg_internal(m, int32(501695), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L16
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(494039), int32(611), int32(337245))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L16
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgstat_replslot_reset_timestamp_cb(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = l1
	return
}
func F_pgstat_replslot_to_serialized_name_cb(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v12 = *(*int32)(unsafe.Add(mBase, _consts[544]))
	v14 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v18 = F_LWLockAcquire(m, v14+int32(4736), int32(1))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		v22 = v12 + v10*int32(288)
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
		if v23 == int32(1) {
			v29 = F_strncpy(m, l2, v22+int32(24), int32(64))
			mBase = m.M
			v30 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v29)+63)) = uint8(v30)
		} else {
		}
		v33 = *(*int32)(unsafe.Add(mBase, _consts[29]))
		F_LWLockRelease(m, v33+int32(4736))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return
		} else {
			if v23 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v8))) = v44
					F_errmsg_internal(m, int32(37785), v8)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						F_errfinish(m, int32(489937), int32(198), int32(500326))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		}
	}
}
func F_pgstat_report_plan_id(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v14 int64
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	v5 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	if v5 == int32(0) {
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, _consts[34])))
		if v9 != int32(1) {
		} else {
			v14 = *(*int64)(unsafe.Add(mBase, uint32(v5)+400))
			if base.B2i32(l1 == int32(0))&base.B2i32(v14 != int64(0)) != 0 {
			} else {
				v18 = int32(4481700)
				v20 = *(*int32)(unsafe.Add(mBase, _consts[7]))
				v21 = int32(1)
				*(*int32)(unsafe.Add(mBase, _consts[7])) = v20 + v21
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v24 + v21
				*(*int64)(unsafe.Add(mBase, uint32(v5)+400)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v24 + int32(2)
				v35 = *(*int32)(unsafe.Add(mBase, _consts[7]))
				*(*int32)(unsafe.Add(mBase, _consts[7])) = v35 - v21
			}
		}
	}
	return
}
func F_pgstat_report_xact_timestamp(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _consts[34])))
	if v4 != int32(1) {
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[31]))
		if v8 == int32(0) {
		} else {
			v11 = int32(4481700)
			v13 = *(*int32)(unsafe.Add(mBase, _consts[7]))
			v14 = int32(1)
			*(*int32)(unsafe.Add(mBase, _consts[7])) = v13 + v14
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v17 + v14
			*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v17 + int32(2)
			v28 = *(*int32)(unsafe.Add(mBase, _consts[7]))
			*(*int32)(unsafe.Add(mBase, _consts[7])) = v28 - v14
		}
	}
	return
}
func F_pgstat_request_entry_refs_gc(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, _consts[454]))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v3)+16)) = v4 + int64(1)
	return
}
func F_pgstat_reset_wait_event_storage(m *base.Module) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[165])) = int32(4470276)
	return
}
func F_pgstat_slru_reset_all_cb(m *base.Module, l0 int64) {
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
	var v13 int64
	_ = v13
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int64
	_ = v49
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int64
	_ = v89
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int64
	_ = v129
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int64
	_ = v169
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int64
	_ = v209
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int64
	_ = v249
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int64
	_ = v289
	var v319 int32
	_ = v319
	v5 = *(*int32)(unsafe.Add(mBase, _consts[454]))
	v7 = v5 + int32(52744)
	v9 = F_LWLockAcquire(m, v7, int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v13 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v5)+uint32(_consts[701]))) = v13
		*(*int64)(unsafe.Add(mBase, uint32(v5)+uint32(_consts[702]))) = v13
		*(*int64)(unsafe.Add(mBase, uint32(v5)+uint32(_consts[703]))) = v13
		*(*int64)(unsafe.Add(mBase, uint32(v5)+uint32(_consts[704]))) = v13
		*(*int64)(unsafe.Add(mBase, uint32(v5)+uint32(_consts[705]))) = v13
		*(*int64)(unsafe.Add(mBase, uint32(v5)+uint32(_consts[706]))) = v13
		*(*int64)(unsafe.Add(mBase, uint32(v5)+uint32(_consts[707]))) = v13
		*(*int64)(unsafe.Add(mBase, uint32(v5)+uint32(_consts[708]))) = l0
		F_LWLockRelease(m, v7)
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, _consts[454]))
			v43 = v41 + int32(52744)
			v45 = F_LWLockAcquire(m, v43, int32(0))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				v49 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v41)+uint32(_consts[709]))) = v49
				*(*int64)(unsafe.Add(mBase, uint32(v41)+uint32(_consts[710]))) = v49
				*(*int64)(unsafe.Add(mBase, uint32(v41)+uint32(_consts[711]))) = v49
				*(*int64)(unsafe.Add(mBase, uint32(v41)+uint32(_consts[712]))) = v49
				*(*int64)(unsafe.Add(mBase, uint32(v41)+uint32(_consts[713]))) = v49
				*(*int64)(unsafe.Add(mBase, uint32(v41)+uint32(_consts[714]))) = v49
				*(*int64)(unsafe.Add(mBase, uint32(v41)+uint32(_consts[715]))) = v49
				*(*int64)(unsafe.Add(mBase, uint32(v41)+uint32(_consts[716]))) = l0
				F_LWLockRelease(m, v43)
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return
				} else {
					v81 = *(*int32)(unsafe.Add(mBase, _consts[454]))
					v83 = v81 + int32(52744)
					v85 = F_LWLockAcquire(m, v83, int32(0))
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return
					} else {
						v89 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[717]))) = v89
						*(*int64)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[718]))) = v89
						*(*int64)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[719]))) = v89
						*(*int64)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[720]))) = v89
						*(*int64)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[721]))) = v89
						*(*int64)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[722]))) = v89
						*(*int64)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[723]))) = v89
						*(*int64)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[724]))) = l0
						F_LWLockRelease(m, v83)
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
							return
						} else {
							v121 = *(*int32)(unsafe.Add(mBase, _consts[454]))
							v123 = v121 + int32(52744)
							v125 = F_LWLockAcquire(m, v123, int32(0))
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return
							} else {
								v129 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[725]))) = v129
								*(*int64)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[726]))) = v129
								*(*int64)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[727]))) = v129
								*(*int64)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[728]))) = v129
								*(*int64)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[729]))) = v129
								*(*int64)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[730]))) = v129
								*(*int64)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[731]))) = v129
								*(*int64)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[732]))) = l0
								F_LWLockRelease(m, v123)
								mBase = m.M
								v159 = m.ExcPending
								if v159 != 0 {
									return
								} else {
									v161 = *(*int32)(unsafe.Add(mBase, _consts[454]))
									v163 = v161 + int32(52744)
									v165 = F_LWLockAcquire(m, v163, int32(0))
									mBase = m.M
									v166 = m.ExcPending
									if v166 != 0 {
										return
									} else {
										v169 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v161)+uint32(_consts[733]))) = v169
										*(*int64)(unsafe.Add(mBase, uint32(v161)+uint32(_consts[734]))) = v169
										*(*int64)(unsafe.Add(mBase, uint32(v161)+uint32(_consts[735]))) = v169
										*(*int64)(unsafe.Add(mBase, uint32(v161)+uint32(_consts[736]))) = v169
										*(*int64)(unsafe.Add(mBase, uint32(v161)+uint32(_consts[737]))) = v169
										*(*int64)(unsafe.Add(mBase, uint32(v161)+uint32(_consts[738]))) = v169
										*(*int64)(unsafe.Add(mBase, uint32(v161)+uint32(_consts[739]))) = v169
										*(*int64)(unsafe.Add(mBase, uint32(v161)+uint32(_consts[740]))) = l0
										F_LWLockRelease(m, v163)
										mBase = m.M
										v199 = m.ExcPending
										if v199 != 0 {
											return
										} else {
											v201 = *(*int32)(unsafe.Add(mBase, _consts[454]))
											v203 = v201 + int32(52744)
											v205 = F_LWLockAcquire(m, v203, int32(0))
											mBase = m.M
											v206 = m.ExcPending
											if v206 != 0 {
												return
											} else {
												v209 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v201)+uint32(_consts[741]))) = v209
												*(*int64)(unsafe.Add(mBase, uint32(v201)+uint32(_consts[742]))) = v209
												*(*int64)(unsafe.Add(mBase, uint32(v201)+uint32(_consts[743]))) = v209
												*(*int64)(unsafe.Add(mBase, uint32(v201)+uint32(_consts[744]))) = v209
												*(*int64)(unsafe.Add(mBase, uint32(v201)+uint32(_consts[745]))) = v209
												*(*int64)(unsafe.Add(mBase, uint32(v201)+uint32(_consts[746]))) = v209
												*(*int64)(unsafe.Add(mBase, uint32(v201)+uint32(_consts[747]))) = v209
												*(*int64)(unsafe.Add(mBase, uint32(v201)+uint32(_consts[748]))) = l0
												F_LWLockRelease(m, v203)
												mBase = m.M
												v239 = m.ExcPending
												if v239 != 0 {
													return
												} else {
													v241 = *(*int32)(unsafe.Add(mBase, _consts[454]))
													v243 = v241 + int32(52744)
													v245 = F_LWLockAcquire(m, v243, int32(0))
													mBase = m.M
													v246 = m.ExcPending
													if v246 != 0 {
														return
													} else {
														v249 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(v241)+uint32(_consts[749]))) = v249
														*(*int64)(unsafe.Add(mBase, uint32(v241)+uint32(_consts[750]))) = v249
														*(*int64)(unsafe.Add(mBase, uint32(v241)+uint32(_consts[751]))) = v249
														*(*int64)(unsafe.Add(mBase, uint32(v241)+uint32(_consts[752]))) = v249
														*(*int64)(unsafe.Add(mBase, uint32(v241)+uint32(_consts[753]))) = v249
														*(*int64)(unsafe.Add(mBase, uint32(v241)+uint32(_consts[754]))) = v249
														*(*int64)(unsafe.Add(mBase, uint32(v241)+uint32(_consts[755]))) = v249
														*(*int64)(unsafe.Add(mBase, uint32(v241)+uint32(_consts[756]))) = l0
														F_LWLockRelease(m, v243)
														mBase = m.M
														v279 = m.ExcPending
														if v279 != 0 {
															return
														} else {
															v281 = *(*int32)(unsafe.Add(mBase, _consts[454]))
															v283 = v281 + int32(52744)
															v285 = F_LWLockAcquire(m, v283, int32(0))
															mBase = m.M
															v286 = m.ExcPending
															if v286 != 0 {
																return
															} else {
																v289 = int64(0)
																*(*int64)(unsafe.Add(mBase, uint32(v281)+uint32(_consts[757]))) = v289
																*(*int64)(unsafe.Add(mBase, uint32(v281)+uint32(_consts[758]))) = v289
																*(*int64)(unsafe.Add(mBase, uint32(v281)+uint32(_consts[759]))) = v289
																*(*int64)(unsafe.Add(mBase, uint32(v281)+uint32(_consts[760]))) = v289
																*(*int64)(unsafe.Add(mBase, uint32(v281)+uint32(_consts[761]))) = v289
																*(*int64)(unsafe.Add(mBase, uint32(v281)+uint32(_consts[762]))) = v289
																*(*int64)(unsafe.Add(mBase, uint32(v281)+uint32(_consts[763]))) = v289
																*(*int64)(unsafe.Add(mBase, uint32(v281)+uint32(_consts[764]))) = l0
																F_LWLockRelease(m, v283)
																mBase = m.M
																v319 = m.ExcPending
																if v319 != 0 {
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
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pgstat_subscription_reset_timestamp_cb(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, uint32(l0)+96)) = l1
	return
}
func F_pgstat_twophase_postabort(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	v9 = *(*int32)(unsafe.Add(mBase, _consts[226]))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+52)))
	if v10 != 0 {
		v11 = int32(0)
	} else {
		v11 = v9
	}
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v15 = F_pgstat_prep_pending_entry(m, int32(2), v11, base.I64_extend_i32_u(v12), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
		*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)) = uint8(v10)
		*(*int32)(unsafe.Add(mBase, uint32(v17))) = v12
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+53)))
		if v20 == int32(0) {
			v23 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
			v30 = v23
		} else {
			v24 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
			*(*int64)(unsafe.Add(mBase, uint32(l2))) = v24
			v26 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
			*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v26
			v28 = *(*int64)(unsafe.Add(mBase, uint32(l2)+40))
			*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v28
			v30 = v24
		}
		v31 = *(*int64)(unsafe.Add(mBase, uint32(v17)+40))
		*(*int64)(unsafe.Add(mBase, uint32(v17)+40)) = v31 + v30
		v34 = *(*int64)(unsafe.Add(mBase, uint32(v17)+48))
		v35 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v17)+48)) = v34 + v35
		v38 = *(*int64)(unsafe.Add(mBase, uint32(v17)+56))
		v39 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v17)+56)) = v38 + v39
		v42 = *(*int64)(unsafe.Add(mBase, uint32(v17)+96))
		v43 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
		v44 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
		*(*int64)(unsafe.Add(mBase, uint32(v17)+96)) = v42 + (v43 + v44)
		return
	}
}
