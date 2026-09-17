package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_u_isupper(m *base.Module, l0 int32) int32 {
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
	v17 = int32(655)
	v18 = int32(0)
	goto L5
L3:
	;
	goto L4
L4:
	;
	v53 = base.B2i32(base.Ui32(l0-int32(65)) < base.Ui32(int32(26)))
	goto L1
L5:
	;
	v23 = base.I32_div_s(v17+v18, int32(2))
	v25 = v23 << (uint(int32(3)) % 32)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_pg_u_isupper[0])))
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
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_pg_u_isupper[1])))
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
func F_pg_u_prop_case_ignorable(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v58 int32
	_ = v58
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v58
L2:
	;
	v19 = int32(505)
	v20 = int32(0)
	goto L5
L3:
	;
	goto L4
L4:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_u_prop_case_ignorable[0]))))
	v58 = int32(base.Ui32(v45&int32(16)) >> (uint(int32(4)) % 32))
	goto L1
L5:
	;
	v25 = base.I32_div_s(v19+v20, int32(2))
	v27 = v25 << (uint(int32(3)) % 32)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_pg_u_prop_case_ignorable[1])))
	if base.Ui32(v29) < base.Ui32(l0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v58 = int32(0)
	goto L1
L7:
	;
	if v40 <= v39 {
		v19 = v39
		v20 = v40
		goto L5
	} else {
		goto L14
	}
L8:
	;
	v39 = v19
	v40 = v25 + int32(1)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_pg_u_prop_case_ignorable[2])))
	if base.Ui32(v34) <= base.Ui32(l0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v58 = int32(1)
	goto L1
L12:
	;
	goto L13
L13:
	;
	v39 = v25 - int32(1)
	v40 = v20
	goto L7
L14:
	;
	goto L6
}
func F_pg_u_prop_cased(m *base.Module, l0 int32) int32 {
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
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v52 = int32(689)
	v53 = int32(0)
	goto L15
L2:
	;
	v10 = int32(3367)
	v11 = int32(0)
	goto L5
L3:
	;
	goto L4
L4:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_u_prop_cased[0]))))
	return int32(base.Ui32(v41&int32(8)) >> (uint(int32(3)) % 32))
L5:
	;
	v16 = base.I32_div_s(v10+v11, int32(2))
	v18 = v16 * int32(12)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_pg_u_prop_cased[1])))
	if base.Ui32(v21) < base.Ui32(l0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_pg_u_prop_cased[2]))))
	if v34 != int32(3) {
		goto L1
	} else {
		goto L14
	}
L7:
	;
	goto L6
L8:
	;
	if v32 <= v31 {
		v10 = v31
		v11 = v32
		goto L5
	} else {
		goto L13
	}
L9:
	;
	v31 = v10
	v32 = v16 + int32(1)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_pg_u_prop_cased[3])))
	if base.Ui32(v27) <= base.Ui32(l0) {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v31 = v16 - int32(1)
	v32 = v11
	goto L8
L13:
	;
	goto L1
L14:
	;
	return int32(1)
L15:
	;
	v58 = base.I32_div_s(v52+v53, int32(2))
	v60 = v58 << (uint(int32(3)) % 32)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+uint32(_c_F_pg_u_prop_cased[4])))
	if base.Ui32(v63) < base.Ui32(l0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v81 = int32(655)
	v82 = int32(0)
	goto L25
L17:
	;
	if v76 <= v75 {
		v52 = v75
		v53 = v76
		goto L15
	} else {
		goto L24
	}
L18:
	;
	v75 = v52
	v76 = v58 + int32(1)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v60)+uint32(_c_F_pg_u_prop_cased[5])))
	if base.Ui32(v69) <= base.Ui32(l0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(1)
L22:
	;
	goto L23
L23:
	;
	v75 = v58 - int32(1)
	v76 = v53
	goto L17
L24:
	;
	goto L16
L25:
	;
	v87 = base.I32_div_s(v81+v82, int32(2))
	v89 = v87 << (uint(int32(3)) % 32)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v89)+uint32(_c_F_pg_u_prop_cased[6])))
	if base.Ui32(v92) < base.Ui32(l0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	return int32(0)
L27:
	;
	if v105 <= v104 {
		v81 = v104
		v82 = v105
		goto L25
	} else {
		goto L34
	}
L28:
	;
	v104 = v81
	v105 = v87 + int32(1)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)+uint32(_c_F_pg_u_prop_cased[7])))
	if base.Ui32(v98) <= base.Ui32(l0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	return int32(1)
L32:
	;
	goto L33
L33:
	;
	v104 = v87 - int32(1)
	v105 = v82
	goto L27
L34:
	;
	goto L26
}
