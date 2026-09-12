package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DatumGetExpandedArray(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v2 == int32(1) {
		v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		if v5 == int32(3) {
			v15 = l0
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+2))
			return v16
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, _consts[9]))
			v11 = F_expand_array(m, l0, v9, int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = v11
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+2))
				return v16
			}
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[9]))
		v11 = F_expand_array(m, l0, v9, int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = v11
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+2))
			return v16
		}
	}
}
func F_datum_write(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v26 int32
	_ = v26
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	v2 = l1
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l2 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L14
	} else {
		goto L74
	}
L2:
	;
	m.G0 = v10 + int32(16)
	return v185 + v187
L3:
	;
	switch l3 - int32(99) {
	case 0:
		v26 = l0
		goto L6
	case 1:
		goto L8
	default:
		goto L7
	case 6:
		goto L9
	}
L4:
	;
	goto L5
L5:
	;
	switch l4 + int32(2) {
	case 0:
		goto L19
	case 1:
		goto L20
	default:
		goto L18
	}
L6:
	;
	switch l4 - int32(1) {
	case 0:
		goto L13
	case 1:
		goto L12
	default:
		goto L10
	case 3:
		goto L11
	}
L7:
	;
	v26 = (l0 + int32(1)) & int32(-2)
	goto L6
L8:
	;
	v26 = (l0 + int32(7)) & int32(-8)
	goto L6
L9:
	;
	v26 = (l0 + int32(3)) & int32(-4)
	goto L6
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v2
	v185 = v26
	v187 = int32(4)
	goto L2
L12:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v26))) = uint16(v2)
	v185 = v26
	v187 = int32(2)
	goto L2
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v26))) = uint8(v2)
	v185 = v26
	v187 = int32(1)
	goto L2
L14:
	;
	return int32(0)
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l4
	F_errmsg_internal(m, int32(475879), v10)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(321683), int32(230), int32(304025))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	switch l3 - int32(99) {
	case 0:
		v182 = l0
		goto L66
	case 1:
		goto L68
	default:
		goto L67
	case 6:
		goto L69
	}
L19:
	;
	if v2&int32(3) == int32(0) {
		v130 = v2
		goto L47
	} else {
		goto L48
	}
L20:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
	if v52 == int32(1) {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v52&int32(1) != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v58 = int32(base.Ui32(v52) >> (uint(int32(1)) % 32))
	if v58 != 0 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	goto L24
L24:
	;
	if l5 == int32(112) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v185 = l0
	v187 = v58
	goto L2
L26:
	;
	v59 = F__emscripten_memcpy_bulkmem(m, l0, v2, v58)
	mBase = m.M
	goto L28
L27:
	;
	goto L28
L28:
	;
	goto L25
L29:
	;
	switch l3 - int32(99) {
	case 0:
		v101 = l0
		goto L37
	case 1:
		goto L39
	default:
		goto L38
	case 6:
		goto L40
	}
L30:
	;
	if v52&int32(2) != 0 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	v67 = int32(base.Ui32(v65) >> (uint(int32(2)) % 32))
	v69 = v67 - int32(3)
	if base.Ui32(int32(127)) < base.Ui32(v69) {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v72 = int32(1)
	v75 = v69<<(uint(v72)%32) | v72
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v75)
	v79 = int32(4)
	v82 = v67 - v79
	if v82 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v185 = l0
	v187 = v69
	goto L2
L34:
	;
	v83 = F__emscripten_memcpy_bulkmem(m, l0+v72, v2+v79, v82)
	mBase = m.M
	goto L36
L35:
	;
	goto L36
L36:
	;
	goto L33
L37:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	v104 = int32(base.Ui32(v102) >> (uint(int32(2)) % 32))
	if v104 != 0 {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	v101 = (l0 + int32(1)) & int32(-2)
	goto L37
L39:
	;
	v101 = (l0 + int32(7)) & int32(-8)
	goto L37
L40:
	;
	v101 = (l0 + int32(3)) & int32(-4)
	goto L37
L41:
	;
	v185 = v101
	v187 = v104
	goto L2
L42:
	;
	v105 = F__emscripten_memcpy_bulkmem(m, v101, v2, v104)
	mBase = m.M
	goto L44
L43:
	;
	goto L44
L44:
	;
	goto L41
L45:
	;
	v165 = v163 + int32(1)
	if v165 != 0 {
		goto L63
	} else {
		goto L64
	}
L46:
	;
	v163 = v155 - v2
	goto L45
L47:
	;
	v134 = v130
	goto L56
L48:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
	if v114 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v163 = int32(0)
	goto L45
L50:
	;
	goto L51
L51:
	;
	v119 = v2
	goto L52
L52:
	;
	v123 = v119 + int32(1)
	if v123&int32(3) == int32(0) {
		v130 = v123
		goto L47
	} else {
		goto L54
	}
L53:
	;
	v155 = v123
	goto L46
L54:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if v128 != 0 {
		v119 = v123
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	v143 = int32(-2139062144)
	if (int32(16843008)-v140|v140)&v143 == v143 {
		v134 = v134 + int32(4)
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v149 = v134
	goto L59
L58:
	;
	goto L57
L59:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v153 != 0 {
		v149 = v149 + int32(1)
		goto L59
	} else {
		goto L61
	}
L60:
	;
	v155 = v149
	goto L46
L61:
	;
	goto L60
L62:
	;
	v185 = l0
	v187 = v165
	goto L2
L63:
	;
	v166 = F__emscripten_memcpy_bulkmem(m, l0, v2, v165)
	mBase = m.M
	goto L65
L64:
	;
	goto L65
L65:
	;
	goto L62
L66:
	;
	if l4 != 0 {
		goto L71
	} else {
		goto L72
	}
L67:
	;
	v182 = (l0 + int32(1)) & int32(-2)
	goto L66
L68:
	;
	v182 = (l0 + int32(7)) & int32(-8)
	goto L66
L69:
	;
	v182 = (l0 + int32(3)) & int32(-4)
	goto L66
L70:
	;
	v185 = v182
	v187 = l4
	goto L2
L71:
	;
	v183 = F__emscripten_memcpy_bulkmem(m, v182, v2, l4)
	mBase = m.M
	goto L73
L72:
	;
	goto L73
L73:
	;
	goto L70
L74:
	;
	F_errmsg_internal(m, int32(396814), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L14
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(486576), int32(2796), int32(344338))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L14
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
