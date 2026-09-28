package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BufFileSeekBlock(m *base.Module, l0 int32, l1 int64) int32 {
	var v5 int64
	_ = v5
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v5 = base.I64_div_s(l1, int64(131072))
	v13 = F_BufFileSeek(m, l0, base.I32_wrap_i64(v5), (l1-v5<<(uint(int64(17))%64))<<(uint(int64(13))%64), int32(0))
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		return v13
	}
}
func F_BufFileWrite(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v33 int64
	_ = v33
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v56 int32
	_ = v56
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = l0 + int32(40)
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = l1
	v16 = l2
	v17 = v13
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	if v17 < int64(8192) {
		v38 = v17
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v40 = base.I32_wrap_i64(v38)
	v41 = int32(_a_F_BufFileWrite_0) - v40
	if base.Ui32(v41) < base.Ui32(v16) {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	if v24 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_BufFileDumpBuffer(m, l0)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v30 + v17
	v33 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v33
	v38 = v33
	goto L6
L11:
	;
	return
L12:
	;
	v29 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v38 = v29
	goto L6
L13:
	;
	v43 = v41
	goto L15
L14:
	;
	v43 = v16
	goto L15
L15:
	;
	if v43 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	base.MemoryCopy(m, l0+int32(56)+v40, v15, v43)
	goto L18
L17:
	;
	goto L18
L18:
	;
	v46 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)) = uint8(v46)
	v48 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v50 = v48 + base.I64_extend_i32_u(v43)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v50
	v52 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	if v52 < v50 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v50
	goto L21
L20:
	;
	goto L21
L21:
	;
	v56 = v16 - v43
	if v56 != 0 {
		v15 = v15 + v43
		v16 = v56
		v17 = v50
		goto L4
	} else {
		goto L22
	}
L22:
	;
	goto L5
}
func F_BufTableInsert(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_BufTableInsert[0]))
	v14 = F_hash_search_with_hash_value(m, v10, l0, l1, int32(1), v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
			v24 = v21
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l2
			v24 = int32(-1)
		}
		m.G0 = v7 + int32(16)
		return v24
	}
}
