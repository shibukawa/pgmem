package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_appendStringInfoStringQuoted(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var __phi55 int32
	_ = __phi55
	var v58 int32
	_ = v58
	var __phi58 int32
	_ = __phi58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = F_strlen(m, l1)
	mBase = m.M
	v16 = base.B2i32(l2 < v12) & base.B2i32(v4 <= l2)
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = F_pg_mbcliplen(m, l1, v12, l2)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v21 = l1
	v22 = v4
	goto L3
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v23 <= v24+int32(1) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	return
L5:
	;
	v19 = F_pnstrdup(m, l1, v17)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v21 = v19
	v22 = v19
	goto L3
L7:
	;
	v44 = int32(39)
	v45 = F___strchrnul(m, v21, v44)
	mBase = m.M
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v47 == v44 {
		goto L14
	} else {
		goto L15
	}
L8:
	;
	F_appendStringInfoChar(m, l0, int32(39))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v31+v24))) = uint8(v33)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v37 = v35 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v41 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v39+v37))) = uint8(v41)
	goto L7
L11:
	;
	goto L7
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v78
	if v16 != 0 {
		goto L28
	} else {
		goto L29
	}
L13:
	;
	if v51 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v51 = v45
	goto L16
L15:
	;
	v51 = int32(0)
	goto L16
L16:
	;
	goto L13
L17:
	;
	v78 = v21
	goto L12
L18:
	;
	goto L19
L19:
	;
	__phi55 = v21
	__phi58 = v51
	v55 = __phi55
	v58 = __phi58
	goto L20
L20:
	;
	F_appendBinaryStringInfoNT(m, l0, v55, v58-v55+int32(1))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L22
	}
L21:
	;
	v78 = v58
	goto L12
L22:
	;
	v68 = int32(39)
	v69 = F___strchrnul(m, v58+int32(1), v68)
	mBase = m.M
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v71 == v68 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v75 != 0 {
		__phi55 = v58
		__phi58 = v75
		v55 = __phi55
		v58 = __phi58
		goto L20
	} else {
		goto L27
	}
L24:
	;
	v75 = v69
	goto L26
L25:
	;
	v75 = int32(0)
	goto L26
L26:
	;
	goto L23
L27:
	;
	goto L21
L28:
	;
	v86 = int32(719751)
	goto L30
L29:
	;
	v86 = int32(717072)
	goto L30
L30:
	;
	F_appendStringInfo(m, l0, v86, v10)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	if v22 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_pfree(m, v22)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	m.G0 = v10 + int32(16)
	return
L35:
	;
	goto L34
}
func F_string_hash(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
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
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
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
	var v135 int32
	_ = v135
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
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
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
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
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
	var v245 int32
	_ = v245
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
	var v267 int32
	_ = v267
	v3 = F_strlen(m, l0)
	mBase = m.M
	v5 = l1 - int32(1)
	if base.Ui32(v3) < base.Ui32(v5) {
		v7 = v3
	} else {
		v7 = v5
	}
	v13 = v7 - int32(1636608432)
	if l0&int32(3) != 0 {
		if base.Ui32(int32(11)) < base.Ui32(v7) {
			v122 = l0
			v123 = v7
			v124 = v13
			v125 = v13
			v126 = v13
			for {
				v128 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
				v129 = v128 + v125
				v130 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
				v132 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
				v133 = v132 + v126
				v135 = int32(4)
				v137 = v130 + v124 - v133 ^ base.I32_rotl(v133, v135)
				v141 = v129 - v137 ^ base.I32_rotl(v137, int32(6))
				v142 = v133 + v129
				v143 = v137 + v142
				v144 = v141 + v143
				v148 = v142 - v141 ^ base.I32_rotl(v141, int32(8))
				v152 = v143 - v148 ^ base.I32_rotl(v148, int32(16))
				v156 = v144 - v152 ^ base.I32_rotl(v152, int32(19))
				v157 = v148 + v144
				v158 = v152 + v157
				v159 = v156 + v158
				v163 = v157 - v156 ^ base.I32_rotl(v156, v135)
				v164 = int32(12)
				v165 = v122 + v164
				v167 = v123 - v164
				if base.Ui32(int32(11)) < base.Ui32(v167) {
					v122 = v165
					v123 = v167
					v124 = v158
					v125 = v159
					v126 = v163
					continue
				} else {
					break
				}
				break
			}
			v170 = v165
			v171 = v167
			v172 = v158
			v173 = v159
			v174 = v163
		} else {
			v170 = l0
			v171 = v7
			v172 = v13
			v173 = v13
			v174 = v13
		}
		switch v171 - int32(1) {
		case 0:
			v233 = v172
			v234 = v173
			v235 = v174
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
			v240 = v233 + v236
			v241 = v234
			v242 = v235
		case 1:
			v226 = v172
			v227 = v173
			v228 = v174
			v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+1)))
			v233 = v229<<(uint(int32(8))%32) + v226
			v234 = v227
			v235 = v228
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
			v240 = v233 + v236
			v241 = v234
			v242 = v235
		case 2:
			v219 = v172
			v220 = v173
			v221 = v174
			v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+2)))
			v226 = v222<<(uint(int32(16))%32) + v219
			v227 = v220
			v228 = v221
			v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+1)))
			v233 = v229<<(uint(int32(8))%32) + v226
			v234 = v227
			v235 = v228
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
			v240 = v233 + v236
			v241 = v234
			v242 = v235
		case 3:
			v213 = v173
			v214 = v174
			v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+3)))
			v219 = v215<<(uint(int32(24))%32) + v172
			v220 = v213
			v221 = v214
			v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+2)))
			v226 = v222<<(uint(int32(16))%32) + v219
			v227 = v220
			v228 = v221
			v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+1)))
			v233 = v229<<(uint(int32(8))%32) + v226
			v234 = v227
			v235 = v228
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
			v240 = v233 + v236
			v241 = v234
			v242 = v235
		case 4:
			v209 = v173
			v210 = v174
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+4)))
			v213 = v209 + v211
			v214 = v210
			v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+3)))
			v219 = v215<<(uint(int32(24))%32) + v172
			v220 = v213
			v221 = v214
			v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+2)))
			v226 = v222<<(uint(int32(16))%32) + v219
			v227 = v220
			v228 = v221
			v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+1)))
			v233 = v229<<(uint(int32(8))%32) + v226
			v234 = v227
			v235 = v228
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
			v240 = v233 + v236
			v241 = v234
			v242 = v235
		case 5:
			v203 = v173
			v204 = v174
			v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+5)))
			v209 = v205<<(uint(int32(8))%32) + v203
			v210 = v204
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+4)))
			v213 = v209 + v211
			v214 = v210
			v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+3)))
			v219 = v215<<(uint(int32(24))%32) + v172
			v220 = v213
			v221 = v214
			v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+2)))
			v226 = v222<<(uint(int32(16))%32) + v219
			v227 = v220
			v228 = v221
			v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+1)))
			v233 = v229<<(uint(int32(8))%32) + v226
			v234 = v227
			v235 = v228
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
			v240 = v233 + v236
			v241 = v234
			v242 = v235
		case 6:
			v197 = v173
			v198 = v174
			v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+6)))
			v203 = v199<<(uint(int32(16))%32) + v197
			v204 = v198
			v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+5)))
			v209 = v205<<(uint(int32(8))%32) + v203
			v210 = v204
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+4)))
			v213 = v209 + v211
			v214 = v210
			v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+3)))
			v219 = v215<<(uint(int32(24))%32) + v172
			v220 = v213
			v221 = v214
			v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+2)))
			v226 = v222<<(uint(int32(16))%32) + v219
			v227 = v220
			v228 = v221
			v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+1)))
			v233 = v229<<(uint(int32(8))%32) + v226
			v234 = v227
			v235 = v228
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
			v240 = v233 + v236
			v241 = v234
			v242 = v235
		case 7:
			v192 = v174
			v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+7)))
			v197 = v193<<(uint(int32(24))%32) + v173
			v198 = v192
			v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+6)))
			v203 = v199<<(uint(int32(16))%32) + v197
			v204 = v198
			v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+5)))
			v209 = v205<<(uint(int32(8))%32) + v203
			v210 = v204
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+4)))
			v213 = v209 + v211
			v214 = v210
			v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+3)))
			v219 = v215<<(uint(int32(24))%32) + v172
			v220 = v213
			v221 = v214
			v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+2)))
			v226 = v222<<(uint(int32(16))%32) + v219
			v227 = v220
			v228 = v221
			v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+1)))
			v233 = v229<<(uint(int32(8))%32) + v226
			v234 = v227
			v235 = v228
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
			v240 = v233 + v236
			v241 = v234
			v242 = v235
		case 8:
			v187 = v174
			v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+8)))
			v192 = v188<<(uint(int32(8))%32) + v187
			v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+7)))
			v197 = v193<<(uint(int32(24))%32) + v173
			v198 = v192
			v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+6)))
			v203 = v199<<(uint(int32(16))%32) + v197
			v204 = v198
			v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+5)))
			v209 = v205<<(uint(int32(8))%32) + v203
			v210 = v204
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+4)))
			v213 = v209 + v211
			v214 = v210
			v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+3)))
			v219 = v215<<(uint(int32(24))%32) + v172
			v220 = v213
			v221 = v214
			v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+2)))
			v226 = v222<<(uint(int32(16))%32) + v219
			v227 = v220
			v228 = v221
			v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+1)))
			v233 = v229<<(uint(int32(8))%32) + v226
			v234 = v227
			v235 = v228
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
			v240 = v233 + v236
			v241 = v234
			v242 = v235
		case 9:
			v182 = v174
			v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+9)))
			v187 = v183<<(uint(int32(16))%32) + v182
			v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+8)))
			v192 = v188<<(uint(int32(8))%32) + v187
			v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+7)))
			v197 = v193<<(uint(int32(24))%32) + v173
			v198 = v192
			v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+6)))
			v203 = v199<<(uint(int32(16))%32) + v197
			v204 = v198
			v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+5)))
			v209 = v205<<(uint(int32(8))%32) + v203
			v210 = v204
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+4)))
			v213 = v209 + v211
			v214 = v210
			v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+3)))
			v219 = v215<<(uint(int32(24))%32) + v172
			v220 = v213
			v221 = v214
			v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+2)))
			v226 = v222<<(uint(int32(16))%32) + v219
			v227 = v220
			v228 = v221
			v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+1)))
			v233 = v229<<(uint(int32(8))%32) + v226
			v234 = v227
			v235 = v228
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
			v240 = v233 + v236
			v241 = v234
			v242 = v235
		case 10:
			v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+10)))
			v182 = v178<<(uint(int32(24))%32) + v174
			v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+9)))
			v187 = v183<<(uint(int32(16))%32) + v182
			v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+8)))
			v192 = v188<<(uint(int32(8))%32) + v187
			v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+7)))
			v197 = v193<<(uint(int32(24))%32) + v173
			v198 = v192
			v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+6)))
			v203 = v199<<(uint(int32(16))%32) + v197
			v204 = v198
			v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+5)))
			v209 = v205<<(uint(int32(8))%32) + v203
			v210 = v204
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+4)))
			v213 = v209 + v211
			v214 = v210
			v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+3)))
			v219 = v215<<(uint(int32(24))%32) + v172
			v220 = v213
			v221 = v214
			v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+2)))
			v226 = v222<<(uint(int32(16))%32) + v219
			v227 = v220
			v228 = v221
			v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+1)))
			v233 = v229<<(uint(int32(8))%32) + v226
			v234 = v227
			v235 = v228
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
			v240 = v233 + v236
			v241 = v234
			v242 = v235
		default:
			v240 = v172
			v241 = v173
			v242 = v174
		}
	} else {
		if base.Ui32(v7) < base.Ui32(int32(12)) {
			v68 = l0
			v69 = v7
			v70 = v13
			v71 = v13
			v72 = v13
		} else {
			v20 = l0
			v21 = v7
			v22 = v13
			v23 = v13
			v24 = v13
			for {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
				v27 = v26 + v23
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
				v31 = v30 + v24
				v33 = int32(4)
				v35 = v28 + v22 - v31 ^ base.I32_rotl(v31, v33)
				v39 = v27 - v35 ^ base.I32_rotl(v35, int32(6))
				v40 = v31 + v27
				v41 = v35 + v40
				v42 = v39 + v41
				v46 = v40 - v39 ^ base.I32_rotl(v39, int32(8))
				v50 = v41 - v46 ^ base.I32_rotl(v46, int32(16))
				v54 = v42 - v50 ^ base.I32_rotl(v50, int32(19))
				v55 = v46 + v42
				v56 = v50 + v55
				v57 = v54 + v56
				v61 = v55 - v54 ^ base.I32_rotl(v54, v33)
				v62 = int32(12)
				v63 = v20 + v62
				v65 = v21 - v62
				if base.Ui32(int32(11)) < base.Ui32(v65) {
					v20 = v63
					v21 = v65
					v22 = v56
					v23 = v57
					v24 = v61
					continue
				} else {
					break
				}
				break
			}
			v68 = v63
			v69 = v65
			v70 = v56
			v71 = v57
			v72 = v61
		}
		switch v69 - int32(1) {
		case 0:
			v119 = v70
			v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
			v240 = v119 + v120
			v241 = v71
			v242 = v72
		case 1:
			v114 = v70
			v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+1)))
			v119 = v115<<(uint(int32(8))%32) + v114
			v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
			v240 = v119 + v120
			v241 = v71
			v242 = v72
		case 2:
			v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+2)))
			v114 = v110<<(uint(int32(16))%32) + v70
			v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+1)))
			v119 = v115<<(uint(int32(8))%32) + v114
			v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
			v240 = v119 + v120
			v241 = v71
			v242 = v72
		case 3:
			v107 = v71
			v108 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
			v240 = v108 + v70
			v241 = v107
			v242 = v72
		case 4:
			v104 = v71
			v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+4)))
			v107 = v104 + v105
			v108 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
			v240 = v108 + v70
			v241 = v107
			v242 = v72
		case 5:
			v99 = v71
			v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+5)))
			v104 = v100<<(uint(int32(8))%32) + v99
			v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+4)))
			v107 = v104 + v105
			v108 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
			v240 = v108 + v70
			v241 = v107
			v242 = v72
		case 6:
			v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+6)))
			v99 = v95<<(uint(int32(16))%32) + v71
			v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+5)))
			v104 = v100<<(uint(int32(8))%32) + v99
			v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+4)))
			v107 = v104 + v105
			v108 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
			v240 = v108 + v70
			v241 = v107
			v242 = v72
		case 7:
			v90 = v72
			v91 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
			v93 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
			v240 = v91 + v70
			v241 = v93 + v71
			v242 = v90
		case 8:
			v85 = v72
			v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+8)))
			v90 = v86<<(uint(int32(8))%32) + v85
			v91 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
			v93 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
			v240 = v91 + v70
			v241 = v93 + v71
			v242 = v90
		case 9:
			v80 = v72
			v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+9)))
			v85 = v81<<(uint(int32(16))%32) + v80
			v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+8)))
			v90 = v86<<(uint(int32(8))%32) + v85
			v91 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
			v93 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
			v240 = v91 + v70
			v241 = v93 + v71
			v242 = v90
		case 10:
			v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+10)))
			v80 = v76<<(uint(int32(24))%32) + v72
			v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+9)))
			v85 = v81<<(uint(int32(16))%32) + v80
			v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+8)))
			v90 = v86<<(uint(int32(8))%32) + v85
			v91 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
			v93 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
			v240 = v91 + v70
			v241 = v93 + v71
			v242 = v90
		default:
			v240 = v70
			v241 = v71
			v242 = v72
		}
	}
	v245 = int32(14)
	v247 = v241 ^ v242 - base.I32_rotl(v241, v245)
	v251 = v247 ^ v240 - base.I32_rotl(v247, int32(11))
	v255 = v251 ^ v241 - base.I32_rotl(v251, int32(25))
	v259 = v255 ^ v247 - base.I32_rotl(v255, int32(16))
	v263 = v259 ^ v251 - base.I32_rotl(v259, int32(4))
	v267 = v263 ^ v255 - base.I32_rotl(v263, v245)
	return v267 ^ v259 - base.I32_rotl(v267, int32(24))
}
func F_transform_string_values_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	if l2 == int32(1) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v9 = F_strlen(m, l1)
		mBase = m.M
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v11 = m.T0[v10].(func(*base.Module, int32, int32, int32) int32)(m, v8, l1, v9)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v16 = F_pg_detoast_datum_packed(m, v11)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
				if v18 == int32(1) {
					v21 = int32(4)
					v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
					if v23&int32(254) == int32(2) {
						v32 = v21
					} else {
						v32 = base.B2i32(v23 == int32(18)) << (uint(v21) % 32)
					}
					if v23 == int32(1) {
						v35 = v21
					} else {
						v35 = v32
					}
					v48 = v35
				} else {
					v36 = int32(1)
					if v18&v36 != 0 {
						v48 = int32(base.Ui32(v18)>>(uint(v36)%32)) - v36
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
						v48 = int32(base.Ui32(v42)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v49 = int32(1)
				if v18&v49 != 0 {
					v53 = v49
				} else {
					v53 = int32(4)
				}
				F_escape_json_with_len(m, v15, v16+v53, v48)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					if v16 != v11 {
						F_pfree(m, v16)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					} else {
						return int32(0)
					}
				}
			}
		}
	} else {
		v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_appendStringInfoString(m, v62, l1)
		mBase = m.M
		v64 = m.ExcPending
		if v64 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	}
}
