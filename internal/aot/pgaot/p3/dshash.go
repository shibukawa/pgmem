package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dshash_find(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
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
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
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
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = m.T0[v9].(func(*base.Module, int32, int32, int32) int32)(m, l1, v7, v8)
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
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v16 = int32(base.Ui32(v10) >> (uint(int32(25)) % 32))
	v24 = F_LWLockAcquire(m, v14+v16*int32(20)+int32(8), l2^int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+2572))
	if v26 == v28 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(base.Ui32(v10)>>(uint(int32(32)-v39)%32))<<(uint(int32(2))%32))))
	if v47 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v39 = v26
	v40 = v30
	goto L4
L6:
	;
	goto L7
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+2576))
	v33 = F_dsa_get_address(m, v31, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v33
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+2572))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v37
	v39 = v37
	v40 = v33
	goto L4
L9:
	;
	return v84
L10:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_LWLockRelease(m, v73+v16*int32(20)+int32(8))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L21
	}
L11:
	;
	v53 = v47
	goto L12
L12:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v57 = F_dsa_get_address(m, v56, v53)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	if v57 != 0 {
		v84 = v60
		goto L9
	} else {
		goto L20
	}
L14:
	;
	v60 = v57 + int32(8)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v64 = m.T0[v63].(func(*base.Module, int32, int32, int32, int32) int32)(m, l1, v60, v61, v62)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v64 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v66 != 0 {
		v53 = v66
		goto L12
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	goto L13
L19:
	;
	goto L10
L20:
	;
	goto L10
L21:
	;
	v84 = int32(0)
	goto L9
}
func F_dshash_memcpy(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	var v5 int32
	_ = v5
	if l2 != 0 {
		v5 = F__emscripten_memcpy_bulkmem(m, l0, l1, l2)
	} else {
	}
	return
}
func F_dshash_strcmp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v8 == int32(0) {
		v27 = v7
		v28 = v8
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v28 - v27
L2:
	;
	goto L1
L3:
	;
	if v7 != v8 {
		v27 = v7
		v28 = v8
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v12 = l0
	v13 = l1
	goto L5
L5:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v17 == int32(0) {
		v27 = v16
		v28 = v17
		goto L2
	} else {
		goto L7
	}
L6:
	;
	v27 = v16
	v28 = v17
	goto L2
L7:
	;
	v20 = int32(1)
	if v16 == v17 {
		v12 = v12 + v20
		v13 = v13 + v20
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
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
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v48 = v40 + v44
	v52 = int32(-2139062144)
	if (v46|(int32(16843008)-v46))&v52 == v52 {
		v40 = v48
		v41 = v46
		v42 = v45
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v57 = v48
	v58 = v46
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
