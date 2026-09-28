package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_try_partial_nestloop_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int64, l7 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 float64
	_ = v279
	var v280 float64
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	v12 = m.G0
	v14 = v12 - int32(96)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v16 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(96)
	return
L2:
	;
	F_initial_cost_nestloop(m, l0, v14, l5, l6, l2, l3, l7)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L93
	} else {
		goto L94
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+252))
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = v21
	goto L6
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v23 = v22
	goto L6
L6:
	;
	v24 = int32(0)
	if v19 == v24 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v77 == int32(0) {
		goto L1
	} else {
		goto L21
	}
L8:
	;
	v77 = int32(1)
	goto L7
L9:
	;
	goto L10
L10:
	;
	if v23 == int32(0) {
		v70 = v24
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v77 = v70
	goto L7
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v34 < v33 {
		v70 = v24
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v36 = int32(1)
	if v33 <= v36 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v39 = v36
	goto L16
L15:
	;
	v39 = v33
	goto L16
L16:
	;
	v40 = int32(8)
	v45 = int32(0)
	goto L17
L17:
	;
	v52 = v45 << (uint(int32(2)) % 32)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v19+v40+v52)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v23+v40+v52)))
	v59 = v54 & (v56 ^ int32(-1))
	v61 = base.B2i32(v59 == int32(0))
	if v59 != 0 {
		v70 = v61
		goto L11
	} else {
		goto L19
	}
L18:
	;
	v70 = v61
	goto L11
L19:
	;
	v63 = v45 + int32(1)
	if v63 != v39 {
		v45 = v63
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v80 == int32(0) {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+252))
	v86 = int32(0)
	if base.B2i32(v83 == v86)|base.B2i32(v85 == v86) != 0 {
		v131 = v86
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v131 == int32(0) {
		goto L2
	} else {
		goto L36
	}
L24:
	;
	goto L23
L25:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v96 < v97 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v99 = v96
	goto L28
L27:
	;
	v99 = v97
	goto L28
L28:
	;
	if v99 <= int32(1) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v102 = int32(1)
	goto L31
L30:
	;
	v102 = v99
	goto L31
L31:
	;
	v103 = int32(8)
	v108 = int32(0)
	goto L32
L32:
	;
	v115 = v108 << (uint(int32(2)) % 32)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v85+v103+v115)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v83+v103+v115)))
	v120 = v117 & v119
	v122 = base.B2i32(v120 != int32(0))
	if v120 != 0 {
		v131 = v122
		goto L24
	} else {
		goto L34
	}
L33:
	;
	v131 = v122
	goto L24
L34:
	;
	v124 = v108 + int32(1)
	if v124 != v102 {
		v108 = v124
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v137 = int32(1)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v138 == int32(0) {
		v265 = v137
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v271 == int32(0) {
		goto L1
	} else {
		goto L92
	}
L38:
	;
	v271 = v265
	goto L37
L39:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v134)+252))
	v143 = F_bms_overlap(m, v141, v142)
	mBase = m.M
	if v143 == int32(0) {
		v265 = v137
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v146 = int32(0)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	switch v147 - int32(282) {
	case 0, 1:
		goto L41
	default:
		v265 = v146
		goto L38
	case 3:
		goto L51
	case 4:
		goto L50
	case 5:
		goto L49
	case 9:
		goto L48
	case 10:
		goto L47
	case 11:
		goto L45
	case 14:
		goto L44
	case 15:
		goto L43
	case 16:
		goto L42
	case 18, 19, 20:
		goto L46
	}
L41:
	;
	v265 = int32(1)
	goto L38
L42:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v255 = F_path_is_reparameterizable_by_child(m, v254, v134)
	mBase = m.M
	if v255 == int32(0) {
		v265 = v146
		goto L38
	} else {
		goto L91
	}
L43:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v253 = F_path_is_reparameterizable_by_child(m, v252, v134)
	mBase = m.M
	if v253 != 0 {
		goto L41
	} else {
		goto L90
	}
L44:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v251 = F_path_is_reparameterizable_by_child(m, v250, v134)
	mBase = m.M
	if v251 != 0 {
		goto L41
	} else {
		goto L89
	}
L45:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	if v228 == int32(0) {
		goto L41
	} else {
		goto L81
	}
L46:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l3)+80))
	v223 = F_path_is_reparameterizable_by_child(m, v222, v134)
	mBase = m.M
	if v223 == int32(0) {
		v265 = v146
		goto L38
	} else {
		goto L79
	}
L47:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l3)+76))
	if v200 == int32(0) {
		goto L41
	} else {
		goto L71
	}
L48:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	if v196 == int32(0) {
		goto L41
	} else {
		goto L69
	}
L49:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	if v174 == int32(0) {
		goto L41
	} else {
		goto L61
	}
L50:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	if v152 == int32(0) {
		goto L41
	} else {
		goto L53
	}
L51:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v151 = F_path_is_reparameterizable_by_child(m, v150, v134)
	mBase = m.M
	if v151 != 0 {
		goto L41
	} else {
		goto L52
	}
L52:
	;
	v265 = v146
	goto L38
L53:
	;
	v155 = int32(0)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	if v156 <= v155 {
		goto L41
	} else {
		goto L54
	}
L54:
	;
	v159 = v155
	goto L55
L55:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v163+v159<<(uint(int32(2))%32))))
	v168 = F_path_is_reparameterizable_by_child(m, v167, v134)
	mBase = m.M
	if v168 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v271 = int32(0)
	goto L37
L57:
	;
	v170 = v159 + int32(1)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	if v170 < v171 {
		v159 = v170
		goto L55
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	goto L56
L60:
	;
	goto L41
L61:
	;
	v177 = int32(0)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
	if v178 <= v177 {
		goto L41
	} else {
		goto L62
	}
L62:
	;
	v181 = v177
	goto L63
L63:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v174)+12))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v185+v181<<(uint(int32(2))%32))))
	v190 = F_path_is_reparameterizable_by_child(m, v189, v134)
	mBase = m.M
	if v190 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v271 = int32(0)
	goto L37
L65:
	;
	v192 = v181 + int32(1)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
	if v192 < v193 {
		v181 = v192
		goto L63
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	goto L64
L68:
	;
	goto L41
L69:
	;
	v199 = F_path_is_reparameterizable_by_child(m, v196, v134)
	mBase = m.M
	if v199 != 0 {
		goto L41
	} else {
		goto L70
	}
L70:
	;
	v265 = v146
	goto L38
L71:
	;
	v203 = int32(0)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	if v204 <= v203 {
		goto L41
	} else {
		goto L72
	}
L72:
	;
	v207 = v203
	goto L73
L73:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v200)+12))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v211+v207<<(uint(int32(2))%32))))
	v216 = F_path_is_reparameterizable_by_child(m, v215, v134)
	mBase = m.M
	if v216 != 0 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v271 = int32(0)
	goto L37
L75:
	;
	v218 = v207 + int32(1)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	if v218 < v219 {
		v207 = v218
		goto L73
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	goto L74
L78:
	;
	goto L41
L79:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	v227 = F_path_is_reparameterizable_by_child(m, v226, v134)
	mBase = m.M
	if v227 != 0 {
		goto L41
	} else {
		goto L80
	}
L80:
	;
	v265 = v146
	goto L38
L81:
	;
	v231 = int32(0)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	if v232 <= v231 {
		goto L41
	} else {
		goto L82
	}
L82:
	;
	v235 = v231
	goto L83
L83:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v228)+12))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v239+v235<<(uint(int32(2))%32))))
	v244 = F_path_is_reparameterizable_by_child(m, v243, v134)
	mBase = m.M
	if v244 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v271 = int32(0)
	goto L37
L85:
	;
	v246 = v235 + int32(1)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	if v246 < v247 {
		v235 = v246
		goto L83
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	goto L84
L88:
	;
	goto L41
L89:
	;
	v265 = v146
	goto L38
L90:
	;
	v265 = v146
	goto L38
L91:
	;
	goto L41
L92:
	;
	goto L2
L93:
	;
	return
L94:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v279 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
	v280 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
	v281 = F_add_partial_path_precheck(m, l1, v278, v279, v280, l4)
	mBase = m.M
	if v281 == int32(0) {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v286 = F_create_nestloop_path(m, l0, l1, l5, v14, l7, l2, l3, v284, l4, int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L93
	} else {
		goto L96
	}
L96:
	;
	F_add_partial_path(m, l1, v286)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L93
	} else {
		goto L97
	}
L97:
	;
	goto L1
}
