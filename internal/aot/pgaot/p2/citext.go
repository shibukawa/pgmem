package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_citext_hash_extended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
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
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
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
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
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
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(1)
		v14 = v9 + v13
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
		v19 = v17 & v13
		if v19 != 0 {
			v20 = v14
		} else {
			v20 = v9 + int32(4)
		}
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v22 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
		if v17 == int32(1) {
			v25 = int32(4)
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			if v27&int32(254) == int32(2) {
				v36 = v25
			} else {
				v36 = base.B2i32(v27 == int32(18)) << (uint(v25) % 32)
			}
			if v27 == int32(1) {
				v39 = v25
			} else {
				v39 = v36
			}
			v50 = v39
		} else {
			v40 = int32(1)
			if v19 != 0 {
				v50 = int32(base.Ui32(v17)>>(uint(v40)%32)) - v40
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v52 = F_str_tolower(m, v20, v50, int32(100))
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return int32(0)
		} else {
			v54 = F_strlen(m, v52)
			mBase = m.M
			v60 = v54 - int32(1636608432)
			if v22 == int64(0) {
				v97 = v60
				v99 = v60
				v101 = v60
			} else {
				v64 = v60 + base.I32_wrap_i64(v22)
				v65 = v64 + v60
				v69 = int32(4)
				v71 = base.I32_wrap_i64(int64(base.Ui64(v22)>>(uint(int64(32))%64))) ^ base.I32_rotl(v60, v69)
				v75 = v64 - v71 ^ base.I32_rotl(v71, int32(6))
				v79 = v65 - v75 ^ base.I32_rotl(v75, int32(8))
				v80 = v71 + v65
				v81 = v75 + v80
				v82 = v79 + v81
				v86 = v80 - v79 ^ base.I32_rotl(v79, int32(16))
				v90 = v81 - v86 ^ base.I32_rotl(v86, int32(19))
				v95 = v86 + v82
				v97 = v95
				v99 = v82 - v90 ^ base.I32_rotl(v90, v69)
				v101 = v90 + v95
			}
			if v52&int32(3) != 0 {
				if base.Ui32(int32(11)) < base.Ui32(v54) {
					v106 = v52
					v107 = v54
					v109 = v97
					v110 = v101
					v111 = v99
					for {
						v113 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
						v114 = v113 + v110
						v115 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
						v117 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
						v118 = v117 + v111
						v120 = int32(4)
						v122 = v115 + v109 - v118 ^ base.I32_rotl(v118, v120)
						v126 = v114 - v122 ^ base.I32_rotl(v122, int32(6))
						v127 = v118 + v114
						v128 = v122 + v127
						v129 = v126 + v128
						v133 = v127 - v126 ^ base.I32_rotl(v126, int32(8))
						v137 = v128 - v133 ^ base.I32_rotl(v133, int32(16))
						v141 = v129 - v137 ^ base.I32_rotl(v137, int32(19))
						v142 = v133 + v129
						v143 = v137 + v142
						v144 = v141 + v143
						v148 = v142 - v141 ^ base.I32_rotl(v141, v120)
						v149 = int32(12)
						v150 = v106 + v149
						v152 = v107 - v149
						if base.Ui32(int32(11)) < base.Ui32(v152) {
							v106 = v150
							v107 = v152
							v109 = v143
							v110 = v144
							v111 = v148
							continue
						} else {
							break
						}
						break
					}
					v155 = v150
					v156 = v152
					v158 = v143
					v159 = v144
					v160 = v148
				} else {
					v155 = v52
					v156 = v54
					v158 = v97
					v159 = v101
					v160 = v99
				}
				switch v156 - int32(1) {
				case 0:
					v325 = v158
					v326 = v159
					v327 = v160
					v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
					v333 = v325 + v328
					v334 = v326
					v335 = v327
				case 1:
					v318 = v158
					v319 = v159
					v320 = v160
					v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+1)))
					v325 = v321<<(uint(int32(8))%32) + v318
					v326 = v319
					v327 = v320
					v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
					v333 = v325 + v328
					v334 = v326
					v335 = v327
				case 2:
					v311 = v158
					v312 = v159
					v313 = v160
					v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+2)))
					v318 = v314<<(uint(int32(16))%32) + v311
					v319 = v312
					v320 = v313
					v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+1)))
					v325 = v321<<(uint(int32(8))%32) + v318
					v326 = v319
					v327 = v320
					v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
					v333 = v325 + v328
					v334 = v326
					v335 = v327
				case 3:
					v305 = v159
					v306 = v160
					v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+3)))
					v311 = v307<<(uint(int32(24))%32) + v158
					v312 = v305
					v313 = v306
					v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+2)))
					v318 = v314<<(uint(int32(16))%32) + v311
					v319 = v312
					v320 = v313
					v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+1)))
					v325 = v321<<(uint(int32(8))%32) + v318
					v326 = v319
					v327 = v320
					v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
					v333 = v325 + v328
					v334 = v326
					v335 = v327
				case 4:
					v301 = v159
					v302 = v160
					v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+4)))
					v305 = v301 + v303
					v306 = v302
					v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+3)))
					v311 = v307<<(uint(int32(24))%32) + v158
					v312 = v305
					v313 = v306
					v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+2)))
					v318 = v314<<(uint(int32(16))%32) + v311
					v319 = v312
					v320 = v313
					v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+1)))
					v325 = v321<<(uint(int32(8))%32) + v318
					v326 = v319
					v327 = v320
					v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
					v333 = v325 + v328
					v334 = v326
					v335 = v327
				case 5:
					v295 = v159
					v296 = v160
					v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+5)))
					v301 = v297<<(uint(int32(8))%32) + v295
					v302 = v296
					v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+4)))
					v305 = v301 + v303
					v306 = v302
					v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+3)))
					v311 = v307<<(uint(int32(24))%32) + v158
					v312 = v305
					v313 = v306
					v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+2)))
					v318 = v314<<(uint(int32(16))%32) + v311
					v319 = v312
					v320 = v313
					v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+1)))
					v325 = v321<<(uint(int32(8))%32) + v318
					v326 = v319
					v327 = v320
					v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
					v333 = v325 + v328
					v334 = v326
					v335 = v327
				case 6:
					v289 = v159
					v290 = v160
					v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+6)))
					v295 = v291<<(uint(int32(16))%32) + v289
					v296 = v290
					v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+5)))
					v301 = v297<<(uint(int32(8))%32) + v295
					v302 = v296
					v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+4)))
					v305 = v301 + v303
					v306 = v302
					v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+3)))
					v311 = v307<<(uint(int32(24))%32) + v158
					v312 = v305
					v313 = v306
					v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+2)))
					v318 = v314<<(uint(int32(16))%32) + v311
					v319 = v312
					v320 = v313
					v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+1)))
					v325 = v321<<(uint(int32(8))%32) + v318
					v326 = v319
					v327 = v320
					v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
					v333 = v325 + v328
					v334 = v326
					v335 = v327
				case 7:
					v284 = v160
					v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+7)))
					v289 = v285<<(uint(int32(24))%32) + v159
					v290 = v284
					v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+6)))
					v295 = v291<<(uint(int32(16))%32) + v289
					v296 = v290
					v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+5)))
					v301 = v297<<(uint(int32(8))%32) + v295
					v302 = v296
					v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+4)))
					v305 = v301 + v303
					v306 = v302
					v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+3)))
					v311 = v307<<(uint(int32(24))%32) + v158
					v312 = v305
					v313 = v306
					v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+2)))
					v318 = v314<<(uint(int32(16))%32) + v311
					v319 = v312
					v320 = v313
					v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+1)))
					v325 = v321<<(uint(int32(8))%32) + v318
					v326 = v319
					v327 = v320
					v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
					v333 = v325 + v328
					v334 = v326
					v335 = v327
				case 8:
					v279 = v160
					v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+8)))
					v284 = v280<<(uint(int32(8))%32) + v279
					v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+7)))
					v289 = v285<<(uint(int32(24))%32) + v159
					v290 = v284
					v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+6)))
					v295 = v291<<(uint(int32(16))%32) + v289
					v296 = v290
					v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+5)))
					v301 = v297<<(uint(int32(8))%32) + v295
					v302 = v296
					v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+4)))
					v305 = v301 + v303
					v306 = v302
					v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+3)))
					v311 = v307<<(uint(int32(24))%32) + v158
					v312 = v305
					v313 = v306
					v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+2)))
					v318 = v314<<(uint(int32(16))%32) + v311
					v319 = v312
					v320 = v313
					v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+1)))
					v325 = v321<<(uint(int32(8))%32) + v318
					v326 = v319
					v327 = v320
					v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
					v333 = v325 + v328
					v334 = v326
					v335 = v327
				case 9:
					v274 = v160
					v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+9)))
					v279 = v275<<(uint(int32(16))%32) + v274
					v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+8)))
					v284 = v280<<(uint(int32(8))%32) + v279
					v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+7)))
					v289 = v285<<(uint(int32(24))%32) + v159
					v290 = v284
					v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+6)))
					v295 = v291<<(uint(int32(16))%32) + v289
					v296 = v290
					v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+5)))
					v301 = v297<<(uint(int32(8))%32) + v295
					v302 = v296
					v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+4)))
					v305 = v301 + v303
					v306 = v302
					v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+3)))
					v311 = v307<<(uint(int32(24))%32) + v158
					v312 = v305
					v313 = v306
					v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+2)))
					v318 = v314<<(uint(int32(16))%32) + v311
					v319 = v312
					v320 = v313
					v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+1)))
					v325 = v321<<(uint(int32(8))%32) + v318
					v326 = v319
					v327 = v320
					v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
					v333 = v325 + v328
					v334 = v326
					v335 = v327
				case 10:
					v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+10)))
					v274 = v270<<(uint(int32(24))%32) + v160
					v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+9)))
					v279 = v275<<(uint(int32(16))%32) + v274
					v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+8)))
					v284 = v280<<(uint(int32(8))%32) + v279
					v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+7)))
					v289 = v285<<(uint(int32(24))%32) + v159
					v290 = v284
					v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+6)))
					v295 = v291<<(uint(int32(16))%32) + v289
					v296 = v290
					v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+5)))
					v301 = v297<<(uint(int32(8))%32) + v295
					v302 = v296
					v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+4)))
					v305 = v301 + v303
					v306 = v302
					v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+3)))
					v311 = v307<<(uint(int32(24))%32) + v158
					v312 = v305
					v313 = v306
					v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+2)))
					v318 = v314<<(uint(int32(16))%32) + v311
					v319 = v312
					v320 = v313
					v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+1)))
					v325 = v321<<(uint(int32(8))%32) + v318
					v326 = v319
					v327 = v320
					v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
					v333 = v325 + v328
					v334 = v326
					v335 = v327
				default:
					v333 = v158
					v334 = v159
					v335 = v160
				}
			} else {
				if base.Ui32(int32(12)) <= base.Ui32(v54) {
					v166 = v52
					v167 = v54
					v169 = v97
					v170 = v101
					v171 = v99
					for {
						v173 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
						v174 = v173 + v170
						v175 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
						v177 = *(*int32)(unsafe.Add(mBase, uint32(v166)+8))
						v178 = v177 + v171
						v180 = int32(4)
						v182 = v175 + v169 - v178 ^ base.I32_rotl(v178, v180)
						v186 = v174 - v182 ^ base.I32_rotl(v182, int32(6))
						v187 = v178 + v174
						v188 = v182 + v187
						v189 = v186 + v188
						v193 = v187 - v186 ^ base.I32_rotl(v186, int32(8))
						v197 = v188 - v193 ^ base.I32_rotl(v193, int32(16))
						v201 = v189 - v197 ^ base.I32_rotl(v197, int32(19))
						v202 = v193 + v189
						v203 = v197 + v202
						v204 = v201 + v203
						v208 = v202 - v201 ^ base.I32_rotl(v201, v180)
						v209 = int32(12)
						v210 = v166 + v209
						v212 = v167 - v209
						if base.Ui32(int32(11)) < base.Ui32(v212) {
							v166 = v210
							v167 = v212
							v169 = v203
							v170 = v204
							v171 = v208
							continue
						} else {
							break
						}
						break
					}
					v215 = v210
					v216 = v212
					v218 = v203
					v219 = v204
					v220 = v208
				} else {
					v215 = v52
					v216 = v54
					v218 = v97
					v219 = v101
					v220 = v99
				}
				switch v216 - int32(1) {
				case 0:
					v267 = v218
					v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
					v333 = v267 + v268
					v334 = v219
					v335 = v220
				case 1:
					v262 = v218
					v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+1)))
					v267 = v263<<(uint(int32(8))%32) + v262
					v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
					v333 = v267 + v268
					v334 = v219
					v335 = v220
				case 2:
					v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+2)))
					v262 = v258<<(uint(int32(16))%32) + v218
					v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+1)))
					v267 = v263<<(uint(int32(8))%32) + v262
					v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
					v333 = v267 + v268
					v334 = v219
					v335 = v220
				case 3:
					v255 = v219
					v256 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
					v333 = v256 + v218
					v334 = v255
					v335 = v220
				case 4:
					v252 = v219
					v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+4)))
					v255 = v252 + v253
					v256 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
					v333 = v256 + v218
					v334 = v255
					v335 = v220
				case 5:
					v247 = v219
					v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+5)))
					v252 = v248<<(uint(int32(8))%32) + v247
					v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+4)))
					v255 = v252 + v253
					v256 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
					v333 = v256 + v218
					v334 = v255
					v335 = v220
				case 6:
					v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+6)))
					v247 = v243<<(uint(int32(16))%32) + v219
					v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+5)))
					v252 = v248<<(uint(int32(8))%32) + v247
					v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+4)))
					v255 = v252 + v253
					v256 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
					v333 = v256 + v218
					v334 = v255
					v335 = v220
				case 7:
					v238 = v220
					v239 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
					v241 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
					v333 = v239 + v218
					v334 = v241 + v219
					v335 = v238
				case 8:
					v233 = v220
					v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+8)))
					v238 = v234<<(uint(int32(8))%32) + v233
					v239 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
					v241 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
					v333 = v239 + v218
					v334 = v241 + v219
					v335 = v238
				case 9:
					v228 = v220
					v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+9)))
					v233 = v229<<(uint(int32(16))%32) + v228
					v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+8)))
					v238 = v234<<(uint(int32(8))%32) + v233
					v239 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
					v241 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
					v333 = v239 + v218
					v334 = v241 + v219
					v335 = v238
				case 10:
					v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+10)))
					v228 = v224<<(uint(int32(24))%32) + v220
					v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+9)))
					v233 = v229<<(uint(int32(16))%32) + v228
					v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+8)))
					v238 = v234<<(uint(int32(8))%32) + v233
					v239 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
					v241 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
					v333 = v239 + v218
					v334 = v241 + v219
					v335 = v238
				default:
					v333 = v218
					v334 = v219
					v335 = v220
				}
			}
			v338 = int32(14)
			v340 = v334 ^ v335 - base.I32_rotl(v334, v338)
			v344 = v340 ^ v333 - base.I32_rotl(v340, int32(11))
			v348 = v344 ^ v334 - base.I32_rotl(v344, int32(25))
			v352 = v348 ^ v340 - base.I32_rotl(v348, int32(16))
			v356 = v352 ^ v344 - base.I32_rotl(v352, int32(4))
			v360 = v356 ^ v348 - base.I32_rotl(v356, v338)
			v370 = F_Int64GetDatum(m, base.I64_extend_i32_u(v360)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v352^v360-base.I32_rotl(v360, int32(24))))
			mBase = m.M
			v371 = m.ExcPending
			if v371 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v52)
				mBase = m.M
				v373 = m.ExcPending
				if v373 != 0 {
					return int32(0)
				} else {
					v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v374 != v9 {
						F_pfree(m, v9)
						mBase = m.M
						v377 = m.ExcPending
						if v377 != 0 {
							return int32(0)
						} else {
							return v370
						}
					} else {
						return v370
					}
				}
			}
		}
	}
}
func F_citext_lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v14 = F_citextcmp(m, v6, v11, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v16 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v20 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v23 = m.ExcPending
							if v23 != 0 {
								return int32(0)
							} else {
								return int32(base.Ui32(v14) >> (uint(int32(31)) % 32))
							}
						} else {
							return int32(base.Ui32(v14) >> (uint(int32(31)) % 32))
						}
					}
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v20 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							return int32(base.Ui32(v14) >> (uint(int32(31)) % 32))
						}
					} else {
						return int32(base.Ui32(v14) >> (uint(int32(31)) % 32))
					}
				}
			}
		}
	}
}
func F_citext_pattern_gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_internal_citext_pattern_cmp(m, v6, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v15 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int32(0)
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v19 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return int32(0)
							} else {
								return base.B2i32(int32(0) < v13)
							}
						} else {
							return base.B2i32(int32(0) < v13)
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							return base.B2i32(int32(0) < v13)
						}
					} else {
						return base.B2i32(int32(0) < v13)
					}
				}
			}
		}
	}
}
func F_citext_pattern_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_internal_citext_pattern_cmp(m, v6, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v15 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int32(0)
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v19 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return int32(0)
							} else {
								return base.B2i32(v13 <= int32(0))
							}
						} else {
							return base.B2i32(v13 <= int32(0))
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v13 <= int32(0))
						}
					} else {
						return base.B2i32(v13 <= int32(0))
					}
				}
			}
		}
	}
}
