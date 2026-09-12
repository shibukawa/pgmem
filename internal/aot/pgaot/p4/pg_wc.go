package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_wc_isprint(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v219 int32
	_ = v219
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v255 int32
	_ = v255
	v3 = *(*int32)(unsafe.Add(mBase, _consts[808]))
	switch v3 - int32(1) {
	case 0:
		goto L3
	case 1:
		goto L2
	case 2:
		goto L1
	default:
		goto L4
	}
L1:
	;
	if base.Ui32(l0) <= base.Ui32(int32(255)) {
		goto L68
	} else {
		goto L69
	}
L2:
	;
	if base.Ui32(l0) <= base.Ui32(int32(254)) {
		goto L61
	} else {
		goto L62
	}
L3:
	;
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	return base.B2i32(base.Ui32(l0-int32(32)) < base.Ui32(int32(95)))
L5:
	;
	return v205
L6:
	;
	if base.Ui32(int32(128)) <= base.Ui32(l0) {
		goto L26
	} else {
		goto L27
	}
L7:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v55 != int32(15) {
		goto L6
	} else {
		goto L20
	}
L8:
	;
	v54 = v27 + int32(1906584)
	goto L7
L9:
	;
	v19 = int32(3367)
	v20 = int32(0)
	goto L12
L10:
	;
	goto L11
L11:
	;
	v54 = l0<<(uint(int32(1))%32) + int32(1881968)
	goto L7
L12:
	;
	v25 = base.I32_div_s(v19+v20, int32(2))
	v27 = v25 * int32(12)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[809])))
	if base.Ui32(v30) < base.Ui32(l0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L6
L14:
	;
	if v41 <= v40 {
		v19 = v40
		v20 = v41
		goto L12
	} else {
		goto L19
	}
L15:
	;
	v40 = v19
	v41 = v25 + int32(1)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[810])))
	if base.Ui32(v36) <= base.Ui32(l0) {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v40 = v25 - int32(1)
	v41 = v20
	goto L14
L19:
	;
	goto L13
L20:
	;
	v205 = int32(0)
	goto L5
L21:
	;
	v205 = v197
	goto L5
L22:
	;
	v197 = base.B2i32(v187&int32(255) == int32(12))
	goto L21
L23:
	;
	if int32(1)<<(uint(v119)%32)&int32(294913) != 0 {
		goto L41
	} else {
		goto L42
	}
L24:
	;
	if l0 == int32(9) {
		v197 = v94
		goto L21
	} else {
		goto L39
	}
L25:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+uint32(_consts[811]))))
	v119 = v108
	goto L23
L26:
	;
	v69 = int32(3367)
	v70 = int32(0)
	goto L29
L27:
	;
	goto L28
L28:
	;
	v94 = int32(1)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(v94)%32))+uint32(_consts[812]))))
	if v94<<(uint(v100)%32)&int32(294913) == int32(0) {
		goto L24
	} else {
		goto L37
	}
L29:
	;
	v75 = base.I32_div_s(v69+v70, int32(2))
	v77 = v75 * int32(12)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)+uint32(_consts[809])))
	if base.Ui32(v80) < base.Ui32(l0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v119 = int32(0)
	goto L23
L31:
	;
	if v91 <= v90 {
		v69 = v90
		v70 = v91
		goto L29
	} else {
		goto L36
	}
L32:
	;
	v90 = v69
	v91 = v75 + int32(1)
	goto L31
L33:
	;
	goto L34
L34:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v77)+uint32(_consts[810])))
	if base.Ui32(v86) <= base.Ui32(l0) {
		goto L25
	} else {
		goto L35
	}
L35:
	;
	v90 = v75 - int32(1)
	v91 = v70
	goto L31
L36:
	;
	goto L30
L37:
	;
	if l0 != int32(9) {
		v187 = v100
		goto L22
	} else {
		goto L38
	}
L38:
	;
	v197 = v94
	goto L21
L39:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_consts[813]))))
	if v115&int32(32) != 0 {
		v187 = v100
		goto L22
	} else {
		goto L40
	}
L40:
	;
	v197 = v94
	goto L21
L41:
	;
	v160 = int32(3367)
	v161 = int32(0)
	goto L52
L42:
	;
	v127 = int32(10)
	v128 = int32(0)
	goto L43
L43:
	;
	v133 = base.I32_div_s(v127+v128, int32(2))
	v135 = v133 << (uint(int32(3)) % 32)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v135)+uint32(_consts[814])))
	if base.Ui32(v138) < base.Ui32(l0) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v205 = int32(1)
	goto L5
L45:
	;
	if v149 <= v148 {
		v127 = v148
		v128 = v149
		goto L43
	} else {
		goto L50
	}
L46:
	;
	v148 = v127
	v149 = v133 + int32(1)
	goto L45
L47:
	;
	goto L48
L48:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v135)+uint32(_consts[815])))
	if base.Ui32(v144) <= base.Ui32(l0) {
		goto L41
	} else {
		goto L49
	}
L49:
	;
	v148 = v133 - int32(1)
	v149 = v128
	goto L45
L50:
	;
	goto L44
L51:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+uint32(_consts[811]))))
	v187 = v185
	goto L22
L52:
	;
	v166 = base.I32_div_s(v160+v161, int32(2))
	v168 = v166 * int32(12)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v168)+uint32(_consts[809])))
	if base.Ui32(v171) < base.Ui32(l0) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v187 = int32(0)
	goto L22
L54:
	;
	if v182 <= v181 {
		v160 = v181
		v161 = v182
		goto L52
	} else {
		goto L59
	}
L55:
	;
	v181 = v160
	v182 = v166 + int32(1)
	goto L54
L56:
	;
	goto L57
L57:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v168)+uint32(_consts[810])))
	if base.Ui32(v177) <= base.Ui32(l0) {
		goto L51
	} else {
		goto L58
	}
L58:
	;
	v181 = v166 - int32(1)
	v182 = v161
	goto L54
L59:
	;
	goto L53
L60:
	;
	return v241
L61:
	;
	v241 = base.B2i32(base.Ui32(int32(32)) < base.Ui32((l0+int32(1))&int32(127)))
	goto L60
L62:
	;
	goto L63
L63:
	;
	v219 = int32(1)
	if base.Ui32(l0-int32(57344)) < base.Ui32(int32(8185)) {
		v239 = v219
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v241 = v239
	goto L60
L65:
	;
	if base.Ui32(l0) < base.Ui32(int32(8232)) {
		v239 = v219
		goto L64
	} else {
		goto L66
	}
L66:
	;
	if base.Ui32(l0-int32(8234)) < base.Ui32(int32(47062)) {
		v239 = v219
		goto L64
	} else {
		goto L67
	}
L67:
	;
	v230 = int32(65534)
	v239 = base.B2i32(l0&v230 != v230) & base.B2i32(base.Ui32(l0-int32(65532)) < base.Ui32(int32(1048580)))
	goto L64
L68:
	;
	goto L71
L69:
	;
	v255 = int32(0)
	goto L70
L70:
	;
	return v255
L71:
	;
	v255 = base.B2i32(base.B2i32(base.Ui32(l0-int32(32)) < base.Ui32(int32(95))) != int32(0))
	goto L70
}
