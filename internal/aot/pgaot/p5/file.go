package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FileSetDelete(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
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
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
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
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
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
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int64
	_ = v287
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	v6 = m.G0
	v8 = v6 - int32(2080)
	m.G0 = v8
	v11 = v8 + int32(1056)
	v12 = F_strlen(m, l1)
	mBase = m.M
	v18 = v12 - int32(1636608432)
	if l1&int32(3) != 0 {
		if base.Ui32(int32(11)) < base.Ui32(v12) {
			v127 = l1
			v128 = v12
			v129 = v18
			v130 = v18
			v131 = v18
			for {
				v133 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
				v134 = v133 + v130
				v135 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
				v137 = *(*int32)(unsafe.Add(mBase, uint32(v127)+8))
				v138 = v137 + v131
				v140 = int32(4)
				v142 = v135 + v129 - v138 ^ base.I32_rotl(v138, v140)
				v146 = v134 - v142 ^ base.I32_rotl(v142, int32(6))
				v147 = v138 + v134
				v148 = v142 + v147
				v149 = v146 + v148
				v153 = v147 - v146 ^ base.I32_rotl(v146, int32(8))
				v157 = v148 - v153 ^ base.I32_rotl(v153, int32(16))
				v161 = v149 - v157 ^ base.I32_rotl(v157, int32(19))
				v162 = v153 + v149
				v163 = v157 + v162
				v164 = v161 + v163
				v168 = v162 - v161 ^ base.I32_rotl(v161, v140)
				v169 = int32(12)
				v170 = v127 + v169
				v172 = v128 - v169
				if base.Ui32(int32(11)) < base.Ui32(v172) {
					v127 = v170
					v128 = v172
					v129 = v163
					v130 = v164
					v131 = v168
					continue
				} else {
					break
				}
				break
			}
			v175 = v170
			v176 = v172
			v177 = v163
			v178 = v164
			v179 = v168
		} else {
			v175 = l1
			v176 = v12
			v177 = v18
			v178 = v18
			v179 = v18
		}
		switch v176 - int32(1) {
		case 0:
			v238 = v177
			v239 = v178
			v240 = v179
			v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
			v245 = v238 + v241
			v246 = v239
			v247 = v240
		case 1:
			v231 = v177
			v232 = v178
			v233 = v179
			v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
			v238 = v234<<(uint(int32(8))%32) + v231
			v239 = v232
			v240 = v233
			v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
			v245 = v238 + v241
			v246 = v239
			v247 = v240
		case 2:
			v224 = v177
			v225 = v178
			v226 = v179
			v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+2)))
			v231 = v227<<(uint(int32(16))%32) + v224
			v232 = v225
			v233 = v226
			v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
			v238 = v234<<(uint(int32(8))%32) + v231
			v239 = v232
			v240 = v233
			v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
			v245 = v238 + v241
			v246 = v239
			v247 = v240
		case 3:
			v218 = v178
			v219 = v179
			v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+3)))
			v224 = v220<<(uint(int32(24))%32) + v177
			v225 = v218
			v226 = v219
			v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+2)))
			v231 = v227<<(uint(int32(16))%32) + v224
			v232 = v225
			v233 = v226
			v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
			v238 = v234<<(uint(int32(8))%32) + v231
			v239 = v232
			v240 = v233
			v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
			v245 = v238 + v241
			v246 = v239
			v247 = v240
		case 4:
			v214 = v178
			v215 = v179
			v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+4)))
			v218 = v214 + v216
			v219 = v215
			v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+3)))
			v224 = v220<<(uint(int32(24))%32) + v177
			v225 = v218
			v226 = v219
			v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+2)))
			v231 = v227<<(uint(int32(16))%32) + v224
			v232 = v225
			v233 = v226
			v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
			v238 = v234<<(uint(int32(8))%32) + v231
			v239 = v232
			v240 = v233
			v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
			v245 = v238 + v241
			v246 = v239
			v247 = v240
		case 5:
			v208 = v178
			v209 = v179
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+5)))
			v214 = v210<<(uint(int32(8))%32) + v208
			v215 = v209
			v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+4)))
			v218 = v214 + v216
			v219 = v215
			v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+3)))
			v224 = v220<<(uint(int32(24))%32) + v177
			v225 = v218
			v226 = v219
			v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+2)))
			v231 = v227<<(uint(int32(16))%32) + v224
			v232 = v225
			v233 = v226
			v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
			v238 = v234<<(uint(int32(8))%32) + v231
			v239 = v232
			v240 = v233
			v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
			v245 = v238 + v241
			v246 = v239
			v247 = v240
		case 6:
			v202 = v178
			v203 = v179
			v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+6)))
			v208 = v204<<(uint(int32(16))%32) + v202
			v209 = v203
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+5)))
			v214 = v210<<(uint(int32(8))%32) + v208
			v215 = v209
			v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+4)))
			v218 = v214 + v216
			v219 = v215
			v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+3)))
			v224 = v220<<(uint(int32(24))%32) + v177
			v225 = v218
			v226 = v219
			v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+2)))
			v231 = v227<<(uint(int32(16))%32) + v224
			v232 = v225
			v233 = v226
			v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
			v238 = v234<<(uint(int32(8))%32) + v231
			v239 = v232
			v240 = v233
			v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
			v245 = v238 + v241
			v246 = v239
			v247 = v240
		case 7:
			v197 = v179
			v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+7)))
			v202 = v198<<(uint(int32(24))%32) + v178
			v203 = v197
			v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+6)))
			v208 = v204<<(uint(int32(16))%32) + v202
			v209 = v203
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+5)))
			v214 = v210<<(uint(int32(8))%32) + v208
			v215 = v209
			v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+4)))
			v218 = v214 + v216
			v219 = v215
			v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+3)))
			v224 = v220<<(uint(int32(24))%32) + v177
			v225 = v218
			v226 = v219
			v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+2)))
			v231 = v227<<(uint(int32(16))%32) + v224
			v232 = v225
			v233 = v226
			v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
			v238 = v234<<(uint(int32(8))%32) + v231
			v239 = v232
			v240 = v233
			v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
			v245 = v238 + v241
			v246 = v239
			v247 = v240
		case 8:
			v192 = v179
			v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+8)))
			v197 = v193<<(uint(int32(8))%32) + v192
			v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+7)))
			v202 = v198<<(uint(int32(24))%32) + v178
			v203 = v197
			v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+6)))
			v208 = v204<<(uint(int32(16))%32) + v202
			v209 = v203
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+5)))
			v214 = v210<<(uint(int32(8))%32) + v208
			v215 = v209
			v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+4)))
			v218 = v214 + v216
			v219 = v215
			v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+3)))
			v224 = v220<<(uint(int32(24))%32) + v177
			v225 = v218
			v226 = v219
			v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+2)))
			v231 = v227<<(uint(int32(16))%32) + v224
			v232 = v225
			v233 = v226
			v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
			v238 = v234<<(uint(int32(8))%32) + v231
			v239 = v232
			v240 = v233
			v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
			v245 = v238 + v241
			v246 = v239
			v247 = v240
		case 9:
			v187 = v179
			v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+9)))
			v192 = v188<<(uint(int32(16))%32) + v187
			v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+8)))
			v197 = v193<<(uint(int32(8))%32) + v192
			v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+7)))
			v202 = v198<<(uint(int32(24))%32) + v178
			v203 = v197
			v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+6)))
			v208 = v204<<(uint(int32(16))%32) + v202
			v209 = v203
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+5)))
			v214 = v210<<(uint(int32(8))%32) + v208
			v215 = v209
			v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+4)))
			v218 = v214 + v216
			v219 = v215
			v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+3)))
			v224 = v220<<(uint(int32(24))%32) + v177
			v225 = v218
			v226 = v219
			v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+2)))
			v231 = v227<<(uint(int32(16))%32) + v224
			v232 = v225
			v233 = v226
			v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
			v238 = v234<<(uint(int32(8))%32) + v231
			v239 = v232
			v240 = v233
			v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
			v245 = v238 + v241
			v246 = v239
			v247 = v240
		case 10:
			v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+10)))
			v187 = v183<<(uint(int32(24))%32) + v179
			v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+9)))
			v192 = v188<<(uint(int32(16))%32) + v187
			v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+8)))
			v197 = v193<<(uint(int32(8))%32) + v192
			v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+7)))
			v202 = v198<<(uint(int32(24))%32) + v178
			v203 = v197
			v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+6)))
			v208 = v204<<(uint(int32(16))%32) + v202
			v209 = v203
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+5)))
			v214 = v210<<(uint(int32(8))%32) + v208
			v215 = v209
			v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+4)))
			v218 = v214 + v216
			v219 = v215
			v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+3)))
			v224 = v220<<(uint(int32(24))%32) + v177
			v225 = v218
			v226 = v219
			v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+2)))
			v231 = v227<<(uint(int32(16))%32) + v224
			v232 = v225
			v233 = v226
			v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
			v238 = v234<<(uint(int32(8))%32) + v231
			v239 = v232
			v240 = v233
			v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
			v245 = v238 + v241
			v246 = v239
			v247 = v240
		default:
			v245 = v177
			v246 = v178
			v247 = v179
		}
	} else {
		if base.Ui32(v12) < base.Ui32(int32(12)) {
			v73 = l1
			v74 = v12
			v75 = v18
			v76 = v18
			v77 = v18
		} else {
			v25 = l1
			v26 = v12
			v27 = v18
			v28 = v18
			v29 = v18
			for {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
				v32 = v31 + v28
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
				v36 = v35 + v29
				v38 = int32(4)
				v40 = v33 + v27 - v36 ^ base.I32_rotl(v36, v38)
				v44 = v32 - v40 ^ base.I32_rotl(v40, int32(6))
				v45 = v36 + v32
				v46 = v40 + v45
				v47 = v44 + v46
				v51 = v45 - v44 ^ base.I32_rotl(v44, int32(8))
				v55 = v46 - v51 ^ base.I32_rotl(v51, int32(16))
				v59 = v47 - v55 ^ base.I32_rotl(v55, int32(19))
				v60 = v51 + v47
				v61 = v55 + v60
				v62 = v59 + v61
				v66 = v60 - v59 ^ base.I32_rotl(v59, v38)
				v67 = int32(12)
				v68 = v25 + v67
				v70 = v26 - v67
				if base.Ui32(int32(11)) < base.Ui32(v70) {
					v25 = v68
					v26 = v70
					v27 = v61
					v28 = v62
					v29 = v66
					continue
				} else {
					break
				}
				break
			}
			v73 = v68
			v74 = v70
			v75 = v61
			v76 = v62
			v77 = v66
		}
		switch v74 - int32(1) {
		case 0:
			v124 = v75
			v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
			v245 = v124 + v125
			v246 = v76
			v247 = v77
		case 1:
			v119 = v75
			v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
			v124 = v120<<(uint(int32(8))%32) + v119
			v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
			v245 = v124 + v125
			v246 = v76
			v247 = v77
		case 2:
			v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+2)))
			v119 = v115<<(uint(int32(16))%32) + v75
			v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
			v124 = v120<<(uint(int32(8))%32) + v119
			v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
			v245 = v124 + v125
			v246 = v76
			v247 = v77
		case 3:
			v112 = v76
			v113 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
			v245 = v113 + v75
			v246 = v112
			v247 = v77
		case 4:
			v109 = v76
			v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+4)))
			v112 = v109 + v110
			v113 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
			v245 = v113 + v75
			v246 = v112
			v247 = v77
		case 5:
			v104 = v76
			v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+5)))
			v109 = v105<<(uint(int32(8))%32) + v104
			v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+4)))
			v112 = v109 + v110
			v113 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
			v245 = v113 + v75
			v246 = v112
			v247 = v77
		case 6:
			v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+6)))
			v104 = v100<<(uint(int32(16))%32) + v76
			v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+5)))
			v109 = v105<<(uint(int32(8))%32) + v104
			v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+4)))
			v112 = v109 + v110
			v113 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
			v245 = v113 + v75
			v246 = v112
			v247 = v77
		case 7:
			v95 = v77
			v96 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
			v98 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
			v245 = v96 + v75
			v246 = v98 + v76
			v247 = v95
		case 8:
			v90 = v77
			v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+8)))
			v95 = v91<<(uint(int32(8))%32) + v90
			v96 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
			v98 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
			v245 = v96 + v75
			v246 = v98 + v76
			v247 = v95
		case 9:
			v85 = v77
			v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+9)))
			v90 = v86<<(uint(int32(16))%32) + v85
			v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+8)))
			v95 = v91<<(uint(int32(8))%32) + v90
			v96 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
			v98 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
			v245 = v96 + v75
			v246 = v98 + v76
			v247 = v95
		case 10:
			v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+10)))
			v85 = v81<<(uint(int32(24))%32) + v77
			v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+9)))
			v90 = v86<<(uint(int32(16))%32) + v85
			v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+8)))
			v95 = v91<<(uint(int32(8))%32) + v90
			v96 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
			v98 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
			v245 = v96 + v75
			v246 = v98 + v76
			v247 = v95
		default:
			v245 = v75
			v246 = v76
			v247 = v77
		}
	}
	v250 = int32(14)
	v252 = v246 ^ v247 - base.I32_rotl(v246, v250)
	v256 = v252 ^ v245 - base.I32_rotl(v252, int32(11))
	v260 = v256 ^ v246 - base.I32_rotl(v256, int32(25))
	v264 = v260 ^ v252 - base.I32_rotl(v260, int32(16))
	v268 = v264 ^ v256 - base.I32_rotl(v264, int32(4))
	v272 = v268 ^ v260 - base.I32_rotl(v268, v250)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v278 = base.I32_rem_u_s(v272^v264-base.I32_rotl(v272, int32(24)), v277)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0+v278<<(uint(int32(2))%32))+12))
	F_TempTablespacePath(m, v11, v282)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		return int32(0)
	} else {
		v287 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_FileSetDelete_0)
		*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v287
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v11
		v293 = v8 + int32(32)
		v298 = F_pg_snprintf(m, v293, int32(1024), int32(_a_F_FileSetDelete_1), v8+int32(16))
		mBase = m.M
		v299 = m.ExcPending
		if v299 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v293
			v304 = F_pg_snprintf(m, v11, int32(1024), int32(_a_F_FileSetDelete_2), v8)
			mBase = m.M
			v305 = m.ExcPending
			if v305 != 0 {
				return int32(0)
			} else {
				v307 = F_PathNameDeleteTemporaryFile(m, v11, int32(1))
				mBase = m.M
				v308 = m.ExcPending
				if v308 != 0 {
					return int32(0)
				} else {
					m.G0 = v8 + int32(2080)
					return v307
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v55 int64
	_ = v55
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int64
	_ = v88
	var v98 int32
	_ = v98
	var v103 int64
	_ = v103
	var v112 int32
	_ = v112
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int64
	_ = v128
	var v138 int64
	_ = v138
	var v141 int64
	_ = v141
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v206 int64
	_ = v206
	var v207 int64
	_ = v207
	var v209 int32
	_ = v209
	var v211 int64
	_ = v211
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	v8 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = F_FileAccess(m, l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L3
	} else {
		goto L37
	}
L2:
	;
	m.G0 = v18 + int32(16)
	return v220
L3:
	;
	return int32(0)
L4:
	;
	if v20 < int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v220 = int32(-1)
	goto L2
L6:
	;
	goto L7
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_FileWriteV[0]))
	v31 = v28 + l0*int32(48)
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_FileWriteV[1]))
	if v33 < int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	goto L25
L9:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+4)))
	if v36&int32(4) == int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if l2 <= int32(0) {
		v128 = l3
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v138 = *(*int64)(unsafe.Add(mBase, uint32(v31)+24))
	if v128 <= v138 {
		goto L8
	} else {
		goto L23
	}
L12:
	;
	v44 = l2 & int32(3)
	v45 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(l2) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v50 = v45
	v55 = l3
	v63 = v8
	goto L16
L14:
	;
	v83 = v45
	v88 = l3
	goto L15
L15:
	;
	v98 = v83
	v103 = v88
	v112 = v8
	goto L20
L16:
	;
	v67 = l1 + v50<<(uint(int32(3))%32)
	v68 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v67)+4)))
	v70 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v67)+12)))
	v72 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v67)+20)))
	v74 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v67)+28)))
	v75 = v55 + v68 + v70 + v72 + v74
	v76 = int32(4)
	v77 = v50 + v76
	v79 = v63 + v76
	if v79 != l2&int32(2147483644) {
		v50 = v77
		v55 = v75
		v63 = v79
		goto L16
	} else {
		goto L18
	}
L17:
	;
	if v44 == int32(0) {
		v128 = v75
		goto L11
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	v83 = v77
	v88 = v75
	goto L15
L20:
	;
	v116 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1+v98<<(uint(int32(3))%32))+4)))
	v117 = v103 + v116
	v118 = int32(1)
	v121 = v112 + v118
	if v121 != v44 {
		v98 = v98 + v118
		v103 = v117
		v112 = v121
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v128 = v117
	goto L11
L22:
	;
	goto L21
L23:
	;
	v141 = *(*int64)(unsafe.Add(mBase, _c_F_FileWriteV[2]))
	if base.Ui64(base.I64_extend_i32_u(v33)<<(uint(int64(10))%64)) < base.Ui64(v141+(v128-v138)) {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L8
L25:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_FileWriteV[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = l4
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if base.B2i32(l2 != int32(1)) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v220 = v190
	goto L2
L27:
	;
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_FileWriteV[3]))
	v193 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v192))) = v193
	if v193 <= v190 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v188 = F_pwrite(m, v183, v186, v187, l3)
	mBase = m.M
	v190 = v188
	goto L27
L29:
	;
	goto L30
L30:
	;
	v189 = F_pwritev(m, v183, l1, l2, l3)
	mBase = m.M
	v190 = v189
	goto L27
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FileWriteV[4])) = int32(51)
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+4)))
	if v200&int32(4) == int32(0) {
		v220 = v190
		goto L2
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_FileWriteV[4]))
	if v217 == int32(27) {
		goto L25
	} else {
		goto L36
	}
L34:
	;
	v206 = l3 + base.I64_extend_i32_u(v190)
	v207 = *(*int64)(unsafe.Add(mBase, uint32(v31)+24))
	if v206 <= v207 {
		v220 = v190
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v209 = int32(_a_F_FileWriteV_0)
	v211 = *(*int64)(unsafe.Add(mBase, _c_F_FileWriteV[2]))
	*(*int64)(unsafe.Add(mBase, _c_F_FileWriteV[2])) = v211 + (v206 - v207)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+24)) = v206
	v220 = v190
	goto L2
L36:
	;
	goto L26
L37:
	;
	F_errcode(m, int32(_a_F_FileWriteV_1))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_FileWriteV[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v247
	F_errmsg(m, int32(_a_F_FileWriteV_2), v18)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_FileWriteV_3), int32(2295), int32(_a_F_FileWriteV_4))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
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
