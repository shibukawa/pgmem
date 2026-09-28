package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_wc_isalnum_builtin(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	v7 = (v3 ^ int32(-1)) & int32(1)
	if base.Ui32(int32(128)) <= base.Ui32(l0) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	return v117
L2:
	;
	v117 = v110
	goto L1
L3:
	;
	v110 = base.B2i32(v99&int32(255) == int32(9))
	goto L2
L4:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+uint32(_c_F_wc_isalnum_builtin[0]))))
	v99 = v92
	goto L3
L5:
	;
	v117 = base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10)))
	goto L1
L6:
	;
	v17 = int32(1201)
	v18 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v72 = int32(1)
	v74 = l0 << (uint(v72) % 32)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+uint32(_c_F_wc_isalnum_builtin[1]))))
	if v75&v72 != 0 {
		v110 = v72
		goto L2
	} else {
		goto L29
	}
L9:
	;
	v23 = base.I32_div_s(v17+v18, int32(2))
	v25 = v23 << (uint(int32(3)) % 32)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_wc_isalnum_builtin[2])))
	if base.Ui32(v28) < base.Ui32(l0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	if v7 != 0 {
		goto L5
	} else {
		goto L19
	}
L11:
	;
	if v40 <= v39 {
		v17 = v39
		v18 = v40
		goto L9
	} else {
		goto L18
	}
L12:
	;
	v39 = v17
	v40 = v23 + int32(1)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_wc_isalnum_builtin[3])))
	if base.Ui32(v34) <= base.Ui32(l0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v117 = int32(1)
	goto L1
L16:
	;
	goto L17
L17:
	;
	v39 = v23 - int32(1)
	v40 = v18
	goto L11
L18:
	;
	goto L10
L19:
	;
	v46 = int32(3408)
	v47 = int32(0)
	goto L21
L20:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+uint32(_c_F_wc_isalnum_builtin[4]))))
	v99 = v71
	goto L3
L21:
	;
	v52 = base.I32_div_s(v46+v47, int32(2))
	v54 = v52 * int32(12)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)+uint32(_c_F_wc_isalnum_builtin[5])))
	if base.Ui32(v57) < base.Ui32(l0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v99 = int32(0)
	goto L3
L23:
	;
	if v68 <= v67 {
		v46 = v67
		v47 = v68
		goto L21
	} else {
		goto L28
	}
L24:
	;
	v67 = v46
	v68 = v52 + int32(1)
	goto L23
L25:
	;
	goto L26
L26:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v54)+uint32(_c_F_wc_isalnum_builtin[6])))
	if base.Ui32(v63) <= base.Ui32(l0) {
		goto L20
	} else {
		goto L27
	}
L27:
	;
	v67 = v52 - int32(1)
	v68 = v47
	goto L23
L28:
	;
	goto L22
L29:
	;
	if v7 == int32(0) {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	goto L5
}
func F_wc_iscased_libc_utf8(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	v4 = F_towlower(m, l0)
	if v4 != l0 {
		v12 = int32(1)
	} else {
		v8 = F_towupper(m, l0)
		v12 = base.B2i32(base.B2i32(v8 != l0) != int32(0))
	}
	return v12
}
func F_wc_isprint_builtin(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return v192
L2:
	;
	if base.Ui32(int32(128)) <= base.Ui32(l0) {
		goto L22
	} else {
		goto L23
	}
L3:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v46 != int32(15) {
		goto L2
	} else {
		goto L16
	}
L4:
	;
	v45 = v18 + int32(_a_F_wc_isprint_builtin_0)
	goto L3
L5:
	;
	v10 = int32(3408)
	v11 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	v45 = l0<<(uint(int32(1))%32) + int32(_a_F_wc_isprint_builtin_1)
	goto L3
L8:
	;
	v16 = base.I32_div_s(v10+v11, int32(2))
	v18 = v16 * int32(12)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_wc_isprint_builtin[0])))
	if base.Ui32(v21) < base.Ui32(l0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L2
L10:
	;
	if v32 <= v31 {
		v10 = v31
		v11 = v32
		goto L8
	} else {
		goto L15
	}
L11:
	;
	v31 = v10
	v32 = v16 + int32(1)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_wc_isprint_builtin[1])))
	if base.Ui32(v27) <= base.Ui32(l0) {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v31 = v16 - int32(1)
	v32 = v11
	goto L10
L15:
	;
	goto L9
L16:
	;
	v192 = int32(0)
	goto L1
L17:
	;
	v192 = v184
	goto L1
L18:
	;
	v184 = base.B2i32(v174&int32(255) == int32(12))
	goto L17
L19:
	;
	if int32(1)<<(uint(v106)%32)&int32(_a_F_wc_isprint_builtin_2) != 0 {
		goto L37
	} else {
		goto L38
	}
L20:
	;
	if l0 == int32(9) {
		v184 = v85
		goto L17
	} else {
		goto L35
	}
L21:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+uint32(_c_F_wc_isprint_builtin[2]))))
	v106 = v97
	goto L19
L22:
	;
	v60 = int32(3408)
	v61 = int32(0)
	goto L25
L23:
	;
	goto L24
L24:
	;
	v85 = int32(1)
	v88 = l0 << (uint(v85) % 32)
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_wc_isprint_builtin[3]))))
	if v85<<(uint(v89)%32)&int32(_a_F_wc_isprint_builtin_2) == int32(0) {
		goto L20
	} else {
		goto L33
	}
L25:
	;
	v66 = base.I32_div_s(v60+v61, int32(2))
	v68 = v66 * int32(12)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+uint32(_c_F_wc_isprint_builtin[0])))
	if base.Ui32(v71) < base.Ui32(l0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v106 = int32(0)
	goto L19
L27:
	;
	if v82 <= v81 {
		v60 = v81
		v61 = v82
		goto L25
	} else {
		goto L32
	}
L28:
	;
	v81 = v60
	v82 = v66 + int32(1)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+uint32(_c_F_wc_isprint_builtin[1])))
	if base.Ui32(v77) <= base.Ui32(l0) {
		goto L21
	} else {
		goto L31
	}
L31:
	;
	v81 = v66 - int32(1)
	v82 = v61
	goto L27
L32:
	;
	goto L26
L33:
	;
	if l0 != int32(9) {
		v174 = v89
		goto L18
	} else {
		goto L34
	}
L34:
	;
	v184 = v85
	goto L17
L35:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_wc_isprint_builtin[4]))))
	if v102&int32(32) != 0 {
		v174 = v89
		goto L18
	} else {
		goto L36
	}
L36:
	;
	v184 = v85
	goto L17
L37:
	;
	v147 = int32(3408)
	v148 = int32(0)
	goto L48
L38:
	;
	v114 = int32(10)
	v115 = int32(0)
	goto L39
L39:
	;
	v120 = base.I32_div_s(v114+v115, int32(2))
	v122 = v120 << (uint(int32(3)) % 32)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v122)+uint32(_c_F_wc_isprint_builtin[5])))
	if base.Ui32(v125) < base.Ui32(l0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v192 = int32(1)
	goto L1
L41:
	;
	if v136 <= v135 {
		v114 = v135
		v115 = v136
		goto L39
	} else {
		goto L46
	}
L42:
	;
	v135 = v114
	v136 = v120 + int32(1)
	goto L41
L43:
	;
	goto L44
L44:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v122)+uint32(_c_F_wc_isprint_builtin[6])))
	if base.Ui32(v131) <= base.Ui32(l0) {
		goto L37
	} else {
		goto L45
	}
L45:
	;
	v135 = v120 - int32(1)
	v136 = v115
	goto L41
L46:
	;
	goto L40
L47:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+uint32(_c_F_wc_isprint_builtin[2]))))
	v174 = v172
	goto L18
L48:
	;
	v153 = base.I32_div_s(v147+v148, int32(2))
	v155 = v153 * int32(12)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v155)+uint32(_c_F_wc_isprint_builtin[0])))
	if base.Ui32(v158) < base.Ui32(l0) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v174 = int32(0)
	goto L18
L50:
	;
	if v169 <= v168 {
		v147 = v168
		v148 = v169
		goto L48
	} else {
		goto L55
	}
L51:
	;
	v168 = v147
	v169 = v153 + int32(1)
	goto L50
L52:
	;
	goto L53
L53:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v155)+uint32(_c_F_wc_isprint_builtin[1])))
	if base.Ui32(v164) <= base.Ui32(l0) {
		goto L47
	} else {
		goto L54
	}
L54:
	;
	v168 = v153 - int32(1)
	v169 = v148
	goto L50
L55:
	;
	goto L49
}
func F_wc_ispunct_libc_mb(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	if base.Ui32(l0) <= base.Ui32(int32(_a_F_wc_ispunct_libc_mb_0)) {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(8))%32)))+uint32(_c_F_wc_ispunct_libc_mb[0]))))
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(3))%32))&int32(31)|v12<<(uint(int32(5))%32))+uint32(_c_F_wc_ispunct_libc_mb[0]))))
		v23 = int32(base.Ui32(v16)>>(uint(l0&int32(7))%32)) & int32(1)
	} else {
		v23 = int32(0)
	}
	return base.B2i32(v23 != int32(0))
}
func F_wc_isxdigit_builtin(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if (v6^int32(-1))&int32(1) != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v122
L2:
	;
	v114 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_isxdigit_builtin[0]))))
	v122 = base.B2i32(v114 < int32(0))
	goto L1
L3:
	;
	v122 = v111
	goto L1
L4:
	;
	if base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(l0-int32(65)) < base.Ui32(int32(6))) != 0 {
		v111 = int32(1)
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v122 = base.B2i32(base.Ui32(l0-int32(97)) < base.Ui32(int32(6)))
	goto L1
L8:
	;
	if base.Ui32(l0) <= base.Ui32(int32(127)) {
		goto L2
	} else {
		goto L23
	}
L9:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v66 != int32(9) {
		goto L8
	} else {
		goto L22
	}
L10:
	;
	v65 = v38 + int32(_a_F_wc_isxdigit_builtin_0)
	goto L9
L11:
	;
	v30 = int32(0)
	v31 = int32(3408)
	goto L14
L12:
	;
	goto L13
L13:
	;
	v65 = l0<<(uint(int32(1))%32) + int32(_a_F_wc_isxdigit_builtin_1)
	goto L9
L14:
	;
	v36 = base.I32_div_s(v30+v31, int32(2))
	v38 = v36 * int32(12)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_wc_isxdigit_builtin[1])))
	if base.Ui32(v41) < base.Ui32(l0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L8
L16:
	;
	if v51 <= v52 {
		v30 = v51
		v31 = v52
		goto L14
	} else {
		goto L21
	}
L17:
	;
	v51 = v36 + int32(1)
	v52 = v31
	goto L16
L18:
	;
	goto L19
L19:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_wc_isxdigit_builtin[2])))
	if base.Ui32(v47) <= base.Ui32(l0) {
		goto L10
	} else {
		goto L20
	}
L20:
	;
	v51 = v30
	v52 = v36 - int32(1)
	goto L16
L21:
	;
	goto L15
L22:
	;
	v122 = int32(1)
	goto L1
L23:
	;
	v80 = int32(0)
	v81 = int32(5)
	goto L24
L24:
	;
	v86 = base.I32_div_s(v80+v81, int32(2))
	v88 = v86 << (uint(int32(3)) % 32)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_wc_isxdigit_builtin[3])))
	if base.Ui32(v91) < base.Ui32(l0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v111 = int32(0)
	goto L3
L26:
	;
	if v102 <= v103 {
		v80 = v102
		v81 = v103
		goto L24
	} else {
		goto L31
	}
L27:
	;
	v102 = v86 + int32(1)
	v103 = v81
	goto L26
L28:
	;
	goto L29
L29:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_wc_isxdigit_builtin[4])))
	if base.Ui32(v98) <= base.Ui32(l0) {
		v122 = int32(1)
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v102 = v80
	v103 = v86 - int32(1)
	goto L26
L31:
	;
	goto L25
}
func F_wc_toupper_builtin(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	if base.Ui32(l0) <= base.Ui32(int32(127)) {
		v146 = l0<<(uint(int32(2))%32) + int32(_a_F_wc_toupper_builtin_0)
	} else {
		v9 = int32(0)
		if base.Ui32(l0) <= base.Ui32(int32(1415)) {
			v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_toupper_builtin[0]))))
			v141 = v14
		} else {
			if base.Ui32(l0) <= base.Ui32(int32(_a_F_wc_toupper_builtin_1)) {
				if base.Ui32(l0) <= base.Ui32(int32(_a_F_wc_toupper_builtin_2)) {
					if base.Ui32(l0-int32(_a_F_wc_toupper_builtin_3)) <= base.Ui32(int32(95)) {
						v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_toupper_builtin[1]))))
						v141 = v27
					} else {
						if base.Ui32(l0) < base.Ui32(int32(_a_F_wc_toupper_builtin_4)) {
							v141 = v9
						} else {
							if base.Ui32(l0) <= base.Ui32(int32(_a_F_wc_toupper_builtin_5)) {
								v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_toupper_builtin[2]))))
								v141 = v36
							} else {
								if base.Ui32(l0) < base.Ui32(int32(_a_F_wc_toupper_builtin_6)) {
									v141 = v9
								} else {
									v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_toupper_builtin[3]))))
									v141 = v43
								}
							}
						}
					}
				} else {
					if base.Ui32(l0) < base.Ui32(int32(_a_F_wc_toupper_builtin_7)) {
						v141 = v9
					} else {
						if base.Ui32(l0) <= base.Ui32(int32(_a_F_wc_toupper_builtin_8)) {
							if base.Ui32(l0) <= base.Ui32(int32(_a_F_wc_toupper_builtin_9)) {
								v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_toupper_builtin[4]))))
								v141 = v54
							} else {
								if base.Ui32(l0) < base.Ui32(int32(_a_F_wc_toupper_builtin_10)) {
									v141 = v9
								} else {
									v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_toupper_builtin[5]))))
									v141 = v61
								}
							}
						} else {
							if base.Ui32(l0) < base.Ui32(int32(_a_F_wc_toupper_builtin_11)) {
								v141 = v9
							} else {
								if base.Ui32(l0) <= base.Ui32(int32(_a_F_wc_toupper_builtin_12)) {
									v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_toupper_builtin[6]))))
									v141 = v70
								} else {
									if base.Ui32(l0) < base.Ui32(int32(_a_F_wc_toupper_builtin_13)) {
										v141 = v9
									} else {
										v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_toupper_builtin[7]))))
										v141 = v77
									}
								}
							}
						}
					}
				}
			} else {
				if base.Ui32(l0) < base.Ui32(int32(_a_F_wc_toupper_builtin_14)) {
					v141 = v9
				} else {
					if base.Ui32(l0) <= base.Ui32(int32(_a_F_wc_toupper_builtin_15)) {
						if base.Ui32(l0) <= base.Ui32(int32(_a_F_wc_toupper_builtin_16)) {
							if base.Ui32(l0) <= base.Ui32(int32(_a_F_wc_toupper_builtin_17)) {
								v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_toupper_builtin[8]))))
								v141 = v90
							} else {
								if base.Ui32(l0) < base.Ui32(int32(_a_F_wc_toupper_builtin_18)) {
									v141 = v9
								} else {
									v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_toupper_builtin[9]))))
									v141 = v97
								}
							}
						} else {
							if base.Ui32(l0) < base.Ui32(int32(_a_F_wc_toupper_builtin_19)) {
								v141 = v9
							} else {
								if base.Ui32(l0) <= base.Ui32(int32(_a_F_wc_toupper_builtin_20)) {
									v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_toupper_builtin[10]))))
									v141 = v106
								} else {
									if base.Ui32(l0) < base.Ui32(int32(_a_F_wc_toupper_builtin_21)) {
										v141 = v9
									} else {
										v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_toupper_builtin[11]))))
										v141 = v113
									}
								}
							}
						}
					} else {
						if base.Ui32(l0) < base.Ui32(int32(_a_F_wc_toupper_builtin_22)) {
							v141 = v9
						} else {
							if base.Ui32(l0) <= base.Ui32(int32(_a_F_wc_toupper_builtin_23)) {
								if base.Ui32(l0) <= base.Ui32(int32(_a_F_wc_toupper_builtin_24)) {
									v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_toupper_builtin[12]))))
									v141 = v124
								} else {
									if base.Ui32(l0) < base.Ui32(int32(_a_F_wc_toupper_builtin_25)) {
										v141 = v9
									} else {
										v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_toupper_builtin[13]))))
										v141 = v131
									}
								}
							} else {
								if base.Ui32(int32(67)) < base.Ui32(l0-int32(_a_F_wc_toupper_builtin_26)) {
									v141 = v9
								} else {
									v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_toupper_builtin[14]))))
									v141 = v140
								}
							}
						}
					}
				}
			}
		}
		v146 = v141<<(uint(int32(2))%32) + int32(_a_F_wc_toupper_builtin_27)
	}
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	if v147 != 0 {
		v148 = v147
	} else {
		v148 = l0
	}
	return v148
}
