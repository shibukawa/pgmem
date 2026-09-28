package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gtrgm_options(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(0)
	F_add_local_int_reloption(m, v2, int32(_a_F_gtrgm_options_0), int32(_a_F_gtrgm_options_1), int32(12), int32(1), int32(2024))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_gtrgm_same(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	v2 = int32(0)
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v14 == v2 {
		v31 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v31&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	if v18 == int32(0) {
		v31 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v21 != int32(7) {
		v31 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v24 != int32(17) {
		v31 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+32)))
	v31 = v27 ^ int32(1)
	goto L2
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v35 = F_get_fn_opclass_options(m, v34)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v40 = int32(12)
	goto L9
L9:
	;
	v43 = base.I32_wrap_i64(v10)
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+4)))
	if v44&int32(2) != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	return int64(0)
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v40 = v39
	goto L9
L12:
	;
	return v10 & int64(4294967295)
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v43))) = uint8(v139)
	goto L12
L14:
	;
	v48 = v44 & int32(4)
	if v48 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v87 = int32(2)
	v89 = int32(5)
	v90 = int32(base.Ui32(v86)>>(uint(v87)%32)) - v89
	v91 = int32(3)
	v92 = base.I32_div_u_s(v90, v91)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v99 = base.I32_div_u_s(int32(base.Ui32(v93)>>(uint(v87)%32))-v89, v91)
	if v92 != v99 {
		v139 = v2
		goto L13
	} else {
		goto L29
	}
L17:
	;
	if v48 != 0 {
		v139 = v2
		goto L13
	} else {
		goto L20
	}
L18:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)))
	if v51&int32(4) == int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v139 = int32(1)
	goto L13
L20:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)))
	if v57&int32(4) != 0 {
		v139 = v2
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v60 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43))) = uint8(v60)
	v62 = int32(0)
	if v40 <= v62 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	v65 = int32(5)
	v69 = v62
	goto L23
L23:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69+(v12+v65)))))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69+(v11+v65)))))
	if v79 == v81 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v139 = v2
	goto L13
L25:
	;
	v84 = v69 + int32(1)
	if v40 != v84 {
		v69 = v84
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	goto L12
L29:
	;
	v101 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43))) = uint8(v101)
	if base.Ui32(v90) < base.Ui32(int32(3)) {
		goto L12
	} else {
		goto L30
	}
L30:
	;
	v105 = int32(5)
	v109 = int32(1)
	if base.Ui32(v92) <= base.Ui32(v109) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v112 = v109
	goto L33
L32:
	;
	v112 = v92
	goto L33
L33:
	;
	v114 = int32(0)
	goto L34
L34:
	;
	v124 = v114 * int32(3)
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_gtrgm_same[0]))
	v129 = m.T0[v128].(func(*base.Module, int32, int32) int32)(m, v12+v105+v124, v11+v105+v124)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L10
	} else {
		goto L36
	}
L35:
	;
	goto L12
L36:
	;
	if v129 != 0 {
		v139 = v2
		goto L13
	} else {
		goto L37
	}
L37:
	;
	v132 = v114 + int32(1)
	if v112 != v132 {
		v114 = v132
		goto L34
	} else {
		goto L38
	}
L38:
	;
	goto L35
}
