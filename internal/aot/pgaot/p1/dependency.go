package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_makeDependencyGraphWalker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v112 int32
	_ = v112
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	v3 = int32(0)
	if l0 == v3 {
		v237 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v237
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v13 - int32(110) {
	case 0:
		v237 = v3
		goto L1
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26:
		goto L3
	case 27:
		goto L7
	case 28:
		goto L6
	case 29:
		goto L5
	case 30:
		goto L4
	case 31:
		goto L8
	default:
		goto L9
	}
L3:
	;
	v233 = F_raw_expression_tree_walker_impl(m, l0, int32(484), l1)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L58
	} else {
		goto L71
	}
L4:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v224 == int32(0) {
		goto L3
	} else {
		goto L69
	}
L5:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v217 == int32(0) {
		goto L3
	} else {
		goto L67
	}
L6:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v210 == int32(0) {
		goto L3
	} else {
		goto L65
	}
L7:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v203 == int32(0) {
		goto L3
	} else {
		goto L63
	}
L8:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v196 == int32(0) {
		goto L3
	} else {
		goto L61
	}
L9:
	;
	if v13 != int32(3) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v18 != 0 {
		v237 = v3
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v19 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v124 <= int32(0) {
		v237 = v3
		goto L1
	} else {
		goto L41
	}
L13:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v22 <= int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v25 = int32(0)
	if v25 < v22 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v28 = v22
	goto L17
L16:
	;
	v28 = v25
	goto L17
L17:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v35 = v3
	goto L18
L18:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v29+v35<<(uint(int32(2))%32))))
	if v43 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L12
L20:
	;
	v112 = v35 + int32(1)
	if v112 != v28 {
		v35 = v112
		goto L18
	} else {
		goto L40
	}
L21:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v46 <= int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v49 = int32(0)
	if v49 < v46 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v52 = v46
	goto L25
L24:
	;
	v52 = v49
	goto L25
L25:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v59 = int32(0)
	goto L26
L26:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v54+v59<<(uint(int32(2))%32))))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v74 == int32(0) {
		v93 = v73
		v94 = v74
		goto L29
	} else {
		goto L30
	}
L27:
	;
	return int32(0)
L28:
	;
	if v94-v93 != 0 {
		goto L36
	} else {
		goto L37
	}
L29:
	;
	goto L28
L30:
	;
	if v73 != v74 {
		v93 = v73
		v94 = v74
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v78 = v53
	v79 = v70
	goto L32
L32:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	if v83 == int32(0) {
		v93 = v82
		v94 = v83
		goto L29
	} else {
		goto L34
	}
L33:
	;
	v93 = v82
	v94 = v83
	goto L29
L34:
	;
	v86 = int32(1)
	if v82 == v83 {
		v78 = v78 + v86
		v79 = v79 + v86
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v97 = v59 + int32(1)
	if v52 != v97 {
		v59 = v97
		goto L26
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	goto L27
L39:
	;
	goto L20
L40:
	;
	goto L19
L41:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v130 = int32(0)
	goto L42
L42:
	;
	v142 = v128 + v130*int32(12)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v148 == int32(0) {
		v167 = v147
		v168 = v148
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v237 = int32(0)
	goto L1
L44:
	;
	if v168-v167 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L45:
	;
	goto L44
L46:
	;
	if v147 != v148 {
		v167 = v147
		v168 = v148
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v152 = v127
	v153 = v144
	goto L48
L48:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+1)))
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+1)))
	if v157 == int32(0) {
		v167 = v156
		v168 = v157
		goto L45
	} else {
		goto L50
	}
L49:
	;
	v167 = v156
	v168 = v157
	goto L45
L50:
	;
	v160 = int32(1)
	if v156 == v157 {
		v152 = v152 + v160
		v153 = v153 + v160
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v172 != v130 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	v194 = v130 + int32(1)
	if v194 != v124 {
		v130 = v194
		goto L42
	} else {
		goto L60
	}
L55:
	;
	v175 = v172 * int32(12)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v128+v175)+8))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	v179 = F_bms_add_member(m, v177, v178)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	v188 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v143)+32)) = uint8(v188)
	return int32(0)
L58:
	;
	return int32(0)
L59:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v183+v175)+8)) = v179
	return int32(0)
L60:
	;
	goto L43
L61:
	;
	F_WalkInnerWith(m, l0, v196, l1)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L58
	} else {
		goto L62
	}
L62:
	;
	return int32(0)
L63:
	;
	F_WalkInnerWith(m, l0, v203, l1)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L58
	} else {
		goto L64
	}
L64:
	;
	return int32(0)
L65:
	;
	F_WalkInnerWith(m, l0, v210, l1)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L58
	} else {
		goto L66
	}
L66:
	;
	return int32(0)
L67:
	;
	F_WalkInnerWith(m, l0, v217, l1)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L58
	} else {
		goto L68
	}
L68:
	;
	return int32(0)
L69:
	;
	F_WalkInnerWith(m, l0, v224, l1)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L58
	} else {
		goto L70
	}
L70:
	;
	return int32(0)
L71:
	;
	v237 = v233
	goto L1
}
