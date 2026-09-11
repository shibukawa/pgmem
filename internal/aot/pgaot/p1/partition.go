package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_PartitionDirectoryLookup(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = F_hash_search(m, v11, v7+int32(12), int32(1), v7+int32(11))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+11)))
		if v21 == int32(1) {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
			v32 = v24
			m.G0 = v7 + int32(16)
			return v32
		} else {
			F_RelationIncrementReferenceCount(m, l1)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = l1
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
				v29 = F_RelationGetPartitionDesc(m, l1, v28)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v29
					v32 = v29
					m.G0 = v7 + int32(16)
					return v32
				}
			}
		}
	}
}
func F__equalPartitionCmd(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6 = F_equal(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v21 = v3
			return v21
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v14 = F_equal(m, v12, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if v14 == int32(0) {
					v21 = v3
				} else {
					v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
					v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
					v21 = base.B2i32(v18 == v19)
				}
				return v21
			}
		}
	}
}
func F_compute_partition_hash_value(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v51 int64
	_ = v51
	var v54 int32
	_ = v54
	var v61 int64
	_ = v61
	v6 = int64(0)
	v11 = F_Int64GetDatum(m, int64(8816678312871386365))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	if int32(0) < l0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v22 = v6
	v23 = int32(0)
	goto L6
L4:
	;
	v61 = v6
	goto L5
L5:
	;
	return v61
L6:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v23))))
	if v27 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v61 = v51
	goto L5
L8:
	;
	v34 = v23 << (uint(int32(2)) % 32)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2+v34)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l3+v34)))
	v39 = F_FunctionCall2Coll(m, l1+v23*int32(28), v36, v38, v11)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v51 = v22
	goto L10
L10:
	;
	v54 = v23 + int32(1)
	if v54 != l0 {
		v22 = v51
		v23 = v54
		goto L6
	} else {
		goto L12
	}
L11:
	;
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
	v51 = v41 + (v22<<(uint(int64(54))%64) + int64(base.Ui64(v22)>>(uint(int64(7))%64))) + int64(5305509591434766563) ^ v22
	goto L10
L12:
	;
	goto L7
}
