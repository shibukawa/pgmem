package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pointerhash_delete(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
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
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var __phi52 int32
	_ = __phi52
	var v54 int32
	_ = v54
	var __phi54 int32
	_ = __phi54
	var v55 int32
	_ = v55
	var __phi55 int32
	_ = __phi55
	var v56 int32
	_ = v56
	var __phi56 int32
	_ = __phi56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v74 int64
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	v8 = int32(16)
	v12 = (int32(base.Ui32(l1)>>(uint(v8)%32)) ^ l1) * int32(-2048144789)
	v17 = (int32(base.Ui32(v12)>>(uint(int32(13))%32)) ^ v12) * int32(-1028477387)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v27 = int32(base.Ui32(v17)>>(uint(v8)%32)) ^ v17
	goto L1
L1:
	;
	v30 = v27 & v22
	v33 = v21 + v30<<(uint(int32(3))%32)
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
	switch v34 {
	case 0:
		v102 = int32(0)
		goto L4
	case 1:
		goto L5
	default:
		goto L3
	}
L3:
	;
	v27 = v30 + int32(1)
	goto L1
L4:
	;
	return v102
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v35 != l1 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v38 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v37 - v38
	v44 = v22 & (v30 + v38)
	v47 = v21 + v44<<(uint(int32(3))%32)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+4)))
	if v48 != v38 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v94 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+4)) = uint8(v94)
	v102 = v38
	goto L4
L8:
	;
	v89 = v33
	goto L7
L9:
	;
	goto L10
L10:
	;
	__phi52 = v44
	__phi54 = v47
	__phi55 = v33
	__phi56 = v22
	v52 = __phi52
	v54 = __phi54
	v55 = __phi55
	v56 = __phi56
	goto L11
L11:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v59 = int32(16)
	v63 = (int32(base.Ui32(v58)>>(uint(v59)%32)) ^ v58) * int32(-2048144789)
	v68 = (int32(base.Ui32(v63)>>(uint(int32(13))%32)) ^ v63) * int32(-1028477387)
	if v52 == (int32(base.Ui32(v68)>>(uint(v59)%32))^v68)&v56 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v89 = v54
	goto L7
L13:
	;
	v89 = v55
	goto L7
L14:
	;
	goto L15
L15:
	;
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v54)))
	*(*int64)(unsafe.Add(mBase, uint32(v55))) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v78 = int32(1)
	v80 = v77 & (v52 + v78)
	v83 = v76 + v80<<(uint(int32(3))%32)
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+4)))
	if v84 == v78 {
		__phi52 = v80
		__phi54 = v83
		__phi55 = v54
		__phi56 = v77
		v52 = __phi52
		v54 = __phi54
		v55 = __phi55
		v56 = __phi56
		goto L11
	} else {
		goto L16
	}
L16:
	;
	goto L12
}
func F_pointerhash_delete_item(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var __phi26 int32
	_ = __phi26
	var v27 int32
	_ = v27
	var __phi27 int32
	_ = __phi27
	var v29 int32
	_ = v29
	var __phi29 int32
	_ = __phi29
	var v30 int32
	_ = v30
	var __phi30 int32
	_ = __phi30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v47 int64
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v7 - v8
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v14 = int32(3)
	v18 = v12 & ((l1-v11)>>(uint(v14)%32) + v8)
	v21 = v11 + v18<<(uint(v14)%32)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)))
	if v22 != v8 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v66 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)) = uint8(v66)
	return
L2:
	;
	v63 = l1
	goto L1
L3:
	;
	goto L4
L4:
	;
	__phi26 = l1
	__phi27 = v21
	__phi29 = v18
	__phi30 = v12
	v26 = __phi26
	v27 = __phi27
	v29 = __phi29
	v30 = __phi30
	goto L5
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v32 = int32(16)
	v36 = (int32(base.Ui32(v31)>>(uint(v32)%32)) ^ v31) * int32(-2048144789)
	v41 = (int32(base.Ui32(v36)>>(uint(int32(13))%32)) ^ v36) * int32(-1028477387)
	if v29 == (int32(base.Ui32(v41)>>(uint(v32)%32))^v41)&v30 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v63 = v27
	goto L1
L7:
	;
	v63 = v26
	goto L1
L8:
	;
	goto L9
L9:
	;
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v27)))
	*(*int64)(unsafe.Add(mBase, uint32(v26))) = v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v51 = int32(1)
	v53 = v50 & (v29 + v51)
	v56 = v49 + v53<<(uint(int32(3))%32)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+4)))
	if v57 == v51 {
		__phi26 = v27
		__phi27 = v56
		__phi29 = v53
		__phi30 = v50
		v26 = __phi26
		v27 = __phi27
		v29 = __phi29
		v30 = __phi30
		goto L5
	} else {
		goto L10
	}
L10:
	;
	goto L6
}
func F_pointerhash_insert(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v4 = int32(16)
	v8 = (int32(base.Ui32(l1)>>(uint(v4)%32)) ^ l1) * int32(-2048144789)
	v13 = (int32(base.Ui32(v8)>>(uint(int32(13))%32)) ^ v8) * int32(-1028477387)
	v17 = F_pointerhash_insert_hash_internal(m, l0, l1, int32(base.Ui32(v13)>>(uint(v4)%32))^v13, l2)
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		return v17
	}
}
func F_pointerhash_lookup_hash(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v8 = v7 & l2
	v11 = v6 + v8<<(uint(int32(3))%32)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+4)))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v17 = v11
	v19 = v8
	goto L5
L4:
	;
	return v17
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v22 == l1 {
		goto L4
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v26 = (v19 + int32(1)) & v7
	v29 = v6 + v26<<(uint(int32(3))%32)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+4)))
	if v30 != 0 {
		v17 = v29
		v19 = v26
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
}
