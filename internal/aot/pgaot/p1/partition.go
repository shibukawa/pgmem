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
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v36 int32
	_ = v36
	var v46 int64
	_ = v46
	var v48 int32
	_ = v48
	var v56 int64
	_ = v56
	v6 = int32(0)
	v7 = int64(0)
	if v6 < l0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = v6
	v16 = v7
	goto L4
L2:
	;
	v56 = v7
	goto L3
L3:
	;
	return v56
L4:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v15))))
	if v18 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v56 = v46
	goto L3
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l2+v15<<(uint(int32(2))%32))))
	v31 = *(*int64)(unsafe.Add(mBase, uint32(l3+v15<<(uint(int32(3))%32))))
	v33 = F_FunctionCall2Coll(m, l1+v15*int32(28), v27, v31, int64(8816678312871386365))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v46 = v16
	goto L8
L8:
	;
	v48 = v15 + int32(1)
	if v48 != l0 {
		v15 = v48
		v16 = v46
		goto L4
	} else {
		goto L11
	}
L9:
	;
	return int64(0)
L10:
	;
	v46 = v33 + (v16<<(uint(int64(54))%64) + int64(base.Ui64(v16)>>(uint(int64(7))%64))) + int64(5305509591434766563) ^ v16
	goto L8
L11:
	;
	goto L5
}
