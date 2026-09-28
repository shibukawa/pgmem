package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_wc_isalpha_builtin(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v46
L2:
	;
	v10 = int32(1201)
	v11 = int32(0)
	goto L5
L3:
	;
	goto L4
L4:
	;
	v36 = int32(1)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(v36)%32))+uint32(_c_F_wc_isalpha_builtin[0]))))
	v46 = v38 & v36
	goto L1
L5:
	;
	v16 = base.I32_div_s(v10+v11, int32(2))
	v18 = v16 << (uint(int32(3)) % 32)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_wc_isalpha_builtin[1])))
	if base.Ui32(v21) < base.Ui32(l0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v46 = int32(0)
	goto L1
L7:
	;
	if v33 <= v32 {
		v10 = v32
		v11 = v33
		goto L5
	} else {
		goto L12
	}
L8:
	;
	v32 = v10
	v33 = v16 + int32(1)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_wc_isalpha_builtin[2])))
	if base.Ui32(v28) <= base.Ui32(l0) {
		v46 = int32(1)
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v32 = v16 - int32(1)
	v33 = v11
	goto L7
L12:
	;
	goto L6
}
func F_wc_isdigit_libc_sb(m *base.Module, l0 int32, l1 int32) int32 {
	var v13 int32
	_ = v13
	if base.Ui32(l0) <= base.Ui32(int32(255)) {
		v13 = base.B2i32(base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10))) != int32(0))
	} else {
		v13 = int32(0)
	}
	return v13
}
func F_wc_isgraph_builtin(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v88 int32
	_ = v88
	var v102 int32
	_ = v102
	if base.Ui32(int32(128)) <= base.Ui32(l0) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	return v102
L2:
	;
	v102 = int32(0)
	goto L1
L3:
	;
	v102 = v88
	goto L1
L4:
	;
	v54 = int32(0)
	if int32(1)<<(uint(v53)%32)&int32(_a_F_wc_isgraph_builtin_0) != 0 {
		v88 = v54
		goto L3
	} else {
		goto L19
	}
L5:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_wc_isgraph_builtin[0]))))
	v53 = v50
	goto L4
L6:
	;
	v10 = int32(3408)
	v11 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v35 = int32(1)
	v38 = l0 << (uint(v35) % 32)
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_wc_isgraph_builtin[1]))))
	if v35<<(uint(v39)%32)&int32(_a_F_wc_isgraph_builtin_0) != 0 {
		goto L2
	} else {
		goto L17
	}
L9:
	;
	v16 = base.I32_div_s(v10+v11, int32(2))
	v18 = v16 * int32(12)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_wc_isgraph_builtin[2])))
	if base.Ui32(v21) < base.Ui32(l0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v53 = int32(0)
	goto L4
L11:
	;
	if v32 <= v31 {
		v10 = v31
		v11 = v32
		goto L9
	} else {
		goto L16
	}
L12:
	;
	v31 = v10
	v32 = v16 + int32(1)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_wc_isgraph_builtin[3])))
	if base.Ui32(v27) <= base.Ui32(l0) {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v31 = v16 - int32(1)
	v32 = v11
	goto L11
L16:
	;
	goto L10
L17:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_wc_isgraph_builtin[4]))))
	if v45&int32(32) == int32(0) {
		v88 = v35
		goto L3
	} else {
		goto L18
	}
L18:
	;
	goto L2
L19:
	;
	v61 = int32(10)
	v62 = v54
	goto L20
L20:
	;
	v67 = base.I32_div_s(v61+v62, int32(2))
	v69 = v67 << (uint(int32(3)) % 32)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)+uint32(_c_F_wc_isgraph_builtin[5])))
	if base.Ui32(v72) < base.Ui32(l0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v88 = int32(1)
	goto L3
L22:
	;
	if v83 <= v82 {
		v61 = v82
		v62 = v83
		goto L20
	} else {
		goto L27
	}
L23:
	;
	v82 = v61
	v83 = v67 + int32(1)
	goto L22
L24:
	;
	goto L25
L25:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v69)+uint32(_c_F_wc_isgraph_builtin[6])))
	if base.Ui32(v78) <= base.Ui32(l0) {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v82 = v67 - int32(1)
	v83 = v62
	goto L22
L27:
	;
	goto L21
}
func F_wc_isprint_libc_mb(m *base.Module, l0 int32, l1 int32) int32 {
	var v14 int32
	_ = v14
	var v37 int32
	_ = v37
	if base.Ui32(l0) <= base.Ui32(int32(254)) {
		v37 = base.B2i32(base.Ui32(int32(32)) < base.Ui32((l0+int32(1))&int32(127)))
	} else {
		v14 = int32(_a_F_wc_isprint_libc_mb_0)
		v37 = base.B2i32(l0&v14 != v14)&base.B2i32(base.Ui32(l0-int32(_a_F_wc_isprint_libc_mb_1)) < base.Ui32(int32(_a_F_wc_isprint_libc_mb_2))) | (base.B2i32(base.Ui32(l0-int32(_a_F_wc_isprint_libc_mb_3)) < base.Ui32(int32(_a_F_wc_isprint_libc_mb_4))) | base.B2i32(base.Ui32(l0) < base.Ui32(int32(_a_F_wc_isprint_libc_mb_5))) | base.B2i32(base.Ui32(l0-int32(_a_F_wc_isprint_libc_mb_6)) < base.Ui32(int32(_a_F_wc_isprint_libc_mb_7))))
	}
	return base.B2i32(v37 != int32(0))
}
func F_wc_ispunct_builtin(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v140 int32
	_ = v140
	v3 = int32(0)
	v6 = int32(1)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if (v8^int32(-1))&v6 != 0 {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	return v140
L2:
	;
	v140 = base.B2i32(v6<<(uint(v129)%32)&int32(1073217536) != int32(0))
	goto L1
L3:
	;
	v116 = int32(1)
	v117 = l0 << (uint(v116) % 32)
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+uint32(_c_F_wc_ispunct_builtin[0]))))
	if v118&v116 != 0 {
		v140 = int32(0)
		goto L1
	} else {
		goto L37
	}
L4:
	;
	v140 = base.B2i32(v6<<(uint(v109)%32)&int32(821559296) != int32(0))
	goto L1
L5:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+uint32(_c_F_wc_ispunct_builtin[1]))))
	v109 = v103
	goto L4
L6:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_ispunct_builtin[2]))))
	v109 = v102
	goto L4
L7:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+uint32(_c_F_wc_ispunct_builtin[1]))))
	v129 = v99
	goto L2
L8:
	;
	if base.Ui32(l0) < base.Ui32(int32(128)) {
		goto L3
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	if base.Ui32(l0) <= base.Ui32(int32(127)) {
		goto L6
	} else {
		goto L28
	}
L11:
	;
	v17 = int32(1201)
	v18 = v3
	goto L12
L12:
	;
	v23 = base.I32_div_s(v17+v18, int32(2))
	v25 = v23 << (uint(int32(3)) % 32)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_wc_ispunct_builtin[3])))
	if base.Ui32(v28) < base.Ui32(l0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v45 = int32(3408)
	v46 = int32(0)
	goto L20
L14:
	;
	if v40 <= v39 {
		v17 = v39
		v18 = v40
		goto L12
	} else {
		goto L19
	}
L15:
	;
	v39 = v17
	v40 = v23 + int32(1)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_wc_ispunct_builtin[4])))
	if base.Ui32(v35) <= base.Ui32(l0) {
		v140 = int32(0)
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v39 = v23 - int32(1)
	v40 = v18
	goto L14
L19:
	;
	goto L13
L20:
	;
	v51 = base.I32_div_s(v45+v46, int32(2))
	v53 = v51 * int32(12)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)+uint32(_c_F_wc_ispunct_builtin[5])))
	if base.Ui32(v56) < base.Ui32(l0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v129 = int32(0)
	goto L2
L22:
	;
	if v67 <= v66 {
		v45 = v66
		v46 = v67
		goto L20
	} else {
		goto L27
	}
L23:
	;
	v66 = v45
	v67 = v51 + int32(1)
	goto L22
L24:
	;
	goto L25
L25:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v53)+uint32(_c_F_wc_ispunct_builtin[6])))
	if base.Ui32(v62) <= base.Ui32(l0) {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	v66 = v51 - int32(1)
	v67 = v46
	goto L22
L27:
	;
	goto L21
L28:
	;
	v74 = int32(3408)
	v75 = v3
	goto L29
L29:
	;
	v80 = base.I32_div_s(v74+v75, int32(2))
	v82 = v80 * int32(12)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)+uint32(_c_F_wc_ispunct_builtin[5])))
	if base.Ui32(v85) < base.Ui32(l0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v109 = int32(0)
	goto L4
L31:
	;
	if v96 <= v95 {
		v74 = v95
		v75 = v96
		goto L29
	} else {
		goto L36
	}
L32:
	;
	v95 = v74
	v96 = v80 + int32(1)
	goto L31
L33:
	;
	goto L34
L34:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v82)+uint32(_c_F_wc_ispunct_builtin[6])))
	if base.Ui32(v91) <= base.Ui32(l0) {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v95 = v80 - int32(1)
	v96 = v75
	goto L31
L36:
	;
	goto L30
L37:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+uint32(_c_F_wc_ispunct_builtin[2]))))
	v129 = v123
	goto L2
}
func F_wc_isupper_builtin(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v53 int32
	_ = v53
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v53
L2:
	;
	goto L1
L3:
	;
	v13 = int32(659)
	v18 = int32(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	v53 = base.B2i32(base.Ui32(l0-int32(65)) < base.Ui32(int32(26)))
	goto L2
L6:
	;
	v23 = base.I32_div_s(v13+v18, int32(2))
	v25 = v23 << (uint(int32(3)) % 32)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_wc_isupper_builtin[0])))
	if base.Ui32(v27) < base.Ui32(l0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v53 = int32(0)
	goto L2
L8:
	;
	if v38 <= v37 {
		v13 = v37
		v18 = v38
		goto L6
	} else {
		goto L13
	}
L9:
	;
	v37 = v13
	v38 = v23 + int32(1)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_wc_isupper_builtin[1])))
	if base.Ui32(v33) <= base.Ui32(l0) {
		v53 = int32(1)
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v37 = v23 - int32(1)
	v38 = v18
	goto L8
L13:
	;
	goto L7
}
func F_wc_isxdigit_libc_mb(m *base.Module, l0 int32, l1 int32) int32 {
	return base.B2i32(base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(l0|int32(32)-int32(97)) < base.Ui32(int32(6))) != int32(0))
}
