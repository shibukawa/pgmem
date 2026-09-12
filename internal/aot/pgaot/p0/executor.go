package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateExecutorState(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v31 int32
	_ = v31
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_CreateExecutorState[0]))
	v10 = F_AllocSetContextCreateInternal(m, v5, int32(_a_F_CreateExecutorState_0), int32(0), int32(_a_F_CreateExecutorState_1), int32(_a_F_CreateExecutorState_2))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = int32(_a_F_CreateExecutorState_3)
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_CreateExecutorState[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_CreateExecutorState[0])) = v10
		v19 = F_palloc0(m, int32(200))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v19)+188)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(v19))) = int64(4294967685)
			*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(v19)+60)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = v21
			v31 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v31
			*(*int32)(unsafe.Add(mBase, uint32(v19)+100)) = v10
			*(*int64)(unsafe.Add(mBase, uint32(v19)+176)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(v19)+164)) = v21
			*(*uint8)(unsafe.Add(mBase, uint32(v19)+160)) = uint8(v31)
			*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v31
			*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(v19)+32)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v31
			*(*int64)(unsafe.Add(mBase, uint32(v19)+68)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(v19)+88)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v31
			*(*uint8)(unsafe.Add(mBase, uint32(v19)+136)) = uint8(v31)
			*(*int64)(unsafe.Add(mBase, uint32(v19)+128)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(v19)+120)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(v19)+112)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(v19)+140)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(v19)+148)) = v21
			*(*int32)(unsafe.Add(mBase, _c_F_CreateExecutorState[0])) = v15
			return v19
		}
	}
}
func F_ExecutorRun(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_ExecutorRun[0]))
	if v6 != 0 {
		m.T0[v6].(func(*base.Module, int32, int32, int64))(m, l0, l1, l2)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	} else {
		F_standard_ExecutorRun(m, l0, l1, l2)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			return
		}
	}
}
func F_FreeExecutorState(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = v5
	goto L4
L2:
	;
	goto L3
L3:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v52 != 0 {
		goto L22
	} else {
		goto L23
	}
L4:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+68))
	if v12 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v13 = int32(_a_F_FreeExecutorState_0)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_FreeExecutorState[0]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_FreeExecutorState[0])) = v16
	v19 = v12
	goto L9
L7:
	;
	goto L8
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	F_MemoryContextDelete(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L11
	} else {
		goto L15
	}
L9:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	m.T0[v25].(func(*base.Module, int32))(m, v24)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FreeExecutorState[0])) = v14
	goto L8
L11:
	;
	return
L12:
	;
	F_pfree(m, v19)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+68))
	if v30 != 0 {
		v19 = v30
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	if v40 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+140))
	v42 = F_list_delete_ptr(m, v41, v11)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L11
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	F_pfree(m, v11)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L11
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+140)) = v42
	goto L18
L20:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v47 != 0 {
		v7 = v47
		goto L4
	} else {
		goto L21
	}
L21:
	;
	goto L5
L22:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FreeExecutorState[1])))
	if v54 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v63 != 0 {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_FreeExecutorState[2]))
	m.T0[v56].(func(*base.Module, int32))(m, v52)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L11
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	F_pfree(m, v52)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L11
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = int32(0)
	goto L24
L30:
	;
	F_DestroyPartitionDirectory(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L11
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	F_MemoryContextDelete(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L11
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = int32(0)
	goto L32
L34:
	;
	return
}
