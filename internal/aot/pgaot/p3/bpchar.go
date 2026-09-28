package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bpchar_smaller(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = int32(1)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v22 = v20 & v18
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = v18
	goto L6
L5:
	;
	v23 = int32(4)
	goto L6
L6:
	;
	v24 = v23 + v11
	if v20 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v57 = v51
	goto L18
L8:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v30 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v41 = int32(1)
	if v22 != 0 {
		v51 = int32(base.Ui32(v20)>>(uint(v41)%32)) - v41
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v33 = int32(16)
	goto L13
L12:
	;
	v33 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v30-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v40 = int32(4)
	goto L16
L15:
	;
	v40 = v33
	goto L16
L16:
	;
	v51 = v40
	goto L7
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	if v57 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v74 = int32(1)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v78 = v76 & v74
	if v78 != 0 {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	goto L19
L21:
	;
	v73 = v51 & (v51 >> (uint(int32(31)) % 32))
	goto L20
L22:
	;
	goto L23
L23:
	;
	v67 = v57 - int32(1)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v67))))
	if v69 == int32(32) {
		v57 = v67
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v73 = v57
	goto L20
L25:
	;
	v79 = v74
	goto L27
L26:
	;
	v79 = int32(4)
	goto L27
L27:
	;
	v80 = v79 + v16
	if v76 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v113 = v107
	goto L39
L29:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v86 == int32(18) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v97 = int32(1)
	if v78 != 0 {
		v107 = int32(base.Ui32(v76)>>(uint(v97)%32)) - v97
		goto L28
	} else {
		goto L38
	}
L32:
	;
	v89 = int32(16)
	goto L34
L33:
	;
	v89 = int32(0)
	goto L34
L34:
	;
	if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v96 = int32(4)
	goto L37
L36:
	;
	v96 = v89
	goto L37
L37:
	;
	v107 = v96
	goto L28
L38:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
	goto L28
L39:
	;
	if v113 <= int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v131 = F_varstr_cmp(m, v24, v73, v80, v128, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L46
	}
L41:
	;
	goto L40
L42:
	;
	v128 = v107 & (v107 >> (uint(int32(31)) % 32))
	goto L41
L43:
	;
	goto L44
L44:
	;
	v123 = v113 - int32(1)
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+v123))))
	if v125 == int32(32) {
		v113 = v123
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v128 = v113
	goto L41
L46:
	;
	if v131 <= int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v135 = v11
	goto L49
L48:
	;
	v135 = v16
	goto L49
L49:
	;
	return base.I64_extend_i32_u(v135)
}
