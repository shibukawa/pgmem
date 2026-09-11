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
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
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
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v171 int32
	_ = v171
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
	return v171
L2:
	;
	if base.Ui32(int32(255)) < base.Ui32(l0) {
		v171 = v2
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
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wc_isalnum[1]))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+16)))
	v23 = (v19 ^ int32(-1)) & int32(1)
	if base.Ui32(int32(128)) <= base.Ui32(l0) {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		v171 = v2
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_pg_wc_isalnum[2]))))
	return base.B2i32(v11&int32(3) != int32(0))
L7:
	;
	return v137
L8:
	;
	v137 = v130
	goto L7
L9:
	;
	v130 = base.B2i32(v119&int32(255) == int32(9))
	goto L8
L10:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_isalnum[3]))))
	v119 = v112
	goto L9
L11:
	;
	v137 = base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10)))
	goto L7
L12:
	;
	v33 = int32(1178)
	v34 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	v88 = int32(1)
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(v88)%32))+uint32(_c_F_pg_wc_isalnum[4]))))
	if v93&v88 != 0 {
		v130 = v88
		goto L8
	} else {
		goto L35
	}
L15:
	;
	v39 = base.I32_div_s(v33+v34, int32(2))
	v41 = v39 << (uint(int32(3)) % 32)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_pg_wc_isalnum[5])))
	if base.Ui32(v44) < base.Ui32(l0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if v23 != 0 {
		goto L11
	} else {
		goto L25
	}
L17:
	;
	if v56 <= v55 {
		v33 = v55
		v34 = v56
		goto L15
	} else {
		goto L24
	}
L18:
	;
	v55 = v33
	v56 = v39 + int32(1)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_pg_wc_isalnum[6])))
	if base.Ui32(v50) <= base.Ui32(l0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v137 = int32(1)
	goto L7
L22:
	;
	goto L23
L23:
	;
	v55 = v39 - int32(1)
	v56 = v34
	goto L17
L24:
	;
	goto L16
L25:
	;
	v62 = int32(3367)
	v63 = int32(0)
	goto L27
L26:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+uint32(_c_F_pg_wc_isalnum[7]))))
	v119 = v87
	goto L9
L27:
	;
	v68 = base.I32_div_s(v62+v63, int32(2))
	v70 = v68 * int32(12)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v70)+uint32(_c_F_pg_wc_isalnum[8])))
	if base.Ui32(v73) < base.Ui32(l0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v119 = int32(0)
	goto L9
L29:
	;
	if v84 <= v83 {
		v62 = v83
		v63 = v84
		goto L27
	} else {
		goto L34
	}
L30:
	;
	v83 = v62
	v84 = v68 + int32(1)
	goto L29
L31:
	;
	goto L32
L32:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v70)+uint32(_c_F_pg_wc_isalnum[9])))
	if base.Ui32(v79) <= base.Ui32(l0) {
		goto L26
	} else {
		goto L33
	}
L33:
	;
	v83 = v68 - int32(1)
	v84 = v63
	goto L29
L34:
	;
	goto L28
L35:
	;
	if v23 == int32(0) {
		goto L10
	} else {
		goto L36
	}
L36:
	;
	goto L11
L37:
	;
	return v151
L38:
	;
	v148 = F_iswalpha(m, l0)
	mBase = m.M
	v151 = base.B2i32(v148 != int32(0))
	goto L40
L39:
	;
	v151 = int32(1)
	goto L40
L40:
	;
	goto L37
L41:
	;
	goto L42
L42:
	;
	v171 = base.B2i32(base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(l0|int32(32)-int32(97)) < base.Ui32(int32(26))) != int32(0))
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v98 int32
	_ = v98
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
	return v76
L6:
	;
	v76 = base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10)))
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
	v76 = base.B2i32(v66&int32(255) == int32(9))
	goto L5
L10:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_pg_wc_isdigit[2]))))
	v66 = v60
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
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_wc_isdigit[3]))))
	v66 = v59
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
	v66 = int32(0)
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
	v98 = int32(0)
	goto L25
L25:
	;
	return v98
L26:
	;
	v98 = base.B2i32(base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10))) != int32(0))
	goto L25
}
