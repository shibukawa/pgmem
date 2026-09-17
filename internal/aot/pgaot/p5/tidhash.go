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
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var __phi79 int32
	_ = __phi79
	var v82 int32
	_ = v82
	var __phi82 int32
	_ = __phi82
	var v83 int32
	_ = v83
	var __phi83 int32
	_ = __phi83
	var v84 int32
	_ = v84
	var __phi84 int32
	_ = __phi84
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v95 int64
	_ = v95
	var v100 int64
	_ = v100
	var v107 int64
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
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
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v41 = v38 + v33<<(uint(int32(3))%32)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+6)))
	switch v42 {
	case 0:
		v137 = int32(0)
		goto L4
	case 1:
		goto L5
	default:
		v139 = v34
		goto L3
	}
L3:
	;
	v33 = v139 & (v33 + int32(1))
	v34 = v139
	goto L1
L4:
	;
	return v137
L5:
	;
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+2)))
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41))))
	v45 = int32(16)
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	if v43|v44<<(uint(v45)%32) == v48|v49<<(uint(v45)%32) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v59 == int32(0) {
		v139 = v60
		goto L3
	} else {
		goto L12
	}
L7:
	;
	goto L6
L8:
	;
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+4)))
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v55 == v56 {
		v59 = int32(1)
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v59 = int32(0)
	goto L7
L11:
	;
	goto L10
L12:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v64 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v63 - v64
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v71 = v60 & (v33 + v64)
	v74 = v68 + v71<<(uint(int32(3))%32)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+6)))
	if v75 != v64 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v128 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v123)+6)) = uint8(v128)
	v137 = v64
	goto L4
L14:
	;
	v123 = v41
	goto L13
L15:
	;
	goto L16
L16:
	;
	__phi79 = v71
	__phi82 = v60
	__phi83 = v41
	__phi84 = v74
	v79 = __phi79
	v82 = __phi82
	v83 = __phi83
	v84 = __phi84
	goto L17
L17:
	;
	v86 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v84)+4)))
	v88 = v86 << (uint(int64(32)) % 64)
	v89 = int64(33)
	v91 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v84))))
	v95 = (int64(base.Ui64(v88)>>(uint(v89)%64)) ^ (v91 | v88)) * int64(-49064778989728563)
	v100 = (int64(base.Ui64(v95)>>(uint(v89)%64)) ^ v95) * int64(-4265267296055464877)
	if v79 == v82&base.I32_wrap_i64(int64(base.Ui64(v100)>>(uint(v89)%64))^v100) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v123 = v84
	goto L13
L19:
	;
	v123 = v83
	goto L13
L20:
	;
	goto L21
L21:
	;
	v107 = *(*int64)(unsafe.Add(mBase, uint32(v84)))
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v107
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v111 = int32(1)
	v113 = v110 & (v79 + v111)
	v116 = v109 + v113<<(uint(int32(3))%32)
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+6)))
	if v117 == v111 {
		__phi79 = v113
		__phi82 = v110
		__phi83 = v84
		__phi84 = v116
		v79 = __phi79
		v82 = __phi82
		v83 = __phi83
		v84 = __phi84
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
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v9)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v11
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
