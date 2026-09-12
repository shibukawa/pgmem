package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_wc_ispunct(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v162 int32
	_ = v162
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[632]))
	switch v4 - int32(1) {
	case 0:
		goto L4
	case 1:
		goto L3
	case 2:
		goto L2
	default:
		goto L5
	}
L1:
	;
	return v208
L2:
	;
	if base.Ui32(int32(255)) < base.Ui32(l0) {
		v208 = v2
		goto L1
	} else {
		goto L52
	}
L3:
	;
	if base.Ui32(l0) <= base.Ui32(int32(131071)) {
		goto L49
	} else {
		goto L50
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[639]))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+16)))
	v22 = int32(1)
	if (v19^int32(-1))&v22 != 0 {
		goto L14
	} else {
		goto L15
	}
L5:
	;
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		v208 = v2
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[633]))))
	return int32(base.Ui32(v11)>>(uint(int32(6))%32)) & int32(1)
L7:
	;
	return v162
L8:
	;
	v162 = base.B2i32(v22<<(uint(v151)%32)&int32(1073217536) != int32(0))
	goto L7
L9:
	;
	v135 = int32(1)
	v136 = l0 << (uint(v135) % 32)
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+uint32(_consts[634]))))
	if v139&v135 != 0 {
		goto L45
	} else {
		goto L46
	}
L10:
	;
	v162 = base.B2i32(v22<<(uint(v129)%32)&int32(821559296) != int32(0))
	goto L7
L11:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+uint32(_consts[640]))))
	v129 = v123
	goto L10
L12:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_consts[641]))))
	v129 = v122
	goto L10
L13:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+uint32(_consts[640]))))
	v151 = v117
	goto L8
L14:
	;
	if base.Ui32(l0) < base.Ui32(int32(128)) {
		goto L9
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if base.Ui32(l0) <= base.Ui32(int32(127)) {
		goto L12
	} else {
		goto L36
	}
L17:
	;
	v34 = int32(0)
	v35 = int32(1178)
	goto L18
L18:
	;
	v40 = base.I32_div_s(v34+v35, int32(2))
	v42 = v40 << (uint(int32(3)) % 32)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+uint32(_consts[635])))
	if base.Ui32(v45) < base.Ui32(l0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v62 = int32(0)
	v63 = int32(3367)
	goto L28
L20:
	;
	if v56 <= v57 {
		v34 = v56
		v35 = v57
		goto L18
	} else {
		goto L27
	}
L21:
	;
	v56 = v40 + int32(1)
	v57 = v35
	goto L20
L22:
	;
	goto L23
L23:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v42)+uint32(_consts[636])))
	if base.Ui32(v51) <= base.Ui32(l0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v162 = int32(0)
	goto L7
L25:
	;
	goto L26
L26:
	;
	v56 = v34
	v57 = v40 - int32(1)
	goto L20
L27:
	;
	goto L19
L28:
	;
	v68 = base.I32_div_s(v62+v63, int32(2))
	v70 = v68 * int32(12)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v70)+uint32(_consts[642])))
	if base.Ui32(v73) < base.Ui32(l0) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v151 = int32(0)
	goto L8
L30:
	;
	if v83 <= v84 {
		v62 = v83
		v63 = v84
		goto L28
	} else {
		goto L35
	}
L31:
	;
	v83 = v68 + int32(1)
	v84 = v63
	goto L30
L32:
	;
	goto L33
L33:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v70)+uint32(_consts[643])))
	if base.Ui32(v79) <= base.Ui32(l0) {
		goto L13
	} else {
		goto L34
	}
L34:
	;
	v83 = v62
	v84 = v68 - int32(1)
	goto L30
L35:
	;
	goto L29
L36:
	;
	v92 = int32(0)
	v93 = int32(3367)
	goto L37
L37:
	;
	v98 = base.I32_div_s(v92+v93, int32(2))
	v100 = v98 * int32(12)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v100)+uint32(_consts[642])))
	if base.Ui32(v103) < base.Ui32(l0) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v129 = int32(0)
	goto L10
L39:
	;
	if v113 <= v114 {
		v92 = v113
		v93 = v114
		goto L37
	} else {
		goto L44
	}
L40:
	;
	v113 = v98 + int32(1)
	v114 = v93
	goto L39
L41:
	;
	goto L42
L42:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v100)+uint32(_consts[643])))
	if base.Ui32(v109) <= base.Ui32(l0) {
		goto L11
	} else {
		goto L43
	}
L43:
	;
	v113 = v92
	v114 = v98 - int32(1)
	goto L39
L44:
	;
	goto L38
L45:
	;
	v162 = int32(0)
	goto L7
L46:
	;
	goto L47
L47:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+uint32(_consts[641]))))
	v151 = v145
	goto L8
L48:
	;
	return v190
L49:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(8))%32)))+uint32(_consts[644]))))
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(3))%32))&int32(31)|v178<<(uint(int32(5))%32))+uint32(_consts[644]))))
	v190 = int32(base.Ui32(v184)>>(uint(l0&int32(7))%32)) & int32(1)
	goto L51
L50:
	;
	v190 = int32(0)
	goto L51
L51:
	;
	goto L48
L52:
	;
	if base.Ui32(l0-int32(33)) <= base.Ui32(int32(93)) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v208 = base.B2i32(v205 != int32(0))
	goto L1
L54:
	;
	v202 = F_isalnum(m, l0)
	mBase = m.M
	v205 = base.B2i32(v202 == int32(0))
	goto L56
L55:
	;
	v205 = int32(0)
	goto L56
L56:
	;
	goto L53
}
func F_pg_wc_isspace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v127 int32
	_ = v127
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[632]))
	switch v4 - int32(1) {
	case 0:
		goto L4
	case 1:
		goto L3
	case 2:
		goto L2
	default:
		goto L5
	}
L1:
	;
	return v127
L2:
	;
	if base.Ui32(int32(255)) < base.Ui32(l0) {
		v127 = v2
		goto L1
	} else {
		goto L41
	}
L3:
	;
	if l0 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L4:
	;
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		v127 = v2
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[633]))))
	return int32(base.Ui32(v11) >> (uint(int32(7)) % 32))
L7:
	;
	return v63
L8:
	;
	v23 = int32(0)
	v24 = int32(10)
	goto L11
L9:
	;
	goto L10
L10:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_consts[634]))))
	v63 = int32(base.Ui32(v53&int32(32)) >> (uint(int32(5)) % 32))
	goto L7
L11:
	;
	v29 = base.I32_div_s(v23+v24, int32(2))
	v31 = v29 << (uint(int32(3)) % 32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+uint32(_consts[645])))
	if base.Ui32(v34) < base.Ui32(l0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v63 = int32(0)
	goto L7
L13:
	;
	if v45 <= v46 {
		v23 = v45
		v24 = v46
		goto L11
	} else {
		goto L20
	}
L14:
	;
	v45 = v29 + int32(1)
	v46 = v24
	goto L13
L15:
	;
	goto L16
L16:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v31)+uint32(_consts[646])))
	if base.Ui32(v40) <= base.Ui32(l0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v63 = int32(1)
	goto L7
L18:
	;
	goto L19
L19:
	;
	v45 = v23
	v46 = v29 - int32(1)
	goto L13
L20:
	;
	goto L12
L21:
	;
	return v111
L22:
	;
	v111 = int32(0)
	goto L21
L23:
	;
	goto L24
L24:
	;
	if l0 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v111 = base.B2i32(v104 != int32(0))
	goto L21
L26:
	;
	v77 = int32(4032880)
	goto L29
L27:
	;
	goto L28
L28:
	;
	v87 = int32(4032880)
	goto L38
L29:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	if v79 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	if v79 != 0 {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	if l0 != v79 {
		v77 = v77 + int32(4)
		goto L29
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	goto L30
L34:
	;
	goto L33
L35:
	;
	v85 = v77
	goto L37
L36:
	;
	v85 = int32(0)
	goto L37
L37:
	;
	v104 = v85
	goto L25
L38:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v93 != 0 {
		v87 = v87 + int32(4)
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v94 = int32(4032880)
	v104 = (v87-v94)&int32(-4) + v94
	goto L25
L40:
	;
	goto L39
L41:
	;
	goto L42
L42:
	;
	v127 = base.B2i32(base.B2i32(l0 == int32(32))|base.B2i32(base.Ui32(l0-int32(9)) < base.Ui32(int32(5))) != int32(0))
	goto L1
}
