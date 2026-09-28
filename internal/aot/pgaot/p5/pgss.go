package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_pgss_shmem_init(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int64
	_ = v124
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v141 int64
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
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
	var v180 int32
	_ = v180
	var v182 int64
	_ = v182
	var v184 int64
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int64
	_ = v215
	var v217 int64
	_ = v217
	var v224 int64
	_ = v224
	var v225 int64
	_ = v225
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(512)
	m.G0 = v14
	v17 = F_LWLockNewTrancheId(m, int32(_a_F_pgss_shmem_init_0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_init[0]))
	F_LWLockInitialize(m, v20, v17)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_init[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+136)) = int32(1024)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+128)) = int64(4621819117588971520)
	v29 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v24)+140)), uint32(v29))
	v32 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+160)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v24)+152)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v24)+144)) = v32
	v41 = m.G0
	v42 = int32(16)
	v43 = v41 - v42
	m.G0 = v43
	F_gettimeofday(m, v43)
	mBase = m.M
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
	v47 = int64(*(*int32)(unsafe.Add(mBase, uint32(v43)+8)))
	m.G0 = v43 + v42
	goto L4
L4:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_init[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v57)+168)) = v47 + v46*int64(1000000) - int64(946684800000000)
	F_on_shmem_exit(m, int32(_a_F_pgss_shmem_init_1), int64(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v63 = int32(_a_F_pgss_shmem_init_2)
	v64 = F_unlink(m, v63)
	mBase = m.M
	v69 = F_AllocateFile(m, v63, int32(_a_F_pgss_shmem_init_3))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L10
	}
L6:
	;
	m.G0 = v14 + int32(512)
	return
L7:
	;
	if v354 != 0 {
		goto L75
	} else {
		goto L76
	}
L8:
	;
	F_errfinish(m, int32(_a_F_pgss_shmem_init_4), v347, int32(_a_F_pgss_shmem_init_5))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L74
	}
L9:
	;
	v324 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L70
	}
L10:
	;
	if v69 == int32(0) {
		v312 = int32(0)
		v314 = v2
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgss_shmem_init[1])))
	if v74 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v77 = F_FreeFile(m, v69)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v81 = F_AllocateFile(m, int32(_a_F_pgss_shmem_init_6), int32(_a_F_pgss_shmem_init_7))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L18
	}
L15:
	;
	goto L6
L16:
	;
	v296 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L66
	}
L17:
	;
	v269 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L62
	}
L18:
	;
	if v81 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_init[2]))
	if v86 != int32(44) {
		v259 = v2
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v92 = F_palloc(m, int32(2048))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	v89 = F_FreeFile(m, v69)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	goto L6
L24:
	;
	v98 = F_fread(m, v14+int32(508), int32(4), int32(1), v81)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v98 != int32(1) {
		v259 = v92
		goto L17
	} else {
		goto L26
	}
L26:
	;
	v106 = F_fread(m, v14+int32(492), int32(4), int32(1), v81)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v106 != int32(1) {
		v259 = v92
		goto L17
	} else {
		goto L28
	}
L28:
	;
	v114 = F_fread(m, v14+int32(496), int32(8), int32(1), v81)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if v114 != int32(1) {
		v259 = v92
		goto L17
	} else {
		goto L30
	}
L30:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v14)+508))
	if v118 != int32(539297585) {
		v286 = v92
		goto L16
	} else {
		goto L31
	}
L31:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v14)+492))
	if v121 != int32(1900) {
		v286 = v92
		goto L16
	} else {
		goto L32
	}
L32:
	;
	v124 = *(*int64)(unsafe.Add(mBase, uint32(v14)+496))
	if int64(0) < v124 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v134 = v92
	v136 = int32(2048)
	v141 = int64(0)
	goto L36
L34:
	;
	v230 = v92
	goto L35
L35:
	;
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_init[0]))
	v244 = F_fread(m, v239+int32(160), int32(16), int32(1), v81)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L57
	}
L36:
	;
	v146 = F_fread(m, v14+int32(40), int32(448), int32(1), v81)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L38
	}
L37:
	;
	v230 = v168
	goto L35
L38:
	;
	if v146 != int32(1) {
		v259 = v134
		goto L17
	} else {
		goto L39
	}
L39:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v14)+456))
	if base.B2i32(base.Ui32(int32(34)) < base.Ui32(v150))|base.B2i32(v150 == int32(7)) != 0 {
		v286 = v134
		goto L16
	} else {
		goto L40
	}
L40:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v14)+452))
	if v136 <= v156 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v158 = int32(1)
	v159 = v136 << (uint(v158) % 32)
	v161 = v156 + v158
	if v161 < v159 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v167 = v156
	v168 = v134
	v169 = v136
	goto L43
L43:
	;
	v170 = int32(1)
	v173 = F_fread(m, v168, v170, v167+v170, v81)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L48
	}
L44:
	;
	v163 = v159
	goto L46
L45:
	;
	v163 = v161
	goto L46
L46:
	;
	v164 = F_repalloc(m, v134, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v14)+452))
	v167 = v166
	v168 = v164
	v169 = v163
	goto L43
L48:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v14)+452))
	if v173 != v175+int32(1) {
		v259 = v168
		goto L17
	} else {
		goto L49
	}
L49:
	;
	v180 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v175+v168))) = uint8(v180)
	v182 = *(*int64)(unsafe.Add(mBase, uint32(v14)+64))
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v14)+72))
	if v182 != int64(0)-v184 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_init[0]))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+144))
	v190 = int32(1)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v14)+452))
	v194 = F_fwrite(m, v168, v190, v191+v190, v69)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v224 = v141 + int64(1)
	v225 = *(*int64)(unsafe.Add(mBase, uint32(v14)+496))
	if v224 < v225 {
		v134 = v168
		v136 = v169
		v141 = v224
		goto L36
	} else {
		goto L56
	}
L53:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v14)+452))
	if v194 != v196+int32(1) {
		v312 = v81
		v314 = v168
		goto L9
	} else {
		goto L54
	}
L54:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_init[0]))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v201)+144)) = v202 + v194
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v14)+456))
	v209 = F_entry_alloc(m, v14+int32(40), v189, v196, v207, int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	base.MemoryCopy(m, v209+int32(24), v14-int32(-64), int32(384))
	v215 = *(*int64)(unsafe.Add(mBase, uint32(v14)+464))
	*(*int64)(unsafe.Add(mBase, uint32(v209)+424)) = v215
	v217 = *(*int64)(unsafe.Add(mBase, uint32(v14)+472))
	*(*int64)(unsafe.Add(mBase, uint32(v209)+432)) = v217
	goto L52
L56:
	;
	goto L37
L57:
	;
	if v244 != int32(1) {
		v259 = v230
		goto L17
	} else {
		goto L58
	}
L58:
	;
	F_pfree(m, v230)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v250 = F_FreeFile(m, v81)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v252 = F_FreeFile(m, v69)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v255 = F_unlink(m, int32(_a_F_pgss_shmem_init_6))
	mBase = m.M
	goto L6
L62:
	;
	if v269 == int32(0) {
		v352 = v81
		v354 = v259
		goto L7
	} else {
		goto L63
	}
L63:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_pgss_shmem_init_6)
	F_errmsg(m, int32(_a_F_pgss_shmem_init_8), v14+int32(16))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v337 = v81
	v339 = v259
	v347 = int32(697)
	goto L8
L66:
	;
	if v296 == int32(0) {
		v352 = v81
		v354 = v286
		goto L7
	} else {
		goto L67
	}
L67:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(_a_F_pgss_shmem_init_6)
	F_errmsg(m, int32(_a_F_pgss_shmem_init_9), v14+int32(32))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v337 = v81
	v339 = v286
	v347 = int32(703)
	goto L8
L70:
	;
	if v324 == int32(0) {
		v352 = v312
		v354 = v314
		goto L7
	} else {
		goto L71
	}
L71:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(_a_F_pgss_shmem_init_2)
	F_errmsg(m, int32(_a_F_pgss_shmem_init_10), v14)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v337 = v312
	v339 = v314
	v347 = int32(709)
	goto L8
L74:
	;
	v352 = v337
	v354 = v339
	goto L7
L75:
	;
	F_pfree(m, v354)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	if v352 != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L77
L79:
	;
	v364 = F_FreeFile(m, v352)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	if v69 != 0 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	goto L81
L83:
	;
	v366 = F_FreeFile(m, v69)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v369 = F_unlink(m, int32(_a_F_pgss_shmem_init_6))
	mBase = m.M
	goto L6
L86:
	;
	goto L85
}
func F_pgss_shmem_shutdown(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v121 int64
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v270 int32
	_ = v270
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v3
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v12 + int32(48)
	return
L2:
	;
	v16 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_shutdown[0]))
	if v18 == v16 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_shutdown[1]))
	if v22 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgss_shmem_shutdown[2])))
	if v26&int32(1) == int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v33 = F_AllocateFile(m, int32(_a_F_pgss_shmem_shutdown_0), int32(_a_F_pgss_shmem_shutdown_1))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v270 = F_unlink(m, int32(_a_F_pgss_shmem_shutdown_2))
	mBase = m.M
	goto L1
L7:
	;
	v240 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L8
	} else {
		goto L51
	}
L8:
	;
	return
L9:
	;
	if v33 == int32(0) {
		v229 = v16
		v231 = v3
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v40 = F_fwrite(m, int32(_a_F_pgss_shmem_shutdown_3), int32(4), int32(1), v33)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	v229 = v220
	v231 = v33
	goto L7
L12:
	;
	v220 = int32(0)
	goto L11
L13:
	;
	if v40 != int32(1) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v47 = F_fwrite(m, int32(_a_F_pgss_shmem_shutdown_4), int32(4), int32(1), v33)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	if v47 != int32(1) {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_shutdown[1]))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v54)+8))
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v54)+808))
	if v56 != int64(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v121
	v127 = F_fwrite(m, v12+int32(16), int32(8), int32(1), v33)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L8
	} else {
		goto L21
	}
L18:
	;
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v54)+752))
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v54)+728))
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v54)+704))
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v54)+680))
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v54)+656))
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v54)+632))
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v54)+608))
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v54)+584))
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v54)+560))
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v54)+536))
	v69 = *(*int64)(unsafe.Add(mBase, uint32(v54)+512))
	v70 = *(*int64)(unsafe.Add(mBase, uint32(v54)+488))
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v54)+464))
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v54)+440))
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v54)+416))
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v54)+392))
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v54)+368))
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v54)+344))
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v54)+320))
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v54)+296))
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v54)+272))
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v54)+248))
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v54)+224))
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v54)+200))
	v83 = *(*int64)(unsafe.Add(mBase, uint32(v54)+176))
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v54)+152))
	v85 = *(*int64)(unsafe.Add(mBase, uint32(v54)+128))
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v54)+104))
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v54)+80))
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v54)+56))
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v54)+32))
	v121 = v59 + (v60 + (v61 + (v62 + (v63 + (v64 + (v65 + (v66 + (v67 + (v68 + (v69 + (v70 + (v71 + (v72 + (v73 + (v74 + (v75 + (v76 + (v77 + (v78 + (v79 + (v80 + (v81 + (v82 + (v83 + (v84 + (v85 + (v86 + (v87 + (v88 + (v89 + v55))))))))))))))))))))))))))))))
	goto L20
L19:
	;
	v121 = v55
	goto L20
L20:
	;
	goto L17
L21:
	;
	if v127 != int32(1) {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	v133 = F_qtext_load_file(m, v12+int32(44))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	if v133 == int32(0) {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	v138 = v12 + int32(24)
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_shutdown[1]))
	F_hash_seq_init(m, v138, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	v143 = F_hash_seq_search(m, v138)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	if v143 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v148 = v143
	goto L30
L28:
	;
	goto L29
L29:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_shutdown[0]))
	v203 = F_fwrite(m, v198+int32(160), int32(16), int32(1), v33)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L8
	} else {
		goto L45
	}
L30:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v148)+412))
	if v155 < int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L29
L32:
	;
	v186 = F_hash_seq_search(m, v12+int32(24))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L8
	} else {
		goto L43
	}
L33:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v148)+408))
	v159 = v158 + v155
	if base.Ui32(v145) <= base.Ui32(v159) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133+v159))))
	if v162 != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v165 = F_fwrite(m, v148, int32(448), int32(1), v33)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	if v165 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v170 = int32(1)
	v172 = v155 + v170
	v173 = F_fwrite(m, v133+v158, v170, v172, v33)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L8
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	F_hash_seq_term(m, v12+int32(24))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L8
	} else {
		goto L42
	}
L40:
	;
	if v173 == v172 {
		goto L32
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v220 = v133
	goto L11
L43:
	;
	if v186 != 0 {
		v148 = v186
		goto L30
	} else {
		goto L44
	}
L44:
	;
	goto L31
L45:
	;
	if v203 != int32(1) {
		v220 = v133
		goto L11
	} else {
		goto L46
	}
L46:
	;
	F_pfree(m, v133)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L8
	} else {
		goto L47
	}
L47:
	;
	v209 = int32(0)
	v211 = F_FreeFile(m, v33)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L8
	} else {
		goto L48
	}
L48:
	;
	if v211 != 0 {
		v229 = v209
		v231 = v209
		goto L7
	} else {
		goto L49
	}
L49:
	;
	v216 = F_durable_rename(m, int32(_a_F_pgss_shmem_shutdown_0), int32(_a_F_pgss_shmem_shutdown_5), int32(15))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L8
	} else {
		goto L50
	}
L50:
	;
	goto L6
L51:
	;
	if v240 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L8
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	if v229 != 0 {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_pgss_shmem_shutdown_0)
	F_errmsg(m, int32(_a_F_pgss_shmem_shutdown_6), v12)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L8
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_pgss_shmem_shutdown_7), int32(820), int32(_a_F_pgss_shmem_shutdown_8))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L8
	} else {
		goto L57
	}
L57:
	;
	goto L54
L58:
	;
	F_pfree(m, v229)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L8
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v231 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L60
L62:
	;
	v256 = F_FreeFile(m, v231)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L8
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v259 = F_unlink(m, int32(_a_F_pgss_shmem_shutdown_0))
	mBase = m.M
	goto L6
L65:
	;
	goto L64
}
