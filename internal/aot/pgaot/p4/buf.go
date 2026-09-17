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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int64
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = l1
	v12 = l2
	v13 = v9
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	if v13 < int32(_a_F_BufFileWrite_0) {
		v31 = v13
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v33 = int32(_a_F_BufFileWrite_0) - v31
	if base.Ui32(v33) < base.Ui32(v12) {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	if v18 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_BufFileDumpBuffer(m, l0)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = int64(0)
	v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v26 + base.I64_extend_i32_u(v13)
	v31 = int32(0)
	goto L6
L11:
	;
	return
L12:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v31 = v23
	goto L6
L13:
	;
	v35 = v33
	goto L15
L14:
	;
	v35 = v12
	goto L15
L15:
	;
	if v35 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	base.MemoryCopy(m, v31+(l0+int32(48)), v11, v35)
	goto L18
L17:
	;
	goto L18
L18:
	;
	v38 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)) = uint8(v38)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v41 = v40 + v35
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v41
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v43 < v41 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v41
	goto L21
L20:
	;
	goto L21
L21:
	;
	v47 = v12 - v35
	if v47 != 0 {
		v11 = v11 + v35
		v12 = v47
		v13 = v41
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
