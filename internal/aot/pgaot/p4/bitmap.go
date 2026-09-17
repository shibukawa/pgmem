package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecEndBitmapAnd(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if int32(0) < v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v10 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v8+v10<<(uint(int32(2))%32))))
	if v17 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	F_ExecEndNode(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v21 = v10 + int32(1)
	if v21 != v5 {
		v10 = v21
		goto L4
	} else {
		goto L11
	}
L9:
	;
	return
L10:
	;
	goto L8
L11:
	;
	goto L5
}
func F_bitmap_hash(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3 == int32(0) {
		return int32(0)
	} else {
		v9 = v3 + int32(8)
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
		v12 = v10 << (uint(int32(2)) % 32)
		v18 = v12 - int32(1636608432)
		if v9&int32(3) != 0 {
			if base.Ui32(int32(11)) < base.Ui32(v12) {
				v127 = v9
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
				v175 = v9
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
				v73 = v9
				v74 = v12
				v75 = v18
				v76 = v18
				v77 = v18
			} else {
				v25 = v9
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
		return v272 ^ v264 - base.I32_rotl(v272, int32(24))
	}
}
func F_create_bitmap_subplan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v102 float64
	_ = v102
	var v104 float64
	_ = v104
	var v106 float64
	_ = v106
	var v107 int32
	_ = v107
	var v108 float64
	_ = v108
	var v109 float64
	_ = v109
	var v110 float64
	_ = v110
	var v119 float64
	_ = v119
	var v123 float64
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
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
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
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
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int64
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int64
	_ = v210
	var v221 float64
	_ = v221
	var v223 float64
	_ = v223
	var v224 int32
	_ = v224
	var v225 float64
	_ = v225
	var v226 float64
	_ = v226
	var v227 float64
	_ = v227
	var v236 float64
	_ = v236
	var v240 float64
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
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
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int64
	_ = v439
	var v446 float64
	_ = v446
	var v448 float64
	_ = v448
	var v450 float64
	_ = v450
	var v451 int32
	_ = v451
	var v452 float64
	_ = v452
	var v453 float64
	_ = v453
	var v454 float64
	_ = v454
	var v463 float64
	_ = v463
	var v467 float64
	_ = v467
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	v6 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(48)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v21 - int32(280) {
	case 0:
		goto L11
	default:
		goto L10
	case 3:
		goto L13
	case 4:
		goto L12
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v579
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v581
	m.G0 = v19 + int32(48)
	return v576
L2:
	;
	v569 = int32(0)
	v576 = v482
	v579 = v569
	v581 = v569
	goto L1
L3:
	;
	v553 = int32(0)
	if v544 == v553 {
		goto L102
	} else {
		goto L103
	}
L4:
	;
	v525 = F_make_orclause(m, v488)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L19
	} else {
		goto L99
	}
L5:
	;
	v517 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v517
	if v513&int32(1) == v517 {
		v542 = v506
		v544 = v508
		goto L3
	} else {
		goto L98
	}
L6:
	;
	if v488 != 0 {
		goto L93
	} else {
		goto L94
	}
L7:
	;
	v437 = F_palloc0(m, int32(80))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L19
	} else {
		goto L87
	}
L8:
	;
	if v401 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L9:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)+88))
	if v325 == int32(0) {
		v383 = v316
		v384 = v317
		goto L69
	} else {
		goto L70
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L19
	} else {
		goto L66
	}
L11:
	;
	v199 = int32(0)
	v202 = F_create_indexscan_plan(m, l0, l1, v199, v199, v199)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L19
	} else {
		goto L48
	}
L12:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v132 == int32(0) {
		v427 = v6
		v429 = v6
		v430 = v6
		v431 = v6
		v432 = v6
		goto L7
	} else {
		goto L31
	}
L13:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v24 == int32(0) {
		v84 = v6
		v86 = v6
		v87 = v6
		v88 = v6
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v93 = F_palloc0(m, int32(80))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L19
	} else {
		goto L26
	}
L15:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v27 <= int32(0) {
		v84 = v6
		v86 = v6
		v87 = v6
		v88 = v6
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v38 = v6
	v39 = v6
	v40 = v6
	v41 = v6
	v42 = v6
	goto L17
L17:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v39<<(uint(int32(2))%32))))
	v57 = F_create_bitmap_subplan(m, l0, v50, v19+int32(44), v19+int32(40), v19+int32(36))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v84 = v67
	v86 = v70
	v87 = v61
	v88 = v64
	goto L14
L19:
	;
	return int32(0)
L20:
	;
	v61 = F_lappend(m, v41, v57)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	v64 = F_list_concat_unique(m, v42, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	v67 = F_list_concat_unique(m, v38, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	v70 = F_list_concat(m, v40, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v73 = v39 + int32(1)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v73 < v74 {
		v38 = v67
		v39 = v73
		v40 = v70
		v41 = v61
		v42 = v64
		goto L17
	} else {
		goto L25
	}
L25:
	;
	goto L18
L26:
	;
	v95 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v93)+44)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = int32(337)
	*(*int64)(unsafe.Add(mBase, uint32(v93)+52)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v93)+72)) = v87
	v102 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v93)+8)) = v102
	v104 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v93)+16)) = v104
	v106 = *(*float64)(unsafe.Add(mBase, uint32(l1)+80))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v108 = *(*float64)(unsafe.Add(mBase, uint32(v107)+120))
	v109 = base.F64_mul(v106, v108)
	v110 = float64(1e+100)
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v109)&int64(9223372036854775807)))|base.F64_gt(v109, v110) != 0 {
		v123 = v110
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v124 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+36)) = uint8(v124)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+32)) = v124
	*(*float64)(unsafe.Add(mBase, uint32(v93)+24)) = v123
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+37)) = uint8(v129)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v88
	v576 = v93
	v579 = v84
	v581 = v86
	goto L1
L28:
	;
	goto L27
L29:
	;
	v119 = float64(1)
	if base.F64_le(v109, v119) != 0 {
		v123 = v119
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v123 = base.F64_nearest(v109)
	goto L28
L31:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v135 <= int32(0) {
		v399 = v6
		v401 = v6
		v402 = v6
		v403 = v6
		v404 = v6
		goto L8
	} else {
		goto L32
	}
L32:
	;
	v145 = v6
	v146 = v6
	v147 = v6
	v148 = v6
	v149 = v6
	v150 = v6
	goto L33
L33:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v154+v146<<(uint(int32(2))%32))))
	v165 = F_create_bitmap_subplan(m, l0, v158, v19+int32(44), v19+int32(40), v19+int32(36))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L19
	} else {
		goto L35
	}
L34:
	;
	v399 = v194
	v401 = v167
	v402 = v172
	v403 = v181
	v404 = v185
	goto L8
L35:
	;
	v167 = F_lappend(m, v147, v165)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	v170 = int32(0)
	v172 = v148 | base.B2i32(v169 == v170)
	if v172&int32(1) == v170 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v177 = F_make_ands_explicit(m, v169)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L19
	} else {
		goto L40
	}
L38:
	;
	v181 = v149
	goto L39
L39:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	v183 = int32(0)
	v185 = v150 | base.B2i32(v182 == v183)
	if v185&int32(1) == v183 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v179 = F_lappend(m, v149, v177)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L19
	} else {
		goto L41
	}
L41:
	;
	v181 = v179
	goto L39
L42:
	;
	v190 = F_make_ands_explicit(m, v182)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L19
	} else {
		goto L45
	}
L43:
	;
	v194 = v145
	goto L44
L44:
	;
	v196 = v146 + int32(1)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v196 < v197 {
		v145 = v194
		v146 = v196
		v147 = v167
		v148 = v172
		v149 = v181
		v150 = v185
		goto L33
	} else {
		goto L47
	}
L45:
	;
	v192 = F_lappend(m, v145, v190)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L19
	} else {
		goto L46
	}
L46:
	;
	v194 = v192
	goto L44
L47:
	;
	goto L34
L48:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v202)+72))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v202)+80))
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v202)+84))
	v208 = F_palloc0(m, int32(96))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L19
	} else {
		goto L49
	}
L49:
	;
	v210 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v208)+44)) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v208))) = int32(343)
	*(*int64)(unsafe.Add(mBase, uint32(v208)+52)) = v210
	*(*int64)(unsafe.Add(mBase, uint32(v208)+88)) = v206
	*(*int32)(unsafe.Add(mBase, uint32(v208)+80)) = v205
	*(*int32)(unsafe.Add(mBase, uint32(v208)+72)) = v204
	*(*int64)(unsafe.Add(mBase, uint32(v208)+8)) = v210
	v221 = *(*float64)(unsafe.Add(mBase, uint32(l1)+96))
	*(*float64)(unsafe.Add(mBase, uint32(v208)+16)) = v221
	v223 = *(*float64)(unsafe.Add(mBase, uint32(l1)+104))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v225 = *(*float64)(unsafe.Add(mBase, uint32(v224)+120))
	v226 = base.F64_mul(v223, v225)
	v227 = float64(1e+100)
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v226)&int64(9223372036854775807)))|base.F64_gt(v226, v227) != 0 {
		v240 = v227
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v241 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v208)+36)) = uint8(v241)
	*(*int32)(unsafe.Add(mBase, uint32(v208)+32)) = v241
	*(*float64)(unsafe.Add(mBase, uint32(v208)+24)) = v240
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v208)+37)) = uint8(v246)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v249 == v241 {
		v316 = v241
		v317 = v6
		v318 = v6
		goto L9
	} else {
		goto L54
	}
L51:
	;
	goto L50
L52:
	;
	v236 = float64(1)
	if base.F64_le(v226, v236) != 0 {
		v240 = v236
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v240 = base.F64_nearest(v226)
	goto L51
L54:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	if v252 <= int32(0) {
		v316 = v241
		v317 = v6
		v318 = v6
		goto L9
	} else {
		goto L55
	}
L55:
	;
	v256 = int32(0)
	v264 = v241
	v265 = v6
	v266 = v6
	goto L56
L56:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v249)+12))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v272+v256<<(uint(int32(2))%32))))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	v279 = F_lappend(m, v265, v278)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L19
	} else {
		goto L58
	}
L57:
	;
	v316 = v284
	v317 = v279
	v318 = v289
	goto L9
L58:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v276)+8))
	v282 = F_get_actual_clauses(m, v281)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L19
	} else {
		goto L59
	}
L59:
	;
	v284 = F_list_concat(m, v264, v282)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L19
	} else {
		goto L60
	}
L60:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v277)+60))
	if v286 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v287 = F_lappend(m, v266, v286)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L19
	} else {
		goto L64
	}
L62:
	;
	v289 = v266
	goto L63
L63:
	;
	v291 = v256 + int32(1)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	if v291 < v292 {
		v256 = v291
		v264 = v284
		v265 = v279
		v266 = v289
		goto L56
	} else {
		goto L65
	}
L64:
	;
	v289 = v287
	goto L63
L65:
	;
	goto L57
L66:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v298
	F_errmsg_internal(m, int32(_a_F_create_bitmap_subplan_0), v19)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L19
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_create_bitmap_subplan_1), int32(3520), int32(_a_F_create_bitmap_subplan_2))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L19
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v384
	v576 = v208
	v579 = v383
	v581 = v318
	goto L1
L70:
	;
	v328 = int32(0)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v325)+4))
	if v329 <= v328 {
		v383 = v316
		v384 = v317
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v332 = v328
	v340 = v316
	v341 = v317
	goto L72
L72:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v325)+12))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v348+v332<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v352
	v358 = F_list_make1_impl(m, int32(1), v19+int32(20))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L19
	} else {
		goto L74
	}
L73:
	;
	v383 = v369
	v384 = v370
	goto L69
L74:
	;
	v361 = F_predicate_implied_by(m, v358, v341, int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L19
	} else {
		goto L75
	}
L75:
	;
	if v361 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v365 = F_lappend(m, v341, v352)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L19
	} else {
		goto L79
	}
L77:
	;
	v369 = v340
	v370 = v341
	goto L78
L78:
	;
	v372 = v332 + int32(1)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v325)+4))
	if v372 < v373 {
		v332 = v372
		v340 = v369
		v341 = v370
		goto L72
	} else {
		goto L81
	}
L79:
	;
	v367 = F_lappend(m, v340, v352)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L19
	} else {
		goto L80
	}
L80:
	;
	v369 = v367
	v370 = v365
	goto L78
L81:
	;
	goto L73
L82:
	;
	v427 = v399
	v429 = int32(0)
	v430 = v402
	v431 = v403
	v432 = v404
	goto L7
L83:
	;
	goto L84
L84:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v401)+4))
	if v411 != int32(1) {
		v427 = v399
		v429 = v401
		v430 = v402
		v431 = v403
		v432 = v404
		goto L7
	} else {
		goto L85
	}
L85:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v401)+12))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v414)))
	if v402&int32(1) == int32(0) {
		v482 = v415
		v484 = v399
		v488 = v403
		v489 = v404
		goto L6
	} else {
		goto L86
	}
L86:
	;
	v506 = v415
	v508 = v399
	v513 = v404
	goto L5
L87:
	;
	v439 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v437)+44)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v437))) = int32(338)
	*(*int64)(unsafe.Add(mBase, uint32(v437)+52)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v437)+76)) = v429
	v446 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v437)+8)) = v446
	v448 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v437)+16)) = v448
	v450 = *(*float64)(unsafe.Add(mBase, uint32(l1)+80))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v452 = *(*float64)(unsafe.Add(mBase, uint32(v451)+120))
	v453 = base.F64_mul(v450, v452)
	v454 = float64(1e+100)
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v453)&int64(9223372036854775807)))|base.F64_gt(v453, v454) != 0 {
		v467 = v454
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v468 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v437)+36)) = uint8(v468)
	*(*int32)(unsafe.Add(mBase, uint32(v437)+32)) = v468
	*(*float64)(unsafe.Add(mBase, uint32(v437)+24)) = v467
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v437)+37)) = uint8(v473)
	if v430&int32(1) != 0 {
		v506 = v437
		v508 = v427
		v513 = v432
		goto L5
	} else {
		goto L92
	}
L89:
	;
	goto L88
L90:
	;
	v463 = float64(1)
	if base.F64_le(v453, v463) != 0 {
		v467 = v463
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v467 = base.F64_nearest(v453)
	goto L89
L92:
	;
	v482 = v437
	v484 = v427
	v488 = v431
	v489 = v432
	goto L6
L93:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v488)+4))
	if int32(1) < v493 {
		goto L4
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v488
	if v489&int32(1) == int32(0) {
		v542 = v482
		v544 = v484
		goto L3
	} else {
		goto L97
	}
L96:
	;
	goto L95
L97:
	;
	goto L2
L98:
	;
	v576 = v506
	v579 = v517
	v581 = v517
	goto L1
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v525
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v525
	v532 = F_list_make1_impl(m, int32(1), v19+int32(16))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L19
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v532
	if v489&int32(1) != 0 {
		goto L2
	} else {
		goto L101
	}
L101:
	;
	v542 = v482
	v544 = v484
	goto L3
L102:
	;
	v576 = v542
	v579 = v544
	v581 = v553
	goto L1
L103:
	;
	goto L104
L104:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v544)+4))
	if v556 < int32(2) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v576 = v542
	v579 = v544
	v581 = v553
	goto L1
L106:
	;
	goto L107
L107:
	;
	v559 = F_make_orclause(m, v544)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L19
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v559
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v559
	v566 = F_list_make1_impl(m, int32(1), v19+int32(12))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L19
	} else {
		goto L109
	}
L109:
	;
	v576 = v542
	v579 = v566
	v581 = v553
	goto L1
}
func F_storeBitmap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	F_tbm_add_tuples(m, v9, l1, int32(1), l5)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+200))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+200)) = v13 + int64(1)
		return
	}
}
