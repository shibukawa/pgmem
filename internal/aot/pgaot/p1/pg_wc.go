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
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v156 int32
	_ = v156
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wc_ispunct[0]))
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
	return v198
L2:
	;
	if base.Ui32(int32(255)) < base.Ui32(l0) {
		v198 = v2
		goto L1
	} else {
		goto L52
	}
L3:
	;
	if base.Ui32(l0) <= base.Ui32(int32(_a_F_pg_wc_ispunct_0)) {
		goto L49
	} else {
		goto L50
	}
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wc_ispunct[1]))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+16)))
	v20 = int32(1)
	if (v17^int32(-1))&v20 != 0 {
		goto L14
	} else {
		goto L15
	}
L5:
	;
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		v198 = v2
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_pg_wc_ispunct[2]))))
	return int32(base.Ui32(v9)>>(uint(int32(6))%32)) & int32(1)
L7:
	;
	return v156
L8:
	;
	v156 = base.B2i32(v20<<(uint(v145)%32)&int32(1073217536) != int32(0))
	goto L7
L9:
	;
	v131 = int32(1)
	v132 = l0 << (uint(v131) % 32)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+uint32(_c_F_pg_wc_ispunct[3]))))
	if v133&v131 != 0 {
		goto L45
	} else {
		goto L46
	}
L10:
	;
	v156 = base.B2i32(v20<<(uint(v125)%32)&int32(821559296) != int32(0))
	goto L7
L11:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+uint32(_c_F_pg_wc_ispunct[4]))))
	v125 = v119
	goto L10
L12:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_ispunct[5]))))
	v125 = v118
	goto L10
L13:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+uint32(_c_F_pg_wc_ispunct[4]))))
	v145 = v115
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
	v32 = int32(0)
	v33 = int32(1178)
	goto L18
L18:
	;
	v38 = base.I32_div_s(v32+v33, int32(2))
	v40 = v38 << (uint(int32(3)) % 32)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+uint32(_c_F_pg_wc_ispunct[6])))
	if base.Ui32(v43) < base.Ui32(l0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v60 = int32(0)
	v61 = int32(3367)
	goto L28
L20:
	;
	if v54 <= v55 {
		v32 = v54
		v33 = v55
		goto L18
	} else {
		goto L27
	}
L21:
	;
	v54 = v38 + int32(1)
	v55 = v33
	goto L20
L22:
	;
	goto L23
L23:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)+uint32(_c_F_pg_wc_ispunct[7])))
	if base.Ui32(v49) <= base.Ui32(l0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v156 = int32(0)
	goto L7
L25:
	;
	goto L26
L26:
	;
	v54 = v32
	v55 = v38 - int32(1)
	goto L20
L27:
	;
	goto L19
L28:
	;
	v66 = base.I32_div_s(v60+v61, int32(2))
	v68 = v66 * int32(12)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+uint32(_c_F_pg_wc_ispunct[8])))
	if base.Ui32(v71) < base.Ui32(l0) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v145 = int32(0)
	goto L8
L30:
	;
	if v81 <= v82 {
		v60 = v81
		v61 = v82
		goto L28
	} else {
		goto L35
	}
L31:
	;
	v81 = v66 + int32(1)
	v82 = v61
	goto L30
L32:
	;
	goto L33
L33:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+uint32(_c_F_pg_wc_ispunct[9])))
	if base.Ui32(v77) <= base.Ui32(l0) {
		goto L13
	} else {
		goto L34
	}
L34:
	;
	v81 = v60
	v82 = v66 - int32(1)
	goto L30
L35:
	;
	goto L29
L36:
	;
	v90 = int32(0)
	v91 = int32(3367)
	goto L37
L37:
	;
	v96 = base.I32_div_s(v90+v91, int32(2))
	v98 = v96 * int32(12)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v98)+uint32(_c_F_pg_wc_ispunct[8])))
	if base.Ui32(v101) < base.Ui32(l0) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v125 = int32(0)
	goto L10
L39:
	;
	if v111 <= v112 {
		v90 = v111
		v91 = v112
		goto L37
	} else {
		goto L44
	}
L40:
	;
	v111 = v96 + int32(1)
	v112 = v91
	goto L39
L41:
	;
	goto L42
L42:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+uint32(_c_F_pg_wc_ispunct[9])))
	if base.Ui32(v107) <= base.Ui32(l0) {
		goto L11
	} else {
		goto L43
	}
L43:
	;
	v111 = v90
	v112 = v96 - int32(1)
	goto L39
L44:
	;
	goto L38
L45:
	;
	v156 = int32(0)
	goto L7
L46:
	;
	goto L47
L47:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+uint32(_c_F_pg_wc_ispunct[5]))))
	v145 = v139
	goto L8
L48:
	;
	return v180
L49:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(8))%32)))+uint32(_c_F_pg_wc_ispunct[10]))))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(3))%32))&int32(31)|v169<<(uint(int32(5))%32))+uint32(_c_F_pg_wc_ispunct[10]))))
	v180 = int32(base.Ui32(v173)>>(uint(l0&int32(7))%32)) & int32(1)
	goto L51
L50:
	;
	v180 = int32(0)
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
	v198 = base.B2i32(base.B2i32(v193 == int32(0)) != int32(0))
	goto L1
L54:
	;
	v191 = F_isalnum(m, l0)
	mBase = m.M
	v193 = v191
	goto L56
L55:
	;
	v193 = int32(1)
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
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v74 int32
	_ = v74
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wc_isspace[0]))
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
	return v74
L2:
	;
	if base.Ui32(int32(255)) < base.Ui32(l0) {
		v74 = v2
		goto L1
	} else {
		goto L25
	}
L3:
	;
	if l0 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v18 = Fn13991(m, l0, int32(5), int32(32), int32(_a_F_pg_wc_isspace_0), int32(_a_F_pg_wc_isspace_1), int32(10))
	mBase = m.M
	goto L7
L5:
	;
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		v74 = v2
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_pg_wc_isspace[1]))))
	return int32(base.Ui32(v9) >> (uint(int32(7)) % 32))
L7:
	;
	return v18
L8:
	;
	return v58
L9:
	;
	v58 = int32(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	if l0 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v58 = base.B2i32(v51 != int32(0))
	goto L8
L13:
	;
	v31 = int32(_a_F_pg_wc_isspace_2)
	goto L16
L14:
	;
	goto L15
L15:
	;
	v41 = int32(_a_F_pg_wc_isspace_2)
	v42 = F_wcslen(m, v41)
	mBase = m.M
	v51 = v42<<(uint(int32(2))%32) + v41
	goto L12
L16:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v34 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v34 != 0 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	if l0 != v34 {
		v31 = v31 + int32(4)
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	goto L20
L22:
	;
	v40 = v31
	goto L24
L23:
	;
	v40 = int32(0)
	goto L24
L24:
	;
	v51 = v40
	goto L12
L25:
	;
	goto L26
L26:
	;
	v74 = base.B2i32(base.B2i32(l0 == int32(32))|base.B2i32(base.Ui32(l0-int32(9)) < base.Ui32(int32(5))) != int32(0))
	goto L1
}
