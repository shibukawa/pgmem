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
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v216 int32
	_ = v216
	var v239 int32
	_ = v239
	var v253 int32
	_ = v253
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wc_isprint[0]))
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
		goto L64
	} else {
		goto L65
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
	return v201
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
	v54 = v27 + int32(_a_F_pg_wc_isprint_0)
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
	v54 = l0<<(uint(int32(1))%32) + int32(_a_F_pg_wc_isprint_1)
	goto L7
L12:
	;
	v25 = base.I32_div_s(v19+v20, int32(2))
	v27 = v25 * int32(12)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_pg_wc_isprint[1])))
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
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_pg_wc_isprint[2])))
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
	v201 = int32(0)
	goto L5
L21:
	;
	v201 = v193
	goto L5
L22:
	;
	v193 = base.B2i32(v183&int32(255) == int32(12))
	goto L21
L23:
	;
	if int32(1)<<(uint(v115)%32)&int32(_a_F_pg_wc_isprint_2) != 0 {
		goto L41
	} else {
		goto L42
	}
L24:
	;
	if l0 == int32(9) {
		v193 = v94
		goto L21
	} else {
		goto L39
	}
L25:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+uint32(_c_F_pg_wc_isprint[3]))))
	v115 = v106
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
	v97 = l0 << (uint(v94) % 32)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+uint32(_c_F_pg_wc_isprint[4]))))
	if v94<<(uint(v98)%32)&int32(_a_F_pg_wc_isprint_2) == int32(0) {
		goto L24
	} else {
		goto L37
	}
L29:
	;
	v75 = base.I32_div_s(v69+v70, int32(2))
	v77 = v75 * int32(12)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)+uint32(_c_F_pg_wc_isprint[1])))
	if base.Ui32(v80) < base.Ui32(l0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v115 = int32(0)
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
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v77)+uint32(_c_F_pg_wc_isprint[2])))
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
		v183 = v98
		goto L22
	} else {
		goto L38
	}
L38:
	;
	v193 = v94
	goto L21
L39:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+uint32(_c_F_pg_wc_isprint[5]))))
	if v111&int32(32) != 0 {
		v183 = v98
		goto L22
	} else {
		goto L40
	}
L40:
	;
	v193 = v94
	goto L21
L41:
	;
	v156 = int32(3367)
	v157 = int32(0)
	goto L52
L42:
	;
	v123 = int32(10)
	v124 = int32(0)
	goto L43
L43:
	;
	v129 = base.I32_div_s(v123+v124, int32(2))
	v131 = v129 << (uint(int32(3)) % 32)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v131)+uint32(_c_F_pg_wc_isprint[6])))
	if base.Ui32(v134) < base.Ui32(l0) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v201 = int32(1)
	goto L5
L45:
	;
	if v145 <= v144 {
		v123 = v144
		v124 = v145
		goto L43
	} else {
		goto L50
	}
L46:
	;
	v144 = v123
	v145 = v129 + int32(1)
	goto L45
L47:
	;
	goto L48
L48:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v131)+uint32(_c_F_pg_wc_isprint[7])))
	if base.Ui32(v140) <= base.Ui32(l0) {
		goto L41
	} else {
		goto L49
	}
L49:
	;
	v144 = v129 - int32(1)
	v145 = v124
	goto L45
L50:
	;
	goto L44
L51:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+uint32(_c_F_pg_wc_isprint[3]))))
	v183 = v181
	goto L22
L52:
	;
	v162 = base.I32_div_s(v156+v157, int32(2))
	v164 = v162 * int32(12)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v164)+uint32(_c_F_pg_wc_isprint[1])))
	if base.Ui32(v167) < base.Ui32(l0) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v183 = int32(0)
	goto L22
L54:
	;
	if v178 <= v177 {
		v156 = v177
		v157 = v178
		goto L52
	} else {
		goto L59
	}
L55:
	;
	v177 = v156
	v178 = v162 + int32(1)
	goto L54
L56:
	;
	goto L57
L57:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v164)+uint32(_c_F_pg_wc_isprint[2])))
	if base.Ui32(v173) <= base.Ui32(l0) {
		goto L51
	} else {
		goto L58
	}
L58:
	;
	v177 = v162 - int32(1)
	v178 = v157
	goto L54
L59:
	;
	goto L53
L60:
	;
	return v239
L61:
	;
	v239 = base.B2i32(base.Ui32(int32(32)) < base.Ui32((l0+int32(1))&int32(127)))
	goto L60
L62:
	;
	goto L63
L63:
	;
	v216 = int32(_a_F_pg_wc_isprint_3)
	v239 = base.B2i32(l0&v216 != v216)&base.B2i32(base.Ui32(l0-int32(_a_F_pg_wc_isprint_4)) < base.Ui32(int32(_a_F_pg_wc_isprint_5))) | (base.B2i32(base.Ui32(l0-int32(_a_F_pg_wc_isprint_6)) < base.Ui32(int32(_a_F_pg_wc_isprint_7))) | base.B2i32(base.Ui32(l0) < base.Ui32(int32(_a_F_pg_wc_isprint_8))) | base.B2i32(base.Ui32(l0-int32(_a_F_pg_wc_isprint_9)) < base.Ui32(int32(_a_F_pg_wc_isprint_10))))
	goto L60
L64:
	;
	goto L67
L65:
	;
	v253 = int32(0)
	goto L66
L66:
	;
	return v253
L67:
	;
	v253 = base.B2i32(base.B2i32(base.Ui32(l0-int32(32)) < base.Ui32(int32(95))) != int32(0))
	goto L66
}
