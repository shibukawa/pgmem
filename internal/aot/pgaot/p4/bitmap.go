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
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int64
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int64
	_ = v154
	var v165 float64
	_ = v165
	var v167 float64
	_ = v167
	var v168 int32
	_ = v168
	var v169 float64
	_ = v169
	var v170 float64
	_ = v170
	var v172 float64
	_ = v172
	var v180 float64
	_ = v180
	var v184 float64
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int64
	_ = v382
	var v389 float64
	_ = v389
	var v391 float64
	_ = v391
	var v393 float64
	_ = v393
	var v394 int32
	_ = v394
	var v395 float64
	_ = v395
	var v396 float64
	_ = v396
	var v398 float64
	_ = v398
	var v406 float64
	_ = v406
	var v410 float64
	_ = v410
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v487 int32
	_ = v487
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int64
	_ = v535
	var v542 float64
	_ = v542
	var v544 float64
	_ = v544
	var v546 float64
	_ = v546
	var v547 int32
	_ = v547
	var v548 float64
	_ = v548
	var v549 float64
	_ = v549
	var v551 float64
	_ = v551
	var v559 float64
	_ = v559
	var v563 float64
	_ = v563
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
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
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v578
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v580
	m.G0 = v19 + int32(48)
	return v577
L2:
	;
	v533 = F_palloc0(m, int32(80))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L20
	} else {
		goto L115
	}
L3:
	;
	v514 = int32(0)
	v577 = v451
	v578 = v514
	v580 = v514
	goto L1
L4:
	;
	v498 = int32(0)
	if v495 == v498 {
		goto L107
	} else {
		goto L108
	}
L5:
	;
	if v458 != 0 {
		goto L99
	} else {
		goto L100
	}
L6:
	;
	v438 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v438
	if v431&int32(1) == v438 {
		v487 = v427
		v495 = v435
		goto L4
	} else {
		goto L97
	}
L7:
	;
	v380 = F_palloc0(m, int32(80))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L20
	} else {
		goto L90
	}
L8:
	;
	if v347 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L9:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+88))
	if v270 == int32(0) {
		v326 = v259
		v330 = v263
		goto L72
	} else {
		goto L73
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L20
	} else {
		goto L69
	}
L11:
	;
	v143 = int32(0)
	v146 = F_create_indexscan_plan(m, l0, l1, v143, v143, v143)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L20
	} else {
		goto L48
	}
L12:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v76 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L13:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v24 == int32(0) {
		v522 = v6
		v524 = v6
		v525 = v6
		v527 = v6
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v27 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v522 = v6
	v524 = v6
	v525 = v6
	v527 = v6
	goto L2
L16:
	;
	goto L17
L17:
	;
	v36 = v6
	v38 = v6
	v39 = v6
	v40 = v6
	v41 = v6
	goto L18
L18:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v40<<(uint(int32(2))%32))))
	v57 = F_create_bitmap_subplan(m, l0, v50, v19+int32(44), v19+int32(40), v19+int32(36))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v522 = v67
	v524 = v70
	v525 = v64
	v527 = v61
	goto L2
L20:
	;
	return int32(0)
L21:
	;
	v61 = F_lappend(m, v41, v57)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	v64 = F_list_concat_unique(m, v39, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	v67 = F_list_concat_unique(m, v36, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	v70 = F_list_concat(m, v38, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	v73 = v40 + int32(1)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v73 < v74 {
		v36 = v67
		v38 = v70
		v39 = v64
		v40 = v73
		v41 = v61
		goto L18
	} else {
		goto L26
	}
L26:
	;
	goto L19
L27:
	;
	v371 = v6
	v372 = v6
	v373 = v6
	v375 = v6
	v376 = v6
	goto L7
L28:
	;
	goto L29
L29:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v79 <= int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v345 = v6
	v346 = v6
	v347 = v6
	v349 = v6
	v350 = v6
	goto L8
L31:
	;
	goto L32
L32:
	;
	v88 = v6
	v90 = v6
	v91 = v6
	v92 = v6
	v94 = v6
	v95 = v6
	goto L33
L33:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v98+v88<<(uint(int32(2))%32))))
	v109 = F_create_bitmap_subplan(m, l0, v102, v19+int32(44), v19+int32(40), v19+int32(36))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L20
	} else {
		goto L35
	}
L34:
	;
	v345 = v116
	v346 = v129
	v347 = v111
	v349 = v125
	v350 = v138
	goto L8
L35:
	;
	v111 = F_lappend(m, v92, v109)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L20
	} else {
		goto L36
	}
L36:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	v114 = int32(0)
	v116 = v90 | base.B2i32(v113 == v114)
	if v116&int32(1) == v114 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v121 = F_make_ands_explicit(m, v113)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L20
	} else {
		goto L40
	}
L38:
	;
	v125 = v94
	goto L39
L39:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	v127 = int32(0)
	v129 = v91 | base.B2i32(v126 == v127)
	if v129&int32(1) == v127 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v123 = F_lappend(m, v94, v121)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L20
	} else {
		goto L41
	}
L41:
	;
	v125 = v123
	goto L39
L42:
	;
	v134 = F_make_ands_explicit(m, v126)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L20
	} else {
		goto L45
	}
L43:
	;
	v138 = v95
	goto L44
L44:
	;
	v140 = v88 + int32(1)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v140 < v141 {
		v88 = v140
		v90 = v116
		v91 = v129
		v92 = v111
		v94 = v125
		v95 = v138
		goto L33
	} else {
		goto L47
	}
L45:
	;
	v136 = F_lappend(m, v95, v134)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L20
	} else {
		goto L46
	}
L46:
	;
	v138 = v136
	goto L44
L47:
	;
	goto L34
L48:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v146)+72))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v146)+80))
	v150 = *(*int64)(unsafe.Add(mBase, uint32(v146)+84))
	v152 = F_palloc0(m, int32(96))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L20
	} else {
		goto L49
	}
L49:
	;
	v154 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v152)+44)) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v152))) = int32(343)
	*(*int64)(unsafe.Add(mBase, uint32(v152)+88)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v152)+80)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v152)+72)) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v152)+8)) = v154
	*(*int64)(unsafe.Add(mBase, uint32(v152)+52)) = v154
	v165 = *(*float64)(unsafe.Add(mBase, uint32(l1)+96))
	*(*float64)(unsafe.Add(mBase, uint32(v152)+16)) = v165
	v167 = *(*float64)(unsafe.Add(mBase, uint32(l1)+104))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v169 = *(*float64)(unsafe.Add(mBase, uint32(v168)+120))
	v170 = base.F64_mul(v167, v169)
	v172 = float64(1e+100)
	if base.F64_gt(v170, v172) != 0 {
		v184 = v172
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v185 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v152)+36)) = uint8(v185)
	*(*int32)(unsafe.Add(mBase, uint32(v152)+32)) = v185
	*(*float64)(unsafe.Add(mBase, uint32(v152)+24)) = v184
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v152)+37)) = uint8(v190)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v194 == v185 {
		v259 = v185
		v261 = v185
		v263 = v6
		goto L9
	} else {
		goto L55
	}
L51:
	;
	goto L50
L52:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v170)&int64(9223372036854775807)) {
		v184 = v172
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v180 = float64(1)
	if base.F64_le(v170, v180) != 0 {
		v184 = v180
		goto L51
	} else {
		goto L54
	}
L54:
	;
	v184 = base.F64_nearest(v170)
	goto L51
L55:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v194)+4))
	if v197 <= int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v259 = v185
	v261 = v185
	v263 = v6
	goto L9
L57:
	;
	goto L58
L58:
	;
	v201 = int32(0)
	v207 = v185
	v209 = v185
	v211 = v6
	goto L59
L59:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v194)+12))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v217+v201<<(uint(int32(2))%32))))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	v224 = F_lappend(m, v211, v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L20
	} else {
		goto L61
	}
L60:
	;
	v259 = v229
	v261 = v234
	v263 = v224
	goto L9
L61:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v221)+8))
	v227 = F_get_actual_clauses(m, v226)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L20
	} else {
		goto L62
	}
L62:
	;
	v229 = F_list_concat(m, v207, v227)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L20
	} else {
		goto L63
	}
L63:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v222)+60))
	if v231 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v232 = F_lappend(m, v209, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L20
	} else {
		goto L67
	}
L65:
	;
	v234 = v209
	goto L66
L66:
	;
	v236 = v201 + int32(1)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v194)+4))
	if v236 < v237 {
		v201 = v236
		v207 = v229
		v209 = v234
		v211 = v224
		goto L59
	} else {
		goto L68
	}
L67:
	;
	v234 = v232
	goto L66
L68:
	;
	goto L60
L69:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v243
	F_errmsg_internal(m, int32(_a_F_create_bitmap_subplan_0), v19)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L20
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_create_bitmap_subplan_1), int32(3520), int32(_a_F_create_bitmap_subplan_2))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L20
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v330
	v577 = v152
	v578 = v326
	v580 = v261
	goto L1
L73:
	;
	v273 = int32(0)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v270)+4))
	if v274 <= v273 {
		v326 = v259
		v330 = v263
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v277 = v273
	v283 = v259
	v287 = v263
	goto L75
L75:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v293+v277<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v297
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v297
	v303 = F_list_make1_impl(m, int32(1), v19+int32(20))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L20
	} else {
		goto L77
	}
L76:
	;
	v326 = v314
	v330 = v315
	goto L72
L77:
	;
	v306 = F_predicate_implied_by(m, v303, v287, int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L20
	} else {
		goto L78
	}
L78:
	;
	if v306 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v310 = F_lappend(m, v287, v297)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L20
	} else {
		goto L82
	}
L80:
	;
	v314 = v283
	v315 = v287
	goto L81
L81:
	;
	v317 = v277 + int32(1)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v270)+4))
	if v317 < v318 {
		v277 = v317
		v283 = v314
		v287 = v315
		goto L75
	} else {
		goto L84
	}
L82:
	;
	v312 = F_lappend(m, v283, v297)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L20
	} else {
		goto L83
	}
L83:
	;
	v314 = v312
	v315 = v310
	goto L81
L84:
	;
	goto L76
L85:
	;
	v371 = v345
	v372 = v346
	v373 = int32(0)
	v375 = v349
	v376 = v350
	goto L7
L86:
	;
	goto L87
L87:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	if v356 != int32(1) {
		v371 = v345
		v372 = v346
		v373 = v347
		v375 = v349
		v376 = v350
		goto L7
	} else {
		goto L88
	}
L88:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)))
	if v345&int32(1) != 0 {
		v427 = v360
		v431 = v346
		v435 = v350
		goto L6
	} else {
		goto L89
	}
L89:
	;
	v451 = v360
	v455 = v346
	v458 = v349
	v459 = v350
	goto L5
L90:
	;
	v382 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v380)+44)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v380))) = int32(338)
	*(*int32)(unsafe.Add(mBase, uint32(v380)+76)) = v373
	*(*int64)(unsafe.Add(mBase, uint32(v380)+52)) = v382
	v389 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v380)+8)) = v389
	v391 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v380)+16)) = v391
	v393 = *(*float64)(unsafe.Add(mBase, uint32(l1)+80))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v395 = *(*float64)(unsafe.Add(mBase, uint32(v394)+120))
	v396 = base.F64_mul(v393, v395)
	v398 = float64(1e+100)
	if base.F64_gt(v396, v398) != 0 {
		v410 = v398
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v411 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v380)+36)) = uint8(v411)
	*(*int32)(unsafe.Add(mBase, uint32(v380)+32)) = v411
	*(*float64)(unsafe.Add(mBase, uint32(v380)+24)) = v410
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v380)+37)) = uint8(v416)
	if v371&int32(1) == v411 {
		v451 = v380
		v455 = v372
		v458 = v375
		v459 = v376
		goto L5
	} else {
		goto L96
	}
L92:
	;
	goto L91
L93:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v396)&int64(9223372036854775807)) {
		v410 = v398
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v406 = float64(1)
	if base.F64_le(v396, v406) != 0 {
		v410 = v406
		goto L92
	} else {
		goto L95
	}
L95:
	;
	v410 = base.F64_nearest(v396)
	goto L92
L96:
	;
	v427 = v380
	v431 = v372
	v435 = v376
	goto L6
L97:
	;
	v577 = v427
	v578 = v438
	v580 = v438
	goto L1
L98:
	;
	v470 = F_make_orclause(m, v458)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L20
	} else {
		goto L104
	}
L99:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v458)+4))
	if int32(1) < v462 {
		goto L98
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v458
	if v455&int32(1) == int32(0) {
		v487 = v451
		v495 = v459
		goto L4
	} else {
		goto L103
	}
L102:
	;
	goto L101
L103:
	;
	goto L3
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v470
	v477 = F_list_make1_impl(m, int32(1), v19+int32(16))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L20
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v477
	if v455&int32(1) != 0 {
		goto L3
	} else {
		goto L106
	}
L106:
	;
	v487 = v451
	v495 = v459
	goto L4
L107:
	;
	v577 = v487
	v578 = v495
	v580 = v498
	goto L1
L108:
	;
	goto L109
L109:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v495)+4))
	if v501 < int32(2) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v577 = v487
	v578 = v495
	v580 = v498
	goto L1
L111:
	;
	goto L112
L112:
	;
	v504 = F_make_orclause(m, v495)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L20
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v504
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v504
	v511 = F_list_make1_impl(m, int32(1), v19+int32(12))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L20
	} else {
		goto L114
	}
L114:
	;
	v577 = v487
	v578 = v511
	v580 = v498
	goto L1
L115:
	;
	v535 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v533)+44)) = v535
	*(*int32)(unsafe.Add(mBase, uint32(v533))) = int32(337)
	*(*int32)(unsafe.Add(mBase, uint32(v533)+72)) = v527
	*(*int64)(unsafe.Add(mBase, uint32(v533)+52)) = v535
	v542 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v533)+8)) = v542
	v544 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v533)+16)) = v544
	v546 = *(*float64)(unsafe.Add(mBase, uint32(l1)+80))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v548 = *(*float64)(unsafe.Add(mBase, uint32(v547)+120))
	v549 = base.F64_mul(v546, v548)
	v551 = float64(1e+100)
	if base.F64_gt(v549, v551) != 0 {
		v563 = v551
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v564 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v533)+36)) = uint8(v564)
	*(*int32)(unsafe.Add(mBase, uint32(v533)+32)) = v564
	*(*float64)(unsafe.Add(mBase, uint32(v533)+24)) = v563
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v533)+37)) = uint8(v569)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v525
	v577 = v533
	v578 = v522
	v580 = v524
	goto L1
L117:
	;
	goto L116
L118:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v549)&int64(9223372036854775807)) {
		v563 = v551
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v559 = float64(1)
	if base.F64_le(v549, v559) != 0 {
		v563 = v559
		goto L117
	} else {
		goto L120
	}
L120:
	;
	v563 = base.F64_nearest(v549)
	goto L117
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
