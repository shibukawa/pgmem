package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tidhash_create(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14019(m, l0, l1, l2, int32(_a_F_tidhash_create_0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_tidhash_insert_hash(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+12)) = uint16(v10)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v12
	v16 = F_tidhash_insert_hash_internal(m, l0, v8+int32(8), l2, l3)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(16)
		return v16
	}
}
func F_tidhash_lookup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v28 int64
	_ = v28
	var v33 int64
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = l1 + int32(4)
	v13 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v12))))
	v14 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1))))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12))))
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+12)) = uint16(v15)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v22 = v13 << (uint(int64(32)) % 64)
	v23 = int64(33)
	v28 = (int64(base.Ui64(v22)>>(uint(v23)%64)) ^ (v22 | v14)) * int64(-49064778989728563)
	v33 = (int64(base.Ui64(v28)>>(uint(v23)%64)) ^ v28) * int64(-4265267296055464877)
	v38 = v20 & base.I32_wrap_i64(int64(base.Ui64(v33)>>(uint(v23)%64))^v33)
	v41 = v19 + v38<<(uint(int32(3))%32)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+6)))
	if v42 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v85
L2:
	;
	v44 = v41
	v48 = v38
	goto L5
L3:
	;
	goto L4
L4:
	;
	v85 = int32(0)
	goto L1
L5:
	;
	v50 = v9 + int32(8)
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+2)))
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44))))
	v53 = int32(16)
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+2)))
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50))))
	if v51|v52<<(uint(v53)%32) == v56|v57<<(uint(v53)%32) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L4
L7:
	;
	if v67 != 0 {
		v85 = v44
		goto L1
	} else {
		goto L13
	}
L8:
	;
	goto L7
L9:
	;
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+4)))
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+4)))
	if v63 == v64 {
		v67 = int32(1)
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v67 = int32(0)
	goto L8
L12:
	;
	goto L11
L13:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v72 = v69 & (v48 + int32(1))
	v75 = v68 + v72<<(uint(int32(3))%32)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+6)))
	if v76 != 0 {
		v44 = v75
		v48 = v72
		goto L5
	} else {
		goto L14
	}
L14:
	;
	goto L6
}
