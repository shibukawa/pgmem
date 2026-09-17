package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dshash_find(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v11 = m.T0[v10].(func(*base.Module, int32, int32, int32) int32)(m, l1, v8, v9)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v17 = int32(base.Ui32(v11) >> (uint(int32(25)) % 32))
	v25 = F_LWLockAcquire(m, v15+v17*int32(20)+int32(8), l2^int32(1))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+2572))
	if v27 == v29 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(base.Ui32(v11)>>(uint(int32(32)-v41)%32))<<(uint(int32(2))%32))))
	if v49 != 0 {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v40 = v28
	v41 = v27
	v42 = v31
	goto L4
L6:
	;
	goto L7
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+2576))
	v34 = F_dsa_get_address(m, v32, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v34
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+2572))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v38
	v40 = v37
	v41 = v38
	v42 = v34
	goto L4
L9:
	;
	return v89
L10:
	;
	v53 = v49
	goto L13
L11:
	;
	v78 = v40
	goto L12
L12:
	;
	F_LWLockRelease(m, v78+v17*int32(20)+int32(8))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L19
	}
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v58 = F_dsa_get_address(m, v57, v53)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v78 = v70
	goto L12
L15:
	;
	v61 = v58 + int32(8)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v65 = m.T0[v64].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v61, v62, v63)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if v65 == int32(0) {
		v89 = v61
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if v69 != 0 {
		v53 = v69
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	v89 = int32(0)
	goto L9
}
func F_dshash_memcpy(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	if l2 != 0 {
		base.MemoryCopy(m, l0, l1, l2)
	} else {
	}
	return
}
func F_dshash_strcmp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v7 == int32(0))|base.B2i32(v7 != v10) != 0 {
		v28 = v7
		v29 = v10
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v28 - v29
L2:
	;
	goto L1
L3:
	;
	v13 = l0
	v14 = l1
	goto L4
L4:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v18 == int32(0) {
		v28 = v18
		v29 = v17
		goto L2
	} else {
		goto L6
	}
L5:
	;
	v28 = v18
	v29 = v17
	goto L2
L6:
	;
	v21 = int32(1)
	if v18 == v17 {
		v13 = v13 + v21
		v14 = v14 + v21
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
}
func F_dshash_strcpy(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	if (l1^l0)&int32(3) != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return
L2:
	;
	goto L1
L3:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v59))) = uint8(v58)
	if v58&int32(255) == int32(0) {
		goto L2
	} else {
		goto L18
	}
L4:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v57 = l1
	v58 = v10
	v59 = l0
	goto L3
L5:
	;
	goto L6
L6:
	;
	if l1&int32(3) != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v14 = l1
	v16 = l0
	goto L10
L8:
	;
	v28 = l1
	v30 = l0
	goto L9
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v35 = int32(-2139062144)
	if (int32(16843008)-v32|v32)&v35 != v35 {
		v57 = v28
		v58 = v32
		v59 = v30
		goto L3
	} else {
		goto L14
	}
L10:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v17)
	if v17 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L11:
	;
	v28 = v24
	v30 = v22
	goto L9
L12:
	;
	v21 = int32(1)
	v22 = v16 + v21
	v24 = v14 + v21
	if v24&int32(3) != 0 {
		v14 = v24
		v16 = v22
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v40 = v28
	v41 = v32
	v42 = v30
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v41
	v44 = int32(4)
	v45 = v42 + v44
	v47 = v40 + v44
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v52 = int32(-2139062144)
	if (int32(16843008)-v49|v49)&v52 == v52 {
		v40 = v47
		v41 = v49
		v42 = v45
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v57 = v47
	v58 = v49
	v59 = v45
	goto L3
L17:
	;
	goto L16
L18:
	;
	v66 = v57
	v68 = v59
	goto L19
L19:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v68)+1)) = uint8(v69)
	v71 = int32(1)
	if v69 != 0 {
		v66 = v66 + v71
		v68 = v68 + v71
		goto L19
	} else {
		goto L21
	}
L20:
	;
	goto L2
L21:
	;
	goto L20
}
