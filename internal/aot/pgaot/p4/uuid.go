package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_uuid_cmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v94 int64
	_ = v94
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v151 int64
	_ = v151
	var v155 int64
	_ = v155
	var v156 int64
	_ = v156
	var v160 int64
	_ = v160
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v7 = int64(56)
	v9 = int64(65280)
	v11 = int64(40)
	v14 = int64(16711680)
	v16 = int64(24)
	v18 = int64(4278190080)
	v20 = int64(8)
	v41 = v6<<(uint(v7)%64) | v6&v9<<(uint(v11)%64) | (v6&v14<<(uint(v16)%64) | v6&v18<<(uint(v20)%64)) | (int64(base.Ui64(v6)>>(uint(v20)%64))&v18 | int64(base.Ui64(v6)>>(uint(v16)%64))&v14 | (int64(base.Ui64(v6)>>(uint(v11)%64))&v9 | int64(base.Ui64(v6)>>(uint(v7)%64))))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
	v78 = v43<<(uint(v7)%64) | v43&v9<<(uint(v11)%64) | (v43&v14<<(uint(v16)%64) | v43&v18<<(uint(v20)%64)) | (int64(base.Ui64(v43)>>(uint(v20)%64))&v18 | int64(base.Ui64(v43)>>(uint(v16)%64))&v14 | (int64(base.Ui64(v43)>>(uint(v11)%64))&v9 | int64(base.Ui64(v43)>>(uint(v7)%64))))
	if v41 != v78 {
		v155 = v78
		v156 = v41
		if base.Ui64(v156) < base.Ui64(v155) {
			v160 = int64(-1)
		} else {
			v160 = int64(1)
		}
		return v160
	} else {
		v80 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
		v81 = int64(56)
		v83 = int64(65280)
		v85 = int64(40)
		v88 = int64(16711680)
		v90 = int64(24)
		v92 = int64(4278190080)
		v94 = int64(8)
		v115 = v80<<(uint(v81)%64) | v80&v83<<(uint(v85)%64) | (v80&v88<<(uint(v90)%64) | v80&v92<<(uint(v94)%64)) | (int64(base.Ui64(v80)>>(uint(v94)%64))&v92 | int64(base.Ui64(v80)>>(uint(v90)%64))&v88 | (int64(base.Ui64(v80)>>(uint(v85)%64))&v83 | int64(base.Ui64(v80)>>(uint(v81)%64))))
		v116 = *(*int64)(unsafe.Add(mBase, uint32(v42)+8))
		v151 = v116<<(uint(v81)%64) | v116&v83<<(uint(v85)%64) | (v116&v88<<(uint(v90)%64) | v116&v92<<(uint(v94)%64)) | (int64(base.Ui64(v116)>>(uint(v94)%64))&v92 | int64(base.Ui64(v116)>>(uint(v90)%64))&v88 | (int64(base.Ui64(v116)>>(uint(v85)%64))&v83 | int64(base.Ui64(v116)>>(uint(v81)%64))))
		if v115 != v151 {
			v155 = v151
			v156 = v115
			if base.Ui64(v156) < base.Ui64(v155) {
				v160 = int64(-1)
			} else {
				v160 = int64(1)
			}
			return v160
		} else {
			return int64(0)
		}
	}
}
func F_uuid_eq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v3)+8))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
	return base.I64_extend_i32_u(base.B2i32(v4^v6|(v8^v9) == int64(0)))
}
func F_uuid_extract_timestamp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v21 int64
	_ = v21
	var v25 int64
	_ = v25
	var v30 int64
	_ = v30
	var v34 int64
	_ = v34
	var v44 int64
	_ = v44
	var v49 int64
	_ = v49
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v58 int64
	_ = v58
	var v62 int64
	_ = v62
	var v66 int64
	_ = v66
	var v70 int64
	_ = v70
	var v79 int32
	_ = v79
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = int32(*(*int8)(unsafe.Add(mBase, uint32(v4)+8)))
	if int32(-64) <= v5 {
		v8 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v8)
		return int64(0)
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+6)))
		switch int32(base.Ui32(v12)>>(uint(int32(4))%32)) - int32(1) {
		case 0:
			v17 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v4)+3)))
			v18 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v4)+1)))
			v21 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
			v25 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v4)+2)))
			v30 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v4)+4)))
			v34 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v4)+5)))
			v44 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v4)+7)))
			v49 = base.I64_div_u_s(v17|(v18<<(uint(int64(16))%64)|v21<<(uint(int64(24))%64)|v25<<(uint(int64(8))%64))|v30<<(uint(int64(40))%64)|v34<<(uint(int64(32))%64)+base.I64_extend_i32_u(v12)&int64(15)<<(uint(int64(56))%64)+v44<<(uint(int64(48))%64), int64(10))
			return v49 - int64(13165977600000000)
		default:
			v79 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v79)
			return int64(0)
		case 6:
			v53 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v4)+5)))
			v54 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v4)+4)))
			v58 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v4)+3)))
			v62 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v4)+2)))
			v66 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v4)+1)))
			v70 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
			return (v53|v54<<(uint(int64(8))%64)|v58<<(uint(int64(16))%64)|v62<<(uint(int64(24))%64)|v66<<(uint(int64(32))%64)|v70<<(uint(int64(40))%64))*int64(1000) - int64(946684800000000)
		}
	}
}
func F_uuid_hash_extended(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
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
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
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
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
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
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
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
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = int32(16)
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = int32(-1636608416)
	if v4 == int64(0) {
		v47 = v10
		v49 = v10
		v51 = v10
	} else {
		v13 = base.I32_wrap_i64(v4)
		v15 = v13 + int32(1021750464)
		v19 = int32(4)
		v21 = base.I32_wrap_i64(int64(base.Ui64(v4)>>(uint(int64(32))%64))) ^ base.I32_rotl(v10, v19)
		v25 = v10 + v13 - v21 ^ base.I32_rotl(v21, int32(6))
		v29 = v15 - v25 ^ base.I32_rotl(v25, int32(8))
		v30 = v15 + v21
		v31 = v25 + v30
		v32 = v29 + v31
		v36 = v30 - v29 ^ base.I32_rotl(v29, int32(16))
		v40 = v31 - v36 ^ base.I32_rotl(v36, int32(19))
		v45 = v32 + v36
		v47 = v45
		v49 = v32 - v40 ^ base.I32_rotl(v40, v19)
		v51 = v40 + v45
	}
	if v2&int32(3) != 0 {
		v56 = v2
		v57 = v3
		v59 = v47
		v60 = v51
		v61 = v49
		for {
			v63 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
			v64 = v63 + v60
			v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
			v67 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
			v68 = v67 + v61
			v70 = int32(4)
			v72 = v65 + v59 - v68 ^ base.I32_rotl(v68, v70)
			v76 = v64 - v72 ^ base.I32_rotl(v72, int32(6))
			v77 = v68 + v64
			v78 = v72 + v77
			v79 = v76 + v78
			v83 = v77 - v76 ^ base.I32_rotl(v76, int32(8))
			v87 = v78 - v83 ^ base.I32_rotl(v83, int32(16))
			v91 = v79 - v87 ^ base.I32_rotl(v87, int32(19))
			v92 = v83 + v79
			v93 = v87 + v92
			v94 = v91 + v93
			v98 = v92 - v91 ^ base.I32_rotl(v91, v70)
			v99 = int32(12)
			v100 = v56 + v99
			v102 = v57 - v99
			if base.Ui32(int32(11)) < base.Ui32(v102) {
				v56 = v100
				v57 = v102
				v59 = v93
				v60 = v94
				v61 = v98
				continue
			} else {
				break
			}
			break
		}
		switch v102 - int32(1) {
		case 0:
			v275 = v93
			v276 = v94
			v277 = v98
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 1:
			v268 = v93
			v269 = v94
			v270 = v98
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 2:
			v261 = v93
			v262 = v94
			v263 = v98
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 3:
			v255 = v94
			v256 = v98
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+3)))
			v261 = v257<<(uint(int32(24))%32) + v93
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 4:
			v251 = v94
			v252 = v98
			v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+4)))
			v255 = v251 + v253
			v256 = v252
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+3)))
			v261 = v257<<(uint(int32(24))%32) + v93
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 5:
			v245 = v94
			v246 = v98
			v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+5)))
			v251 = v247<<(uint(int32(8))%32) + v245
			v252 = v246
			v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+4)))
			v255 = v251 + v253
			v256 = v252
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+3)))
			v261 = v257<<(uint(int32(24))%32) + v93
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 6:
			v239 = v94
			v240 = v98
			v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+6)))
			v245 = v241<<(uint(int32(16))%32) + v239
			v246 = v240
			v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+5)))
			v251 = v247<<(uint(int32(8))%32) + v245
			v252 = v246
			v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+4)))
			v255 = v251 + v253
			v256 = v252
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+3)))
			v261 = v257<<(uint(int32(24))%32) + v93
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 7:
			v234 = v98
			v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+7)))
			v239 = v235<<(uint(int32(24))%32) + v94
			v240 = v234
			v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+6)))
			v245 = v241<<(uint(int32(16))%32) + v239
			v246 = v240
			v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+5)))
			v251 = v247<<(uint(int32(8))%32) + v245
			v252 = v246
			v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+4)))
			v255 = v251 + v253
			v256 = v252
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+3)))
			v261 = v257<<(uint(int32(24))%32) + v93
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 8:
			v229 = v98
			v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+8)))
			v234 = v230<<(uint(int32(8))%32) + v229
			v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+7)))
			v239 = v235<<(uint(int32(24))%32) + v94
			v240 = v234
			v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+6)))
			v245 = v241<<(uint(int32(16))%32) + v239
			v246 = v240
			v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+5)))
			v251 = v247<<(uint(int32(8))%32) + v245
			v252 = v246
			v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+4)))
			v255 = v251 + v253
			v256 = v252
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+3)))
			v261 = v257<<(uint(int32(24))%32) + v93
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 9:
			v224 = v98
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+9)))
			v229 = v225<<(uint(int32(16))%32) + v224
			v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+8)))
			v234 = v230<<(uint(int32(8))%32) + v229
			v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+7)))
			v239 = v235<<(uint(int32(24))%32) + v94
			v240 = v234
			v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+6)))
			v245 = v241<<(uint(int32(16))%32) + v239
			v246 = v240
			v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+5)))
			v251 = v247<<(uint(int32(8))%32) + v245
			v252 = v246
			v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+4)))
			v255 = v251 + v253
			v256 = v252
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+3)))
			v261 = v257<<(uint(int32(24))%32) + v93
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 10:
			v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+10)))
			v224 = v220<<(uint(int32(24))%32) + v98
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+9)))
			v229 = v225<<(uint(int32(16))%32) + v224
			v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+8)))
			v234 = v230<<(uint(int32(8))%32) + v229
			v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+7)))
			v239 = v235<<(uint(int32(24))%32) + v94
			v240 = v234
			v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+6)))
			v245 = v241<<(uint(int32(16))%32) + v239
			v246 = v240
			v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+5)))
			v251 = v247<<(uint(int32(8))%32) + v245
			v252 = v246
			v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+4)))
			v255 = v251 + v253
			v256 = v252
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+3)))
			v261 = v257<<(uint(int32(24))%32) + v93
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		default:
			v283 = v93
			v284 = v94
			v285 = v98
		}
	} else {
		v116 = v2
		v117 = v3
		v119 = v47
		v120 = v51
		v121 = v49
		for {
			v123 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
			v124 = v123 + v120
			v125 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
			v127 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
			v128 = v127 + v121
			v130 = int32(4)
			v132 = v125 + v119 - v128 ^ base.I32_rotl(v128, v130)
			v136 = v124 - v132 ^ base.I32_rotl(v132, int32(6))
			v137 = v128 + v124
			v138 = v132 + v137
			v139 = v136 + v138
			v143 = v137 - v136 ^ base.I32_rotl(v136, int32(8))
			v147 = v138 - v143 ^ base.I32_rotl(v143, int32(16))
			v151 = v139 - v147 ^ base.I32_rotl(v147, int32(19))
			v152 = v143 + v139
			v153 = v147 + v152
			v154 = v151 + v153
			v158 = v152 - v151 ^ base.I32_rotl(v151, v130)
			v159 = int32(12)
			v160 = v116 + v159
			v162 = v117 - v159
			if base.Ui32(int32(11)) < base.Ui32(v162) {
				v116 = v160
				v117 = v162
				v119 = v153
				v120 = v154
				v121 = v158
				continue
			} else {
				break
			}
			break
		}
		switch v162 - int32(1) {
		case 0:
			v217 = v153
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
			v283 = v217 + v218
			v284 = v154
			v285 = v158
		case 1:
			v212 = v153
			v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+1)))
			v217 = v213<<(uint(int32(8))%32) + v212
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
			v283 = v217 + v218
			v284 = v154
			v285 = v158
		case 2:
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+2)))
			v212 = v208<<(uint(int32(16))%32) + v153
			v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+1)))
			v217 = v213<<(uint(int32(8))%32) + v212
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
			v283 = v217 + v218
			v284 = v154
			v285 = v158
		case 3:
			v205 = v154
			v206 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
			v283 = v206 + v153
			v284 = v205
			v285 = v158
		case 4:
			v202 = v154
			v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+4)))
			v205 = v202 + v203
			v206 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
			v283 = v206 + v153
			v284 = v205
			v285 = v158
		case 5:
			v197 = v154
			v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+5)))
			v202 = v198<<(uint(int32(8))%32) + v197
			v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+4)))
			v205 = v202 + v203
			v206 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
			v283 = v206 + v153
			v284 = v205
			v285 = v158
		case 6:
			v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+6)))
			v197 = v193<<(uint(int32(16))%32) + v154
			v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+5)))
			v202 = v198<<(uint(int32(8))%32) + v197
			v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+4)))
			v205 = v202 + v203
			v206 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
			v283 = v206 + v153
			v284 = v205
			v285 = v158
		case 7:
			v188 = v158
			v189 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
			v191 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
			v283 = v189 + v153
			v284 = v191 + v154
			v285 = v188
		case 8:
			v183 = v158
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
			v191 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
			v283 = v189 + v153
			v284 = v191 + v154
			v285 = v188
		case 9:
			v178 = v158
			v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+9)))
			v183 = v179<<(uint(int32(16))%32) + v178
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
			v191 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
			v283 = v189 + v153
			v284 = v191 + v154
			v285 = v188
		case 10:
			v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+10)))
			v178 = v174<<(uint(int32(24))%32) + v158
			v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+9)))
			v183 = v179<<(uint(int32(16))%32) + v178
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
			v191 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
			v283 = v189 + v153
			v284 = v191 + v154
			v285 = v188
		default:
			v283 = v153
			v284 = v154
			v285 = v158
		}
	}
	v288 = int32(14)
	v290 = v284 ^ v285 - base.I32_rotl(v284, v288)
	v294 = v290 ^ v283 - base.I32_rotl(v290, int32(11))
	v298 = v294 ^ v284 - base.I32_rotl(v294, int32(25))
	v302 = v298 ^ v290 - base.I32_rotl(v298, int32(16))
	v306 = v302 ^ v294 - base.I32_rotl(v302, int32(4))
	v310 = v306 ^ v298 - base.I32_rotl(v306, v288)
	return base.I64_extend_i32_u(v310)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v310^v302-base.I32_rotl(v310, int32(24)))
}
func F_uuid_out(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_palloc(m, int32(37))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = v8
		v13 = int32(0)
		for {
			switch v13&int32(13) - int32(4) {
			case 0, 4:
				v21 = int32(45)
				*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v21)
				v25 = v12 + int32(1)
			default:
				v25 = v12
			}
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v6))))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27&int32(15))+uint32(_c_F_uuid_out[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)) = uint8(v30)
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v27)>>(uint(int32(4))%32)))+uint32(_c_F_uuid_out[0]))))
			*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v34)
			v37 = v25 + int32(2)
			v39 = v13 + int32(1)
			if v39 != int32(16) {
				v12 = v37
				v13 = v39
				continue
			} else {
				break
			}
			break
		}
		v42 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v37))) = uint8(v42)
		return base.I64_extend_i32_u(v8)
	}
}
