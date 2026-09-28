package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_nameeq(m *base.Module, l0 int32) int64 {
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 == int32(950) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return base.I64_extend_i32_u(base.B2i32(v61 == int32(0)))
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
	return int64(0)
L19:
	;
	v61 = v57
	goto L1
}
func F_nameeqfast(m *base.Module, l0 int64, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	v3 = base.I32_wrap_i64(l0)
	v4 = base.I32_wrap_i64(l1)
	goto L3
L1:
	;
	return base.B2i32(v42-v43 == int32(0))
L3:
	;
	goto L4
L4:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v11 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v12 = v3
	v13 = v4
	v14 = int32(64)
	v15 = v11
	goto L9
L6:
	;
	v38 = v4
	v42 = int32(0)
	goto L7
L7:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	goto L1
L8:
	;
	v38 = v33
	v42 = v35
	goto L7
L9:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if base.B2i32(v15 != v17)|base.B2i32(v17 == int32(0)) != 0 {
		v33 = v13
		v35 = v15
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v33 = v27
	v35 = int32(0)
	goto L8
L11:
	;
	v23 = v14 - int32(1)
	if v23 == int32(0) {
		v33 = v13
		v35 = v15
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v26 = int32(1)
	v27 = v13 + v26
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v28 != 0 {
		v12 = v12 + v26
		v13 = v27
		v14 = v23
		v15 = v28
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
}
func F_namege(m *base.Module, l0 int32) int64 {
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 == int32(950) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return base.I64_extend_i32_u(base.B2i32(int32(0) <= v61))
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
	return int64(0)
L19:
	;
	v61 = v57
	goto L1
}
func F_nameicregexne(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14331(m, l0, int32(27))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_namenetext(m *base.Module, l0 int32) int64 {
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
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	return int64(0)
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
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	F_errfinish(m, int32(_a_F_namenetext_2), int32(1337), int32(_a_F_namenetext_3))
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
	return base.I64_extend_i32_u(v150)
L54:
	;
	goto L53
}
func F_nameregexne(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14331(m, l0, int32(19))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_namesend(m *base.Module, l0 int32) int64 {
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
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_pq_begintypsend(m, v5)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = F_strlen(m, v7)
		mBase = m.M
		F_pq_sendtext(m, v5, v7, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = v17 << (uint(int32(2)) % 32)
			m.G0 = v5 + int32(16)
			return base.I64_extend_i32_u(v16)
		}
	}
}
func F_ndistinct_array_start(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14255(m, l0, int32(_a_F_ndistinct_array_start_0), int32(226), int32(_a_F_ndistinct_array_start_1), int32(_a_F_ndistinct_array_start_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
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
func F_normal_rand(m *base.Module, l0 int32) int64 {
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
	var v26 int64
	_ = v26
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 float64
	_ = v37
	var v39 float64
	_ = v39
	var v40 int32
	_ = v40
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 float64
	_ = v60
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v76 int32
	_ = v76
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v81 int64
	_ = v81
	var v102 float64
	_ = v102
	var v105 float64
	_ = v105
	var v107 int32
	_ = v107
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v133 float64
	_ = v133
	var v136 float64
	_ = v136
	var v138 float64
	_ = v138
	var v143 float64
	_ = v143
	var v147 float64
	_ = v147
	var v151 float64
	_ = v151
	var v152 float64
	_ = v152
	var v158 int64
	_ = v158
	var v161 float64
	_ = v161
	var v171 int64
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	v2 = float64(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v14 == int32(0) {
		v17 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v21 = int32(_a_F_normal_rand_0)
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_normal_rand[0]))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
			*(*int32)(unsafe.Add(mBase, _c_F_normal_rand[0])) = v24
			v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
			if v26&int64(2147483648) != int64(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v194 = m.ExcPending
				if v194 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v197 = m.ExcPending
					if v197 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_normal_rand_1), int32(0))
						mBase = m.M
						v201 = m.ExcPending
						if v201 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_normal_rand_2), int32(208), int32(_a_F_normal_rand_3))
							mBase = m.M
							v206 = m.ExcPending
							if v206 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v26 & int64(2147483647)
				v35 = F_palloc(m, int32(32))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int64(0)
				} else {
					v37 = *(*float64)(unsafe.Add(mBase, uint32(l0)+40))
					*(*float64)(unsafe.Add(mBase, uint32(v35))) = v37
					v39 = *(*float64)(unsafe.Add(mBase, uint32(l0)+56))
					v40 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v35)+24)) = uint8(v40)
					*(*int64)(unsafe.Add(mBase, uint32(v35)+16)) = int64(0)
					*(*float64)(unsafe.Add(mBase, uint32(v35)+8)) = v39
					*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v35
					*(*int32)(unsafe.Add(mBase, _c_F_normal_rand[0])) = v22
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
					v55 = *(*int64)(unsafe.Add(mBase, uint32(v54)))
					v56 = *(*int64)(unsafe.Add(mBase, uint32(v54)+8))
					if base.Ui64(v55) < base.Ui64(v56) {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
						v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+24)))
						if v59 != 0 {
							v60 = *(*float64)(unsafe.Add(mBase, uint32(v58)+16))
							v161 = v60
							v171 = v55
							v172 = int32(0)
						} else {
							v62 = *(*float64)(unsafe.Add(mBase, uint32(v58)+8))
							v63 = *(*float64)(unsafe.Add(mBase, uint32(v58)))
							for {
								v76 = int32(_a_F_normal_rand_4)
								v79 = *(*int64)(unsafe.Add(mBase, _c_F_normal_rand[1]))
								v80 = *(*int64)(unsafe.Add(mBase, _c_F_normal_rand[2]))
								v81 = v79 ^ v80
								*(*int64)(unsafe.Add(mBase, _c_F_normal_rand[2])) = base.I64_rotl(v81, int64(37))
								*(*int64)(unsafe.Add(mBase, _c_F_normal_rand[1])) = v81<<(uint(int64(16))%64) ^ base.I64_rotl(v79, int64(24)) ^ v81
								v102 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v79*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
								mBase = m.M
								v105 = base.F64_add(base.F64_add(v102, v102), float64(-1))
								v107 = int32(_a_F_normal_rand_4)
								v110 = *(*int64)(unsafe.Add(mBase, _c_F_normal_rand[1]))
								v111 = *(*int64)(unsafe.Add(mBase, _c_F_normal_rand[2]))
								v112 = v110 ^ v111
								*(*int64)(unsafe.Add(mBase, _c_F_normal_rand[2])) = base.I64_rotl(v112, int64(37))
								*(*int64)(unsafe.Add(mBase, _c_F_normal_rand[1])) = v112<<(uint(int64(16))%64) ^ base.I64_rotl(v110, int64(24)) ^ v112
								v133 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v110*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
								mBase = m.M
								v136 = base.F64_add(base.F64_add(v133, v133), float64(-1))
								v138 = base.F64_add(base.F64_mul(v105, v105), base.F64_mul(v136, v136))
								if base.F64_ge(v138, float64(1)) != 0 {
									continue
								} else {
									break
								}
								break
							}
							if base.F64_ne(v138, float64(0)) != 0 {
								v143 = F_log(m, v138)
								mBase = m.M
								v147 = base.F64_sqrt(base.F64_div(base.F64_mul(v143, float64(-2)), v138))
								v151 = base.F64_mul(v105, v147)
								v152 = base.F64_mul(v136, v147)
							} else {
								v151 = v2
								v152 = v2
							}
							*(*float64)(unsafe.Add(mBase, uint32(v58)+16)) = base.F64_add(base.F64_mul(v62, v152), v63)
							v158 = *(*int64)(unsafe.Add(mBase, uint32(v54)))
							v161 = base.F64_add(base.F64_mul(v62, v151), v63)
							v171 = v158
							v172 = int32(1)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v58)+24)) = uint8(v172)
						*(*int64)(unsafe.Add(mBase, uint32(v54))) = v171 + int64(1)
						v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v177)+20)) = int32(1)
						return base.I64_reinterpret_f64(v161)
					} else {
						F_end_MultiFuncCall(m, l0)
						mBase = m.M
						v183 = m.ExcPending
						if v183 != 0 {
							return int64(0)
						} else {
							v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v184)+20)) = int32(2)
							v187 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v187)
							return int64(0)
						}
					}
				}
			}
		}
	} else {
		v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
		v55 = *(*int64)(unsafe.Add(mBase, uint32(v54)))
		v56 = *(*int64)(unsafe.Add(mBase, uint32(v54)+8))
		if base.Ui64(v55) < base.Ui64(v56) {
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
			v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+24)))
			if v59 != 0 {
				v60 = *(*float64)(unsafe.Add(mBase, uint32(v58)+16))
				v161 = v60
				v171 = v55
				v172 = int32(0)
			} else {
				v62 = *(*float64)(unsafe.Add(mBase, uint32(v58)+8))
				v63 = *(*float64)(unsafe.Add(mBase, uint32(v58)))
				for {
					v76 = int32(_a_F_normal_rand_4)
					v79 = *(*int64)(unsafe.Add(mBase, _c_F_normal_rand[1]))
					v80 = *(*int64)(unsafe.Add(mBase, _c_F_normal_rand[2]))
					v81 = v79 ^ v80
					*(*int64)(unsafe.Add(mBase, _c_F_normal_rand[2])) = base.I64_rotl(v81, int64(37))
					*(*int64)(unsafe.Add(mBase, _c_F_normal_rand[1])) = v81<<(uint(int64(16))%64) ^ base.I64_rotl(v79, int64(24)) ^ v81
					v102 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v79*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
					mBase = m.M
					v105 = base.F64_add(base.F64_add(v102, v102), float64(-1))
					v107 = int32(_a_F_normal_rand_4)
					v110 = *(*int64)(unsafe.Add(mBase, _c_F_normal_rand[1]))
					v111 = *(*int64)(unsafe.Add(mBase, _c_F_normal_rand[2]))
					v112 = v110 ^ v111
					*(*int64)(unsafe.Add(mBase, _c_F_normal_rand[2])) = base.I64_rotl(v112, int64(37))
					*(*int64)(unsafe.Add(mBase, _c_F_normal_rand[1])) = v112<<(uint(int64(16))%64) ^ base.I64_rotl(v110, int64(24)) ^ v112
					v133 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v110*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
					mBase = m.M
					v136 = base.F64_add(base.F64_add(v133, v133), float64(-1))
					v138 = base.F64_add(base.F64_mul(v105, v105), base.F64_mul(v136, v136))
					if base.F64_ge(v138, float64(1)) != 0 {
						continue
					} else {
						break
					}
					break
				}
				if base.F64_ne(v138, float64(0)) != 0 {
					v143 = F_log(m, v138)
					mBase = m.M
					v147 = base.F64_sqrt(base.F64_div(base.F64_mul(v143, float64(-2)), v138))
					v151 = base.F64_mul(v105, v147)
					v152 = base.F64_mul(v136, v147)
				} else {
					v151 = v2
					v152 = v2
				}
				*(*float64)(unsafe.Add(mBase, uint32(v58)+16)) = base.F64_add(base.F64_mul(v62, v152), v63)
				v158 = *(*int64)(unsafe.Add(mBase, uint32(v54)))
				v161 = base.F64_add(base.F64_mul(v62, v151), v63)
				v171 = v158
				v172 = int32(1)
			}
			*(*uint8)(unsafe.Add(mBase, uint32(v58)+24)) = uint8(v172)
			*(*int64)(unsafe.Add(mBase, uint32(v54))) = v171 + int64(1)
			v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v177)+20)) = int32(1)
			return base.I64_reinterpret_f64(v161)
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v183 = m.ExcPending
			if v183 != 0 {
				return int64(0)
			} else {
				v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v184)+20)) = int32(2)
				v187 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v187)
				return int64(0)
			}
		}
	}
}
func F_norwegian_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v291 int32
	_ = v291
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = v9 + int32(3)
	if v7 < v11 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v9
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v129
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v129 < v131 {
		goto L38
	} else {
		goto L39
	}
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v21 < v20 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v61 < int32(0) {
		goto L1
	} else {
		goto L18
	}
L4:
	;
	v23 = v20
	goto L6
L5:
	;
	v23 = v21
	goto L6
L6:
	;
	v30 = v20
	goto L8
L7:
	;
	v61 = v41
	goto L3
L8:
	;
	if v30 == v23 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v61 = int32(-1)
	goto L3
L11:
	;
	goto L12
L12:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34+v30))))
	if int32(248) < v36 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v53 = v30 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v53
	v30 = v53
	goto L8
L14:
	;
	v38 = v36 - int32(97)
	if v38 < int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v41 = int32(1)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v38)>>(uint(int32(3))%32)))+uint32(_c_F_norwegian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v45)>>(uint(v38&int32(7))%32))&v41 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	goto L13
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v65 = v64 + v61
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v65
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v76 < v65 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v119 < int32(0) {
		goto L1
	} else {
		goto L33
	}
L20:
	;
	v78 = v65
	goto L22
L21:
	;
	v78 = v76
	goto L22
L22:
	;
	v84 = v65
	goto L24
L23:
	;
	v119 = int32(1)
	goto L19
L24:
	;
	if v84 == v78 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v119 = int32(-1)
	goto L19
L27:
	;
	goto L28
L28:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91+v84))))
	if int32(248) < v93 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	v95 = v93 - int32(97)
	if v95 < int32(0) {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v95)>>(uint(int32(3))%32)))+uint32(_c_F_norwegian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v101)>>(uint(v95&int32(7))%32))&int32(1) == int32(0) {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	v110 = v84 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v110
	v84 = v110
	goto L24
L33:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v123 = v122 + v119
	if v11 < v123 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v125 = v123
	goto L36
L35:
	;
	v125 = v11
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v125
	goto L1
L37:
	;
	return v434
L38:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v354
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v354 < v357 {
		v391 = v356
		goto L93
	} else {
		goto L94
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v131
	if v129 <= v131 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v9
	goto L38
L41:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v138 = int32(1)
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136+v129-v138))))
	if base.B2i32(v140&int32(224) != int32(96))|base.B2i32(v138<<(uint(v140)%32)&int32(_a_F_norwegian_ISO_8859_1_stem_0) == int32(0)) != 0 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v155 = F_find_among_b(m, l0, int32(_a_F_norwegian_ISO_8859_1_stem_1), int32(29), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	return int32(0)
L44:
	;
	if v155 == int32(0) {
		goto L40
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v9
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v162
	switch v155 - int32(1) {
	case 0:
		goto L49
	case 1:
		goto L48
	case 2:
		goto L47
	case 3:
		goto L46
	default:
		goto L38
	}
L46:
	;
	v342 = F_slice_from_s(m, l0, int32(2), int32(_a_F_norwegian_ISO_8859_1_stem_2))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L43
	} else {
		goto L91
	}
L47:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L60
L48:
	;
	if v162 <= v9 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v166 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v166 {
		goto L38
	} else {
		goto L50
	}
L50:
	;
	v434 = v166
	goto L37
L51:
	;
	v194 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v194 {
		goto L38
	} else {
		goto L56
	}
L52:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v172 = int32(1)
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170+v162-v172))))
	if base.B2i32(v174&int32(224) != int32(96))|base.B2i32(v172<<(uint(v174)%32)&int32(_a_F_norwegian_ISO_8859_1_stem_3) == int32(0)) != 0 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v189 = F_find_among_b(m, l0, int32(_a_F_norwegian_ISO_8859_1_stem_4), int32(15), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L43
	} else {
		goto L54
	}
L54:
	;
	if v189 != int32(1) {
		goto L38
	} else {
		goto L55
	}
L55:
	;
	goto L51
L56:
	;
	v434 = v194
	goto L37
L57:
	;
	v337 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v337 {
		goto L38
	} else {
		goto L90
	}
L58:
	;
	if v250 == int32(0) {
		goto L57
	} else {
		goto L69
	}
L59:
	;
	v250 = v246
	goto L58
L60:
	;
	if v206 <= v207 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v246 = int32(0)
	goto L59
L62:
	;
	v250 = int32(-1)
	goto L58
L63:
	;
	goto L64
L64:
	;
	v219 = int32(1)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220+v206-v219))))
	if int32(122) < v224 {
		v246 = v219
		goto L59
	} else {
		goto L65
	}
L65:
	;
	v226 = v224 - int32(98)
	if v226 < int32(0) {
		v246 = v219
		goto L59
	} else {
		goto L66
	}
L66:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v226)>>(uint(int32(3))%32)))+uint32(_c_F_norwegian_ISO_8859_1_stem[1]))))
	if int32(base.Ui32(v232)>>(uint(v226&int32(7))%32))&int32(1) == int32(0) {
		v246 = v219
		goto L59
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v206 - int32(1)
	goto L68
L68:
	;
	goto L61
L69:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v255 = v253 + (v162 - v197)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v255
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v255 <= v257 {
		goto L38
	} else {
		goto L70
	}
L70:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v260 = v259 + v255
	v262 = v260 - int32(1)
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262))))
	if v263 == int32(114) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v267 = v255 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v267
	if v267 <= v257 {
		goto L57
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v255
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262))))
	if v277 != int32(107) {
		goto L38
	} else {
		goto L76
	}
L74:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260-int32(2)))))
	if v272 != int32(101) {
		goto L57
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v281 = v255 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v281
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L79
L77:
	;
	if v331 != 0 {
		goto L38
	} else {
		goto L89
	}
L78:
	;
	v331 = v328
	goto L77
L79:
	;
	if v281 <= v291 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v328 = int32(0)
	goto L78
L81:
	;
	v331 = int32(-1)
	goto L77
L82:
	;
	goto L83
L83:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302+v281-int32(1)))))
	if int32(248) < v306 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v281 - int32(1)
	goto L88
L85:
	;
	v308 = v306 - int32(97)
	if v308 < int32(0) {
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v311 = int32(1)
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v308)>>(uint(int32(3))%32)))+uint32(_c_F_norwegian_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v315)>>(uint(v308&int32(7))%32))&v311 != 0 {
		v328 = v311
		goto L78
	} else {
		goto L87
	}
L87:
	;
	goto L84
L88:
	;
	goto L80
L89:
	;
	goto L57
L90:
	;
	v434 = v337
	goto L37
L91:
	;
	if int32(0) <= v342 {
		goto L38
	} else {
		goto L92
	}
L92:
	;
	v434 = v342
	goto L37
L93:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v393
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v393 < v395 {
		v430 = v391
		goto L106
	} else {
		goto L107
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v357
	v362 = v354 - int32(1)
	if v357 < v362 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v373 = F_find_among_b(m, l0, int32(_a_F_norwegian_ISO_8859_1_stem_5), int32(2), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L43
	} else {
		goto L100
	}
L96:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364+v362))))
	if v366 == int32(116) {
		goto L95
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v356
	v391 = v356
	goto L93
L99:
	;
	goto L98
L100:
	;
	if v373 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v356
	v391 = v356
	goto L93
L102:
	;
	goto L103
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v356
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v379
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v381 <= v356 {
		v391 = v356
		goto L93
	} else {
		goto L104
	}
L104:
	;
	v384 = v381 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v384
	v387 = F_slice_del(m, l0)
	mBase = m.M
	if v387 < int32(0) {
		v434 = v387
		goto L37
	} else {
		goto L105
	}
L105:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v391 = v390
	goto L93
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v430
	v434 = int32(1)
	goto L37
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v395
	v400 = v393 - int32(1)
	if v400 <= v395 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v418 = F_find_among_b(m, l0, int32(_a_F_norwegian_ISO_8859_1_stem_6), int32(11), int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L43
	} else {
		goto L113
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v391
	v430 = v391
	goto L106
L110:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402+v400))))
	if v404&int32(224) != int32(96) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	if int32(1)<<(uint(v404)%32)&int32(_a_F_norwegian_ISO_8859_1_stem_7) != 0 {
		goto L108
	} else {
		goto L112
	}
L112:
	;
	goto L109
L113:
	;
	if v418 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v391
	v430 = v391
	goto L106
L115:
	;
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v391
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v424
	v426 = F_slice_del(m, l0)
	mBase = m.M
	if v426 < int32(0) {
		v434 = v426
		goto L37
	} else {
		goto L117
	}
L117:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v430 = v429
	goto L106
}
func F_norwegian_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v138 int32
	_ = v138
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v260 int32
	_ = v260
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v294 int32
	_ = v294
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v505 int32
	_ = v505
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v650 int32
	_ = v650
	var v667 int32
	_ = v667
	var v674 int32
	_ = v674
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v776 int32
	_ = v776
	var v781 int32
	_ = v781
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	goto L4
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v10
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v315
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v315 < v317 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v87 = v10
	goto L25
L23:
	;
	if v182 < int32(0) {
		goto L1
	} else {
		goto L48
	}
L24:
	;
	v182 = v154
	goto L23
L25:
	;
	if v78 <= v87 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v182 = int32(-1)
	goto L23
L28:
	;
	goto L29
L29:
	;
	v94 = int32(1)
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87+v79))))
	if base.Ui32(v96) < base.Ui32(int32(192)) {
		v153 = v96
		v154 = v94
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if int32(248) < v153 {
		goto L43
	} else {
		goto L44
	}
L31:
	;
	v100 = v87 + int32(1)
	if v100 == v78 {
		v153 = v96
		v154 = v94
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+v79))))
	v105 = v103 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v96) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+v79))))
	v121 = v119 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v96) {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	v109 = v87 + int32(2)
	if v109 != v78 {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v153 = v96<<(uint(int32(6))%32)&int32(1984) | v105
	v154 = int32(2)
	goto L30
L37:
	;
	goto L36
L38:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79+v125))))
	v153 = v138&int32(63) | (v96<<(uint(int32(18))%32)&int32(_a_F_norwegian_UTF_8_stem_0) | v105<<(uint(int32(12))%32) | v121<<(uint(int32(6))%32))
	v154 = int32(4)
	goto L30
L39:
	;
	v125 = v87 + int32(3)
	if v125 != v78 {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v153 = v96<<(uint(int32(12))%32)&int32(_a_F_norwegian_UTF_8_stem_1) | v105<<(uint(int32(6))%32) | v121
	v154 = int32(3)
	goto L30
L42:
	;
	goto L41
L43:
	;
	v171 = v154 + v87
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v171
	v87 = v171
	goto L25
L44:
	;
	v158 = v153 - int32(97)
	if v158 < int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v158)>>(uint(int32(3))%32)))+uint32(_c_F_norwegian_UTF_8_stem[0]))))
	if int32(base.Ui32(v164)>>(uint(v158&int32(7))%32))&int32(1) != 0 {
		goto L24
	} else {
		goto L46
	}
L46:
	;
	goto L43
L48:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v186 = v185 + v182
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v186
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v209 = v186
	goto L51
L49:
	;
	if v305 < int32(0) {
		goto L1
	} else {
		goto L73
	}
L50:
	;
	v305 = v276
	goto L49
L51:
	;
	if v200 <= v209 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v305 = int32(-1)
	goto L49
L54:
	;
	goto L55
L55:
	;
	v216 = int32(1)
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209+v201))))
	if base.Ui32(v218) < base.Ui32(int32(192)) {
		v275 = v218
		v276 = v216
		goto L56
	} else {
		goto L57
	}
L56:
	;
	if int32(248) < v275 {
		goto L50
	} else {
		goto L69
	}
L57:
	;
	v222 = v209 + int32(1)
	if v222 == v200 {
		v275 = v218
		v276 = v216
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222+v201))))
	v227 = v225 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v218) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231+v201))))
	v243 = v241 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v218) {
		goto L65
	} else {
		goto L66
	}
L60:
	;
	v231 = v209 + int32(2)
	if v231 != v200 {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v275 = v218<<(uint(int32(6))%32)&int32(1984) | v227
	v276 = int32(2)
	goto L56
L63:
	;
	goto L62
L64:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201+v247))))
	v275 = v260&int32(63) | (v218<<(uint(int32(18))%32)&int32(_a_F_norwegian_UTF_8_stem_0) | v227<<(uint(int32(12))%32) | v243<<(uint(int32(6))%32))
	v276 = int32(4)
	goto L56
L65:
	;
	v247 = v209 + int32(3)
	if v247 != v200 {
		goto L64
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v275 = v218<<(uint(int32(12))%32)&int32(_a_F_norwegian_UTF_8_stem_1) | v227<<(uint(int32(6))%32) | v243
	v276 = int32(3)
	goto L56
L68:
	;
	goto L67
L69:
	;
	v280 = v275 - int32(97)
	if v280 < int32(0) {
		goto L50
	} else {
		goto L70
	}
L70:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v280)>>(uint(int32(3))%32)))+uint32(_c_F_norwegian_UTF_8_stem[0]))))
	if int32(base.Ui32(v286)>>(uint(v280&int32(7))%32))&int32(1) == int32(0) {
		goto L50
	} else {
		goto L71
	}
L71:
	;
	v294 = v276 + v209
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v294
	v209 = v294
	goto L51
L73:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v309 = v308 + v305
	if v62 < v309 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v311 = v309
	goto L76
L75:
	;
	v311 = v62
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v311
	goto L1
L77:
	;
	return v829
L78:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v697
	v699 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v697 < v699 {
		goto L151
	} else {
		goto L152
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v317
	if v315 <= v317 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v10
	goto L78
L81:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v324 = int32(1)
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322+v315-v324))))
	if base.B2i32(v326&int32(224) != int32(96))|base.B2i32(v324<<(uint(v326)%32)&int32(_a_F_norwegian_UTF_8_stem_2) == int32(0)) != 0 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v341 = F_find_among_b(m, l0, int32(_a_F_norwegian_UTF_8_stem_3), int32(29), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	return int32(0)
L84:
	;
	if v341 == int32(0) {
		goto L80
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v10
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v348
	switch v341 - int32(1) {
	case 0:
		goto L89
	case 1:
		goto L88
	case 2:
		goto L87
	case 3:
		goto L86
	default:
		goto L78
	}
L86:
	;
	v685 = F_slice_from_s(m, l0, int32(2), int32(_a_F_norwegian_UTF_8_stem_4))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L83
	} else {
		goto L149
	}
L87:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L100
L88:
	;
	if v348 <= v10 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	v352 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v352 {
		goto L78
	} else {
		goto L90
	}
L90:
	;
	v829 = v352
	goto L77
L91:
	;
	v380 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v380 {
		goto L78
	} else {
		goto L96
	}
L92:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v358 = int32(1)
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356+v348-v358))))
	if base.B2i32(v360&int32(224) != int32(96))|base.B2i32(v358<<(uint(v360)%32)&int32(_a_F_norwegian_UTF_8_stem_5) == int32(0)) != 0 {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v375 = F_find_among_b(m, l0, int32(_a_F_norwegian_UTF_8_stem_6), int32(15), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L83
	} else {
		goto L94
	}
L94:
	;
	if v375 != int32(1) {
		goto L78
	} else {
		goto L95
	}
L95:
	;
	goto L91
L96:
	;
	v829 = v380
	goto L77
L97:
	;
	v680 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v680 {
		goto L78
	} else {
		goto L148
	}
L98:
	;
	if v512 == int32(0) {
		goto L97
	} else {
		goto L121
	}
L99:
	;
	v512 = v505
	goto L98
L100:
	;
	if v396 <= v397 {
		v505 = int32(-1)
		goto L99
	} else {
		goto L102
	}
L101:
	;
	v505 = int32(0)
	goto L99
L102:
	;
	v414 = int32(1)
	v415 = v396 - v414
	v417 = int32(*(*int8)(unsafe.Add(mBase, uint32(v398+v415))))
	v419 = v417 & int32(255)
	if base.B2i32(v415 == v397)|base.B2i32(int32(0) <= v417) != 0 {
		v477 = v419
		v481 = v414
		goto L103
	} else {
		goto L104
	}
L103:
	;
	if int32(122) < v477 {
		goto L111
	} else {
		goto L112
	}
L104:
	;
	v426 = v419 & int32(63)
	v428 = v396 - int32(2)
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398+v428))))
	v432 = v430 << (uint(int32(6)) % 32)
	if base.B2i32(v428 != v397)&base.B2i32(base.Ui32(v430) < base.Ui32(int32(192))) == int32(0) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v477 = v432&int32(1984) | v426
	v481 = int32(2)
	goto L103
L106:
	;
	goto L107
L107:
	;
	v445 = v432&int32(4032) | v426
	v447 = v396 - int32(3)
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398+v447))))
	if base.B2i32(v447 != v397)&base.B2i32(base.Ui32(v449) < base.Ui32(int32(224))) == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v477 = v449<<(uint(int32(12))%32)&int32(_a_F_norwegian_UTF_8_stem_1) | v445
	v481 = int32(3)
	goto L103
L109:
	;
	goto L110
L110:
	;
	v467 = int32(4)
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396+v398-v467))))
	v477 = v449<<(uint(int32(12))%32)&int32(_a_F_norwegian_UTF_8_stem_7) | v469&int32(7)<<(uint(int32(18))%32) | v445
	v481 = v467
	goto L103
L111:
	;
	v512 = v481
	goto L98
L112:
	;
	goto L113
L113:
	;
	v483 = v477 - int32(98)
	if v483 < int32(0) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v512 = v481
	goto L98
L115:
	;
	goto L116
L116:
	;
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v483)>>(uint(int32(3))%32)))+uint32(_c_F_norwegian_UTF_8_stem[1]))))
	if int32(base.Ui32(v489)>>(uint(v483&int32(7))%32))&int32(1) == int32(0) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v512 = v481
	goto L98
L118:
	;
	goto L119
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v396 - v481
	goto L120
L120:
	;
	goto L101
L121:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v517 = v515 + (v348 - v383)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v517
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v517 <= v519 {
		goto L78
	} else {
		goto L122
	}
L122:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v522 = v521 + v517
	v524 = v522 - int32(1)
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524))))
	if v525 == int32(114) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v529 = v517 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v529
	if v529 <= v519 {
		goto L97
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v517
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524))))
	if v539 != int32(107) {
		goto L78
	} else {
		goto L128
	}
L126:
	;
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522-int32(2)))))
	if v534 != int32(101) {
		goto L97
	} else {
		goto L127
	}
L127:
	;
	goto L125
L128:
	;
	v543 = v517 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v543
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L131
L129:
	;
	if v674 != 0 {
		goto L78
	} else {
		goto L147
	}
L130:
	;
	v674 = v667
	goto L129
L131:
	;
	if v543 <= v558 {
		v667 = int32(-1)
		goto L130
	} else {
		goto L133
	}
L132:
	;
	v667 = int32(0)
	goto L130
L133:
	;
	v575 = int32(1)
	v576 = v543 - v575
	v578 = int32(*(*int8)(unsafe.Add(mBase, uint32(v559+v576))))
	v580 = v578 & int32(255)
	if base.B2i32(v576 == v558)|base.B2i32(int32(0) <= v578) != 0 {
		v638 = v580
		v642 = v575
		goto L134
	} else {
		goto L135
	}
L134:
	;
	if int32(248) < v638 {
		goto L142
	} else {
		goto L143
	}
L135:
	;
	v587 = v580 & int32(63)
	v589 = v543 - int32(2)
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v559+v589))))
	v593 = v591 << (uint(int32(6)) % 32)
	if base.B2i32(v589 != v558)&base.B2i32(base.Ui32(v591) < base.Ui32(int32(192))) == int32(0) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v638 = v593&int32(1984) | v587
	v642 = int32(2)
	goto L134
L137:
	;
	goto L138
L138:
	;
	v606 = v593&int32(4032) | v587
	v608 = v543 - int32(3)
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v559+v608))))
	if base.B2i32(v608 != v558)&base.B2i32(base.Ui32(v610) < base.Ui32(int32(224))) == int32(0) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v638 = v610<<(uint(int32(12))%32)&int32(_a_F_norwegian_UTF_8_stem_1) | v606
	v642 = int32(3)
	goto L134
L140:
	;
	goto L141
L141:
	;
	v628 = int32(4)
	v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543+v559-v628))))
	v638 = v610<<(uint(int32(12))%32)&int32(_a_F_norwegian_UTF_8_stem_7) | v630&int32(7)<<(uint(int32(18))%32) | v606
	v642 = v628
	goto L134
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v543 - v642
	goto L146
L143:
	;
	v644 = v638 - int32(97)
	if v644 < int32(0) {
		goto L142
	} else {
		goto L144
	}
L144:
	;
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v644)>>(uint(int32(3))%32)))+uint32(_c_F_norwegian_UTF_8_stem[0]))))
	if int32(base.Ui32(v650)>>(uint(v644&int32(7))%32))&int32(1) == int32(0) {
		goto L142
	} else {
		goto L145
	}
L145:
	;
	v674 = v642
	goto L129
L146:
	;
	goto L132
L147:
	;
	goto L97
L148:
	;
	v829 = v680
	goto L77
L149:
	;
	if int32(0) <= v685 {
		goto L78
	} else {
		goto L150
	}
L150:
	;
	v829 = v685
	goto L77
L151:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v787
	v789 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v790 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v787 < v790 {
		v825 = v789
		goto L179
	} else {
		goto L180
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v697
	v702 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v699
	v705 = v697 - int32(1)
	if v705 <= v699 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v702
	goto L151
L154:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v707+v705))))
	if v709 != int32(116) {
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v715 = F_find_among_b(m, l0, int32(_a_F_norwegian_UTF_8_stem_8), int32(2), int32(0))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L83
	} else {
		goto L156
	}
L156:
	;
	if v715 == int32(0) {
		goto L153
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v702
	v720 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v720
	v722 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v722
	v724 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L160
L158:
	;
	if v776 < int32(0) {
		goto L151
	} else {
		goto L177
	}
L160:
	;
	goto L161
L161:
	;
	goto L162
L162:
	;
	v731 = v722
	v733 = int32(1)
	goto L165
L164:
	;
	v776 = v758
	goto L158
L165:
	;
	if v731 <= v702 {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	goto L164
L167:
	;
	v776 = int32(-1)
	goto L158
L168:
	;
	goto L169
L169:
	;
	v738 = v731 - int32(1)
	v740 = int32(*(*int8)(unsafe.Add(mBase, uint32(v724+v738))))
	if base.B2i32(int32(0) <= v740)|base.B2i32(v738 <= v702) != 0 {
		v758 = v738
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v762 = int32(1)
	if v762 < v733 {
		v731 = v758
		v733 = v733 - v762
		goto L165
	} else {
		goto L176
	}
L171:
	;
	v746 = v738
	goto L172
L172:
	;
	v751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v724+v746))))
	if base.Ui32(int32(191)) < base.Ui32(v751) {
		v758 = v746
		goto L170
	} else {
		goto L174
	}
L173:
	;
	v758 = v702
	goto L170
L174:
	;
	v755 = v746 - int32(1)
	if v702 < v755 {
		v746 = v755
		goto L172
	} else {
		goto L175
	}
L175:
	;
	goto L173
L176:
	;
	goto L166
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v776
	v781 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v781 {
		goto L151
	} else {
		goto L178
	}
L178:
	;
	v829 = v781
	goto L77
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v825
	v829 = int32(1)
	goto L77
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v790
	v795 = v787 - int32(1)
	if v795 <= v790 {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	v813 = F_find_among_b(m, l0, int32(_a_F_norwegian_UTF_8_stem_9), int32(11), int32(0))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L83
	} else {
		goto L186
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v789
	v825 = v789
	goto L179
L183:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v797+v795))))
	if v799&int32(224) != int32(96) {
		goto L182
	} else {
		goto L184
	}
L184:
	;
	if int32(1)<<(uint(v799)%32)&int32(_a_F_norwegian_UTF_8_stem_10) != 0 {
		goto L181
	} else {
		goto L185
	}
L185:
	;
	goto L182
L186:
	;
	if v813 == int32(0) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v789
	v825 = v789
	goto L179
L188:
	;
	goto L189
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v789
	v819 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v819
	v821 = F_slice_del(m, l0)
	mBase = m.M
	if v821 < int32(0) {
		v829 = v821
		goto L77
	} else {
		goto L190
	}
L190:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v825 = v824
	goto L179
}
func F_numerictypmodin(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v15 = F_ArrayGetIntegerTypmods(m, v9, v6+int32(12))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			switch v17 - int32(1) {
			case 0:
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				if base.Ui32(v20-int32(1001)) <= base.Ui32(int32(-1001)) {
					v25 = int32(-1)
					v27 = F_errsave_start(m, int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int64(0)
					} else {
						if v27 == int32(0) {
							v71 = v25
							m.G0 = v6 + int32(16)
							return base.I64_extend_i32_s(v71)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(1000)
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = v20
								F_errmsg(m, int32(_a_F_numerictypmodin_0), v6)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int64(0)
								} else {
									F_errsave_finish(m, int32(0), int32(_a_F_numerictypmodin_1), int32(1321), int32(_a_F_numerictypmodin_2))
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int64(0)
									} else {
										v71 = v25
										m.G0 = v6 + int32(16)
										return base.I64_extend_i32_s(v71)
									}
								}
							}
						}
					}
				} else {
					v71 = v20<<(uint(int32(16))%32) | int32(4)
					m.G0 = v6 + int32(16)
					return base.I64_extend_i32_s(v71)
				}
			case 1:
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
				v69 = F_make_numeric_typmod_safe(m, v66, v67, int32(0))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int64(0)
				} else {
					v71 = v69
					m.G0 = v6 + int32(16)
					return base.I64_extend_i32_s(v71)
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_numerictypmodin_3), int32(0))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_numerictypmodin_1), int32(1352), int32(_a_F_numerictypmodin_4))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int64(0)
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
func F_numrange_subdiff(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	v3 = int32(0)
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = F_DirectFunctionCall2Coll(m, int32(19), v3, v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = F_DirectFunctionCall1Coll(m, int32(18), v3, v8)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			return v12
		}
	}
}
