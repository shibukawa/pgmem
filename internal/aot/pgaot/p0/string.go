package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_StringAt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	v5 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l1 < v5 {
		v82 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v82
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 <= l1 {
		v82 = v5
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l3
	v17 = v15 + l1
	v21 = l3
	goto L4
L4:
	;
	v25 = v21 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v28 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v82 = int32(1)
	goto L1
L6:
	;
	v82 = int32(0)
	goto L1
L7:
	;
	goto L8
L8:
	;
	if l2 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v76 != 0 {
		v21 = v25
		goto L4
	} else {
		goto L22
	}
L10:
	;
	v76 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v37 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v38 = v17
	v39 = v27
	v40 = l2
	v41 = v37
	goto L17
L14:
	;
	v64 = v27
	v68 = int32(0)
	goto L15
L15:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v76 = v68 - v69
	goto L9
L16:
	;
	v64 = v59
	v68 = v61
	goto L15
L17:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if base.B2i32(v41 != v43)|base.B2i32(v43 == int32(0)) != 0 {
		v59 = v39
		v61 = v41
		goto L16
	} else {
		goto L19
	}
L18:
	;
	v59 = v53
	v61 = int32(0)
	goto L16
L19:
	;
	v49 = v40 - int32(1)
	if v49 == int32(0) {
		v59 = v39
		v61 = v41
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v52 = int32(1)
	v53 = v39 + v52
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	if v54 != 0 {
		v38 = v38 + v52
		v39 = v53
		v40 = v49
		v41 = v54
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	goto L5
}
func F_makeStringInfoExt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v5 = F_palloc(m, int32(16))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_palloc(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v9
			v13 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v13)
			*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v13
			*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v13
			return v5
		}
	}
}
func F_resetStringInfo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2))) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3
	return
}
func F_string_agg_finalfn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v6 != 0 {
		v31 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
		v34 = int32(0)
		return v34
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v7 == int32(0) {
			v31 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
			v34 = int32(0)
			return v34
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			v13 = v11 - v12
			v15 = v13 + int32(4)
			v16 = F_palloc(m, v15)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v15 << (uint(int32(2)) % 32)
				if v13 == int32(0) {
					v34 = v16
					return v34
				} else {
					base.MemoryCopy(m, v16+int32(4), v10+v12, v13)
					return v16
				}
			}
		}
	}
}
func F_string_agg_serialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pq_begintypsend(m, v7)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
		F_enlargeStringInfo(m, v7, int32(4))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v23 = int32(16711935)
			*(*int32)(unsafe.Add(mBase, uint32(v18+v19))) = base.I32_rotr(v14, int32(24))&v23 | base.I32_rotr(v14&v23, int32(8))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v18 + int32(4)
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			F_appendBinaryStringInfo(m, v7, v34, v35)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v39))) = v40 << (uint(int32(2)) % 32)
				m.G0 = v7 + int32(16)
				return v39
			}
		}
	}
}
