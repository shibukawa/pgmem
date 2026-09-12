package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_u_isprint(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v190 int32
	_ = v190
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	if base.Ui32(int32(128)) <= base.Ui32(l0) {
		goto L21
	} else {
		goto L22
	}
L2:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v46 != int32(15) {
		goto L1
	} else {
		goto L15
	}
L3:
	;
	v45 = v18 + int32(1906584)
	goto L2
L4:
	;
	v10 = int32(3367)
	v11 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	v45 = l0<<(uint(int32(1))%32) + int32(1881968)
	goto L2
L7:
	;
	v16 = base.I32_div_s(v10+v11, int32(2))
	v18 = v16 * int32(12)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[1224])))
	if base.Ui32(v21) < base.Ui32(l0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L1
L9:
	;
	if v32 <= v31 {
		v10 = v31
		v11 = v32
		goto L7
	} else {
		goto L14
	}
L10:
	;
	v31 = v10
	v32 = v16 + int32(1)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[1225])))
	if base.Ui32(v27) <= base.Ui32(l0) {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v31 = v16 - int32(1)
	v32 = v11
	goto L9
L14:
	;
	goto L8
L15:
	;
	return int32(0)
L16:
	;
	return v190
L17:
	;
	v190 = base.B2i32(v180&int32(255) == int32(12))
	goto L16
L18:
	;
	if int32(1)<<(uint(v111)%32)&int32(294913) != 0 {
		goto L36
	} else {
		goto L37
	}
L19:
	;
	if l0 == int32(9) {
		v190 = v86
		goto L16
	} else {
		goto L34
	}
L20:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+uint32(_consts[1226]))))
	v111 = v100
	goto L18
L21:
	;
	v61 = int32(3367)
	v62 = int32(0)
	goto L24
L22:
	;
	goto L23
L23:
	;
	v86 = int32(1)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(v86)%32))+uint32(_consts[1227]))))
	if v86<<(uint(v92)%32)&int32(294913) == int32(0) {
		goto L19
	} else {
		goto L32
	}
L24:
	;
	v67 = base.I32_div_s(v61+v62, int32(2))
	v69 = v67 * int32(12)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)+uint32(_consts[1224])))
	if base.Ui32(v72) < base.Ui32(l0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v111 = int32(0)
	goto L18
L26:
	;
	if v83 <= v82 {
		v61 = v82
		v62 = v83
		goto L24
	} else {
		goto L31
	}
L27:
	;
	v82 = v61
	v83 = v67 + int32(1)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v69)+uint32(_consts[1225])))
	if base.Ui32(v78) <= base.Ui32(l0) {
		goto L20
	} else {
		goto L30
	}
L30:
	;
	v82 = v67 - int32(1)
	v83 = v62
	goto L26
L31:
	;
	goto L25
L32:
	;
	if l0 != int32(9) {
		v180 = v92
		goto L17
	} else {
		goto L33
	}
L33:
	;
	v190 = v86
	goto L16
L34:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_consts[655]))))
	if v107&int32(32) != 0 {
		v180 = v92
		goto L17
	} else {
		goto L35
	}
L35:
	;
	v190 = v86
	goto L16
L36:
	;
	v153 = int32(3367)
	v154 = int32(0)
	goto L47
L37:
	;
	v119 = int32(10)
	v120 = int32(0)
	goto L38
L38:
	;
	v125 = base.I32_div_s(v119+v120, int32(2))
	v127 = v125 << (uint(int32(3)) % 32)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v127)+uint32(_consts[656])))
	if base.Ui32(v130) < base.Ui32(l0) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	return int32(1)
L40:
	;
	if v141 <= v140 {
		v119 = v140
		v120 = v141
		goto L38
	} else {
		goto L45
	}
L41:
	;
	v140 = v119
	v141 = v125 + int32(1)
	goto L40
L42:
	;
	goto L43
L43:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v127)+uint32(_consts[657])))
	if base.Ui32(v136) <= base.Ui32(l0) {
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v140 = v125 - int32(1)
	v141 = v120
	goto L40
L45:
	;
	goto L39
L46:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+uint32(_consts[1226]))))
	v180 = v178
	goto L17
L47:
	;
	v159 = base.I32_div_s(v153+v154, int32(2))
	v161 = v159 * int32(12)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v161)+uint32(_consts[1224])))
	if base.Ui32(v164) < base.Ui32(l0) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v180 = int32(0)
	goto L17
L49:
	;
	if v175 <= v174 {
		v153 = v174
		v154 = v175
		goto L47
	} else {
		goto L54
	}
L50:
	;
	v174 = v153
	v175 = v159 + int32(1)
	goto L49
L51:
	;
	goto L52
L52:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v161)+uint32(_consts[1225])))
	if base.Ui32(v170) <= base.Ui32(l0) {
		goto L46
	} else {
		goto L53
	}
L53:
	;
	v174 = v159 - int32(1)
	v175 = v154
	goto L49
L54:
	;
	goto L48
}
