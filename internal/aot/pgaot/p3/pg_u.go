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
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v186 int32
	_ = v186
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
	v45 = v18 + int32(_a_F_pg_u_isprint_0)
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
	v45 = l0<<(uint(int32(1))%32) + int32(_a_F_pg_u_isprint_1)
	goto L2
L7:
	;
	v16 = base.I32_div_s(v10+v11, int32(2))
	v18 = v16 * int32(12)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_pg_u_isprint[0])))
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
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_pg_u_isprint[1])))
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
	return v186
L17:
	;
	v186 = base.B2i32(v176&int32(255) == int32(12))
	goto L16
L18:
	;
	if int32(1)<<(uint(v107)%32)&int32(_a_F_pg_u_isprint_2) != 0 {
		goto L36
	} else {
		goto L37
	}
L19:
	;
	if l0 == int32(9) {
		v186 = v86
		goto L16
	} else {
		goto L34
	}
L20:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+uint32(_c_F_pg_u_isprint[2]))))
	v107 = v98
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
	v89 = l0 << (uint(v86) % 32)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+uint32(_c_F_pg_u_isprint[3]))))
	if v86<<(uint(v90)%32)&int32(_a_F_pg_u_isprint_2) == int32(0) {
		goto L19
	} else {
		goto L32
	}
L24:
	;
	v67 = base.I32_div_s(v61+v62, int32(2))
	v69 = v67 * int32(12)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)+uint32(_c_F_pg_u_isprint[0])))
	if base.Ui32(v72) < base.Ui32(l0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v107 = int32(0)
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
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v69)+uint32(_c_F_pg_u_isprint[1])))
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
		v176 = v90
		goto L17
	} else {
		goto L33
	}
L33:
	;
	v186 = v86
	goto L16
L34:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+uint32(_c_F_pg_u_isprint[4]))))
	if v103&int32(32) != 0 {
		v176 = v90
		goto L17
	} else {
		goto L35
	}
L35:
	;
	v186 = v86
	goto L16
L36:
	;
	v149 = int32(3367)
	v150 = int32(0)
	goto L47
L37:
	;
	v115 = int32(10)
	v116 = int32(0)
	goto L38
L38:
	;
	v121 = base.I32_div_s(v115+v116, int32(2))
	v123 = v121 << (uint(int32(3)) % 32)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v123)+uint32(_c_F_pg_u_isprint[5])))
	if base.Ui32(v126) < base.Ui32(l0) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	return int32(1)
L40:
	;
	if v137 <= v136 {
		v115 = v136
		v116 = v137
		goto L38
	} else {
		goto L45
	}
L41:
	;
	v136 = v115
	v137 = v121 + int32(1)
	goto L40
L42:
	;
	goto L43
L43:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v123)+uint32(_c_F_pg_u_isprint[6])))
	if base.Ui32(v132) <= base.Ui32(l0) {
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v136 = v121 - int32(1)
	v137 = v116
	goto L40
L45:
	;
	goto L39
L46:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+uint32(_c_F_pg_u_isprint[2]))))
	v176 = v174
	goto L17
L47:
	;
	v155 = base.I32_div_s(v149+v150, int32(2))
	v157 = v155 * int32(12)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v157)+uint32(_c_F_pg_u_isprint[0])))
	if base.Ui32(v160) < base.Ui32(l0) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v176 = int32(0)
	goto L17
L49:
	;
	if v171 <= v170 {
		v149 = v170
		v150 = v171
		goto L47
	} else {
		goto L54
	}
L50:
	;
	v170 = v149
	v171 = v155 + int32(1)
	goto L49
L51:
	;
	goto L52
L52:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v157)+uint32(_c_F_pg_u_isprint[1])))
	if base.Ui32(v166) <= base.Ui32(l0) {
		goto L46
	} else {
		goto L53
	}
L53:
	;
	v170 = v155 - int32(1)
	v171 = v150
	goto L49
L54:
	;
	goto L48
}
