package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FlushBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int64
	_ = v180
	var v181 int64
	_ = v181
	var v185 int64
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int64
	_ = v200
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v212 int64
	_ = v212
	var v213 int64
	_ = v213
	var v217 int64
	_ = v217
	var v224 int32
	_ = v224
	var v226 int64
	_ = v226
	var v228 int64
	_ = v228
	var v231 int32
	_ = v231
	var v233 int64
	_ = v233
	var v236 int32
	_ = v236
	var v238 int64
	_ = v238
	var v263 int32
	_ = v263
	var v270 int64
	_ = v270
	var v274 int32
	_ = v274
	var v286 int32
	_ = v286
	var v293 int64
	_ = v293
	var v297 int32
	_ = v297
	var v308 int32
	_ = v308
	var v315 int64
	_ = v315
	var v321 int64
	_ = v321
	var v326 int32
	_ = v326
	var v334 int32
	_ = v334
	var v336 int64
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v16 = F_StartBufferIO(m, l0, v4, v4)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v16 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(1074)
	v21 = int32(4418184)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v10 + int32(-36)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v22
	if l1 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	m.G0 = v12 - int32(-64)
	return
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v34
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v37
	v40 = F_smgropen(m, v12, int32(-1))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v42 = l1
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = int32(215570)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = int32(464766)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = int64(0)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v55 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v54 | v55
	if v54&v55 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v42 = v40
	goto L8
L10:
	;
	goto L13
L11:
	;
	v83 = v54
	goto L12
L12:
	;
	v91 = int32(4047244)
	v92 = *(*int32)(unsafe.Add(mBase, _consts[573]))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(-24))+8))
	if v94 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L13:
	;
	F_perform_spin_delay(m, v10+int32(-24))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	v83 = v73
	goto L12
L15:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v74 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v73 | v74
	if v73&v74 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v116 = v112 + v113<<(uint(int32(13))%32)
	v117 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v116)+4)))
	v118 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v116))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v83 & int32(-272629761)
	if v83 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, _consts[573])) = v109
	goto L18
L20:
	;
	if int32(999) < v92 {
		goto L18
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if v92 < int32(11) {
		goto L18
	} else {
		goto L27
	}
L23:
	;
	v99 = int32(900)
	if v99 <= v92 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v102 = v99
	goto L26
L25:
	;
	v102 = v92
	goto L26
L26:
	;
	v109 = v102 + int32(100)
	goto L19
L27:
	;
	v109 = v92 - int32(1)
	goto L19
L28:
	;
	F_XLogFlush(m, v118<<(uint(int64(32))%64)|v117)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v131 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v135 = v131 + v132<<(uint(int32(13))%32)
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v135)+14)))
	if v136 == int32(0) {
		v169 = v135
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L30
L32:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, _consts[267])))
	v174 = m.G0
	v176 = v174 - int32(16)
	m.G0 = v176
	if v171 != 0 {
		goto L45
	} else {
		goto L46
	}
L33:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _consts[178]))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+252))
	goto L34
L34:
	;
	if base.B2i32(v141 != int32(0)) == int32(0) {
		v169 = v135
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _consts[575]))
	if v147 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	v156 = F_MemoryContextAllocAligned(m, v152, int32(8192), int32(4096), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	v159 = v147
	goto L38
L38:
	;
	goto L41
L39:
	;
	*(*int32)(unsafe.Add(mBase, _consts[575])) = v156
	v159 = v156
	goto L38
L40:
	;
	v163 = F_pg_checksum_page(m, v161, v129)
	mBase = m.M
	v165 = *(*int32)(unsafe.Add(mBase, _consts[575]))
	*(*uint16)(unsafe.Add(mBase, uint32(v165)+8)) = uint16(v163)
	v169 = v165
	goto L32
L41:
	;
	v161 = F__emscripten_memcpy_bulkmem(m, v159, v135, int32(8192))
	mBase = m.M
	goto L43
L43:
	;
	goto L40
L44:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v169
	F_smgrwritev(m, v42, v190, v189, v10+int32(-24), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L48
	}
L45:
	;
	F___clock_gettime(m, int32(1), v176)
	mBase = m.M
	v180 = int64(*(*int32)(unsafe.Add(mBase, uint32(v176)+8)))
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v176)))
	v185 = v180 + v181*int64(1000000000)
	goto L47
L46:
	;
	v185 = int64(0)
	goto L47
L47:
	;
	m.G0 = v176 + int32(16)
	goto L44
L48:
	;
	v197 = int32(0)
	v199 = int32(1)
	v200 = int64(8192)
	v204 = m.G0
	v206 = v204 - int32(16)
	m.G0 = v206
	if v185 != int64(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v334 = int32(4323568)
	v336 = *(*int64)(unsafe.Add(mBase, _consts[332]))
	*(*int64)(unsafe.Add(mBase, _consts[332])) = v336 + int64(1)
	v340 = int32(1)
	v341 = int32(0)
	F_TerminateBufferIO(m, l0, v340, v341, v340, v341)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L66
	}
L50:
	;
	F___clock_gettime(m, int32(1), v206)
	mBase = m.M
	v212 = int64(*(*int32)(unsafe.Add(mBase, uint32(v206)+8)))
	v213 = *(*int64)(unsafe.Add(mBase, uint32(v206)))
	v217 = v212 + (v213*int64(1000000000) - v185)
	goto L54
L51:
	;
	goto L52
L52:
	;
	v308 = l2 << (uint(int32(6)) % 32)
	v315 = *(*int64)(unsafe.Add(mBase, uint32(v308)+uint32(_consts[576])))
	*(*int64)(unsafe.Add(mBase, uint32(v308)+uint32(_consts[576]))) = v315 + base.I64_extend_i32_u(v199)
	v321 = *(*int64)(unsafe.Add(mBase, uint32(v308)+uint32(_consts[577])))
	*(*int64)(unsafe.Add(mBase, uint32(v308)+uint32(_consts[577]))) = v321 + v200
	F_pgstat_count_backend_io_op(m, v197, l2, int32(7), v199, v200)
	mBase = m.M
	v326 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[137])) = uint8(v326)
	*(*uint8)(unsafe.Add(mBase, _consts[170])) = uint8(v326)
	m.G0 = v206 + int32(16)
	goto L49
L53:
	;
	v263 = l2 << (uint(int32(6)) % 32)
	v270 = *(*int64)(unsafe.Add(mBase, uint32(v263)+uint32(_consts[578])))
	*(*int64)(unsafe.Add(mBase, uint32(v263)+uint32(_consts[578]))) = v270 + v217
	v274 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	if base.Ui32(int32(16)) < base.Ui32(v274) {
		goto L63
	} else {
		goto L64
	}
L54:
	;
	goto L55
L55:
	;
	v224 = int32(4405008)
	v226 = *(*int64)(unsafe.Add(mBase, _consts[579]))
	v228 = base.I64_div_s(v217, int64(1000))
	*(*int64)(unsafe.Add(mBase, _consts[579])) = v226 + v228
	switch v197 {
	case 0:
		goto L59
	case 1:
		goto L58
	default:
		goto L53
	}
L58:
	;
	v236 = int32(4323648)
	v238 = *(*int64)(unsafe.Add(mBase, _consts[342]))
	*(*int64)(unsafe.Add(mBase, _consts[342])) = v238 + v217
	goto L53
L59:
	;
	v231 = int32(4323632)
	v233 = *(*int64)(unsafe.Add(mBase, _consts[340]))
	*(*int64)(unsafe.Add(mBase, _consts[340])) = v233 + v217
	goto L53
L63:
	;
	goto L52
L64:
	;
	if int32(1)<<(uint(v274)%32)&int32(115186) == int32(0) {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v286 = l2 << (uint(int32(6)) % 32)
	v293 = *(*int64)(unsafe.Add(mBase, uint32(v286)+uint32(_consts[580])))
	*(*int64)(unsafe.Add(mBase, uint32(v286)+uint32(_consts[580]))) = v293 + v217
	v297 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[137])) = uint8(v297)
	*(*uint8)(unsafe.Add(mBase, _consts[174])) = uint8(v297)
	goto L63
L66:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v347
	goto L5
}
func F_UnpinBufferNoOwner(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v84 int32
	_ = v84
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = v11 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v13
	v15 = int32(4341120)
	v17 = *(*int32)(unsafe.Add(mBase, _consts[563]))
	if v13 == v17 {
		v60 = v15
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v63 = v61 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+4)) = v63
	if v63 != 0 {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v19 = int32(4341128)
	v21 = *(*int32)(unsafe.Add(mBase, _consts[564]))
	if v13 == v21 {
		v60 = v19
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = int32(4341136)
	v25 = *(*int32)(unsafe.Add(mBase, _consts[565]))
	if v13 == v25 {
		v60 = v23
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = int32(4341144)
	v29 = *(*int32)(unsafe.Add(mBase, _consts[566]))
	if v13 == v29 {
		v60 = v27
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v31 = int32(4341152)
	v33 = *(*int32)(unsafe.Add(mBase, _consts[567]))
	if v13 == v33 {
		v60 = v31
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v35 = int32(4341160)
	v37 = *(*int32)(unsafe.Add(mBase, _consts[568]))
	if v13 == v37 {
		v60 = v35
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v39 = int32(4341168)
	v41 = *(*int32)(unsafe.Add(mBase, _consts[569]))
	if v13 == v41 {
		v60 = v39
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v43 = int32(4341176)
	v45 = *(*int32)(unsafe.Add(mBase, _consts[570]))
	if v13 == v45 {
		v60 = v43
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v47 = int32(0)
	v49 = *(*int32)(unsafe.Add(mBase, _consts[571]))
	if v49 == v47 {
		v60 = v47
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[572]))
	v56 = int32(0)
	v58 = F_hash_search(m, v53, v9+int32(4), v56, v56)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	v60 = v58
	goto L1
L13:
	;
	m.G0 = v9 + int32(32)
	return
L14:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v67 = v65
	goto L15
L15:
	;
	if v67&int32(4194304) != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v136&int32(536870912) == int32(0) {
		goto L44
	} else {
		goto L45
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(428699)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(6287)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(464766)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(0)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v84&int32(4194304) != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v130 = v67
	goto L19
L19:
	;
	v136 = v130 - int32(1)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v138 = base.B2i32(v130 == v137)
	if v130 == v137 {
		goto L38
	} else {
		goto L39
	}
L20:
	;
	goto L23
L21:
	;
	v101 = v84
	goto L22
L22:
	;
	v109 = int32(4047244)
	v110 = *(*int32)(unsafe.Add(mBase, _consts[573]))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(4))+8))
	if v112 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L23:
	;
	F_perform_spin_delay(m, v9+int32(4))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L11
	} else {
		goto L25
	}
L24:
	;
	v101 = v97
	goto L22
L25:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v97&int32(4194304) != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v130 = v101
	goto L19
L28:
	;
	goto L27
L29:
	;
	*(*int32)(unsafe.Add(mBase, _consts[573])) = v127
	goto L28
L30:
	;
	if int32(999) < v110 {
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if v110 < int32(11) {
		goto L28
	} else {
		goto L37
	}
L33:
	;
	v117 = int32(900)
	if v117 <= v110 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v120 = v117
	goto L36
L35:
	;
	v120 = v110
	goto L36
L36:
	;
	v127 = v120 + int32(100)
	goto L29
L37:
	;
	v127 = v110 - int32(1)
	goto L29
L38:
	;
	v139 = v136
	goto L40
L39:
	;
	v139 = v137
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v139
	if v138 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v67 = v137
	goto L15
L42:
	;
	goto L43
L43:
	;
	goto L16
L44:
	;
	if base.Ui32(v60) < base.Ui32(int32(4341120)) {
		goto L68
	} else {
		goto L69
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(215570)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(464766)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(0)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v158 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v157 | v158
	if v157&v158 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	goto L49
L47:
	;
	v180 = v157
	goto L48
L48:
	;
	v188 = int32(4047244)
	v189 = *(*int32)(unsafe.Add(mBase, _consts[573]))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(4))+8))
	if v191 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L49:
	;
	F_perform_spin_delay(m, v9+int32(4))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L11
	} else {
		goto L51
	}
L50:
	;
	v180 = v173
	goto L48
L51:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v174 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v173 | v174
	if v173&v174 != 0 {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	if v180&int32(537133055) == int32(536870913) {
		goto L64
	} else {
		goto L65
	}
L54:
	;
	goto L53
L55:
	;
	*(*int32)(unsafe.Add(mBase, _consts[573])) = v206
	goto L54
L56:
	;
	if int32(999) < v189 {
		goto L54
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if v189 < int32(11) {
		goto L54
	} else {
		goto L63
	}
L59:
	;
	v196 = int32(900)
	if v196 <= v189 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v199 = v196
	goto L62
L61:
	;
	v199 = v189
	goto L62
L62:
	;
	v206 = v199 + int32(100)
	goto L55
L63:
	;
	v206 = v189 - int32(1)
	goto L55
L64:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v180 & int32(-541327359)
	F_ProcSendSignal(m, v212)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L11
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v180 & int32(-4194305)
	goto L44
L67:
	;
	goto L44
L68:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v235
	v238 = *(*int32)(unsafe.Add(mBase, _consts[572]))
	v244 = F_hash_search(m, v238, v9+int32(4), int32(2), v9+int32(31))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L11
	} else {
		goto L71
	}
L69:
	;
	if base.Ui32(int32(4341184)) <= base.Ui32(v60) {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[574])) = v60
	goto L13
L71:
	;
	v246 = int32(4341080)
	v248 = *(*int32)(unsafe.Add(mBase, _consts[571]))
	*(*int32)(unsafe.Add(mBase, _consts[571])) = v248 - int32(1)
	goto L13
}
func F_buffer_readv_report(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	v11 = m.G0
	v13 = v11 - int32(192)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v23 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+20)))
	if v25&int32(256) != 0 {
		v28 = v23
	} else {
		v28 = int32(-1)
	}
	F_GetRelationPath(m, v13+int32(120), v19, v20, v21, v28, base.I32_extend8_s(v25))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return
	} else {
		v34 = v16 + v15 - int32(1)
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v37 = int32(base.Ui32(v35) >> (uint(int32(25)) % 32))
		v39 = int32(base.Ui32(v35) >> (uint(int32(18)) % 32))
		v41 = int32(base.Ui32(v35) >> (uint(int32(11)) % 32))
		v43 = v41 & int32(127)
		if v35&int32(512) != 0 {
			if v35&int32(1024) != 0 {
				v49 = F_errstart(m, l2, int32(0))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					if v49 == int32(0) {
						m.G0 = v13 + int32(192)
						return
					} else {
						F_errcode(m, int32(16779816))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v13)+112)) = v13 + int32(120)
							*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = v34
							*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v15
							v61 = int32(127)
							v62 = v39 & v61
							*(*int32)(unsafe.Add(mBase, uint32(v13)+100)) = v62
							v65 = v41 & v61
							*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v65
							F_errmsg(m, int32(656943), v13+int32(96))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return
							} else {
								if base.Ui32(int32(2)) <= base.Ui32(v43) {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v15 + v37
									F_errdetail(m, int32(593296), v13+int32(80))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return
									} else {
										v83 = v62 + v65 - int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v83
										F_errhint_plural(m, int32(574579), int32(546121), v83, v13-int32(-64))
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return
										} else {
											v164 = int32(7325)
											F_errfinish(m, int32(464766), v164, int32(75834))
											mBase = m.M
											v172 = m.ExcPending
											if v172 != 0 {
												return
											} else {
												m.G0 = v13 + int32(192)
												return
											}
										}
									}
								} else {
									v83 = v62 + v65 - int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v83
									F_errhint_plural(m, int32(574579), int32(546121), v83, v13-int32(-64))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										v164 = int32(7325)
										F_errfinish(m, int32(464766), v164, int32(75834))
										mBase = m.M
										v172 = m.ExcPending
										if v172 != 0 {
											return
										} else {
											m.G0 = v13 + int32(192)
											return
										}
									}
								}
							}
						}
					}
				}
			} else {
				if v35&int32(448) != int32(256) {
					v114 = v43
					v115 = int32(383419)
					v116 = int32(611454)
					v117 = int32(593296)
					v118 = int32(656877)
				} else {
					v114 = v43
					v115 = int32(657090)
					v116 = int32(611354)
					v117 = int32(593220)
					v118 = int32(656889)
				}
				v120 = F_errstart(m, l2, int32(0))
				mBase = m.M
				v121 = m.ExcPending
				if v121 != 0 {
					return
				} else {
					if v120 == int32(0) {
						m.G0 = v13 + int32(192)
						return
					} else {
						F_errcode(m, int32(16779816))
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return
						} else {
							if v114 == int32(1) {
								*(*int32)(unsafe.Add(mBase, uint32(v13))) = v15 + v37
								*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v13 + int32(120)
								F_errmsg_internal(m, v115, v13)
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
									return
								} else {
									v164 = int32(7368)
									F_errfinish(m, int32(464766), v164, int32(75834))
									mBase = m.M
									v172 = m.ExcPending
									if v172 != 0 {
										return
									} else {
										m.G0 = v13 + int32(192)
										return
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v34
								*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v15
								*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v114
								*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = v13 + int32(120)
								F_errmsg_internal(m, v118, v13+int32(48))
								mBase = m.M
								v146 = m.ExcPending
								if v146 != 0 {
									return
								} else {
									v147 = int32(7368)
									if v114 == int32(0) {
										v164 = v147
										F_errfinish(m, int32(464766), v164, int32(75834))
										mBase = m.M
										v172 = m.ExcPending
										if v172 != 0 {
											return
										} else {
											m.G0 = v13 + int32(192)
											return
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v15 + v37
										F_errdetail_internal(m, v117, v13+int32(32))
										mBase = m.M
										v155 = m.ExcPending
										if v155 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v114 - int32(1)
											F_errhint_internal(m, v116, v13+int32(16))
											mBase = m.M
											v162 = m.ExcPending
											if v162 != 0 {
												return
											} else {
												v164 = v147
												F_errfinish(m, int32(464766), v164, int32(75834))
												mBase = m.M
												v172 = m.ExcPending
												if v172 != 0 {
													return
												} else {
													m.G0 = v13 + int32(192)
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
		} else {
			if v35&int32(448) == int32(256) {
				v114 = v43
				v115 = int32(657090)
				v116 = int32(611354)
				v117 = int32(593220)
				v118 = int32(656889)
			} else {
				v114 = v39 & int32(127)
				v115 = int32(657035)
				v116 = int32(611404)
				v117 = int32(593258)
				v118 = int32(656810)
			}
			v120 = F_errstart(m, l2, int32(0))
			mBase = m.M
			v121 = m.ExcPending
			if v121 != 0 {
				return
			} else {
				if v120 == int32(0) {
					m.G0 = v13 + int32(192)
					return
				} else {
					F_errcode(m, int32(16779816))
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return
					} else {
						if v114 == int32(1) {
							*(*int32)(unsafe.Add(mBase, uint32(v13))) = v15 + v37
							*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v13 + int32(120)
							F_errmsg_internal(m, v115, v13)
							mBase = m.M
							v135 = m.ExcPending
							if v135 != 0 {
								return
							} else {
								v164 = int32(7368)
								F_errfinish(m, int32(464766), v164, int32(75834))
								mBase = m.M
								v172 = m.ExcPending
								if v172 != 0 {
									return
								} else {
									m.G0 = v13 + int32(192)
									return
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v34
							*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v15
							*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v114
							*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = v13 + int32(120)
							F_errmsg_internal(m, v118, v13+int32(48))
							mBase = m.M
							v146 = m.ExcPending
							if v146 != 0 {
								return
							} else {
								v147 = int32(7368)
								if v114 == int32(0) {
									v164 = v147
									F_errfinish(m, int32(464766), v164, int32(75834))
									mBase = m.M
									v172 = m.ExcPending
									if v172 != 0 {
										return
									} else {
										m.G0 = v13 + int32(192)
										return
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v15 + v37
									F_errdetail_internal(m, v117, v13+int32(32))
									mBase = m.M
									v155 = m.ExcPending
									if v155 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v114 - int32(1)
										F_errhint_internal(m, v116, v13+int32(16))
										mBase = m.M
										v162 = m.ExcPending
										if v162 != 0 {
											return
										} else {
											v164 = v147
											F_errfinish(m, int32(464766), v164, int32(75834))
											mBase = m.M
											v172 = m.ExcPending
											if v172 != 0 {
												return
											} else {
												m.G0 = v13 + int32(192)
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
func F_show_buffer_usage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int64
	_ = v24
	var v28 int64
	_ = v28
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v38 int64
	_ = v38
	var v41 int64
	_ = v41
	var v44 int64
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v53 int64
	_ = v53
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v60 int64
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v76 int64
	_ = v76
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v107 int64
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v118 int64
	_ = v118
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v129 int64
	_ = v129
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int64
	_ = v155
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v166 int64
	_ = v166
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v177 int64
	_ = v177
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v188 int64
	_ = v188
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int64
	_ = v213
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v224 int64
	_ = v224
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int64
	_ = v260
	var v263 int32
	_ = v263
	var v272 int32
	_ = v272
	var v274 int64
	_ = v274
	var v277 int32
	_ = v277
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int64
	_ = v301
	var v304 int32
	_ = v304
	var v313 int32
	_ = v313
	var v315 int64
	_ = v315
	var v318 int32
	_ = v318
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int64
	_ = v343
	var v346 int32
	_ = v346
	var v355 int32
	_ = v355
	var v357 int64
	_ = v357
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v381 int64
	_ = v381
	var v383 int32
	_ = v383
	var v386 int64
	_ = v386
	var v388 int32
	_ = v388
	var v391 int64
	_ = v391
	var v393 int32
	_ = v393
	var v396 int64
	_ = v396
	var v398 int32
	_ = v398
	var v401 int64
	_ = v401
	var v403 int32
	_ = v403
	var v406 int64
	_ = v406
	var v408 int32
	_ = v408
	var v411 int64
	_ = v411
	var v413 int32
	_ = v413
	var v416 int64
	_ = v416
	var v418 int32
	_ = v418
	var v421 int64
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v430 int64
	_ = v430
	var v436 int32
	_ = v436
	var v439 int64
	_ = v439
	var v445 int32
	_ = v445
	var v448 int64
	_ = v448
	var v454 int32
	_ = v454
	var v457 int64
	_ = v457
	var v463 int32
	_ = v463
	var v466 int64
	_ = v466
	var v472 int32
	_ = v472
	var v475 int64
	_ = v475
	var v481 int32
	_ = v481
	v11 = m.G0
	v13 = v11 - int32(256)
	m.G0 = v13
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v16 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(256)
	return
L2:
	;
	v19 = int32(1)
	if int64(0) < v15 {
		v34 = v19
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_ExplainPropertyInteger(m, int32(142799), int32(0), v15, l0)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L28
	} else {
		goto L126
	}
L5:
	;
	v35 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	if int64(0) < v35 {
		v47 = v19
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	if int64(0) < v24 {
		v34 = int32(1)
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v28 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	if int64(0) < v28 {
		v34 = int32(1)
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v31 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	v34 = base.B2i32(int64(0) < v31)
	goto L5
L9:
	;
	v48 = int32(1)
	v50 = *(*int64)(unsafe.Add(mBase, uint32(l1)+64))
	if v50 <= int64(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v38 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	if int64(0) < v38 {
		v47 = v19
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v41 = *(*int64)(unsafe.Add(mBase, uint32(l1)+48))
	if int64(0) < v41 {
		v47 = v19
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v44 = *(*int64)(unsafe.Add(mBase, uint32(l1)+56))
	v47 = base.B2i32(int64(0) < v44)
	goto L9
L13:
	;
	v53 = *(*int64)(unsafe.Add(mBase, uint32(l1)+72))
	v56 = base.B2i32(int64(0) < v53)
	goto L15
L14:
	;
	v56 = v48
	goto L15
L15:
	;
	v57 = *(*int64)(unsafe.Add(mBase, uint32(l1)+80))
	if v57 == int64(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v60 = *(*int64)(unsafe.Add(mBase, uint32(l1)+88))
	v63 = base.B2i32(v60 != int64(0))
	goto L18
L17:
	;
	v63 = v48
	goto L18
L18:
	;
	v64 = int32(1)
	v66 = *(*int64)(unsafe.Add(mBase, uint32(l1)+96))
	if v66 == int64(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v69 = *(*int64)(unsafe.Add(mBase, uint32(l1)+104))
	v72 = base.B2i32(v69 != int64(0))
	goto L21
L20:
	;
	v72 = v64
	goto L21
L21:
	;
	v73 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if v73 == int64(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v76 = *(*int64)(unsafe.Add(mBase, uint32(l1)+120))
	v79 = base.B2i32(v76 != int64(0))
	goto L24
L23:
	;
	v79 = v64
	goto L24
L24:
	;
	if (v34|v47|v56)&int32(1) != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	F_ExplainIndentText(m, l0)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	if v79|(v72|v63) == int32(0) {
		goto L1
	} else {
		goto L85
	}
L28:
	;
	return
L29:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoString(m, v86, int32(514908))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	if v34 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v47 != 0 {
		goto L54
	} else {
		goto L55
	}
L32:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoString(m, v92, int32(424572))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v96 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	if int64(0) < v96 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+240)) = v96
	F_appendStringInfo(m, v99, int32(403745), v13+int32(240))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L28
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v107 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	if int64(0) < v107 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+224)) = v107
	F_appendStringInfo(m, v110, int32(403826), v13+int32(224))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L28
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v118 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	if int64(0) < v118 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L40
L42:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+208)) = v118
	F_appendStringInfo(m, v121, int32(403812), v13+int32(208))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L28
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v129 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	if int64(0) < v129 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L44
L46:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+192)) = v129
	F_appendStringInfo(m, v132, int32(403769), v13+int32(192))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L28
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if (v56|v47)&int32(1) == int32(0) {
		goto L31
	} else {
		goto L50
	}
L49:
	;
	goto L48
L50:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoChar(m, v145, int32(44))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L28
	} else {
		goto L51
	}
L51:
	;
	goto L31
L52:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoChar(m, v237, int32(10))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L28
	} else {
		goto L84
	}
L53:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoString(m, v209, int32(221157))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L28
	} else {
		goto L77
	}
L54:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoString(m, v151, int32(294882))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L28
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v56 == int32(0) {
		goto L52
	} else {
		goto L76
	}
L57:
	;
	v155 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	if int64(0) < v155 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+176)) = v155
	F_appendStringInfo(m, v158, int32(403745), v13+int32(176))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L28
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v166 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	if int64(0) < v166 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L60
L62:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+160)) = v166
	F_appendStringInfo(m, v169, int32(403826), v13+int32(160))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L28
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v177 = *(*int64)(unsafe.Add(mBase, uint32(l1)+48))
	if int64(0) < v177 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	goto L64
L66:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+144)) = v177
	F_appendStringInfo(m, v180, int32(403812), v13+int32(144))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L28
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v188 = *(*int64)(unsafe.Add(mBase, uint32(l1)+56))
	if int64(0) < v188 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L68
L70:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+128)) = v188
	F_appendStringInfo(m, v191, int32(403769), v13+int32(128))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L28
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	if v56 == int32(0) {
		goto L52
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoChar(m, v201, int32(44))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L28
	} else {
		goto L75
	}
L75:
	;
	goto L53
L76:
	;
	goto L53
L77:
	;
	v213 = *(*int64)(unsafe.Add(mBase, uint32(l1)+64))
	if int64(0) < v213 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+112)) = v213
	F_appendStringInfo(m, v216, int32(403826), v13+int32(112))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L28
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v224 = *(*int64)(unsafe.Add(mBase, uint32(l1)+72))
	if v224 <= int64(0) {
		goto L52
	} else {
		goto L82
	}
L81:
	;
	goto L80
L82:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v224
	F_appendStringInfo(m, v227, int32(403769), v13+int32(96))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L28
	} else {
		goto L83
	}
L83:
	;
	goto L52
L84:
	;
	goto L27
L85:
	;
	F_ExplainIndentText(m, l0)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L28
	} else {
		goto L86
	}
L86:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoString(m, v250, int32(514953))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L28
	} else {
		goto L87
	}
L87:
	;
	if v63 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	if v72 != 0 {
		goto L103
	} else {
		goto L104
	}
L89:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoString(m, v256, int32(424572))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L28
	} else {
		goto L90
	}
L90:
	;
	v260 = *(*int64)(unsafe.Add(mBase, uint32(l1)+80))
	if v260 != int64(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*float64)(unsafe.Add(mBase, uint32(v13)+80)) = base.F64_div(base.F64_convert_i64_s(v260), float64(1e+06))
	F_appendStringInfo(m, v263, int32(318843), v13+int32(80))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L28
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v274 = *(*int64)(unsafe.Add(mBase, uint32(l1)+88))
	if v274 != int64(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	goto L93
L95:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*float64)(unsafe.Add(mBase, uint32(v13)+64)) = base.F64_div(base.F64_convert_i64_s(v274), float64(1e+06))
	F_appendStringInfo(m, v277, int32(318830), v13-int32(-64))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L28
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	if v79|v72 == int32(0) {
		goto L88
	} else {
		goto L99
	}
L98:
	;
	goto L97
L99:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoChar(m, v291, int32(44))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L28
	} else {
		goto L100
	}
L100:
	;
	goto L88
L101:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoChar(m, v371, int32(10))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L28
	} else {
		goto L125
	}
L102:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoString(m, v339, int32(221157))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L28
	} else {
		goto L118
	}
L103:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoString(m, v297, int32(294882))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L28
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	if v79 == int32(0) {
		goto L101
	} else {
		goto L117
	}
L106:
	;
	v301 = *(*int64)(unsafe.Add(mBase, uint32(l1)+96))
	if v301 != int64(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*float64)(unsafe.Add(mBase, uint32(v13)+48)) = base.F64_div(base.F64_convert_i64_s(v301), float64(1e+06))
	F_appendStringInfo(m, v304, int32(318843), v13+int32(48))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L28
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v315 = *(*int64)(unsafe.Add(mBase, uint32(l1)+104))
	if v315 != int64(0) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	goto L109
L111:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*float64)(unsafe.Add(mBase, uint32(v13)+32)) = base.F64_div(base.F64_convert_i64_s(v315), float64(1e+06))
	F_appendStringInfo(m, v318, int32(318830), v13+int32(32))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L28
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	if v79 == int32(0) {
		goto L101
	} else {
		goto L115
	}
L114:
	;
	goto L113
L115:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendStringInfoChar(m, v331, int32(44))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L28
	} else {
		goto L116
	}
L116:
	;
	goto L102
L117:
	;
	goto L102
L118:
	;
	v343 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if v343 != int64(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*float64)(unsafe.Add(mBase, uint32(v13)+16)) = base.F64_div(base.F64_convert_i64_s(v343), float64(1e+06))
	F_appendStringInfo(m, v346, int32(318843), v13+int32(16))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L28
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v357 = *(*int64)(unsafe.Add(mBase, uint32(l1)+120))
	if v357 == int64(0) {
		goto L101
	} else {
		goto L123
	}
L122:
	;
	goto L121
L123:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*float64)(unsafe.Add(mBase, uint32(v13))) = base.F64_div(base.F64_convert_i64_s(v357), float64(1e+06))
	F_appendStringInfo(m, v360, int32(318830), v13)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L28
	} else {
		goto L124
	}
L124:
	;
	goto L101
L125:
	;
	goto L1
L126:
	;
	v381 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	F_ExplainPropertyInteger(m, int32(143008), int32(0), v381, l0)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L28
	} else {
		goto L127
	}
L127:
	;
	v386 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	F_ExplainPropertyInteger(m, int32(142951), int32(0), v386, l0)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L28
	} else {
		goto L128
	}
L128:
	;
	v391 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	F_ExplainPropertyInteger(m, int32(142908), int32(0), v391, l0)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L28
	} else {
		goto L129
	}
L129:
	;
	v396 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	F_ExplainPropertyInteger(m, int32(142782), int32(0), v396, l0)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L28
	} else {
		goto L130
	}
L130:
	;
	v401 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	F_ExplainPropertyInteger(m, int32(142990), int32(0), v401, l0)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L28
	} else {
		goto L131
	}
L131:
	;
	v406 = *(*int64)(unsafe.Add(mBase, uint32(l1)+48))
	F_ExplainPropertyInteger(m, int32(142930), int32(0), v406, l0)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L28
	} else {
		goto L132
	}
L132:
	;
	v411 = *(*int64)(unsafe.Add(mBase, uint32(l1)+56))
	F_ExplainPropertyInteger(m, int32(142887), int32(0), v411, l0)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L28
	} else {
		goto L133
	}
L133:
	;
	v416 = *(*int64)(unsafe.Add(mBase, uint32(l1)+64))
	F_ExplainPropertyInteger(m, int32(142973), int32(0), v416, l0)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L28
	} else {
		goto L134
	}
L134:
	;
	v421 = *(*int64)(unsafe.Add(mBase, uint32(l1)+72))
	F_ExplainPropertyInteger(m, int32(142867), int32(0), v421, l0)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L28
	} else {
		goto L135
	}
L135:
	;
	v425 = int32(*(*uint8)(unsafe.Add(mBase, _consts[267])))
	if v425 != int32(1) {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v430 = *(*int64)(unsafe.Add(mBase, uint32(l1)+80))
	F_ExplainPropertyFloat(m, int32(353563), int32(140554), base.F64_div(base.F64_convert_i64_s(v430), float64(1e+06)), int32(3), l0)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L28
	} else {
		goto L137
	}
L137:
	;
	v439 = *(*int64)(unsafe.Add(mBase, uint32(l1)+88))
	F_ExplainPropertyFloat(m, int32(353502), int32(140554), base.F64_div(base.F64_convert_i64_s(v439), float64(1e+06)), int32(3), l0)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L28
	} else {
		goto L138
	}
L138:
	;
	v448 = *(*int64)(unsafe.Add(mBase, uint32(l1)+96))
	F_ExplainPropertyFloat(m, int32(353543), int32(140554), base.F64_div(base.F64_convert_i64_s(v448), float64(1e+06)), int32(3), l0)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L28
	} else {
		goto L139
	}
L139:
	;
	v457 = *(*int64)(unsafe.Add(mBase, uint32(l1)+104))
	F_ExplainPropertyFloat(m, int32(353481), int32(140554), base.F64_div(base.F64_convert_i64_s(v457), float64(1e+06)), int32(3), l0)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L28
	} else {
		goto L140
	}
L140:
	;
	v466 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	F_ExplainPropertyFloat(m, int32(353524), int32(140554), base.F64_div(base.F64_convert_i64_s(v466), float64(1e+06)), int32(3), l0)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L28
	} else {
		goto L141
	}
L141:
	;
	v475 = *(*int64)(unsafe.Add(mBase, uint32(l1)+120))
	F_ExplainPropertyFloat(m, int32(353461), int32(140554), base.F64_div(base.F64_convert_i64_s(v475), float64(1e+06)), int32(3), l0)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L28
	} else {
		goto L142
	}
L142:
	;
	goto L1
}
