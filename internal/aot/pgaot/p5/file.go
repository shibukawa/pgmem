package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FileSetDelete(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
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
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
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
	var v213 int32
	_ = v213
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
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int64
	_ = v286
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	v5 = m.G0
	v7 = v5 - int32(2080)
	m.G0 = v7
	v11 = F_strlen(m, l1)
	mBase = m.M
	v17 = v11 - int32(1636608432)
	if l1&int32(3) != 0 {
		if base.Ui32(int32(11)) < base.Ui32(v11) {
			v126 = l1
			v127 = v11
			v128 = v17
			v129 = v17
			v130 = v17
			for {
				v132 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
				v133 = v132 + v129
				v134 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
				v136 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
				v137 = v136 + v130
				v139 = int32(4)
				v141 = v134 + v128 - v137 ^ base.I32_rotl(v137, v139)
				v145 = v133 - v141 ^ base.I32_rotl(v141, int32(6))
				v146 = v137 + v133
				v147 = v141 + v146
				v148 = v145 + v147
				v152 = v146 - v145 ^ base.I32_rotl(v145, int32(8))
				v156 = v147 - v152 ^ base.I32_rotl(v152, int32(16))
				v160 = v148 - v156 ^ base.I32_rotl(v156, int32(19))
				v161 = v152 + v148
				v162 = v156 + v161
				v163 = v160 + v162
				v167 = v161 - v160 ^ base.I32_rotl(v160, v139)
				v168 = int32(12)
				v169 = v126 + v168
				v171 = v127 - v168
				if base.Ui32(int32(11)) < base.Ui32(v171) {
					v126 = v169
					v127 = v171
					v128 = v162
					v129 = v163
					v130 = v167
					continue
				} else {
					break
				}
				break
			}
			v174 = v169
			v175 = v171
			v176 = v162
			v177 = v163
			v178 = v167
		} else {
			v174 = l1
			v175 = v11
			v176 = v17
			v177 = v17
			v178 = v17
		}
		switch v175 - int32(1) {
		case 0:
			v237 = v176
			v238 = v177
			v239 = v178
			v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
			v244 = v237 + v240
			v245 = v238
			v246 = v239
		case 1:
			v230 = v176
			v231 = v177
			v232 = v178
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
			v237 = v233<<(uint(int32(8))%32) + v230
			v238 = v231
			v239 = v232
			v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
			v244 = v237 + v240
			v245 = v238
			v246 = v239
		case 2:
			v223 = v176
			v224 = v177
			v225 = v178
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+2)))
			v230 = v226<<(uint(int32(16))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
			v237 = v233<<(uint(int32(8))%32) + v230
			v238 = v231
			v239 = v232
			v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
			v244 = v237 + v240
			v245 = v238
			v246 = v239
		case 3:
			v217 = v177
			v218 = v178
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+3)))
			v223 = v219<<(uint(int32(24))%32) + v176
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+2)))
			v230 = v226<<(uint(int32(16))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
			v237 = v233<<(uint(int32(8))%32) + v230
			v238 = v231
			v239 = v232
			v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
			v244 = v237 + v240
			v245 = v238
			v246 = v239
		case 4:
			v213 = v177
			v214 = v178
			v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+4)))
			v217 = v213 + v215
			v218 = v214
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+3)))
			v223 = v219<<(uint(int32(24))%32) + v176
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+2)))
			v230 = v226<<(uint(int32(16))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
			v237 = v233<<(uint(int32(8))%32) + v230
			v238 = v231
			v239 = v232
			v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
			v244 = v237 + v240
			v245 = v238
			v246 = v239
		case 5:
			v207 = v177
			v208 = v178
			v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+5)))
			v213 = v209<<(uint(int32(8))%32) + v207
			v214 = v208
			v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+4)))
			v217 = v213 + v215
			v218 = v214
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+3)))
			v223 = v219<<(uint(int32(24))%32) + v176
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+2)))
			v230 = v226<<(uint(int32(16))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
			v237 = v233<<(uint(int32(8))%32) + v230
			v238 = v231
			v239 = v232
			v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
			v244 = v237 + v240
			v245 = v238
			v246 = v239
		case 6:
			v201 = v177
			v202 = v178
			v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+6)))
			v207 = v203<<(uint(int32(16))%32) + v201
			v208 = v202
			v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+5)))
			v213 = v209<<(uint(int32(8))%32) + v207
			v214 = v208
			v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+4)))
			v217 = v213 + v215
			v218 = v214
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+3)))
			v223 = v219<<(uint(int32(24))%32) + v176
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+2)))
			v230 = v226<<(uint(int32(16))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
			v237 = v233<<(uint(int32(8))%32) + v230
			v238 = v231
			v239 = v232
			v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
			v244 = v237 + v240
			v245 = v238
			v246 = v239
		case 7:
			v196 = v178
			v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+7)))
			v201 = v197<<(uint(int32(24))%32) + v177
			v202 = v196
			v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+6)))
			v207 = v203<<(uint(int32(16))%32) + v201
			v208 = v202
			v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+5)))
			v213 = v209<<(uint(int32(8))%32) + v207
			v214 = v208
			v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+4)))
			v217 = v213 + v215
			v218 = v214
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+3)))
			v223 = v219<<(uint(int32(24))%32) + v176
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+2)))
			v230 = v226<<(uint(int32(16))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
			v237 = v233<<(uint(int32(8))%32) + v230
			v238 = v231
			v239 = v232
			v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
			v244 = v237 + v240
			v245 = v238
			v246 = v239
		case 8:
			v191 = v178
			v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+8)))
			v196 = v192<<(uint(int32(8))%32) + v191
			v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+7)))
			v201 = v197<<(uint(int32(24))%32) + v177
			v202 = v196
			v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+6)))
			v207 = v203<<(uint(int32(16))%32) + v201
			v208 = v202
			v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+5)))
			v213 = v209<<(uint(int32(8))%32) + v207
			v214 = v208
			v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+4)))
			v217 = v213 + v215
			v218 = v214
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+3)))
			v223 = v219<<(uint(int32(24))%32) + v176
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+2)))
			v230 = v226<<(uint(int32(16))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
			v237 = v233<<(uint(int32(8))%32) + v230
			v238 = v231
			v239 = v232
			v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
			v244 = v237 + v240
			v245 = v238
			v246 = v239
		case 9:
			v186 = v178
			v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+9)))
			v191 = v187<<(uint(int32(16))%32) + v186
			v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+8)))
			v196 = v192<<(uint(int32(8))%32) + v191
			v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+7)))
			v201 = v197<<(uint(int32(24))%32) + v177
			v202 = v196
			v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+6)))
			v207 = v203<<(uint(int32(16))%32) + v201
			v208 = v202
			v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+5)))
			v213 = v209<<(uint(int32(8))%32) + v207
			v214 = v208
			v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+4)))
			v217 = v213 + v215
			v218 = v214
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+3)))
			v223 = v219<<(uint(int32(24))%32) + v176
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+2)))
			v230 = v226<<(uint(int32(16))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
			v237 = v233<<(uint(int32(8))%32) + v230
			v238 = v231
			v239 = v232
			v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
			v244 = v237 + v240
			v245 = v238
			v246 = v239
		case 10:
			v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+10)))
			v186 = v182<<(uint(int32(24))%32) + v178
			v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+9)))
			v191 = v187<<(uint(int32(16))%32) + v186
			v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+8)))
			v196 = v192<<(uint(int32(8))%32) + v191
			v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+7)))
			v201 = v197<<(uint(int32(24))%32) + v177
			v202 = v196
			v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+6)))
			v207 = v203<<(uint(int32(16))%32) + v201
			v208 = v202
			v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+5)))
			v213 = v209<<(uint(int32(8))%32) + v207
			v214 = v208
			v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+4)))
			v217 = v213 + v215
			v218 = v214
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+3)))
			v223 = v219<<(uint(int32(24))%32) + v176
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+2)))
			v230 = v226<<(uint(int32(16))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
			v237 = v233<<(uint(int32(8))%32) + v230
			v238 = v231
			v239 = v232
			v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
			v244 = v237 + v240
			v245 = v238
			v246 = v239
		default:
			v244 = v176
			v245 = v177
			v246 = v178
		}
	} else {
		if base.Ui32(v11) < base.Ui32(int32(12)) {
			v72 = l1
			v73 = v11
			v74 = v17
			v75 = v17
			v76 = v17
		} else {
			v24 = l1
			v25 = v11
			v26 = v17
			v27 = v17
			v28 = v17
			for {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
				v31 = v30 + v27
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
				v35 = v34 + v28
				v37 = int32(4)
				v39 = v32 + v26 - v35 ^ base.I32_rotl(v35, v37)
				v43 = v31 - v39 ^ base.I32_rotl(v39, int32(6))
				v44 = v35 + v31
				v45 = v39 + v44
				v46 = v43 + v45
				v50 = v44 - v43 ^ base.I32_rotl(v43, int32(8))
				v54 = v45 - v50 ^ base.I32_rotl(v50, int32(16))
				v58 = v46 - v54 ^ base.I32_rotl(v54, int32(19))
				v59 = v50 + v46
				v60 = v54 + v59
				v61 = v58 + v60
				v65 = v59 - v58 ^ base.I32_rotl(v58, v37)
				v66 = int32(12)
				v67 = v24 + v66
				v69 = v25 - v66
				if base.Ui32(int32(11)) < base.Ui32(v69) {
					v24 = v67
					v25 = v69
					v26 = v60
					v27 = v61
					v28 = v65
					continue
				} else {
					break
				}
				break
			}
			v72 = v67
			v73 = v69
			v74 = v60
			v75 = v61
			v76 = v65
		}
		switch v73 - int32(1) {
		case 0:
			v123 = v74
			v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
			v244 = v123 + v124
			v245 = v75
			v246 = v76
		case 1:
			v118 = v74
			v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
			v123 = v119<<(uint(int32(8))%32) + v118
			v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
			v244 = v123 + v124
			v245 = v75
			v246 = v76
		case 2:
			v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+2)))
			v118 = v114<<(uint(int32(16))%32) + v74
			v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
			v123 = v119<<(uint(int32(8))%32) + v118
			v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
			v244 = v123 + v124
			v245 = v75
			v246 = v76
		case 3:
			v111 = v75
			v112 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
			v244 = v112 + v74
			v245 = v111
			v246 = v76
		case 4:
			v108 = v75
			v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+4)))
			v111 = v108 + v109
			v112 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
			v244 = v112 + v74
			v245 = v111
			v246 = v76
		case 5:
			v103 = v75
			v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+5)))
			v108 = v104<<(uint(int32(8))%32) + v103
			v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+4)))
			v111 = v108 + v109
			v112 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
			v244 = v112 + v74
			v245 = v111
			v246 = v76
		case 6:
			v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+6)))
			v103 = v99<<(uint(int32(16))%32) + v75
			v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+5)))
			v108 = v104<<(uint(int32(8))%32) + v103
			v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+4)))
			v111 = v108 + v109
			v112 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
			v244 = v112 + v74
			v245 = v111
			v246 = v76
		case 7:
			v94 = v76
			v95 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
			v97 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
			v244 = v95 + v74
			v245 = v97 + v75
			v246 = v94
		case 8:
			v89 = v76
			v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+8)))
			v94 = v90<<(uint(int32(8))%32) + v89
			v95 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
			v97 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
			v244 = v95 + v74
			v245 = v97 + v75
			v246 = v94
		case 9:
			v84 = v76
			v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+9)))
			v89 = v85<<(uint(int32(16))%32) + v84
			v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+8)))
			v94 = v90<<(uint(int32(8))%32) + v89
			v95 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
			v97 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
			v244 = v95 + v74
			v245 = v97 + v75
			v246 = v94
		case 10:
			v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+10)))
			v84 = v80<<(uint(int32(24))%32) + v76
			v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+9)))
			v89 = v85<<(uint(int32(16))%32) + v84
			v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+8)))
			v94 = v90<<(uint(int32(8))%32) + v89
			v95 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
			v97 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
			v244 = v95 + v74
			v245 = v97 + v75
			v246 = v94
		default:
			v244 = v74
			v245 = v75
			v246 = v76
		}
	}
	v249 = int32(14)
	v251 = v245 ^ v246 - base.I32_rotl(v245, v249)
	v255 = v251 ^ v244 - base.I32_rotl(v251, int32(11))
	v259 = v255 ^ v245 - base.I32_rotl(v255, int32(25))
	v263 = v259 ^ v251 - base.I32_rotl(v259, int32(16))
	v267 = v263 ^ v255 - base.I32_rotl(v263, int32(4))
	v271 = v267 ^ v259 - base.I32_rotl(v267, v249)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v277 = base.I32_rem_u_s(v271^v263-base.I32_rotl(v271, int32(24)), v276)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0+v277<<(uint(int32(2))%32))+12))
	F_TempTablespacePath(m, v7+int32(1056), v281)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		return int32(0)
	} else {
		v286 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = int32(_a_F_FileSetDelete_0)
		*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v286
		*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v7 + int32(1056)
		v299 = F_pg_snprintf(m, v7+int32(32), int32(1024), int32(_a_F_FileSetDelete_1), v7+int32(16))
		mBase = m.M
		v300 = m.ExcPending
		if v300 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 + int32(32)
			v309 = F_pg_snprintf(m, v7+int32(1056), int32(1024), int32(_a_F_FileSetDelete_2), v7)
			mBase = m.M
			v310 = m.ExcPending
			if v310 != 0 {
				return int32(0)
			} else {
				v314 = F_PathNameDeleteTemporaryFile(m, v7+int32(1056), int32(1))
				mBase = m.M
				v315 = m.ExcPending
				if v315 != 0 {
					return int32(0)
				} else {
					m.G0 = v7 + int32(2080)
					return v314
				}
			}
		}
	}
}
func F_FileWriteV(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v52 int64
	_ = v52
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int64
	_ = v82
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v100 int64
	_ = v100
	var v105 int32
	_ = v105
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v124 int64
	_ = v124
	var v133 int64
	_ = v133
	var v136 int64
	_ = v136
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
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
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v199 int64
	_ = v199
	var v200 int64
	_ = v200
	var v202 int32
	_ = v202
	var v204 int64
	_ = v204
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	v8 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = F_FileAccess(m, l0)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L3
	} else {
		goto L37
	}
L2:
	;
	m.G0 = v17 + int32(16)
	return v213
L3:
	;
	return int32(0)
L4:
	;
	if v19 < int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v213 = int32(-1)
	goto L2
L6:
	;
	goto L7
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_FileWriteV[0]))
	v30 = v27 + l0*int32(48)
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_FileWriteV[1]))
	if v32 < int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	goto L25
L9:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+4)))
	if v35&int32(4) == int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if l2 <= int32(0) {
		v124 = l3
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v30)+24))
	if v124 <= v133 {
		goto L8
	} else {
		goto L23
	}
L12:
	;
	v42 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(l2) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v47 = v42
	v52 = l3
	v60 = v8
	goto L16
L14:
	;
	v77 = v42
	v82 = l3
	goto L15
L15:
	;
	v92 = l2 & int32(3)
	if v92 == int32(0) {
		v124 = v82
		goto L11
	} else {
		goto L19
	}
L16:
	;
	v63 = l1 + v47<<(uint(int32(3))%32)
	v64 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v63)+4)))
	v66 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v63)+12)))
	v68 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v63)+20)))
	v70 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v63)+28)))
	v71 = v52 + v64 + v66 + v68 + v70
	v72 = int32(4)
	v73 = v47 + v72
	v75 = v60 + v72
	if v75 != l2&int32(2147483644) {
		v47 = v73
		v52 = v71
		v60 = v75
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v77 = v73
	v82 = v71
	goto L15
L18:
	;
	goto L17
L19:
	;
	v95 = v77
	v100 = v82
	v105 = v8
	goto L20
L20:
	;
	v112 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1+v95<<(uint(int32(3))%32))+4)))
	v113 = v100 + v112
	v114 = int32(1)
	v117 = v105 + v114
	if v117 != v92 {
		v95 = v95 + v114
		v100 = v113
		v105 = v117
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v124 = v113
	goto L11
L22:
	;
	goto L21
L23:
	;
	v136 = *(*int64)(unsafe.Add(mBase, _c_F_FileWriteV[2]))
	if base.Ui64(base.I64_extend_i32_u(v32)<<(uint(int64(10))%64)) < base.Ui64(v136+(v124-v133)) {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L8
L25:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_FileWriteV[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v174))) = l4
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if base.B2i32(l2 != int32(1)) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v213 = v183
	goto L2
L27:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_FileWriteV[3]))
	v186 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v185))) = v186
	if v186 <= v183 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v181 = F_pwrite(m, v176, v179, v180, l3)
	mBase = m.M
	v183 = v181
	goto L27
L29:
	;
	goto L30
L30:
	;
	v182 = F_pwritev(m, v176, l1, l2, l3)
	mBase = m.M
	v183 = v182
	goto L27
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FileWriteV[4])) = int32(51)
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+4)))
	if v193&int32(4) == int32(0) {
		v213 = v183
		goto L2
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_FileWriteV[4]))
	if v210 == int32(27) {
		goto L25
	} else {
		goto L36
	}
L34:
	;
	v199 = l3 + base.I64_extend_i32_u(v183)
	v200 = *(*int64)(unsafe.Add(mBase, uint32(v30)+24))
	if v199 <= v200 {
		v213 = v183
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v202 = int32(_a_F_FileWriteV_0)
	v204 = *(*int64)(unsafe.Add(mBase, _c_F_FileWriteV[2]))
	*(*int64)(unsafe.Add(mBase, _c_F_FileWriteV[2])) = v204 + (v199 - v200)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+24)) = v199
	v213 = v183
	goto L2
L36:
	;
	goto L26
L37:
	;
	F_errcode(m, int32(_a_F_FileWriteV_1))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_FileWriteV[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v239
	F_errmsg(m, int32(_a_F_FileWriteV_2), v17)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_FileWriteV_3), int32(2295), int32(_a_F_FileWriteV_4))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
