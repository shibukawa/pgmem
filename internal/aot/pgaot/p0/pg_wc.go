package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_wc_isalnum(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
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
	var v144 int32
	_ = v144
	var v164 int32
	_ = v164
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wc_isalnum[0]))
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
	return v164
L2:
	;
	if base.Ui32(int32(255)) < base.Ui32(l0) {
		v164 = v2
		goto L1
	} else {
		goto L41
	}
L3:
	;
	if base.Ui32(int32(10)) <= base.Ui32(l0-int32(48)) {
		goto L38
	} else {
		goto L39
	}
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wc_isalnum[1]))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+16)))
	v21 = (v17 ^ int32(-1)) & int32(1)
	if base.Ui32(int32(128)) <= base.Ui32(l0) {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		v164 = v2
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_pg_wc_isalnum[2]))))
	return base.B2i32(v9&int32(3) != int32(0))
L7:
	;
	return v131
L8:
	;
	v131 = v124
	goto L7
L9:
	;
	v124 = base.B2i32(v113&int32(255) == int32(9))
	goto L8
L10:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_pg_wc_isalnum[3]))))
	v113 = v106
	goto L9
L11:
	;
	v131 = base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10)))
	goto L7
L12:
	;
	v31 = int32(1178)
	v32 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	v86 = int32(1)
	v88 = l0 << (uint(v86) % 32)
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_pg_wc_isalnum[4]))))
	if v89&v86 != 0 {
		v124 = v86
		goto L8
	} else {
		goto L35
	}
L15:
	;
	v37 = base.I32_div_s(v31+v32, int32(2))
	v39 = v37 << (uint(int32(3)) % 32)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_pg_wc_isalnum[5])))
	if base.Ui32(v42) < base.Ui32(l0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if v21 != 0 {
		goto L11
	} else {
		goto L25
	}
L17:
	;
	if v54 <= v53 {
		v31 = v53
		v32 = v54
		goto L15
	} else {
		goto L24
	}
L18:
	;
	v53 = v31
	v54 = v37 + int32(1)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_pg_wc_isalnum[6])))
	if base.Ui32(v48) <= base.Ui32(l0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v131 = int32(1)
	goto L7
L22:
	;
	goto L23
L23:
	;
	v53 = v37 - int32(1)
	v54 = v32
	goto L17
L24:
	;
	goto L16
L25:
	;
	v60 = int32(3367)
	v61 = int32(0)
	goto L27
L26:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+uint32(_c_F_pg_wc_isalnum[7]))))
	v113 = v85
	goto L9
L27:
	;
	v66 = base.I32_div_s(v60+v61, int32(2))
	v68 = v66 * int32(12)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+uint32(_c_F_pg_wc_isalnum[8])))
	if base.Ui32(v71) < base.Ui32(l0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v113 = int32(0)
	goto L9
L29:
	;
	if v82 <= v81 {
		v60 = v81
		v61 = v82
		goto L27
	} else {
		goto L34
	}
L30:
	;
	v81 = v60
	v82 = v66 + int32(1)
	goto L29
L31:
	;
	goto L32
L32:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+uint32(_c_F_pg_wc_isalnum[9])))
	if base.Ui32(v77) <= base.Ui32(l0) {
		goto L26
	} else {
		goto L33
	}
L33:
	;
	v81 = v66 - int32(1)
	v82 = v61
	goto L29
L34:
	;
	goto L28
L35:
	;
	if v21 == int32(0) {
		goto L10
	} else {
		goto L36
	}
L36:
	;
	goto L11
L37:
	;
	return v144
L38:
	;
	v140 = F_iswalpha(m, l0)
	mBase = m.M
	v144 = base.B2i32(v140 != int32(0))
	goto L40
L39:
	;
	v144 = int32(1)
	goto L40
L40:
	;
	goto L37
L41:
	;
	goto L42
L42:
	;
	v164 = base.B2i32(base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(l0|int32(32)-int32(97)) < base.Ui32(int32(26))) != int32(0))
	goto L1
}
func F_pg_wc_isdigit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v96 int32
	_ = v96
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wc_isdigit[0]))
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
		goto L23
	} else {
		goto L24
	}
L2:
	;
	goto L22
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wc_isdigit[1]))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)))
	if (v13^int32(-1))&int32(1) != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10)))
L5:
	;
	return v74
L6:
	;
	v74 = base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10)))
	goto L5
L7:
	;
	goto L8
L8:
	;
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v74 = base.B2i32(v64&int32(255) == int32(9))
	goto L5
L10:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_pg_wc_isdigit[2]))))
	v64 = v58
	goto L9
L11:
	;
	v30 = int32(0)
	v31 = int32(3367)
	goto L14
L12:
	;
	goto L13
L13:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_isdigit[3]))))
	v64 = v57
	goto L9
L14:
	;
	v36 = base.I32_div_s(v30+v31, int32(2))
	v38 = v36 * int32(12)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_pg_wc_isdigit[4])))
	if base.Ui32(v41) < base.Ui32(l0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v64 = int32(0)
	goto L9
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
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_pg_wc_isdigit[5])))
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
	return base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10)))
L23:
	;
	goto L26
L24:
	;
	v96 = int32(0)
	goto L25
L25:
	;
	return v96
L26:
	;
	v96 = base.B2i32(base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10))) != int32(0))
	goto L25
}
