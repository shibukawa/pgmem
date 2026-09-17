package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_u_isalnum(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v107 int32
	_ = v107
	if base.Ui32(int32(128)) <= base.Ui32(l0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return v107
L2:
	;
	v107 = base.B2i32(v96&int32(255) == int32(9))
	goto L1
L3:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+uint32(_c_F_pg_u_isalnum[0]))))
	v96 = v89
	goto L2
L4:
	;
	return base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10)))
L5:
	;
	v12 = int32(1178)
	v13 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	v68 = int32(1)
	v70 = l0 << (uint(v68) % 32)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+uint32(_c_F_pg_u_isalnum[1]))))
	if v71&v68 != 0 {
		v107 = v68
		goto L1
	} else {
		goto L28
	}
L8:
	;
	v18 = base.I32_div_s(v12+v13, int32(2))
	v20 = v18 << (uint(int32(3)) % 32)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_pg_u_isalnum[2])))
	if base.Ui32(v23) < base.Ui32(l0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if l1 != 0 {
		goto L4
	} else {
		goto L18
	}
L10:
	;
	if v36 <= v35 {
		v12 = v35
		v13 = v36
		goto L8
	} else {
		goto L17
	}
L11:
	;
	v35 = v12
	v36 = v18 + int32(1)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_c_F_pg_u_isalnum[3])))
	if base.Ui32(v29) <= base.Ui32(l0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(1)
L15:
	;
	goto L16
L16:
	;
	v35 = v18 - int32(1)
	v36 = v13
	goto L10
L17:
	;
	goto L9
L18:
	;
	v42 = int32(3367)
	v43 = int32(0)
	goto L20
L19:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+uint32(_c_F_pg_u_isalnum[4]))))
	v96 = v67
	goto L2
L20:
	;
	v48 = base.I32_div_s(v42+v43, int32(2))
	v50 = v48 * int32(12)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v50)+uint32(_c_F_pg_u_isalnum[5])))
	if base.Ui32(v53) < base.Ui32(l0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v96 = int32(0)
	goto L2
L22:
	;
	if v64 <= v63 {
		v42 = v63
		v43 = v64
		goto L20
	} else {
		goto L27
	}
L23:
	;
	v63 = v42
	v64 = v48 + int32(1)
	goto L22
L24:
	;
	goto L25
L25:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v50)+uint32(_c_F_pg_u_isalnum[6])))
	if base.Ui32(v59) <= base.Ui32(l0) {
		goto L19
	} else {
		goto L26
	}
L26:
	;
	v63 = v48 - int32(1)
	v64 = v43
	goto L22
L27:
	;
	goto L21
L28:
	;
	if l1 == int32(0) {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	goto L4
}
func F_pg_u_isdigit(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10)))
L2:
	;
	goto L3
L3:
	;
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return base.B2i32(v50&int32(255) == int32(9))
L5:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_pg_u_isdigit[0]))))
	v50 = v44
	goto L4
L6:
	;
	v16 = int32(0)
	v17 = int32(3367)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_u_isdigit[1]))))
	v50 = v43
	goto L4
L9:
	;
	v22 = base.I32_div_s(v16+v17, int32(2))
	v24 = v22 * int32(12)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_pg_u_isdigit[2])))
	if base.Ui32(v27) < base.Ui32(l0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v50 = int32(0)
	goto L4
L11:
	;
	if v37 <= v38 {
		v16 = v37
		v17 = v38
		goto L9
	} else {
		goto L16
	}
L12:
	;
	v37 = v22 + int32(1)
	v38 = v17
	goto L11
L13:
	;
	goto L14
L14:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_pg_u_isdigit[3])))
	if base.Ui32(v33) <= base.Ui32(l0) {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v37 = v16
	v38 = v22 - int32(1)
	goto L11
L16:
	;
	goto L10
}
func F_pg_u_islower(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v53 int32
	_ = v53
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v53
L2:
	;
	v17 = int32(689)
	v18 = int32(0)
	goto L5
L3:
	;
	goto L4
L4:
	;
	v53 = base.B2i32(base.Ui32(l0-int32(97)) < base.Ui32(int32(26)))
	goto L1
L5:
	;
	v23 = base.I32_div_s(v17+v18, int32(2))
	v25 = v23 << (uint(int32(3)) % 32)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_pg_u_islower[0])))
	if base.Ui32(v27) < base.Ui32(l0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v53 = int32(0)
	goto L1
L7:
	;
	if v38 <= v37 {
		v17 = v37
		v18 = v38
		goto L5
	} else {
		goto L14
	}
L8:
	;
	v37 = v17
	v38 = v23 + int32(1)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_pg_u_islower[1])))
	if base.Ui32(v32) <= base.Ui32(l0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v53 = int32(1)
	goto L1
L12:
	;
	goto L13
L13:
	;
	v37 = v23 - int32(1)
	v38 = v18
	goto L7
L14:
	;
	goto L6
}
