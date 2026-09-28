package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_qsort_partition_hbound_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v6 < v7 {
		return int32(-1)
	} else {
		if v7 < v6 {
			return int32(1)
		} else {
			if v6 != v7 {
				v24 = v3
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v15 == v16 {
					v24 = v3
				} else {
					if v16 < v15 {
						v21 = int32(1)
					} else {
						v21 = int32(-1)
					}
					v24 = v21
				}
			}
			return v24
		}
	}
}
func F_qsort_partition_rbound_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
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
	var v32 int32
	_ = v32
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int64
	_ = v65
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v131 int32
	_ = v131
	v4 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+12)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+12)))
	v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+4)))
	if v20 <= v4 {
		v75 = v4
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v131
L2:
	;
	v110 = int32(0)
	if v98 < v110 {
		goto L22
	} else {
		goto L23
	}
L3:
	;
	v88 = int32(1)
	if v18&v88 != 0 {
		goto L16
	} else {
		goto L17
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v32 = int32(0)
	goto L5
L5:
	;
	v45 = v32 << (uint(int32(2)) % 32)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v27+v45)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v25)))
	if v47 < v49 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v75 = v20
	goto L3
L7:
	;
	return v32 ^ int32(-1)
L8:
	;
	goto L9
L9:
	;
	v55 = v32 + int32(1)
	if v49 < v47 {
		v131 = v55
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v47 != 0 {
		v75 = v55
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v45+v23)))
	v63 = v32 << (uint(int32(3)) % 32)
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v28+v63)))
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v63+v26)))
	v68 = F_FunctionCall2Coll(m, v24+v32*int32(28), v61, v65, v67)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	v72 = base.I32_wrap_i64(v68)
	if v72 != 0 {
		v97 = v55
		v98 = v72
		goto L2
	} else {
		goto L14
	}
L14:
	;
	if v20 != v55 {
		v32 = v55
		goto L5
	} else {
		goto L15
	}
L15:
	;
	goto L6
L16:
	;
	v92 = v88
	goto L18
L17:
	;
	v92 = int32(-1)
	goto L18
L18:
	;
	if v18 != v16 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v95 = v92
	goto L21
L20:
	;
	v95 = int32(0)
	goto L21
L21:
	;
	v97 = v75
	v98 = v95
	goto L2
L22:
	;
	v114 = v110 - v97
	goto L24
L23:
	;
	v114 = v97
	goto L24
L24:
	;
	if v98 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v116 = v114
	goto L27
L26:
	;
	v116 = int32(0)
	goto L27
L27:
	;
	v131 = v116
	goto L1
}
