package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_u_isalpha(m *base.Module, l0 int32) int32 {
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = int32(0)
	v11 = int32(1178)
	goto L4
L2:
	;
	goto L3
L3:
	;
	v38 = int32(1)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(v38)%32))+uint32(_consts[499]))))
	return v42 & v38
L4:
	;
	v16 = base.I32_div_s(v10+v11, int32(2))
	v18 = v16 << (uint(int32(3)) % 32)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[500])))
	if base.Ui32(v21) < base.Ui32(l0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	return int32(0)
L6:
	;
	if v33 <= v34 {
		v10 = v33
		v11 = v34
		goto L4
	} else {
		goto L13
	}
L7:
	;
	v33 = v16 + int32(1)
	v34 = v11
	goto L6
L8:
	;
	goto L9
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[501])))
	if base.Ui32(v27) <= base.Ui32(l0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(1)
L11:
	;
	goto L12
L12:
	;
	v33 = v10
	v34 = v16 - int32(1)
	goto L6
L13:
	;
	goto L5
}
func F_pg_u_isgraph(m *base.Module, l0 int32) int32 {
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
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	if base.Ui32(int32(128)) <= base.Ui32(l0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return int32(0)
L2:
	;
	return v90
L3:
	;
	v56 = int32(0)
	if int32(1)<<(uint(v55)%32)&int32(294913) != 0 {
		v90 = v56
		goto L2
	} else {
		goto L18
	}
L4:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[502]))))
	v55 = v52
	goto L3
L5:
	;
	v10 = int32(3367)
	v11 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	v35 = int32(1)
	v38 = l0 << (uint(v35) % 32)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+uint32(_consts[498]))))
	if v35<<(uint(v41)%32)&int32(294913) != 0 {
		goto L1
	} else {
		goto L16
	}
L8:
	;
	v16 = base.I32_div_s(v10+v11, int32(2))
	v18 = v16 * int32(12)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[503])))
	if base.Ui32(v21) < base.Ui32(l0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v55 = int32(0)
	goto L3
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
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[504])))
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
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+uint32(_consts[499]))))
	if v47&int32(32) == int32(0) {
		v90 = v35
		goto L2
	} else {
		goto L17
	}
L17:
	;
	goto L1
L18:
	;
	v63 = int32(10)
	v64 = v56
	goto L19
L19:
	;
	v69 = base.I32_div_s(v63+v64, int32(2))
	v71 = v69 << (uint(int32(3)) % 32)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v71)+uint32(_consts[1026])))
	if base.Ui32(v74) < base.Ui32(l0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v90 = int32(1)
	goto L2
L21:
	;
	if v85 <= v84 {
		v63 = v84
		v64 = v85
		goto L19
	} else {
		goto L26
	}
L22:
	;
	v84 = v63
	v85 = v69 + int32(1)
	goto L21
L23:
	;
	goto L24
L24:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v71)+uint32(_consts[1027])))
	if base.Ui32(v80) <= base.Ui32(l0) {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v84 = v69 - int32(1)
	v85 = v64
	goto L21
L26:
	;
	goto L20
}
func F_pg_u_ispunct(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
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
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	v6 = int32(1)
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	return base.B2i32(v6<<(uint(v133)%32)&int32(1073217536) != int32(0))
L2:
	;
	v116 = int32(1)
	v117 = l0 << (uint(v116) % 32)
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+uint32(_consts[499]))))
	if v120&v116 != 0 {
		goto L38
	} else {
		goto L39
	}
L3:
	;
	return base.B2i32(v6<<(uint(v109)%32)&int32(821559296) != int32(0))
L4:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+uint32(_consts[502]))))
	v109 = v103
	goto L3
L5:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_consts[498]))))
	v109 = v102
	goto L3
L6:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+uint32(_consts[502]))))
	v133 = v97
	goto L1
L7:
	;
	if base.Ui32(l0) < base.Ui32(int32(128)) {
		goto L2
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if base.Ui32(l0) <= base.Ui32(int32(127)) {
		goto L5
	} else {
		goto L29
	}
L10:
	;
	v13 = int32(0)
	v14 = int32(1178)
	goto L11
L11:
	;
	v19 = base.I32_div_s(v13+v14, int32(2))
	v21 = v19 << (uint(int32(3)) % 32)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[500])))
	if base.Ui32(v24) < base.Ui32(l0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v42 = int32(0)
	v43 = int32(3367)
	goto L21
L13:
	;
	if v36 <= v37 {
		v13 = v36
		v14 = v37
		goto L11
	} else {
		goto L20
	}
L14:
	;
	v36 = v19 + int32(1)
	v37 = v14
	goto L13
L15:
	;
	goto L16
L16:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[501])))
	if base.Ui32(v30) <= base.Ui32(l0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	goto L19
L19:
	;
	v36 = v13
	v37 = v19 - int32(1)
	goto L13
L20:
	;
	goto L12
L21:
	;
	v48 = base.I32_div_s(v42+v43, int32(2))
	v50 = v48 * int32(12)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v50)+uint32(_consts[503])))
	if base.Ui32(v53) < base.Ui32(l0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v133 = int32(0)
	goto L1
L23:
	;
	if v63 <= v64 {
		v42 = v63
		v43 = v64
		goto L21
	} else {
		goto L28
	}
L24:
	;
	v63 = v48 + int32(1)
	v64 = v43
	goto L23
L25:
	;
	goto L26
L26:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v50)+uint32(_consts[504])))
	if base.Ui32(v59) <= base.Ui32(l0) {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v63 = v42
	v64 = v48 - int32(1)
	goto L23
L28:
	;
	goto L22
L29:
	;
	v72 = int32(0)
	v73 = int32(3367)
	goto L30
L30:
	;
	v78 = base.I32_div_s(v72+v73, int32(2))
	v80 = v78 * int32(12)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v80)+uint32(_consts[503])))
	if base.Ui32(v83) < base.Ui32(l0) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v109 = int32(0)
	goto L3
L32:
	;
	if v93 <= v94 {
		v72 = v93
		v73 = v94
		goto L30
	} else {
		goto L37
	}
L33:
	;
	v93 = v78 + int32(1)
	v94 = v73
	goto L32
L34:
	;
	goto L35
L35:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v80)+uint32(_consts[504])))
	if base.Ui32(v89) <= base.Ui32(l0) {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v93 = v72
	v94 = v78 - int32(1)
	goto L32
L37:
	;
	goto L31
L38:
	;
	return int32(0)
L39:
	;
	goto L40
L40:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+uint32(_consts[498]))))
	v133 = v127
	goto L1
}
func F_pg_u_isspace(m *base.Module, l0 int32) int32 {
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = int32(0)
	v11 = int32(10)
	goto L4
L2:
	;
	goto L3
L3:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_consts[499]))))
	return int32(base.Ui32(v42&int32(32)) >> (uint(int32(5)) % 32))
L4:
	;
	v16 = base.I32_div_s(v10+v11, int32(2))
	v18 = v16 << (uint(int32(3)) % 32)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[1026])))
	if base.Ui32(v21) < base.Ui32(l0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	return int32(0)
L6:
	;
	if v33 <= v34 {
		v10 = v33
		v11 = v34
		goto L4
	} else {
		goto L13
	}
L7:
	;
	v33 = v16 + int32(1)
	v34 = v11
	goto L6
L8:
	;
	goto L9
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[1027])))
	if base.Ui32(v27) <= base.Ui32(l0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(1)
L11:
	;
	goto L12
L12:
	;
	v33 = v10
	v34 = v16 - int32(1)
	goto L6
L13:
	;
	goto L5
}
