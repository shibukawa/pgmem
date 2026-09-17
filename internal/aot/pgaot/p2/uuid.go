package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_uuid_generate_v4(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v3 = m.G0
	v5 = v3 + int32(-64)
	m.G0 = v5
	F_uuid_generate_random(m, v5)
	v9 = v3 + int32(-48)
	F_uuid_unparse(m, v5, v9)
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v16 = F_DirectFunctionCall1Coll(m, int32(3376), int32(0), v9)
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			m.G0 = v5 - int32(-64)
			return v16
		}
	}
}
func F_uuid_hash(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
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
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(16)
	v9 = int32(-1636608416)
	if v2&int32(3) != 0 {
		v118 = v2
		v119 = v3
		v120 = v9
		v121 = v9
		v122 = v9
		for {
			v124 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
			v125 = v124 + v121
			v126 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
			v128 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
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
			v161 = v118 + v160
			v163 = v119 - v160
			if base.Ui32(int32(11)) < base.Ui32(v163) {
				v118 = v161
				v119 = v163
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
			v229 = v154
			v230 = v155
			v231 = v159
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 1:
			v222 = v154
			v223 = v155
			v224 = v159
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 2:
			v215 = v154
			v216 = v155
			v217 = v159
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 3:
			v209 = v155
			v210 = v159
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+3)))
			v215 = v211<<(uint(int32(24))%32) + v154
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 4:
			v205 = v155
			v206 = v159
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+3)))
			v215 = v211<<(uint(int32(24))%32) + v154
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 5:
			v199 = v155
			v200 = v159
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+3)))
			v215 = v211<<(uint(int32(24))%32) + v154
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 6:
			v193 = v155
			v194 = v159
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+3)))
			v215 = v211<<(uint(int32(24))%32) + v154
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 7:
			v188 = v159
			v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+7)))
			v193 = v189<<(uint(int32(24))%32) + v155
			v194 = v188
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+3)))
			v215 = v211<<(uint(int32(24))%32) + v154
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 8:
			v183 = v159
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+7)))
			v193 = v189<<(uint(int32(24))%32) + v155
			v194 = v188
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+3)))
			v215 = v211<<(uint(int32(24))%32) + v154
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 9:
			v178 = v159
			v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+9)))
			v183 = v179<<(uint(int32(16))%32) + v178
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+7)))
			v193 = v189<<(uint(int32(24))%32) + v155
			v194 = v188
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+3)))
			v215 = v211<<(uint(int32(24))%32) + v154
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		case 10:
			v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+10)))
			v178 = v174<<(uint(int32(24))%32) + v159
			v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+9)))
			v183 = v179<<(uint(int32(16))%32) + v178
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+7)))
			v193 = v189<<(uint(int32(24))%32) + v155
			v194 = v188
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+6)))
			v199 = v195<<(uint(int32(16))%32) + v193
			v200 = v194
			v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+5)))
			v205 = v201<<(uint(int32(8))%32) + v199
			v206 = v200
			v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+4)))
			v209 = v205 + v207
			v210 = v206
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+3)))
			v215 = v211<<(uint(int32(24))%32) + v154
			v216 = v209
			v217 = v210
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+2)))
			v222 = v218<<(uint(int32(16))%32) + v215
			v223 = v216
			v224 = v217
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
			v229 = v225<<(uint(int32(8))%32) + v222
			v230 = v223
			v231 = v224
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
			v236 = v229 + v232
			v237 = v230
			v238 = v231
		default:
			v236 = v154
			v237 = v155
			v238 = v159
		}
	} else {
		v16 = v2
		v17 = v3
		v18 = v9
		v19 = v9
		v20 = v9
		for {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
			v23 = v22 + v19
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
			v27 = v26 + v20
			v29 = int32(4)
			v31 = v24 + v18 - v27 ^ base.I32_rotl(v27, v29)
			v35 = v23 - v31 ^ base.I32_rotl(v31, int32(6))
			v36 = v27 + v23
			v37 = v31 + v36
			v38 = v35 + v37
			v42 = v36 - v35 ^ base.I32_rotl(v35, int32(8))
			v46 = v37 - v42 ^ base.I32_rotl(v42, int32(16))
			v50 = v38 - v46 ^ base.I32_rotl(v46, int32(19))
			v51 = v42 + v38
			v52 = v46 + v51
			v53 = v50 + v52
			v57 = v51 - v50 ^ base.I32_rotl(v50, v29)
			v58 = int32(12)
			v59 = v16 + v58
			v61 = v17 - v58
			if base.Ui32(int32(11)) < base.Ui32(v61) {
				v16 = v59
				v17 = v61
				v18 = v52
				v19 = v53
				v20 = v57
				continue
			} else {
				break
			}
			break
		}
		switch v61 - int32(1) {
		case 0:
			v115 = v52
			v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
			v236 = v115 + v116
			v237 = v53
			v238 = v57
		case 1:
			v110 = v52
			v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+1)))
			v115 = v111<<(uint(int32(8))%32) + v110
			v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
			v236 = v115 + v116
			v237 = v53
			v238 = v57
		case 2:
			v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+2)))
			v110 = v106<<(uint(int32(16))%32) + v52
			v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+1)))
			v115 = v111<<(uint(int32(8))%32) + v110
			v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
			v236 = v115 + v116
			v237 = v53
			v238 = v57
		case 3:
			v103 = v53
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
			v236 = v104 + v52
			v237 = v103
			v238 = v57
		case 4:
			v100 = v53
			v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+4)))
			v103 = v100 + v101
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
			v236 = v104 + v52
			v237 = v103
			v238 = v57
		case 5:
			v95 = v53
			v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+5)))
			v100 = v96<<(uint(int32(8))%32) + v95
			v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+4)))
			v103 = v100 + v101
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
			v236 = v104 + v52
			v237 = v103
			v238 = v57
		case 6:
			v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+6)))
			v95 = v91<<(uint(int32(16))%32) + v53
			v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+5)))
			v100 = v96<<(uint(int32(8))%32) + v95
			v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+4)))
			v103 = v100 + v101
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
			v236 = v104 + v52
			v237 = v103
			v238 = v57
		case 7:
			v86 = v57
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
			v236 = v87 + v52
			v237 = v89 + v53
			v238 = v86
		case 8:
			v81 = v57
			v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+8)))
			v86 = v82<<(uint(int32(8))%32) + v81
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
			v236 = v87 + v52
			v237 = v89 + v53
			v238 = v86
		case 9:
			v76 = v57
			v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+9)))
			v81 = v77<<(uint(int32(16))%32) + v76
			v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+8)))
			v86 = v82<<(uint(int32(8))%32) + v81
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
			v236 = v87 + v52
			v237 = v89 + v53
			v238 = v86
		case 10:
			v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+10)))
			v76 = v72<<(uint(int32(24))%32) + v57
			v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+9)))
			v81 = v77<<(uint(int32(16))%32) + v76
			v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+8)))
			v86 = v82<<(uint(int32(8))%32) + v81
			v87 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
			v89 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
			v236 = v87 + v52
			v237 = v89 + v53
			v238 = v86
		default:
			v236 = v52
			v237 = v53
			v238 = v57
		}
	}
	v241 = int32(14)
	v243 = v237 ^ v238 - base.I32_rotl(v237, v241)
	v247 = v243 ^ v236 - base.I32_rotl(v243, int32(11))
	v251 = v247 ^ v237 - base.I32_rotl(v247, int32(25))
	v255 = v251 ^ v243 - base.I32_rotl(v251, int32(16))
	v259 = v255 ^ v247 - base.I32_rotl(v255, int32(4))
	v263 = v259 ^ v251 - base.I32_rotl(v259, v241)
	return v263 ^ v255 - base.I32_rotl(v263, int32(24))
}
func F_uuid_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_palloc(m, int32(16))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v10 = F_pq_getmsgbytes(m, v3, int32(16))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
			*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = v12
			v14 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
			*(*int64)(unsafe.Add(mBase, uint32(v5))) = v14
			return v5
		}
	}
}
func F_uuid_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pq_begintypsend(m, v5)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_appendBinaryStringInfo(m, v5, v7, int32(16))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = v17 << (uint(int32(2)) % 32)
			m.G0 = v5 + int32(16)
			return v16
		}
	}
}
