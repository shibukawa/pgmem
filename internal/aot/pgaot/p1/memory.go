package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetMemoryChunkContext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(8))))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v4&int32(15)*int32(36))+uint32(_consts[1234])))
	v12 = m.T0[v11].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		return v12
	}
}
func F_MemoryContextAlloc(m *base.Module, l0 int32, l1 int32) int32 {
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
	v3 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v3)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = m.T0[v7].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, v3)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_MemoryContextDelete(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	v10 = l0
	goto L1
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v13 != 0 {
		v10 = v13
		goto L1
	} else {
		goto L3
	}
L2:
	;
	return
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = v15
	goto L7
L5:
	;
	v34 = v14
	goto L6
L6:
	;
	if v34 != 0 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	m.T0[v25].(func(*base.Module, int32))(m, v24)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v34 = v29
	goto L6
L9:
	;
	return
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	if v28 != 0 {
		v18 = v28
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	if v37 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = int32(0)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	m.T0[v51].(func(*base.Module, int32))(m, v10)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L9
	} else {
		goto L22
	}
L15:
	;
	if v36 != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+28)) = v36
	goto L15
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v36
	goto L15
L19:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+24)) = v40
	goto L21
L20:
	;
	goto L21
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(0)
	goto L14
L22:
	;
	if l0 != v10 {
		v10 = v14
		goto L1
	} else {
		goto L23
	}
L23:
	;
	goto L2
}
func F_MemoryContextMemConsumed(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int64
	_ = v5
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v3 = int32(0)
	v5 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v5
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v5
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	m.T0[v13].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, v3, v3, l1, v3)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v16 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v20 = v16
	goto L6
L4:
	;
	goto L5
L5:
	;
	return
L6:
	;
	v21 = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	m.T0[v25].(func(*base.Module, int32, int32, int32, int32, int32))(m, v20, v21, v21, l1, v21)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	if v28 != 0 {
		v20 = v28
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v31 = v20
	goto L10
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+28))
	if v33 != 0 {
		v20 = v33
		goto L6
	} else {
		goto L12
	}
L11:
	;
	goto L7
L12:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	if v34 != l0 {
		v31 = v34
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
}
func F_MemoryContextReset(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = v8
	goto L4
L2:
	;
	goto L3
L3:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v75 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L4:
	;
	v20 = v12
	goto L6
L5:
	;
	goto L3
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	if v23 != 0 {
		v20 = v23
		goto L6
	} else {
		goto L8
	}
L7:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v67 != 0 {
		v12 = v67
		goto L4
	} else {
		goto L29
	}
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	if v25 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v28 = v25
	goto L12
L10:
	;
	v46 = v24
	goto L11
L11:
	;
	if v46 != 0 {
		goto L17
	} else {
		goto L18
	}
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	m.T0[v36].(func(*base.Module, int32))(m, v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v46 = v40
	goto L11
L14:
	;
	return
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	if v39 != 0 {
		v28 = v39
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	if v49 != 0 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = int32(0)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	m.T0[v63].(func(*base.Module, int32))(m, v20)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L14
	} else {
		goto L27
	}
L20:
	;
	if v48 != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+28)) = v48
	goto L20
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = v48
	goto L20
L24:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v52
	goto L26
L25:
	;
	goto L26
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = int32(0)
	goto L19
L27:
	;
	if v20 != v12 {
		v20 = v24
		goto L6
	} else {
		goto L28
	}
L28:
	;
	goto L7
L29:
	;
	goto L5
L30:
	;
	goto L33
L31:
	;
	goto L32
L32:
	;
	return
L33:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v85 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	m.T0[v93].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L14
	} else {
		goto L39
	}
L35:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	m.T0[v89].(func(*base.Module, int32))(m, v88)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L14
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	goto L34
L38:
	;
	goto L33
L39:
	;
	v96 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v96)
	goto L32
}
func F_MemoryContextSizeFailure(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
		F_errmsg_internal(m, int32(39014), v5)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			F_errfinish(m, int32(517422), int32(1177), int32(382915))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_MemoryContextStats(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v9 = v6 + int32(40)
	v10 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v10
	*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = v10
	v14 = int32(1)
	v15 = int32(100)
	F_MemoryContextStatsInternal(m, l0, v14, v15, v15, v6+int32(32), v14)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v22 - v23
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = v22
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v6)+32))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v27
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v23
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v6)+36))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v30
		v33 = *(*int32)(unsafe.Add(mBase, _consts[466]))
		v35 = F_pg_fprintf(m, v33, int32(787711), v6)
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			m.G0 = v6 + int32(48)
			return
		}
	}
}
