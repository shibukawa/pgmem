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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
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
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
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
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
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
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			if v28 == int32(18) {
				v31 = int32(16)
			} else {
				v31 = int32(0)
			}
			if base.Ui32((v28-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v38 = int32(4)
			} else {
				v38 = v31
			}
			v49 = v38
		} else {
			v39 = int32(1)
			if v19 != 0 {
				v49 = int32(base.Ui32(v17)>>(uint(v39)%32)) - v39
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				v49 = int32(base.Ui32(v43)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v51 = F_str_tolower(m, v20, v49, int32(100))
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			v53 = F_strlen(m, v51)
			mBase = m.M
			v59 = v53 - int32(1636608432)
			if v22 == int64(0) {
				v96 = v59
				v98 = v59
				v100 = v59
			} else {
				v63 = v59 + base.I32_wrap_i64(v22)
				v64 = v63 + v59
				v68 = int32(4)
				v70 = base.I32_wrap_i64(int64(base.Ui64(v22)>>(uint(int64(32))%64))) ^ base.I32_rotl(v59, v68)
				v74 = v63 - v70 ^ base.I32_rotl(v70, int32(6))
				v78 = v64 - v74 ^ base.I32_rotl(v74, int32(8))
				v79 = v64 + v70
				v80 = v74 + v79
				v81 = v78 + v80
				v85 = v79 - v78 ^ base.I32_rotl(v78, int32(16))
				v89 = v80 - v85 ^ base.I32_rotl(v85, int32(19))
				v94 = v81 + v85
				v96 = v94
				v98 = v81 - v89 ^ base.I32_rotl(v89, v68)
				v100 = v89 + v94
			}
			if v51&int32(3) != 0 {
				if base.Ui32(int32(11)) < base.Ui32(v53) {
					v105 = v51
					v106 = v53
					v108 = v96
					v109 = v100
					v110 = v98
					for {
						v112 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
						v113 = v112 + v109
						v114 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
						v116 = *(*int32)(unsafe.Add(mBase, uint32(v105)+8))
						v117 = v116 + v110
						v119 = int32(4)
						v121 = v114 + v108 - v117 ^ base.I32_rotl(v117, v119)
						v125 = v113 - v121 ^ base.I32_rotl(v121, int32(6))
						v126 = v117 + v113
						v127 = v121 + v126
						v128 = v125 + v127
						v132 = v126 - v125 ^ base.I32_rotl(v125, int32(8))
						v136 = v127 - v132 ^ base.I32_rotl(v132, int32(16))
						v140 = v128 - v136 ^ base.I32_rotl(v136, int32(19))
						v141 = v132 + v128
						v142 = v136 + v141
						v143 = v140 + v142
						v147 = v141 - v140 ^ base.I32_rotl(v140, v119)
						v148 = int32(12)
						v149 = v105 + v148
						v151 = v106 - v148
						if base.Ui32(int32(11)) < base.Ui32(v151) {
							v105 = v149
							v106 = v151
							v108 = v142
							v109 = v143
							v110 = v147
							continue
						} else {
							break
						}
						break
					}
					v154 = v149
					v155 = v151
					v157 = v142
					v158 = v143
					v159 = v147
				} else {
					v154 = v51
					v155 = v53
					v157 = v96
					v158 = v100
					v159 = v98
				}
				switch v155 - int32(1) {
				case 0:
					v324 = v157
					v325 = v158
					v326 = v159
					v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
					v332 = v324 + v327
					v333 = v325
					v334 = v326
				case 1:
					v317 = v157
					v318 = v158
					v319 = v159
					v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)))
					v324 = v320<<(uint(int32(8))%32) + v317
					v325 = v318
					v326 = v319
					v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
					v332 = v324 + v327
					v333 = v325
					v334 = v326
				case 2:
					v310 = v157
					v311 = v158
					v312 = v159
					v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+2)))
					v317 = v313<<(uint(int32(16))%32) + v310
					v318 = v311
					v319 = v312
					v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)))
					v324 = v320<<(uint(int32(8))%32) + v317
					v325 = v318
					v326 = v319
					v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
					v332 = v324 + v327
					v333 = v325
					v334 = v326
				case 3:
					v304 = v158
					v305 = v159
					v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+3)))
					v310 = v306<<(uint(int32(24))%32) + v157
					v311 = v304
					v312 = v305
					v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+2)))
					v317 = v313<<(uint(int32(16))%32) + v310
					v318 = v311
					v319 = v312
					v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)))
					v324 = v320<<(uint(int32(8))%32) + v317
					v325 = v318
					v326 = v319
					v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
					v332 = v324 + v327
					v333 = v325
					v334 = v326
				case 4:
					v300 = v158
					v301 = v159
					v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+4)))
					v304 = v300 + v302
					v305 = v301
					v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+3)))
					v310 = v306<<(uint(int32(24))%32) + v157
					v311 = v304
					v312 = v305
					v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+2)))
					v317 = v313<<(uint(int32(16))%32) + v310
					v318 = v311
					v319 = v312
					v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)))
					v324 = v320<<(uint(int32(8))%32) + v317
					v325 = v318
					v326 = v319
					v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
					v332 = v324 + v327
					v333 = v325
					v334 = v326
				case 5:
					v294 = v158
					v295 = v159
					v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+5)))
					v300 = v296<<(uint(int32(8))%32) + v294
					v301 = v295
					v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+4)))
					v304 = v300 + v302
					v305 = v301
					v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+3)))
					v310 = v306<<(uint(int32(24))%32) + v157
					v311 = v304
					v312 = v305
					v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+2)))
					v317 = v313<<(uint(int32(16))%32) + v310
					v318 = v311
					v319 = v312
					v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)))
					v324 = v320<<(uint(int32(8))%32) + v317
					v325 = v318
					v326 = v319
					v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
					v332 = v324 + v327
					v333 = v325
					v334 = v326
				case 6:
					v288 = v158
					v289 = v159
					v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+6)))
					v294 = v290<<(uint(int32(16))%32) + v288
					v295 = v289
					v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+5)))
					v300 = v296<<(uint(int32(8))%32) + v294
					v301 = v295
					v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+4)))
					v304 = v300 + v302
					v305 = v301
					v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+3)))
					v310 = v306<<(uint(int32(24))%32) + v157
					v311 = v304
					v312 = v305
					v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+2)))
					v317 = v313<<(uint(int32(16))%32) + v310
					v318 = v311
					v319 = v312
					v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)))
					v324 = v320<<(uint(int32(8))%32) + v317
					v325 = v318
					v326 = v319
					v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
					v332 = v324 + v327
					v333 = v325
					v334 = v326
				case 7:
					v283 = v159
					v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+7)))
					v288 = v284<<(uint(int32(24))%32) + v158
					v289 = v283
					v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+6)))
					v294 = v290<<(uint(int32(16))%32) + v288
					v295 = v289
					v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+5)))
					v300 = v296<<(uint(int32(8))%32) + v294
					v301 = v295
					v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+4)))
					v304 = v300 + v302
					v305 = v301
					v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+3)))
					v310 = v306<<(uint(int32(24))%32) + v157
					v311 = v304
					v312 = v305
					v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+2)))
					v317 = v313<<(uint(int32(16))%32) + v310
					v318 = v311
					v319 = v312
					v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)))
					v324 = v320<<(uint(int32(8))%32) + v317
					v325 = v318
					v326 = v319
					v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
					v332 = v324 + v327
					v333 = v325
					v334 = v326
				case 8:
					v278 = v159
					v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+8)))
					v283 = v279<<(uint(int32(8))%32) + v278
					v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+7)))
					v288 = v284<<(uint(int32(24))%32) + v158
					v289 = v283
					v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+6)))
					v294 = v290<<(uint(int32(16))%32) + v288
					v295 = v289
					v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+5)))
					v300 = v296<<(uint(int32(8))%32) + v294
					v301 = v295
					v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+4)))
					v304 = v300 + v302
					v305 = v301
					v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+3)))
					v310 = v306<<(uint(int32(24))%32) + v157
					v311 = v304
					v312 = v305
					v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+2)))
					v317 = v313<<(uint(int32(16))%32) + v310
					v318 = v311
					v319 = v312
					v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)))
					v324 = v320<<(uint(int32(8))%32) + v317
					v325 = v318
					v326 = v319
					v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
					v332 = v324 + v327
					v333 = v325
					v334 = v326
				case 9:
					v273 = v159
					v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+9)))
					v278 = v274<<(uint(int32(16))%32) + v273
					v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+8)))
					v283 = v279<<(uint(int32(8))%32) + v278
					v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+7)))
					v288 = v284<<(uint(int32(24))%32) + v158
					v289 = v283
					v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+6)))
					v294 = v290<<(uint(int32(16))%32) + v288
					v295 = v289
					v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+5)))
					v300 = v296<<(uint(int32(8))%32) + v294
					v301 = v295
					v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+4)))
					v304 = v300 + v302
					v305 = v301
					v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+3)))
					v310 = v306<<(uint(int32(24))%32) + v157
					v311 = v304
					v312 = v305
					v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+2)))
					v317 = v313<<(uint(int32(16))%32) + v310
					v318 = v311
					v319 = v312
					v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)))
					v324 = v320<<(uint(int32(8))%32) + v317
					v325 = v318
					v326 = v319
					v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
					v332 = v324 + v327
					v333 = v325
					v334 = v326
				case 10:
					v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+10)))
					v273 = v269<<(uint(int32(24))%32) + v159
					v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+9)))
					v278 = v274<<(uint(int32(16))%32) + v273
					v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+8)))
					v283 = v279<<(uint(int32(8))%32) + v278
					v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+7)))
					v288 = v284<<(uint(int32(24))%32) + v158
					v289 = v283
					v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+6)))
					v294 = v290<<(uint(int32(16))%32) + v288
					v295 = v289
					v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+5)))
					v300 = v296<<(uint(int32(8))%32) + v294
					v301 = v295
					v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+4)))
					v304 = v300 + v302
					v305 = v301
					v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+3)))
					v310 = v306<<(uint(int32(24))%32) + v157
					v311 = v304
					v312 = v305
					v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+2)))
					v317 = v313<<(uint(int32(16))%32) + v310
					v318 = v311
					v319 = v312
					v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)))
					v324 = v320<<(uint(int32(8))%32) + v317
					v325 = v318
					v326 = v319
					v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
					v332 = v324 + v327
					v333 = v325
					v334 = v326
				default:
					v332 = v157
					v333 = v158
					v334 = v159
				}
			} else {
				if base.Ui32(int32(12)) <= base.Ui32(v53) {
					v165 = v51
					v166 = v53
					v168 = v96
					v169 = v100
					v170 = v98
					for {
						v172 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
						v173 = v172 + v169
						v174 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
						v176 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
						v177 = v176 + v170
						v179 = int32(4)
						v181 = v174 + v168 - v177 ^ base.I32_rotl(v177, v179)
						v185 = v173 - v181 ^ base.I32_rotl(v181, int32(6))
						v186 = v177 + v173
						v187 = v181 + v186
						v188 = v185 + v187
						v192 = v186 - v185 ^ base.I32_rotl(v185, int32(8))
						v196 = v187 - v192 ^ base.I32_rotl(v192, int32(16))
						v200 = v188 - v196 ^ base.I32_rotl(v196, int32(19))
						v201 = v192 + v188
						v202 = v196 + v201
						v203 = v200 + v202
						v207 = v201 - v200 ^ base.I32_rotl(v200, v179)
						v208 = int32(12)
						v209 = v165 + v208
						v211 = v166 - v208
						if base.Ui32(int32(11)) < base.Ui32(v211) {
							v165 = v209
							v166 = v211
							v168 = v202
							v169 = v203
							v170 = v207
							continue
						} else {
							break
						}
						break
					}
					v214 = v209
					v215 = v211
					v217 = v202
					v218 = v203
					v219 = v207
				} else {
					v214 = v51
					v215 = v53
					v217 = v96
					v218 = v100
					v219 = v98
				}
				switch v215 - int32(1) {
				case 0:
					v266 = v217
					v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
					v332 = v266 + v267
					v333 = v218
					v334 = v219
				case 1:
					v261 = v217
					v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+1)))
					v266 = v262<<(uint(int32(8))%32) + v261
					v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
					v332 = v266 + v267
					v333 = v218
					v334 = v219
				case 2:
					v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+2)))
					v261 = v257<<(uint(int32(16))%32) + v217
					v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+1)))
					v266 = v262<<(uint(int32(8))%32) + v261
					v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
					v332 = v266 + v267
					v333 = v218
					v334 = v219
				case 3:
					v254 = v218
					v255 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
					v332 = v255 + v217
					v333 = v254
					v334 = v219
				case 4:
					v251 = v218
					v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+4)))
					v254 = v251 + v252
					v255 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
					v332 = v255 + v217
					v333 = v254
					v334 = v219
				case 5:
					v246 = v218
					v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+5)))
					v251 = v247<<(uint(int32(8))%32) + v246
					v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+4)))
					v254 = v251 + v252
					v255 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
					v332 = v255 + v217
					v333 = v254
					v334 = v219
				case 6:
					v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+6)))
					v246 = v242<<(uint(int32(16))%32) + v218
					v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+5)))
					v251 = v247<<(uint(int32(8))%32) + v246
					v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+4)))
					v254 = v251 + v252
					v255 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
					v332 = v255 + v217
					v333 = v254
					v334 = v219
				case 7:
					v237 = v219
					v238 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
					v240 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
					v332 = v238 + v217
					v333 = v240 + v218
					v334 = v237
				case 8:
					v232 = v219
					v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+8)))
					v237 = v233<<(uint(int32(8))%32) + v232
					v238 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
					v240 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
					v332 = v238 + v217
					v333 = v240 + v218
					v334 = v237
				case 9:
					v227 = v219
					v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+9)))
					v232 = v228<<(uint(int32(16))%32) + v227
					v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+8)))
					v237 = v233<<(uint(int32(8))%32) + v232
					v238 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
					v240 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
					v332 = v238 + v217
					v333 = v240 + v218
					v334 = v237
				case 10:
					v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+10)))
					v227 = v223<<(uint(int32(24))%32) + v219
					v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+9)))
					v232 = v228<<(uint(int32(16))%32) + v227
					v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+8)))
					v237 = v233<<(uint(int32(8))%32) + v232
					v238 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
					v240 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
					v332 = v238 + v217
					v333 = v240 + v218
					v334 = v237
				default:
					v332 = v217
					v333 = v218
					v334 = v219
				}
			}
			v337 = int32(14)
			v339 = v333 ^ v334 - base.I32_rotl(v333, v337)
			v343 = v339 ^ v332 - base.I32_rotl(v339, int32(11))
			v347 = v343 ^ v333 - base.I32_rotl(v343, int32(25))
			v351 = v347 ^ v339 - base.I32_rotl(v347, int32(16))
			v355 = v351 ^ v343 - base.I32_rotl(v351, int32(4))
			v359 = v355 ^ v347 - base.I32_rotl(v355, v337)
			v369 = F_Int64GetDatum(m, base.I64_extend_i32_u(v359)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v359^v351-base.I32_rotl(v359, int32(24))))
			mBase = m.M
			v370 = m.ExcPending
			if v370 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v51)
				mBase = m.M
				v372 = m.ExcPending
				if v372 != 0 {
					return int32(0)
				} else {
					v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v373 != v9 {
						F_pfree(m, v9)
						mBase = m.M
						v376 = m.ExcPending
						if v376 != 0 {
							return int32(0)
						} else {
							return v369
						}
					} else {
						return v369
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
