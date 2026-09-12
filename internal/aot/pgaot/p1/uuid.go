package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_uuid_generate_random(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	v3 = m.Env.Pgmem_random_bytes(m, l0, int32(16))
	mBase = m.M
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	v8 = v4&int32(15) | int32(64)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)) = uint8(v8)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v14 = v10&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v14)
	return
}
func F_uuid_generate_v1mc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	v3 = m.G0
	v5 = v3 - int32(128)
	m.G0 = v5
	F_uuid_generate_random(m, v5)
	mBase = m.M
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+10)))
	v10 = v8 | int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+10)) = uint8(v10)
	F_uuid_unparse(m, v5, v5+int32(16))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = v5 - int32(-64)
	F_uuid_generate_time(m, v19)
	mBase = m.M
	F_uuid_unparse(m, v19, v5+int32(80))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = v5 + int32(103)
	v30 = v5 + int32(40)
	if (v30^v28)&int32(3) != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v109 = F_DirectFunctionCall1Coll(m, int32(3392), int32(0), v5+int32(80))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L25
	}
L5:
	;
	goto L4
L6:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v84)
	if v84&int32(255) == int32(0) {
		goto L5
	} else {
		goto L21
	}
L7:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v83 = v30
	v84 = v36
	v85 = v28
	goto L6
L8:
	;
	goto L9
L9:
	;
	if v30&int32(3) != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v40 = v30
	v42 = v28
	goto L13
L11:
	;
	v54 = v30
	v56 = v28
	goto L12
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v61 = int32(-2139062144)
	if (int32(16843008)-v58|v58)&v61 != v61 {
		v83 = v54
		v84 = v58
		v85 = v56
		goto L6
	} else {
		goto L17
	}
L13:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	*(*uint8)(unsafe.Add(mBase, uint32(v42))) = uint8(v43)
	if v43 == int32(0) {
		goto L5
	} else {
		goto L15
	}
L14:
	;
	v54 = v50
	v56 = v48
	goto L12
L15:
	;
	v47 = int32(1)
	v48 = v42 + v47
	v50 = v40 + v47
	if v50&int32(3) != 0 {
		v40 = v50
		v42 = v48
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v66 = v54
	v67 = v58
	v68 = v56
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v67
	v70 = int32(4)
	v71 = v68 + v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v74 = v66 + v70
	v78 = int32(-2139062144)
	if (v72|(int32(16843008)-v72))&v78 == v78 {
		v66 = v74
		v67 = v72
		v68 = v71
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v83 = v74
	v84 = v72
	v85 = v71
	goto L6
L20:
	;
	goto L19
L21:
	;
	v92 = v83
	v94 = v85
	goto L22
L22:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v94)+1)) = uint8(v95)
	v97 = int32(1)
	if v95 != 0 {
		v92 = v92 + v97
		v94 = v94 + v97
		goto L22
	} else {
		goto L24
	}
L23:
	;
	goto L5
L24:
	;
	goto L23
L25:
	;
	m.G0 = v5 + int32(128)
	return v109
}
func F_uuid_generate_v3(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = int32(1)
		v13 = v8 + v12
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		v18 = v16 & v12
		if v18 != 0 {
			v19 = v13
		} else {
			v19 = v8 + int32(4)
		}
		if v16 == int32(1) {
			v22 = int32(4)
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			if v24&int32(254) == int32(2) {
				v33 = v22
			} else {
				v33 = base.B2i32(v24 == int32(18)) << (uint(v22) % 32)
			}
			if v24 == int32(1) {
				v36 = v22
			} else {
				v36 = v33
			}
			v47 = v36
		} else {
			v37 = int32(1)
			if v18 != 0 {
				v47 = int32(base.Ui32(v16)>>(uint(v37)%32)) - v37
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v47 = int32(base.Ui32(v41)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v48 = F_uuid_generate_internal(m, int32(3), v6, v19, v47)
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return int32(0)
		} else {
			return v48
		}
	}
}
func F_uuid_generate_v5(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = int32(1)
		v13 = v8 + v12
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		v18 = v16 & v12
		if v18 != 0 {
			v19 = v13
		} else {
			v19 = v8 + int32(4)
		}
		if v16 == int32(1) {
			v22 = int32(4)
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			if v24&int32(254) == int32(2) {
				v33 = v22
			} else {
				v33 = base.B2i32(v24 == int32(18)) << (uint(v22) % 32)
			}
			if v24 == int32(1) {
				v36 = v22
			} else {
				v36 = v33
			}
			v47 = v36
		} else {
			v37 = int32(1)
			if v18 != 0 {
				v47 = int32(base.Ui32(v16)>>(uint(v37)%32)) - v37
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v47 = int32(base.Ui32(v41)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v48 = F_uuid_generate_internal(m, int32(5), v6, v19, v47)
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return int32(0)
		} else {
			return v48
		}
	}
}
func F_uuid_gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = int32(16)
	goto L4
L1:
	;
	return base.B2i32(int32(0) < v66)
L2:
	;
	v66 = int32(0)
	goto L1
L3:
	;
	v40 = v35
	v41 = v36
	v42 = v37
	goto L13
L4:
	;
	if (v2|v3)&int32(3) != 0 {
		v35 = v2
		v36 = v3
		v37 = v4
		goto L3
	} else {
		goto L7
	}
L6:
	;
	if v25 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L7:
	;
	v12 = v2
	v13 = v3
	v14 = v4
	goto L8
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v17 != v18 {
		v35 = v12
		v36 = v13
		v37 = v14
		goto L3
	} else {
		goto L10
	}
L9:
	;
	goto L6
L10:
	;
	v20 = int32(4)
	v21 = v13 + v20
	v23 = v12 + v20
	v25 = v14 - v20
	if base.Ui32(int32(3)) < base.Ui32(v25) {
		v12 = v23
		v13 = v21
		v14 = v25
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v35 = v23
	v36 = v21
	v37 = v25
	goto L3
L13:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v45 == v46 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v66 = v45 - v46
	goto L1
L15:
	;
	v48 = int32(1)
	v53 = v42 - v48
	if v53 != 0 {
		v40 = v40 + v48
		v41 = v41 + v48
		v42 = v53
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	goto L14
L18:
	;
	goto L2
}
func F_uuid_ns_dns(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	v7 = int32(581424)
	v8 = *(*int64)(unsafe.Add(mBase, _consts[1451]))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+29)) = v8
	v10 = *(*int64)(unsafe.Add(mBase, _consts[1452]))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+24)) = v10
	v12 = *(*int64)(unsafe.Add(mBase, _consts[1453]))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+16)) = v12
	v14 = *(*int64)(unsafe.Add(mBase, _consts[1454]))
	*(*int64)(unsafe.Add(mBase, uint32(v5))) = v14
	v16 = *(*int64)(unsafe.Add(mBase, _consts[1455]))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = v16
	v20 = F_DirectFunctionCall1Coll(m, int32(3392), int32(0), v5)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		m.G0 = v5 + int32(48)
		return v20
	}
}
func F_uuid_ns_url(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	v7 = int32(581387)
	v8 = *(*int64)(unsafe.Add(mBase, _consts[1456]))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+29)) = v8
	v10 = *(*int64)(unsafe.Add(mBase, _consts[1457]))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+24)) = v10
	v12 = *(*int64)(unsafe.Add(mBase, _consts[1458]))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+16)) = v12
	v14 = *(*int64)(unsafe.Add(mBase, _consts[1459]))
	*(*int64)(unsafe.Add(mBase, uint32(v5))) = v14
	v16 = *(*int64)(unsafe.Add(mBase, _consts[1460]))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = v16
	v20 = F_DirectFunctionCall1Coll(m, int32(3392), int32(0), v5)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		m.G0 = v5 + int32(48)
		return v20
	}
}
