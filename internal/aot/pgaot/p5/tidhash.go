package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tidhash_delete(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v19 int64
	_ = v19
	var v24 int64
	_ = v24
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var __phi80 int32
	_ = __phi80
	var v83 int32
	_ = v83
	var __phi83 int32
	_ = __phi83
	var v84 int32
	_ = v84
	var __phi84 int32
	_ = __phi84
	var v86 int32
	_ = v86
	var __phi86 int32
	_ = __phi86
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v96 int64
	_ = v96
	var v101 int64
	_ = v101
	var v108 int64
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v12 = v10 << (uint(int64(32)) % 64)
	v13 = int64(33)
	v15 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1))))
	v19 = (int64(base.Ui64(v12)>>(uint(v13)%64)) ^ (v15 | v12)) * int64(-49064778989728563)
	v24 = (int64(base.Ui64(v19)>>(uint(v13)%64)) ^ v19) * int64(-4265267296055464877)
	v33 = v9 & base.I32_wrap_i64(int64(base.Ui64(v24)>>(uint(v13)%64))^v24)
	v34 = v9
	goto L1
L1:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v42 = v39 + v33<<(uint(int32(3))%32)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+6)))
	switch v43 {
	case 0:
		v137 = int32(0)
		goto L4
	case 1:
		goto L5
	default:
		v140 = v34
		goto L3
	}
L3:
	;
	v33 = v140 & (v33 + int32(1))
	v34 = v140
	goto L1
L4:
	;
	return v137
L5:
	;
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+2)))
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42))))
	v46 = int32(16)
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	if v44|v45<<(uint(v46)%32) == v49|v50<<(uint(v46)%32) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v60 == int32(0) {
		v140 = v61
		goto L3
	} else {
		goto L12
	}
L7:
	;
	goto L6
L8:
	;
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+4)))
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v56 == v57 {
		v60 = int32(1)
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v60 = int32(0)
	goto L7
L11:
	;
	goto L10
L12:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v65 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v64 - v65
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v72 = v61 & (v33 + v65)
	v75 = v69 + v72<<(uint(int32(3))%32)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+6)))
	if v76 != v65 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v129 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v124)+6)) = uint8(v129)
	v137 = v65
	goto L4
L14:
	;
	v124 = v42
	goto L13
L15:
	;
	goto L16
L16:
	;
	__phi80 = v72
	__phi83 = v61
	__phi84 = v42
	__phi86 = v75
	v80 = __phi80
	v83 = __phi83
	v84 = __phi84
	v86 = __phi86
	goto L17
L17:
	;
	v87 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v86)+4)))
	v89 = v87 << (uint(int64(32)) % 64)
	v90 = int64(33)
	v92 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v86))))
	v96 = (int64(base.Ui64(v89)>>(uint(v90)%64)) ^ (v92 | v89)) * int64(-49064778989728563)
	v101 = (int64(base.Ui64(v96)>>(uint(v90)%64)) ^ v96) * int64(-4265267296055464877)
	if v80 == v83&base.I32_wrap_i64(int64(base.Ui64(v101)>>(uint(v90)%64))^v101) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v124 = v86
	goto L13
L19:
	;
	v124 = v84
	goto L13
L20:
	;
	goto L21
L21:
	;
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
	*(*int64)(unsafe.Add(mBase, uint32(v84))) = v108
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v112 = int32(1)
	v114 = v111 & (v80 + v112)
	v117 = v110 + v114<<(uint(int32(3))%32)
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+6)))
	if v118 == v112 {
		__phi80 = v114
		__phi83 = v111
		__phi84 = v86
		__phi86 = v117
		v80 = __phi80
		v83 = __phi83
		v84 = __phi84
		v86 = __phi86
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L18
}
func F_tidhash_lookup_hash(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v9
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v11)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v15 = v14 & l2
	v18 = v13 + v15<<(uint(int32(3))%32)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+6)))
	if v19 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return v58
L2:
	;
	v21 = v18
	v22 = v15
	goto L5
L3:
	;
	goto L4
L4:
	;
	v58 = int32(0)
	goto L1
L5:
	;
	v25 = v7 + int32(8)
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+2)))
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21))))
	v28 = int32(16)
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+2)))
	v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	if v26|v27<<(uint(v28)%32) == v31|v32<<(uint(v28)%32) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L4
L7:
	;
	if v42 != 0 {
		v58 = v21
		goto L1
	} else {
		goto L13
	}
L8:
	;
	goto L7
L9:
	;
	v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)))
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
	if v38 == v39 {
		v42 = int32(1)
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v42 = int32(0)
	goto L8
L12:
	;
	goto L11
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v47 = v44 & (v22 + int32(1))
	v50 = v43 + v47<<(uint(int32(3))%32)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+6)))
	if v51 != 0 {
		v21 = v50
		v22 = v47
		goto L5
	} else {
		goto L14
	}
L14:
	;
	goto L6
}
func F_tidhash_start_iterate_at(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v5 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)) = uint8(v5)
	v7 = v4 & l2
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v7
	return
}
