package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_NameListToString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = v9 + int32(16)
	F_initStringInfo(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l0 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L26
	}
L4:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	m.G0 = v9 + int32(32)
	return v80
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v19 <= int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v24 != int32(468) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v37 = int32(1)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v38 <= v37 {
		goto L4
	} else {
		goto L14
	}
L8:
	;
	if v24 != int32(77) {
		v87 = v23
		goto L3
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	F_appendStringInfoString(m, v9+int32(16), v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	F_appendStringInfoChar(m, v12, int32(42))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	goto L7
L13:
	;
	goto L7
L14:
	;
	v44 = v37
	goto L15
L15:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47+v44<<(uint(int32(2))%32))))
	v53 = v9 + int32(16)
	F_appendStringInfoChar(m, v53, int32(46))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L4
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v57 != int32(77) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v71 = v44 + int32(1)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v71 < v72 {
		v44 = v71
		goto L15
	} else {
		goto L25
	}
L19:
	;
	if v57 != int32(468) {
		v87 = v51
		goto L3
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	F_appendStringInfoChar(m, v9+int32(16), int32(42))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	F_appendStringInfoString(m, v53, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	goto L18
L24:
	;
	goto L18
L25:
	;
	goto L16
L26:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v95
	F_errmsg_internal(m, int32(_a_F_NameListToString_0), v9)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_NameListToString_1), int32(3617), int32(_a_F_NameListToString_2))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_nameeq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 == int32(950) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	goto L6
L2:
	;
	goto L3
L3:
	;
	v58 = F_strlen(m, v5)
	mBase = m.M
	v59 = F_strlen(m, v4)
	mBase = m.M
	v60 = F_varstr_cmp(m, v5, v58, v4, v59, v6)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	return base.B2i32(v46-v47 == int32(0))
L6:
	;
	goto L7
L7:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v15 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v16 = v5
	v17 = v4
	v18 = int32(64)
	v19 = v15
	goto L12
L9:
	;
	v42 = v4
	v46 = int32(0)
	goto L10
L10:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	goto L4
L11:
	;
	v42 = v37
	v46 = v39
	goto L10
L12:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if base.B2i32(v19 != v21)|base.B2i32(v21 == int32(0)) != 0 {
		v37 = v17
		v39 = v19
		goto L11
	} else {
		goto L14
	}
L13:
	;
	v37 = v31
	v39 = int32(0)
	goto L11
L14:
	;
	v27 = v18 - int32(1)
	if v27 == int32(0) {
		v37 = v17
		v39 = v19
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v30 = int32(1)
	v31 = v17 + v30
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v32 != 0 {
		v16 = v16 + v30
		v17 = v31
		v18 = v27
		v19 = v32
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	return int32(0)
L18:
	;
	return base.B2i32(v60 == int32(0))
}
func F_nameeqfast(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	goto L3
L1:
	;
	return base.B2i32(v40-v41 == int32(0))
L3:
	;
	goto L4
L4:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v9 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v10 = l0
	v11 = l1
	v12 = int32(64)
	v13 = v9
	goto L9
L6:
	;
	v36 = l1
	v40 = int32(0)
	goto L7
L7:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	goto L1
L8:
	;
	v36 = v31
	v40 = v33
	goto L7
L9:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if base.B2i32(v13 != v15)|base.B2i32(v15 == int32(0)) != 0 {
		v31 = v11
		v33 = v13
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v31 = v25
	v33 = int32(0)
	goto L8
L11:
	;
	v21 = v12 - int32(1)
	if v21 == int32(0) {
		v31 = v11
		v33 = v13
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v24 = int32(1)
	v25 = v11 + v24
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v26 != 0 {
		v10 = v10 + v24
		v11 = v25
		v12 = v21
		v13 = v26
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
}
func F_namege(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 == int32(950) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return int32(base.Ui32(v61^int32(-1)) >> (uint(int32(31)) % 32))
L2:
	;
	goto L7
L3:
	;
	goto L4
L4:
	;
	v55 = F_strlen(m, v5)
	mBase = m.M
	v56 = F_strlen(m, v4)
	mBase = m.M
	v57 = F_varstr_cmp(m, v5, v55, v4, v56, v6)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	v61 = v46 - v47
	goto L1
L7:
	;
	goto L8
L8:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v15 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v16 = v5
	v17 = v4
	v18 = int32(64)
	v19 = v15
	goto L13
L10:
	;
	v42 = v4
	v46 = int32(0)
	goto L11
L11:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	goto L5
L12:
	;
	v42 = v37
	v46 = v39
	goto L11
L13:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if base.B2i32(v19 != v21)|base.B2i32(v21 == int32(0)) != 0 {
		v37 = v17
		v39 = v19
		goto L12
	} else {
		goto L15
	}
L14:
	;
	v37 = v31
	v39 = int32(0)
	goto L12
L15:
	;
	v27 = v18 - int32(1)
	if v27 == int32(0) {
		v37 = v17
		v39 = v19
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v30 = int32(1)
	v31 = v17 + v30
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v32 != 0 {
		v16 = v16 + v30
		v17 = v31
		v18 = v27
		v19 = v32
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	return int32(0)
L19:
	;
	v61 = v57
	goto L1
}
func F_nameicregexne(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13950(m, l0, int32(27))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_namenetext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = F_strlen(m, v8)
	mBase = m.M
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v15 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v45 != int32(950) {
		goto L16
	} else {
		goto L17
	}
L4:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v21 == int32(18) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v32 = int32(1)
	if v15&v32 != 0 {
		v44 = int32(base.Ui32(v15)>>(uint(v32)%32)) - v32
		goto L3
	} else {
		goto L13
	}
L7:
	;
	v24 = int32(16)
	goto L9
L8:
	;
	v24 = int32(0)
	goto L9
L9:
	;
	if base.Ui32((v21-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v31 = int32(4)
	goto L12
L11:
	;
	v31 = v24
	goto L12
L12:
	;
	v44 = v31
	goto L3
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v44 = int32(base.Ui32(v38)>>(uint(int32(2))%32)) - int32(4)
	goto L3
L14:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v151 != v10 {
		goto L51
	} else {
		goto L52
	}
L15:
	;
	v140 = int32(1)
	if v15&v140 != 0 {
		goto L47
	} else {
		goto L48
	}
L16:
	;
	if v45 != 0 {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if v44 != v14 {
		v150 = int32(1)
		goto L14
	} else {
		goto L25
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errmsg(m, int32(_a_F_namenetext_0), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_errhint(m, int32(_a_F_namenetext_1), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_namenetext_2), int32(1648), int32(_a_F_namenetext_3))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	v70 = int32(1)
	if v15&v70 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v74 = v70
	goto L28
L27:
	;
	v74 = int32(4)
	goto L28
L28:
	;
	v75 = v10 + v74
	if base.Ui32(int32(4)) <= base.Ui32(v14) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v150 = base.B2i32(v137 != int32(0))
	goto L14
L30:
	;
	v137 = int32(0)
	goto L29
L31:
	;
	v111 = v106
	v112 = v107
	v113 = v108
	goto L41
L32:
	;
	if (v8|v75)&int32(3) != 0 {
		v106 = v8
		v107 = v75
		v108 = v14
		goto L31
	} else {
		goto L35
	}
L33:
	;
	v99 = v8
	v100 = v75
	v101 = v14
	goto L34
L34:
	;
	if v101 == int32(0) {
		goto L30
	} else {
		goto L40
	}
L35:
	;
	v83 = v8
	v84 = v75
	v85 = v14
	goto L36
L36:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	if v88 != v89 {
		v106 = v83
		v107 = v84
		v108 = v85
		goto L31
	} else {
		goto L38
	}
L37:
	;
	v99 = v94
	v100 = v92
	v101 = v96
	goto L34
L38:
	;
	v91 = int32(4)
	v92 = v84 + v91
	v94 = v83 + v91
	v96 = v85 - v91
	if base.Ui32(int32(3)) < base.Ui32(v96) {
		v83 = v94
		v84 = v92
		v85 = v96
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v106 = v99
	v107 = v100
	v108 = v101
	goto L31
L41:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v116 == v117 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v137 = v116 - v117
	goto L29
L43:
	;
	v119 = int32(1)
	v124 = v113 - v119
	if v124 != 0 {
		v111 = v111 + v119
		v112 = v112 + v119
		v113 = v124
		goto L41
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	goto L42
L46:
	;
	goto L30
L47:
	;
	v144 = v140
	goto L49
L48:
	;
	v144 = int32(4)
	goto L49
L49:
	;
	v146 = F_varstr_cmp(m, v8, v14, v10+v144, v44, v45)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v150 = base.B2i32(v146 != int32(0))
	goto L14
L51:
	;
	F_pfree(m, v10)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	return v150
L54:
	;
	goto L53
}
func F_nameregexne(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13950(m, l0, int32(19))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_namesend(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pq_begintypsend(m, v5)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_strlen(m, v7)
		mBase = m.M
		F_pq_sendtext(m, v5, v7, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = v17 << (uint(int32(2)) % 32)
			m.G0 = v5 + int32(16)
			return v16
		}
	}
}
func F_new_intArrayType(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	if l0 <= int32(0) {
		v7 = F_construct_empty_array(m, int32(23))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v7
		}
	} else {
		v15 = l0<<(uint(int32(2))%32) + int32(24)
		v16 = F_palloc0(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(23)
			*(*int64)(unsafe.Add(mBase, uint32(v16)+4)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = v15 << (uint(int32(2)) % 32)
			return v16
		}
	}
}
func F_newstate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int64
	_ = v108
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_newstate[0]))
	if v8 != 0 {
		F_ProcessInterrupts(m)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v13 != 0 {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v14
				v92 = v13
				v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v97 + int32(1)
				v101 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v92)+4)) = uint8(v101)
				*(*int32)(unsafe.Add(mBase, uint32(v92))) = v97
				v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v104 == v101 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v92
				} else {
				}
				v108 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v92)+24)) = v108
				*(*int64)(unsafe.Add(mBase, uint32(v92)+16)) = v108
				*(*int64)(unsafe.Add(mBase, uint32(v92)+8)) = v108
				v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v114 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v114)+28)) = v92
					v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					v118 = v116
				} else {
					v118 = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v92)+32)) = v118
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v92
				return v92
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				if v16 != 0 {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
					if base.Ui32(v18) <= base.Ui32(v17) {
						v36 = l0 + int32(76)
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+136))
						if base.Ui32(v38) < base.Ui32(int32(98000000)) {
							v53 = int32(1024)
							v55 = v18 << (uint(int32(1)) % 32)
							if base.Ui32(v53) <= base.Ui32(v55) {
								v58 = v53
							} else {
								v58 = v55
							}
							v60 = v36
							v61 = v58
							v65 = v61*int32(36) + int32(8)
							v67 = F_palloc_extended(m, v65, int32(2))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								v69 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
								if v67 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v69)+24)) = int32(101)
									v74 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
									if v75 != 0 {
										v77 = v75
									} else {
										v77 = int32(12)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v74)+12)) = v77
									return int32(0)
								} else {
									v81 = *(*int32)(unsafe.Add(mBase, uint32(v69)+136))
									*(*int32)(unsafe.Add(mBase, uint32(v69)+136)) = v81 + v65
									*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = v61
									v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									*(*int32)(unsafe.Add(mBase, uint32(v67))) = v85
									*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v67
									v92 = v67 + int32(8)
									v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v97 + int32(1)
									v101 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v92)+4)) = uint8(v101)
									*(*int32)(unsafe.Add(mBase, uint32(v92))) = v97
									v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									if v104 == v101 {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v92
									} else {
									}
									v108 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v92)+24)) = v108
									*(*int64)(unsafe.Add(mBase, uint32(v92)+16)) = v108
									*(*int64)(unsafe.Add(mBase, uint32(v92)+8)) = v108
									v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									if v114 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v114)+28)) = v92
										v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v118 = v116
									} else {
										v118 = int32(0)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v92)+32)) = v118
									*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v92
									return v92
								}
							}
						} else {
							v41 = v37
							v42 = v36
							*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = int32(101)
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
							if v47 != 0 {
								v49 = v47
							} else {
								v49 = int32(19)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = v49
							return int32(0)
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v17 + int32(1)
						v92 = v16 + v17*int32(36) + int32(8)
						v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v97 + int32(1)
						v101 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v92)+4)) = uint8(v101)
						*(*int32)(unsafe.Add(mBase, uint32(v92))) = v97
						v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v104 == v101 {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v92
						} else {
						}
						v108 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v92)+24)) = v108
						*(*int64)(unsafe.Add(mBase, uint32(v92)+16)) = v108
						*(*int64)(unsafe.Add(mBase, uint32(v92)+8)) = v108
						v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						if v114 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v114)+28)) = v92
							v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v118 = v116
						} else {
							v118 = int32(0)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v92)+32)) = v118
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v92
						return v92
					}
				} else {
					v29 = l0 + int32(76)
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+136))
					if base.Ui32(int32(97999999)) < base.Ui32(v32) {
						v41 = v31
						v42 = v29
						*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = int32(101)
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
						if v47 != 0 {
							v49 = v47
						} else {
							v49 = int32(19)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = v49
						return int32(0)
					} else {
						v60 = v29
						v61 = int32(32)
						v65 = v61*int32(36) + int32(8)
						v67 = F_palloc_extended(m, v65, int32(2))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
							if v67 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(v69)+24)) = int32(101)
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
								if v75 != 0 {
									v77 = v75
								} else {
									v77 = int32(12)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v74)+12)) = v77
								return int32(0)
							} else {
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v69)+136))
								*(*int32)(unsafe.Add(mBase, uint32(v69)+136)) = v81 + v65
								*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = v61
								v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								*(*int32)(unsafe.Add(mBase, uint32(v67))) = v85
								*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v67
								v92 = v67 + int32(8)
								v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v97 + int32(1)
								v101 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+4)) = uint8(v101)
								*(*int32)(unsafe.Add(mBase, uint32(v92))) = v97
								v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								if v104 == v101 {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v92
								} else {
								}
								v108 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v92)+24)) = v108
								*(*int64)(unsafe.Add(mBase, uint32(v92)+16)) = v108
								*(*int64)(unsafe.Add(mBase, uint32(v92)+8)) = v108
								v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								if v114 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v114)+28)) = v92
									v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v118 = v116
								} else {
									v118 = int32(0)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v92)+32)) = v118
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v92
								return v92
							}
						}
					}
				}
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v13 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v14
			v92 = v13
			v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v97 + int32(1)
			v101 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v92)+4)) = uint8(v101)
			*(*int32)(unsafe.Add(mBase, uint32(v92))) = v97
			v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v104 == v101 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v92
			} else {
			}
			v108 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v92)+24)) = v108
			*(*int64)(unsafe.Add(mBase, uint32(v92)+16)) = v108
			*(*int64)(unsafe.Add(mBase, uint32(v92)+8)) = v108
			v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v114 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v114)+28)) = v92
				v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v118 = v116
			} else {
				v118 = int32(0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v92)+32)) = v118
			*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v92
			return v92
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			if v16 != 0 {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
				if base.Ui32(v18) <= base.Ui32(v17) {
					v36 = l0 + int32(76)
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+136))
					if base.Ui32(v38) < base.Ui32(int32(98000000)) {
						v53 = int32(1024)
						v55 = v18 << (uint(int32(1)) % 32)
						if base.Ui32(v53) <= base.Ui32(v55) {
							v58 = v53
						} else {
							v58 = v55
						}
						v60 = v36
						v61 = v58
						v65 = v61*int32(36) + int32(8)
						v67 = F_palloc_extended(m, v65, int32(2))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
							if v67 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(v69)+24)) = int32(101)
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
								if v75 != 0 {
									v77 = v75
								} else {
									v77 = int32(12)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v74)+12)) = v77
								return int32(0)
							} else {
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v69)+136))
								*(*int32)(unsafe.Add(mBase, uint32(v69)+136)) = v81 + v65
								*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = v61
								v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								*(*int32)(unsafe.Add(mBase, uint32(v67))) = v85
								*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v67
								v92 = v67 + int32(8)
								v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v97 + int32(1)
								v101 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+4)) = uint8(v101)
								*(*int32)(unsafe.Add(mBase, uint32(v92))) = v97
								v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								if v104 == v101 {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v92
								} else {
								}
								v108 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v92)+24)) = v108
								*(*int64)(unsafe.Add(mBase, uint32(v92)+16)) = v108
								*(*int64)(unsafe.Add(mBase, uint32(v92)+8)) = v108
								v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								if v114 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v114)+28)) = v92
									v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v118 = v116
								} else {
									v118 = int32(0)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v92)+32)) = v118
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v92
								return v92
							}
						}
					} else {
						v41 = v37
						v42 = v36
						*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = int32(101)
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
						if v47 != 0 {
							v49 = v47
						} else {
							v49 = int32(19)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = v49
						return int32(0)
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v17 + int32(1)
					v92 = v16 + v17*int32(36) + int32(8)
					v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v97 + int32(1)
					v101 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v92)+4)) = uint8(v101)
					*(*int32)(unsafe.Add(mBase, uint32(v92))) = v97
					v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v104 == v101 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v92
					} else {
					}
					v108 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v92)+24)) = v108
					*(*int64)(unsafe.Add(mBase, uint32(v92)+16)) = v108
					*(*int64)(unsafe.Add(mBase, uint32(v92)+8)) = v108
					v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v114 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v114)+28)) = v92
						v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v118 = v116
					} else {
						v118 = int32(0)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v92)+32)) = v118
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v92
					return v92
				}
			} else {
				v29 = l0 + int32(76)
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+136))
				if base.Ui32(int32(97999999)) < base.Ui32(v32) {
					v41 = v31
					v42 = v29
					*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = int32(101)
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
					if v47 != 0 {
						v49 = v47
					} else {
						v49 = int32(19)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = v49
					return int32(0)
				} else {
					v60 = v29
					v61 = int32(32)
					v65 = v61*int32(36) + int32(8)
					v67 = F_palloc_extended(m, v65, int32(2))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
						if v67 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v69)+24)) = int32(101)
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
							if v75 != 0 {
								v77 = v75
							} else {
								v77 = int32(12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v74)+12)) = v77
							return int32(0)
						} else {
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v69)+136))
							*(*int32)(unsafe.Add(mBase, uint32(v69)+136)) = v81 + v65
							*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = v61
							v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v67))) = v85
							*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v67
							v92 = v67 + int32(8)
							v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v97 + int32(1)
							v101 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v92)+4)) = uint8(v101)
							*(*int32)(unsafe.Add(mBase, uint32(v92))) = v97
							v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v104 == v101 {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v92
							} else {
							}
							v108 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v92)+24)) = v108
							*(*int64)(unsafe.Add(mBase, uint32(v92)+16)) = v108
							*(*int64)(unsafe.Add(mBase, uint32(v92)+8)) = v108
							v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							if v114 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v114)+28)) = v92
								v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v118 = v116
							} else {
								v118 = int32(0)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v92)+32)) = v118
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v92
							return v92
						}
					}
				}
			}
		}
	}
}
func F_normal_rand(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 float64
	_ = v2
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 float64
	_ = v35
	var v37 int32
	_ = v37
	var v38 float64
	_ = v38
	var v39 int32
	_ = v39
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 float64
	_ = v58
	var v60 float64
	_ = v60
	var v61 float64
	_ = v61
	var v74 int32
	_ = v74
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v100 float64
	_ = v100
	var v103 float64
	_ = v103
	var v105 int32
	_ = v105
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v131 float64
	_ = v131
	var v134 float64
	_ = v134
	var v136 float64
	_ = v136
	var v141 float64
	_ = v141
	var v145 float64
	_ = v145
	var v149 float64
	_ = v149
	var v150 float64
	_ = v150
	var v156 int64
	_ = v156
	var v159 float64
	_ = v159
	var v169 int64
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	v2 = float64(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v14 == int32(0) {
		v17 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int32(_a_F_normal_rand_0)
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_normal_rand[0]))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
			*(*int32)(unsafe.Add(mBase, _c_F_normal_rand[0])) = v24
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v26 < int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v193 = m.ExcPending
				if v193 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v196 = m.ExcPending
					if v196 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_normal_rand_1), int32(0))
						mBase = m.M
						v200 = m.ExcPending
						if v200 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_normal_rand_2), int32(206), int32(_a_F_normal_rand_3))
							mBase = m.M
							v205 = m.ExcPending
							if v205 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = base.I64_extend_i32_u(v26)
				v32 = F_palloc(m, int32(32))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v35 = *(*float64)(unsafe.Add(mBase, uint32(v34)))
					*(*float64)(unsafe.Add(mBase, uint32(v32))) = v35
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					v38 = *(*float64)(unsafe.Add(mBase, uint32(v37)))
					v39 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v32)+24)) = uint8(v39)
					*(*int64)(unsafe.Add(mBase, uint32(v32)+16)) = int64(0)
					*(*float64)(unsafe.Add(mBase, uint32(v32)+8)) = v38
					*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v32
					*(*int32)(unsafe.Add(mBase, _c_F_normal_rand[0])) = v22
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
					v53 = *(*int64)(unsafe.Add(mBase, uint32(v52)))
					v54 = *(*int64)(unsafe.Add(mBase, uint32(v52)+8))
					if base.Ui64(v53) < base.Ui64(v54) {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
						v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+24)))
						if v57 != 0 {
							v58 = *(*float64)(unsafe.Add(mBase, uint32(v56)+16))
							v159 = v58
							v169 = v53
							v170 = int32(0)
						} else {
							v60 = *(*float64)(unsafe.Add(mBase, uint32(v56)+8))
							v61 = *(*float64)(unsafe.Add(mBase, uint32(v56)))
							for {
								v74 = int32(_a_F_normal_rand_4)
								v77 = *(*int64)(unsafe.Add(mBase, _c_F_normal_rand[1]))
								v78 = *(*int64)(unsafe.Add(mBase, _c_F_normal_rand[2]))
								v79 = v77 ^ v78
								*(*int64)(unsafe.Add(mBase, _c_F_normal_rand[2])) = base.I64_rotl(v79, int64(37))
								*(*int64)(unsafe.Add(mBase, _c_F_normal_rand[1])) = v79<<(uint(int64(16))%64) ^ base.I64_rotl(v77, int64(24)) ^ v79
								v100 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v77*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
								mBase = m.M
								v103 = base.F64_add(base.F64_add(v100, v100), float64(-1))
								v105 = int32(_a_F_normal_rand_4)
								v108 = *(*int64)(unsafe.Add(mBase, _c_F_normal_rand[1]))
								v109 = *(*int64)(unsafe.Add(mBase, _c_F_normal_rand[2]))
								v110 = v108 ^ v109
								*(*int64)(unsafe.Add(mBase, _c_F_normal_rand[2])) = base.I64_rotl(v110, int64(37))
								*(*int64)(unsafe.Add(mBase, _c_F_normal_rand[1])) = v110<<(uint(int64(16))%64) ^ base.I64_rotl(v108, int64(24)) ^ v110
								v131 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v108*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
								mBase = m.M
								v134 = base.F64_add(base.F64_add(v131, v131), float64(-1))
								v136 = base.F64_add(base.F64_mul(v103, v103), base.F64_mul(v134, v134))
								if base.F64_ge(v136, float64(1)) != 0 {
									continue
								} else {
									break
								}
								break
							}
							if base.F64_ne(v136, float64(0)) != 0 {
								v141 = F_log(m, v136)
								mBase = m.M
								v145 = base.F64_sqrt(base.F64_div(base.F64_mul(v141, float64(-2)), v136))
								v149 = base.F64_mul(v103, v145)
								v150 = base.F64_mul(v134, v145)
							} else {
								v149 = v2
								v150 = v2
							}
							*(*float64)(unsafe.Add(mBase, uint32(v56)+16)) = base.F64_add(base.F64_mul(v60, v150), v61)
							v156 = *(*int64)(unsafe.Add(mBase, uint32(v52)))
							v159 = base.F64_add(base.F64_mul(v60, v149), v61)
							v169 = v156
							v170 = int32(1)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v56)+24)) = uint8(v170)
						*(*int64)(unsafe.Add(mBase, uint32(v52))) = v169 + int64(1)
						v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v175)+20)) = int32(1)
						v178 = F_Float8GetDatum(m, v159)
						mBase = m.M
						v179 = m.ExcPending
						if v179 != 0 {
							return int32(0)
						} else {
							return v178
						}
					} else {
						F_end_MultiFuncCall(m, l0)
						mBase = m.M
						v182 = m.ExcPending
						if v182 != 0 {
							return int32(0)
						} else {
							v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v183)+20)) = int32(2)
							v186 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v186)
							return int32(0)
						}
					}
				}
			}
		}
	} else {
		v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
		v53 = *(*int64)(unsafe.Add(mBase, uint32(v52)))
		v54 = *(*int64)(unsafe.Add(mBase, uint32(v52)+8))
		if base.Ui64(v53) < base.Ui64(v54) {
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
			v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+24)))
			if v57 != 0 {
				v58 = *(*float64)(unsafe.Add(mBase, uint32(v56)+16))
				v159 = v58
				v169 = v53
				v170 = int32(0)
			} else {
				v60 = *(*float64)(unsafe.Add(mBase, uint32(v56)+8))
				v61 = *(*float64)(unsafe.Add(mBase, uint32(v56)))
				for {
					v74 = int32(_a_F_normal_rand_4)
					v77 = *(*int64)(unsafe.Add(mBase, _c_F_normal_rand[1]))
					v78 = *(*int64)(unsafe.Add(mBase, _c_F_normal_rand[2]))
					v79 = v77 ^ v78
					*(*int64)(unsafe.Add(mBase, _c_F_normal_rand[2])) = base.I64_rotl(v79, int64(37))
					*(*int64)(unsafe.Add(mBase, _c_F_normal_rand[1])) = v79<<(uint(int64(16))%64) ^ base.I64_rotl(v77, int64(24)) ^ v79
					v100 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v77*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
					mBase = m.M
					v103 = base.F64_add(base.F64_add(v100, v100), float64(-1))
					v105 = int32(_a_F_normal_rand_4)
					v108 = *(*int64)(unsafe.Add(mBase, _c_F_normal_rand[1]))
					v109 = *(*int64)(unsafe.Add(mBase, _c_F_normal_rand[2]))
					v110 = v108 ^ v109
					*(*int64)(unsafe.Add(mBase, _c_F_normal_rand[2])) = base.I64_rotl(v110, int64(37))
					*(*int64)(unsafe.Add(mBase, _c_F_normal_rand[1])) = v110<<(uint(int64(16))%64) ^ base.I64_rotl(v108, int64(24)) ^ v110
					v131 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v108*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
					mBase = m.M
					v134 = base.F64_add(base.F64_add(v131, v131), float64(-1))
					v136 = base.F64_add(base.F64_mul(v103, v103), base.F64_mul(v134, v134))
					if base.F64_ge(v136, float64(1)) != 0 {
						continue
					} else {
						break
					}
					break
				}
				if base.F64_ne(v136, float64(0)) != 0 {
					v141 = F_log(m, v136)
					mBase = m.M
					v145 = base.F64_sqrt(base.F64_div(base.F64_mul(v141, float64(-2)), v136))
					v149 = base.F64_mul(v103, v145)
					v150 = base.F64_mul(v134, v145)
				} else {
					v149 = v2
					v150 = v2
				}
				*(*float64)(unsafe.Add(mBase, uint32(v56)+16)) = base.F64_add(base.F64_mul(v60, v150), v61)
				v156 = *(*int64)(unsafe.Add(mBase, uint32(v52)))
				v159 = base.F64_add(base.F64_mul(v60, v149), v61)
				v169 = v156
				v170 = int32(1)
			}
			*(*uint8)(unsafe.Add(mBase, uint32(v56)+24)) = uint8(v170)
			*(*int64)(unsafe.Add(mBase, uint32(v52))) = v169 + int64(1)
			v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v175)+20)) = int32(1)
			v178 = F_Float8GetDatum(m, v159)
			mBase = m.M
			v179 = m.ExcPending
			if v179 != 0 {
				return int32(0)
			} else {
				return v178
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v182 = m.ExcPending
			if v182 != 0 {
				return int32(0)
			} else {
				v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v183)+20)) = int32(2)
				v186 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v186)
				return int32(0)
			}
		}
	}
}
func F_norwegian_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v253 int32
	_ = v253
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = v11 + int32(3)
	if v9 < v13 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v134
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	if v134 < v137 {
		goto L38
	} else {
		goto L39
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v13
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v11
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v25 < v11 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v65 < int32(0) {
		goto L1
	} else {
		goto L18
	}
L4:
	;
	v27 = v11
	goto L6
L5:
	;
	v27 = v25
	goto L6
L6:
	;
	v34 = v11
	goto L8
L7:
	;
	v65 = v45
	goto L3
L8:
	;
	if v34 == v27 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v65 = int32(-1)
	goto L3
L11:
	;
	goto L12
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v34))))
	if int32(248) < v40 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v57 = v34 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v57
	v34 = v57
	goto L8
L14:
	;
	v42 = v40 - int32(97)
	if v42 < int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v45 = int32(1)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v42)>>(uint(int32(3))%32)))+uint32(_c_F_norwegian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v49)>>(uint(v42&int32(7))%32))&v45 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	goto L13
L18:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v77 < v76 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v120 < int32(0) {
		goto L1
	} else {
		goto L33
	}
L20:
	;
	v79 = v76
	goto L22
L21:
	;
	v79 = v77
	goto L22
L22:
	;
	v85 = v76
	goto L24
L23:
	;
	v120 = int32(1)
	goto L19
L24:
	;
	if v85 == v79 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v120 = int32(-1)
	goto L19
L27:
	;
	goto L28
L28:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+v85))))
	if int32(248) < v94 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	v96 = v94 - int32(97)
	if v96 < int32(0) {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v96)>>(uint(int32(3))%32)))+uint32(_c_F_norwegian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v102)>>(uint(v96&int32(7))%32))&int32(1) == int32(0) {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	v111 = v85 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v111
	v85 = v111
	goto L24
L33:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v124 = v123 + v120
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v124
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	if v127 < v124 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v129 = v124
	goto L36
L35:
	;
	v129 = v127
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126)+4)) = v129
	goto L1
L37:
	;
	return v391
L38:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v309
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+4))
	if v309 < v313 {
		v347 = v311
		goto L84
	} else {
		goto L85
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v137
	if v134 <= v137 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	goto L38
L41:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v144 = int32(1)
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142+v134-v144))))
	if base.B2i32(v146&int32(224) != int32(96))|base.B2i32(v144<<(uint(v146)%32)&int32(_a_F_norwegian_ISO_8859_1_stem_0) == int32(0)) != 0 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v160 = F_find_among_b(m, l0, int32(_a_F_norwegian_ISO_8859_1_stem_1), int32(29))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	return int32(0)
L44:
	;
	if v160 == int32(0) {
		goto L40
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v167
	switch v160 - int32(1) {
	case 0:
		goto L48
	case 1:
		goto L47
	case 2:
		goto L46
	default:
		goto L38
	}
L46:
	;
	v301 = F_slice_from_s(m, l0, int32(2), int32(_a_F_norwegian_ISO_8859_1_stem_2))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L43
	} else {
		goto L82
	}
L47:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L53
L48:
	;
	v171 = F_slice_del(m, l0)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L43
	} else {
		goto L49
	}
L49:
	;
	if int32(0) <= v171 {
		goto L38
	} else {
		goto L50
	}
L50:
	;
	v391 = v171
	goto L37
L51:
	;
	if v228 != 0 {
		goto L62
	} else {
		goto L63
	}
L52:
	;
	v228 = v224
	goto L51
L53:
	;
	if v184 <= v185 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v224 = int32(0)
	goto L52
L55:
	;
	v228 = int32(-1)
	goto L51
L56:
	;
	goto L57
L57:
	;
	v197 = int32(1)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198+v184-v197))))
	if int32(122) < v202 {
		v224 = v197
		goto L52
	} else {
		goto L58
	}
L58:
	;
	v204 = v202 - int32(98)
	if v204 < int32(0) {
		v224 = v197
		goto L52
	} else {
		goto L59
	}
L59:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v204)>>(uint(int32(3))%32)))+uint32(_c_F_norwegian_ISO_8859_1_stem[1]))))
	if int32(base.Ui32(v210)>>(uint(v204&int32(7))%32))&int32(1) == int32(0) {
		v224 = v197
		goto L52
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v184 - int32(1)
	goto L61
L61:
	;
	goto L54
L62:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v231 = v229 + (v167 - v175)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v231
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v231 <= v233 {
		goto L38
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v295 = F_slice_del(m, l0)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L43
	} else {
		goto L80
	}
L65:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235+v231-int32(1)))))
	if v239 != int32(107) {
		goto L38
	} else {
		goto L66
	}
L66:
	;
	v243 = v231 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v243
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L69
L67:
	;
	if v293 != 0 {
		goto L38
	} else {
		goto L79
	}
L68:
	;
	v293 = v290
	goto L67
L69:
	;
	if v243 <= v253 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v290 = int32(0)
	goto L68
L71:
	;
	v293 = int32(-1)
	goto L67
L72:
	;
	goto L73
L73:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264+v243-int32(1)))))
	if int32(248) < v268 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v243 - int32(1)
	goto L78
L75:
	;
	v270 = v268 - int32(97)
	if v270 < int32(0) {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v273 = int32(1)
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v270)>>(uint(int32(3))%32)))+uint32(_c_F_norwegian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v277)>>(uint(v270&int32(7))%32))&v273 != 0 {
		v290 = v273
		goto L68
	} else {
		goto L77
	}
L77:
	;
	goto L74
L78:
	;
	goto L70
L79:
	;
	goto L64
L80:
	;
	if int32(0) <= v295 {
		goto L38
	} else {
		goto L81
	}
L81:
	;
	v391 = v295
	goto L37
L82:
	;
	if int32(0) <= v301 {
		goto L38
	} else {
		goto L83
	}
L83:
	;
	v391 = v301
	goto L37
L84:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v349
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v351)+4))
	if v349 < v352 {
		v387 = v347
		goto L98
	} else {
		goto L99
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v309
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v313
	v318 = v309 - int32(1)
	if v313 < v318 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v328 = F_find_among_b(m, l0, int32(_a_F_norwegian_ISO_8859_1_stem_3), int32(2))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L43
	} else {
		goto L91
	}
L87:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320+v318))))
	if v322 == int32(116) {
		goto L86
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v311
	v347 = v311
	goto L84
L90:
	;
	goto L89
L91:
	;
	if v328 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v311
	v347 = v311
	goto L84
L93:
	;
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v311
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v334
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v336 <= v311 {
		v347 = v311
		goto L84
	} else {
		goto L95
	}
L95:
	;
	v339 = v336 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v339
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v339
	v342 = F_slice_del(m, l0)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L43
	} else {
		goto L96
	}
L96:
	;
	if v342 < int32(0) {
		v391 = v342
		goto L37
	} else {
		goto L97
	}
L97:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v347 = v346
	goto L84
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v387
	v391 = int32(1)
	goto L37
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v349
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v352
	v357 = v349 - int32(1)
	if v357 <= v352 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v374 = F_find_among_b(m, l0, int32(_a_F_norwegian_ISO_8859_1_stem_4), int32(11))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L43
	} else {
		goto L105
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v347
	v387 = v347
	goto L98
L102:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359+v357))))
	if v361&int32(224) != int32(96) {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	if int32(1)<<(uint(v361)%32)&int32(_a_F_norwegian_ISO_8859_1_stem_5) != 0 {
		goto L100
	} else {
		goto L104
	}
L104:
	;
	goto L101
L105:
	;
	if v374 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v347
	v387 = v347
	goto L98
L107:
	;
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v347
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v380
	v382 = F_slice_del(m, l0)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L43
	} else {
		goto L109
	}
L109:
	;
	if v382 < int32(0) {
		v391 = v382
		goto L37
	} else {
		goto L110
	}
L110:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v387 = v386
	goto L98
}
func F_norwegian_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v140 int32
	_ = v140
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v184 int32
	_ = v184
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v259 int32
	_ = v259
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v465 int32
	_ = v465
	var v481 int32
	_ = v481
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v610 int32
	_ = v610
	var v627 int32
	_ = v627
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v699 int32
	_ = v699
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v729 int32
	_ = v729
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	goto L4
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v5
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v318
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)+4))
	if v318 < v321 {
		goto L78
	} else {
		goto L79
	}
L2:
	;
	if v62 < int32(0) {
		goto L1
	} else {
		goto L22
	}
L4:
	;
	goto L5
L5:
	;
	goto L6
L6:
	;
	v17 = v10
	v19 = int32(3)
	goto L9
L8:
	;
	v62 = v47
	goto L2
L9:
	;
	if v7 <= v17 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v62 = int32(-1)
	goto L2
L12:
	;
	goto L13
L13:
	;
	v24 = v17 + int32(1)
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v17))))
	if base.Ui32(v26) < base.Ui32(int32(192)) {
		v47 = v24
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v48 = int32(1)
	if v48 < v19 {
		v17 = v47
		v19 = v19 - v48
		goto L9
	} else {
		goto L21
	}
L15:
	;
	if v7 <= v24 {
		v47 = v24
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v33 = v24
	goto L17
L17:
	;
	v36 = int32(*(*int8)(unsafe.Add(mBase, uint32(v9+v33))))
	if int32(-65) < v36 {
		v47 = v33
		goto L14
	} else {
		goto L19
	}
L18:
	;
	v47 = v7
	goto L14
L19:
	;
	v40 = v33 + int32(1)
	if v40 != v7 {
		v33 = v40
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	goto L10
L22:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v62
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v89 = v10
	goto L25
L23:
	;
	if v184 < int32(0) {
		goto L1
	} else {
		goto L48
	}
L24:
	;
	v184 = v156
	goto L23
L25:
	;
	if v80 <= v89 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v184 = int32(-1)
	goto L23
L28:
	;
	goto L29
L29:
	;
	v96 = int32(1)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89+v81))))
	if base.Ui32(v98) < base.Ui32(int32(192)) {
		v155 = v98
		v156 = v96
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if int32(248) < v155 {
		goto L43
	} else {
		goto L44
	}
L31:
	;
	v102 = v89 + int32(1)
	if v102 == v80 {
		v155 = v98
		v156 = v96
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+v81))))
	v107 = v105 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v98) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111+v81))))
	v123 = v121 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v98) {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	v111 = v89 + int32(2)
	if v111 != v80 {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v155 = v98<<(uint(int32(6))%32)&int32(1984) | v107
	v156 = int32(2)
	goto L30
L37:
	;
	goto L36
L38:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v127))))
	v155 = v140&int32(63) | (v98<<(uint(int32(18))%32)&int32(_a_F_norwegian_UTF_8_stem_0) | v107<<(uint(int32(12))%32) | v123<<(uint(int32(6))%32))
	v156 = int32(4)
	goto L30
L39:
	;
	v127 = v89 + int32(3)
	if v127 != v80 {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v155 = v98<<(uint(int32(12))%32)&int32(_a_F_norwegian_UTF_8_stem_1) | v107<<(uint(int32(6))%32) | v123
	v156 = int32(3)
	goto L30
L42:
	;
	goto L41
L43:
	;
	v173 = v156 + v89
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v173
	v89 = v173
	goto L25
L44:
	;
	v160 = v155 - int32(97)
	if v160 < int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v160)>>(uint(int32(3))%32)))+uint32(_c_F_norwegian_UTF_8_stem[0]))))
	if int32(base.Ui32(v166)>>(uint(v160&int32(7))%32))&int32(1) != 0 {
		goto L24
	} else {
		goto L46
	}
L46:
	;
	goto L43
L48:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v208 = v198
	goto L51
L49:
	;
	if v304 < int32(0) {
		goto L1
	} else {
		goto L73
	}
L50:
	;
	v304 = v275
	goto L49
L51:
	;
	if v199 <= v208 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v304 = int32(-1)
	goto L49
L54:
	;
	goto L55
L55:
	;
	v215 = int32(1)
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208+v200))))
	if base.Ui32(v217) < base.Ui32(int32(192)) {
		v274 = v217
		v275 = v215
		goto L56
	} else {
		goto L57
	}
L56:
	;
	if int32(248) < v274 {
		goto L50
	} else {
		goto L69
	}
L57:
	;
	v221 = v208 + int32(1)
	if v221 == v199 {
		v274 = v217
		v275 = v215
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221+v200))))
	v226 = v224 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v217) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230+v200))))
	v242 = v240 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v217) {
		goto L65
	} else {
		goto L66
	}
L60:
	;
	v230 = v208 + int32(2)
	if v230 != v199 {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v274 = v217<<(uint(int32(6))%32)&int32(1984) | v226
	v275 = int32(2)
	goto L56
L63:
	;
	goto L62
L64:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200+v246))))
	v274 = v259&int32(63) | (v217<<(uint(int32(18))%32)&int32(_a_F_norwegian_UTF_8_stem_0) | v226<<(uint(int32(12))%32) | v242<<(uint(int32(6))%32))
	v275 = int32(4)
	goto L56
L65:
	;
	v246 = v208 + int32(3)
	if v246 != v199 {
		goto L64
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v274 = v217<<(uint(int32(12))%32)&int32(_a_F_norwegian_UTF_8_stem_1) | v226<<(uint(int32(6))%32) | v242
	v275 = int32(3)
	goto L56
L68:
	;
	goto L67
L69:
	;
	v279 = v274 - int32(97)
	if v279 < int32(0) {
		goto L50
	} else {
		goto L70
	}
L70:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v279)>>(uint(int32(3))%32)))+uint32(_c_F_norwegian_UTF_8_stem[0]))))
	if int32(base.Ui32(v285)>>(uint(v279&int32(7))%32))&int32(1) == int32(0) {
		goto L50
	} else {
		goto L71
	}
L71:
	;
	v293 = v275 + v208
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v293
	v208 = v293
	goto L51
L73:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v308 = v307 + v304
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v308
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
	if v311 < v308 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v313 = v308
	goto L76
L75:
	;
	v313 = v311
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v310)+4)) = v313
	goto L1
L77:
	;
	return v784
L78:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v650
	v652 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v652)+4))
	if v650 < v653 {
		goto L142
	} else {
		goto L143
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v318
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v321
	if v318 <= v321 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v5
	goto L78
L81:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v328 = int32(1)
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326+v318-v328))))
	if base.B2i32(v330&int32(224) != int32(96))|base.B2i32(v328<<(uint(v330)%32)&int32(_a_F_norwegian_UTF_8_stem_2) == int32(0)) != 0 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v344 = F_find_among_b(m, l0, int32(_a_F_norwegian_UTF_8_stem_3), int32(29))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	return int32(0)
L84:
	;
	if v344 == int32(0) {
		goto L80
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v5
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v351
	switch v344 - int32(1) {
	case 0:
		goto L88
	case 1:
		goto L87
	case 2:
		goto L86
	default:
		goto L78
	}
L86:
	;
	v642 = F_slice_from_s(m, l0, int32(2), int32(_a_F_norwegian_UTF_8_stem_4))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L83
	} else {
		goto L140
	}
L87:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L93
L88:
	;
	v355 = F_slice_del(m, l0)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L83
	} else {
		goto L89
	}
L89:
	;
	if int32(0) <= v355 {
		goto L78
	} else {
		goto L90
	}
L90:
	;
	v784 = v355
	goto L77
L91:
	;
	if v488 != 0 {
		goto L114
	} else {
		goto L115
	}
L92:
	;
	v488 = v481
	goto L91
L93:
	;
	if v372 <= v373 {
		v481 = int32(-1)
		goto L92
	} else {
		goto L95
	}
L94:
	;
	v481 = int32(0)
	goto L92
L95:
	;
	v390 = int32(1)
	v391 = v372 - v390
	v393 = int32(*(*int8)(unsafe.Add(mBase, uint32(v374+v391))))
	v395 = v393 & int32(255)
	if base.B2i32(v391 == v373)|base.B2i32(int32(0) <= v393) != 0 {
		v453 = v395
		v457 = v390
		goto L96
	} else {
		goto L97
	}
L96:
	;
	if int32(122) < v453 {
		goto L104
	} else {
		goto L105
	}
L97:
	;
	v402 = v395 & int32(63)
	v404 = v372 - int32(2)
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374+v404))))
	v408 = v406 << (uint(int32(6)) % 32)
	if base.B2i32(v404 != v373)&base.B2i32(base.Ui32(v406) < base.Ui32(int32(192))) == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v453 = v408&int32(1984) | v402
	v457 = int32(2)
	goto L96
L99:
	;
	goto L100
L100:
	;
	v421 = v408&int32(4032) | v402
	v423 = v372 - int32(3)
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374+v423))))
	if base.B2i32(v423 != v373)&base.B2i32(base.Ui32(v425) < base.Ui32(int32(224))) == int32(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v453 = v425<<(uint(int32(12))%32)&int32(_a_F_norwegian_UTF_8_stem_1) | v421
	v457 = int32(3)
	goto L96
L102:
	;
	goto L103
L103:
	;
	v443 = int32(4)
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372+v374-v443))))
	v453 = v425<<(uint(int32(12))%32)&int32(_a_F_norwegian_UTF_8_stem_5) | v445&int32(7)<<(uint(int32(18))%32) | v421
	v457 = v443
	goto L96
L104:
	;
	v488 = v457
	goto L91
L105:
	;
	goto L106
L106:
	;
	v459 = v453 - int32(98)
	if v459 < int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v488 = v457
	goto L91
L108:
	;
	goto L109
L109:
	;
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v459)>>(uint(int32(3))%32)))+uint32(_c_F_norwegian_UTF_8_stem[1]))))
	if int32(base.Ui32(v465)>>(uint(v459&int32(7))%32))&int32(1) == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v488 = v457
	goto L91
L111:
	;
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v372 - v457
	goto L113
L113:
	;
	goto L94
L114:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v491 = v489 + (v351 - v359)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v491
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v491 <= v493 {
		goto L78
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v636 = F_slice_del(m, l0)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L83
	} else {
		goto L138
	}
L117:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495+v491-int32(1)))))
	if v499 != int32(107) {
		goto L78
	} else {
		goto L118
	}
L118:
	;
	v503 = v491 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v503
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L121
L119:
	;
	if v634 != 0 {
		goto L78
	} else {
		goto L137
	}
L120:
	;
	v634 = v627
	goto L119
L121:
	;
	if v503 <= v518 {
		v627 = int32(-1)
		goto L120
	} else {
		goto L123
	}
L122:
	;
	v627 = int32(0)
	goto L120
L123:
	;
	v535 = int32(1)
	v536 = v503 - v535
	v538 = int32(*(*int8)(unsafe.Add(mBase, uint32(v519+v536))))
	v540 = v538 & int32(255)
	if base.B2i32(v536 == v518)|base.B2i32(int32(0) <= v538) != 0 {
		v598 = v540
		v602 = v535
		goto L124
	} else {
		goto L125
	}
L124:
	;
	if int32(248) < v598 {
		goto L132
	} else {
		goto L133
	}
L125:
	;
	v547 = v540 & int32(63)
	v549 = v503 - int32(2)
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519+v549))))
	v553 = v551 << (uint(int32(6)) % 32)
	if base.B2i32(v549 != v518)&base.B2i32(base.Ui32(v551) < base.Ui32(int32(192))) == int32(0) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v598 = v553&int32(1984) | v547
	v602 = int32(2)
	goto L124
L127:
	;
	goto L128
L128:
	;
	v566 = v553&int32(4032) | v547
	v568 = v503 - int32(3)
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519+v568))))
	if base.B2i32(v568 != v518)&base.B2i32(base.Ui32(v570) < base.Ui32(int32(224))) == int32(0) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v598 = v570<<(uint(int32(12))%32)&int32(_a_F_norwegian_UTF_8_stem_1) | v566
	v602 = int32(3)
	goto L124
L130:
	;
	goto L131
L131:
	;
	v588 = int32(4)
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503+v519-v588))))
	v598 = v570<<(uint(int32(12))%32)&int32(_a_F_norwegian_UTF_8_stem_5) | v590&int32(7)<<(uint(int32(18))%32) | v566
	v602 = v588
	goto L124
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v503 - v602
	goto L136
L133:
	;
	v604 = v598 - int32(97)
	if v604 < int32(0) {
		goto L132
	} else {
		goto L134
	}
L134:
	;
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v604)>>(uint(int32(3))%32)))+uint32(_c_F_norwegian_UTF_8_stem[0]))))
	if int32(base.Ui32(v610)>>(uint(v604&int32(7))%32))&int32(1) == int32(0) {
		goto L132
	} else {
		goto L135
	}
L135:
	;
	v634 = v602
	goto L119
L136:
	;
	goto L122
L137:
	;
	goto L116
L138:
	;
	if int32(0) <= v636 {
		goto L78
	} else {
		goto L139
	}
L139:
	;
	v784 = v636
	goto L77
L140:
	;
	if int32(0) <= v642 {
		goto L78
	} else {
		goto L141
	}
L141:
	;
	v784 = v642
	goto L77
L142:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v741
	v743 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v744 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v744)+4))
	if v741 < v745 {
		v780 = v743
		goto L171
	} else {
		goto L172
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v650
	v656 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v653
	v659 = v650 - int32(1)
	if v659 <= v653 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v656
	goto L142
L145:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v661+v659))))
	if v663 != int32(116) {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v668 = F_find_among_b(m, l0, int32(_a_F_norwegian_UTF_8_stem_6), int32(2))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L83
	} else {
		goto L147
	}
L147:
	;
	if v668 == int32(0) {
		goto L144
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v656
	v673 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v673
	v675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v675
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L151
L149:
	;
	if v729 < int32(0) {
		goto L142
	} else {
		goto L168
	}
L151:
	;
	goto L152
L152:
	;
	goto L153
L153:
	;
	v684 = v675
	v686 = int32(1)
	goto L156
L155:
	;
	v729 = v711
	goto L149
L156:
	;
	if v684 <= v656 {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	goto L155
L158:
	;
	v729 = int32(-1)
	goto L149
L159:
	;
	goto L160
L160:
	;
	v691 = v684 - int32(1)
	v693 = int32(*(*int8)(unsafe.Add(mBase, uint32(v677+v691))))
	if base.B2i32(int32(0) <= v693)|base.B2i32(v691 <= v656) != 0 {
		v711 = v691
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v715 = int32(1)
	if v715 < v686 {
		v684 = v711
		v686 = v686 - v715
		goto L156
	} else {
		goto L167
	}
L162:
	;
	v699 = v691
	goto L163
L163:
	;
	v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677+v699))))
	if base.Ui32(int32(191)) < base.Ui32(v704) {
		v711 = v699
		goto L161
	} else {
		goto L165
	}
L164:
	;
	v711 = v656
	goto L161
L165:
	;
	v708 = v699 - int32(1)
	if v656 < v708 {
		v699 = v708
		goto L163
	} else {
		goto L166
	}
L166:
	;
	goto L164
L167:
	;
	goto L157
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v729
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v729
	v734 = F_slice_del(m, l0)
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L83
	} else {
		goto L169
	}
L169:
	;
	if int32(0) <= v734 {
		goto L142
	} else {
		goto L170
	}
L170:
	;
	v784 = v734
	goto L77
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v780
	v784 = int32(1)
	goto L77
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v741
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v745
	v750 = v741 - int32(1)
	if v750 <= v745 {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v767 = F_find_among_b(m, l0, int32(_a_F_norwegian_UTF_8_stem_7), int32(11))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L83
	} else {
		goto L178
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v743
	v780 = v743
	goto L171
L175:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v752+v750))))
	if v754&int32(224) != int32(96) {
		goto L174
	} else {
		goto L176
	}
L176:
	;
	if int32(1)<<(uint(v754)%32)&int32(_a_F_norwegian_UTF_8_stem_8) != 0 {
		goto L173
	} else {
		goto L177
	}
L177:
	;
	goto L174
L178:
	;
	if v767 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v743
	v780 = v743
	goto L171
L180:
	;
	goto L181
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v743
	v773 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v773
	v775 = F_slice_del(m, l0)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L83
	} else {
		goto L182
	}
L182:
	;
	if v775 < int32(0) {
		v784 = v775
		goto L77
	} else {
		goto L183
	}
L183:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v780 = v779
	goto L171
}
func F_numerictypmodin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v16 = F_ArrayGetIntegerTypmods(m, v10, v7+int32(44))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
			switch v18 - int32(1) {
			case 0:
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				if base.Ui32(v38-int32(1001)) <= base.Ui32(int32(-1001)) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return int32(0)
						} else {
							v101 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
							*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = int32(1000)
							*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v101
							F_errmsg(m, int32(_a_F_numerictypmodin_0), v7+int32(32))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_numerictypmodin_1), int32(1353), int32(_a_F_numerictypmodin_2))
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v49 = v38<<(uint(int32(16))%32) | int32(4)
					m.G0 = v7 + int32(48)
					return v49
				}
			case 1:
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				if base.Ui32(v21-int32(1001)) <= base.Ui32(int32(-1001)) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(1000)
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v61
							F_errmsg(m, int32(_a_F_numerictypmodin_0), v7)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_numerictypmodin_1), int32(1339), int32(_a_F_numerictypmodin_2))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
					if base.Ui32(v26-int32(1001)) <= base.Ui32(int32(-2002)) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int32(0)
							} else {
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
								*(*int64)(unsafe.Add(mBase, uint32(v7)+20)) = int64(4299262262296)
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v80
								F_errmsg(m, int32(_a_F_numerictypmodin_3), v7+int32(16))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_numerictypmodin_1), int32(1344), int32(_a_F_numerictypmodin_2))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v49 = v26&int32(2047) | v21<<(uint(int32(16))%32) + int32(4)
						m.G0 = v7 + int32(48)
						return v49
					}
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v118 = m.ExcPending
				if v118 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_numerictypmodin_4), int32(0))
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_numerictypmodin_1), int32(1361), int32(_a_F_numerictypmodin_2))
							mBase = m.M
							v130 = m.ExcPending
							if v130 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_numericvar_to_double_no_overflow(m *base.Module, l0 int32) float64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 float64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = F_get_str_from_var(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return float64(0)
	} else {
		v14 = F_strtod(m, v8, v6+int32(12))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return float64(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
			if v17 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return float64(0)
				} else {
					F_errcode(m, int32(33685634))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return float64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v8
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_numericvar_to_double_no_overflow_0)
						F_errmsg(m, int32(_a_F_numericvar_to_double_no_overflow_1), v6)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return float64(0)
						} else {
							F_errfinish(m, int32(_a_F_numericvar_to_double_no_overflow_2), int32(_a_F_numericvar_to_double_no_overflow_3), int32(_a_F_numericvar_to_double_no_overflow_4))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return float64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				F_pfree(m, v8)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return float64(0)
				} else {
					m.G0 = v6 + int32(16)
					return v14
				}
			}
		}
	}
}
func F_numrange_subdiff(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_DirectFunctionCall2Coll(m, int32(18), v3, v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_DirectFunctionCall1Coll(m, int32(17), v3, v8)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
			v15 = F_Float8GetDatum(m, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
