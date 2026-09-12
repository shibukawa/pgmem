package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_uuid_cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = int32(16)
	goto L4
L1:
	;
	return v66
L2:
	;
	v66 = int32(0)
	goto L1
L3:
	;
	v40 = v35
	v41 = v36
	v42 = v37
	goto L13
L4:
	;
	if (v2|v3)&int32(3) != 0 {
		v35 = v2
		v36 = v3
		v37 = v4
		goto L3
	} else {
		goto L7
	}
L6:
	;
	if v25 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L7:
	;
	v12 = v2
	v13 = v3
	v14 = v4
	goto L8
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v17 != v18 {
		v35 = v12
		v36 = v13
		v37 = v14
		goto L3
	} else {
		goto L10
	}
L9:
	;
	goto L6
L10:
	;
	v20 = int32(4)
	v21 = v13 + v20
	v23 = v12 + v20
	v25 = v14 - v20
	if base.Ui32(int32(3)) < base.Ui32(v25) {
		v12 = v23
		v13 = v21
		v14 = v25
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v35 = v23
	v36 = v21
	v37 = v25
	goto L3
L13:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v45 == v46 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v66 = v45 - v46
	goto L1
L15:
	;
	v48 = int32(1)
	v53 = v42 - v48
	if v53 != 0 {
		v40 = v40 + v48
		v41 = v41 + v48
		v42 = v53
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
	goto L2
}
func F_uuid_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = int32(16)
	goto L4
L1:
	;
	return base.B2i32(v66 == int32(0))
L2:
	;
	v66 = int32(0)
	goto L1
L3:
	;
	v40 = v35
	v41 = v36
	v42 = v37
	goto L13
L4:
	;
	if (v2|v3)&int32(3) != 0 {
		v35 = v2
		v36 = v3
		v37 = v4
		goto L3
	} else {
		goto L7
	}
L6:
	;
	if v25 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L7:
	;
	v12 = v2
	v13 = v3
	v14 = v4
	goto L8
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v17 != v18 {
		v35 = v12
		v36 = v13
		v37 = v14
		goto L3
	} else {
		goto L10
	}
L9:
	;
	goto L6
L10:
	;
	v20 = int32(4)
	v21 = v13 + v20
	v23 = v12 + v20
	v25 = v14 - v20
	if base.Ui32(int32(3)) < base.Ui32(v25) {
		v12 = v23
		v13 = v21
		v14 = v25
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v35 = v23
	v36 = v21
	v37 = v25
	goto L3
L13:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v45 == v46 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v66 = v45 - v46
	goto L1
L15:
	;
	v48 = int32(1)
	v53 = v42 - v48
	if v53 != 0 {
		v40 = v40 + v48
		v41 = v41 + v48
		v42 = v53
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
	goto L2
}
func F_uuid_extract_timestamp(m *base.Module, l0 int32) int32 {
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
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v62 int64
	_ = v62
	var v66 int64
	_ = v66
	var v70 int64
	_ = v70
	var v74 int64
	_ = v74
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = int32(*(*int8)(unsafe.Add(mBase, uint32(v4)+8)))
	if int32(-64) <= v5 {
		v8 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v8)
		return int32(0)
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
			v52 = F_Int64GetDatum(m, v49-int64(13165977600000000))
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int32(0)
			} else {
				return v52
			}
		default:
			v85 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v85)
			return int32(0)
		case 6:
			v57 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v4)+5)))
			v58 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v4)+4)))
			v62 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v4)+3)))
			v66 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v4)+2)))
			v70 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v4)+1)))
			v74 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
			v82 = F_Int64GetDatum(m, (v57|v58<<(uint(int64(8))%64)|v62<<(uint(int64(16))%64)|v66<<(uint(int64(24))%64)|v70<<(uint(int64(32))%64)|v74<<(uint(int64(40))%64))*int64(1000)-int64(946684800000000))
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return int32(0)
			} else {
				return v82
			}
		}
	}
}
func F_uuid_hash_extended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
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
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
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
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(16)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v11 = int32(-1636608416)
	if v5 == int64(0) {
		v48 = v11
		v50 = v11
		v52 = v11
	} else {
		v14 = base.I32_wrap_i64(v5)
		v16 = v14 + int32(1021750464)
		v20 = int32(4)
		v22 = base.I32_wrap_i64(int64(base.Ui64(v5)>>(uint(int64(32))%64))) ^ base.I32_rotl(v11, v20)
		v26 = v11 + v14 - v22 ^ base.I32_rotl(v22, int32(6))
		v30 = v16 - v26 ^ base.I32_rotl(v26, int32(8))
		v31 = v22 + v16
		v32 = v26 + v31
		v33 = v30 + v32
		v37 = v31 - v30 ^ base.I32_rotl(v30, int32(16))
		v41 = v32 - v37 ^ base.I32_rotl(v37, int32(19))
		v46 = v37 + v33
		v48 = v46
		v50 = v33 - v41 ^ base.I32_rotl(v41, v20)
		v52 = v41 + v46
	}
	if v2&int32(3) != 0 {
		v57 = v2
		v58 = v3
		v60 = v48
		v61 = v52
		v62 = v50
		for {
			v64 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
			v65 = v64 + v61
			v66 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
			v68 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
			v69 = v68 + v62
			v71 = int32(4)
			v73 = v66 + v60 - v69 ^ base.I32_rotl(v69, v71)
			v77 = v65 - v73 ^ base.I32_rotl(v73, int32(6))
			v78 = v69 + v65
			v79 = v73 + v78
			v80 = v77 + v79
			v84 = v78 - v77 ^ base.I32_rotl(v77, int32(8))
			v88 = v79 - v84 ^ base.I32_rotl(v84, int32(16))
			v92 = v80 - v88 ^ base.I32_rotl(v88, int32(19))
			v93 = v84 + v80
			v94 = v88 + v93
			v95 = v92 + v94
			v99 = v93 - v92 ^ base.I32_rotl(v92, v71)
			v100 = int32(12)
			v101 = v57 + v100
			v103 = v58 - v100
			if base.Ui32(int32(11)) < base.Ui32(v103) {
				v57 = v101
				v58 = v103
				v60 = v94
				v61 = v95
				v62 = v99
				continue
			} else {
				break
			}
			break
		}
		switch v103 - int32(1) {
		case 0:
			v276 = v94
			v277 = v95
			v278 = v99
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 1:
			v269 = v94
			v270 = v95
			v271 = v99
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 2:
			v262 = v94
			v263 = v95
			v264 = v99
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 3:
			v256 = v95
			v257 = v99
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+3)))
			v262 = v258<<(uint(int32(24))%32) + v94
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 4:
			v252 = v95
			v253 = v99
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+3)))
			v262 = v258<<(uint(int32(24))%32) + v94
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 5:
			v246 = v95
			v247 = v99
			v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+5)))
			v252 = v248<<(uint(int32(8))%32) + v246
			v253 = v247
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+3)))
			v262 = v258<<(uint(int32(24))%32) + v94
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 6:
			v240 = v95
			v241 = v99
			v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+6)))
			v246 = v242<<(uint(int32(16))%32) + v240
			v247 = v241
			v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+5)))
			v252 = v248<<(uint(int32(8))%32) + v246
			v253 = v247
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+3)))
			v262 = v258<<(uint(int32(24))%32) + v94
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 7:
			v235 = v99
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+7)))
			v240 = v236<<(uint(int32(24))%32) + v95
			v241 = v235
			v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+6)))
			v246 = v242<<(uint(int32(16))%32) + v240
			v247 = v241
			v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+5)))
			v252 = v248<<(uint(int32(8))%32) + v246
			v253 = v247
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+3)))
			v262 = v258<<(uint(int32(24))%32) + v94
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 8:
			v230 = v99
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+8)))
			v235 = v231<<(uint(int32(8))%32) + v230
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+7)))
			v240 = v236<<(uint(int32(24))%32) + v95
			v241 = v235
			v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+6)))
			v246 = v242<<(uint(int32(16))%32) + v240
			v247 = v241
			v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+5)))
			v252 = v248<<(uint(int32(8))%32) + v246
			v253 = v247
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+3)))
			v262 = v258<<(uint(int32(24))%32) + v94
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 9:
			v225 = v99
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+9)))
			v230 = v226<<(uint(int32(16))%32) + v225
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+8)))
			v235 = v231<<(uint(int32(8))%32) + v230
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+7)))
			v240 = v236<<(uint(int32(24))%32) + v95
			v241 = v235
			v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+6)))
			v246 = v242<<(uint(int32(16))%32) + v240
			v247 = v241
			v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+5)))
			v252 = v248<<(uint(int32(8))%32) + v246
			v253 = v247
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+3)))
			v262 = v258<<(uint(int32(24))%32) + v94
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 10:
			v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+10)))
			v225 = v221<<(uint(int32(24))%32) + v99
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+9)))
			v230 = v226<<(uint(int32(16))%32) + v225
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+8)))
			v235 = v231<<(uint(int32(8))%32) + v230
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+7)))
			v240 = v236<<(uint(int32(24))%32) + v95
			v241 = v235
			v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+6)))
			v246 = v242<<(uint(int32(16))%32) + v240
			v247 = v241
			v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+5)))
			v252 = v248<<(uint(int32(8))%32) + v246
			v253 = v247
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+3)))
			v262 = v258<<(uint(int32(24))%32) + v94
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		default:
			v284 = v94
			v285 = v95
			v286 = v99
		}
	} else {
		v117 = v2
		v118 = v3
		v120 = v48
		v121 = v52
		v122 = v50
		for {
			v124 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
			v125 = v124 + v121
			v126 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
			v128 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
			v129 = v128 + v122
			v131 = int32(4)
			v133 = v126 + v120 - v129 ^ base.I32_rotl(v129, v131)
			v137 = v125 - v133 ^ base.I32_rotl(v133, int32(6))
			v138 = v129 + v125
			v139 = v133 + v138
			v140 = v137 + v139
			v144 = v138 - v137 ^ base.I32_rotl(v137, int32(8))
			v148 = v139 - v144 ^ base.I32_rotl(v144, int32(16))
			v152 = v140 - v148 ^ base.I32_rotl(v148, int32(19))
			v153 = v144 + v140
			v154 = v148 + v153
			v155 = v152 + v154
			v159 = v153 - v152 ^ base.I32_rotl(v152, v131)
			v160 = int32(12)
			v161 = v117 + v160
			v163 = v118 - v160
			if base.Ui32(int32(11)) < base.Ui32(v163) {
				v117 = v161
				v118 = v163
				v120 = v154
				v121 = v155
				v122 = v159
				continue
			} else {
				break
			}
			break
		}
		switch v163 - int32(1) {
		case 0:
			v218 = v154
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
			v284 = v218 + v219
			v285 = v155
			v286 = v159
		case 1:
			v213 = v154
			v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
			v218 = v214<<(uint(int32(8))%32) + v213
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
			v284 = v218 + v219
			v285 = v155
			v286 = v159
		case 2:
			v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+2)))
			v213 = v209<<(uint(int32(16))%32) + v154
			v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
			v218 = v214<<(uint(int32(8))%32) + v213
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
			v284 = v218 + v219
			v285 = v155
			v286 = v159
		case 3:
			v206 = v155
			v207 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
			v284 = v207 + v154
			v285 = v206
			v286 = v159
		case 4:
			v203 = v155
			v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+4)))
			v206 = v203 + v204
			v207 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
			v284 = v207 + v154
			v285 = v206
			v286 = v159
		case 5:
			v198 = v155
			v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+5)))
			v203 = v199<<(uint(int32(8))%32) + v198
			v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+4)))
			v206 = v203 + v204
			v207 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
			v284 = v207 + v154
			v285 = v206
			v286 = v159
		case 6:
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+6)))
			v198 = v194<<(uint(int32(16))%32) + v155
			v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+5)))
			v203 = v199<<(uint(int32(8))%32) + v198
			v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+4)))
			v206 = v203 + v204
			v207 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
			v284 = v207 + v154
			v285 = v206
			v286 = v159
		case 7:
			v189 = v159
			v190 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
			v192 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
			v284 = v190 + v154
			v285 = v192 + v155
			v286 = v189
		case 8:
			v184 = v159
			v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+8)))
			v189 = v185<<(uint(int32(8))%32) + v184
			v190 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
			v192 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
			v284 = v190 + v154
			v285 = v192 + v155
			v286 = v189
		case 9:
			v179 = v159
			v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+9)))
			v184 = v180<<(uint(int32(16))%32) + v179
			v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+8)))
			v189 = v185<<(uint(int32(8))%32) + v184
			v190 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
			v192 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
			v284 = v190 + v154
			v285 = v192 + v155
			v286 = v189
		case 10:
			v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+10)))
			v179 = v175<<(uint(int32(24))%32) + v159
			v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+9)))
			v184 = v180<<(uint(int32(16))%32) + v179
			v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+8)))
			v189 = v185<<(uint(int32(8))%32) + v184
			v190 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
			v192 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
			v284 = v190 + v154
			v285 = v192 + v155
			v286 = v189
		default:
			v284 = v154
			v285 = v155
			v286 = v159
		}
	}
	v289 = int32(14)
	v291 = v285 ^ v286 - base.I32_rotl(v285, v289)
	v295 = v291 ^ v284 - base.I32_rotl(v291, int32(11))
	v299 = v295 ^ v285 - base.I32_rotl(v295, int32(25))
	v303 = v299 ^ v291 - base.I32_rotl(v299, int32(16))
	v307 = v303 ^ v295 - base.I32_rotl(v303, int32(4))
	v311 = v307 ^ v299 - base.I32_rotl(v307, v289)
	v321 = F_Int64GetDatum(m, base.I64_extend_i32_u(v311)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v303^v311-base.I32_rotl(v311, int32(24))))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		return int32(0)
	} else {
		return v321
	}
}
func F_uuid_out(m *base.Module, l0 int32) int32 {
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
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_palloc(m, int32(37))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
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
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27&int32(15))+uint32(_consts[1344]))))
			*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)) = uint8(v32)
			v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v27)>>(uint(int32(4))%32)))+uint32(_consts[1344]))))
			*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v38)
			v41 = v25 + int32(2)
			v43 = v13 + int32(1)
			if v43 != int32(16) {
				v12 = v41
				v13 = v43
				continue
			} else {
				break
			}
			break
		}
		v46 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v46)
		return v8
	}
}
