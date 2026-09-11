package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_timetz_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v12 = int64(1000000)
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	return base.B2i32(v6 == v8) & base.B2i32(v10+base.I64_extend_i32_s(v6)*v12 == v15+base.I64_extend_i32_s(v8)*v12)
}
func F_timetz_hash(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_DirectFunctionCall1Coll(m, int32(1284), int32(0), v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		v14 = int32(711645284)
		v17 = v9 - int32(1636608428) ^ v14 - int32(1455628627)
		v22 = v17 ^ int32(-1636608428) - base.I32_rotl(v17, int32(25))
		v27 = v22 ^ v14 - base.I32_rotl(v22, int32(16))
		v31 = v27 ^ v17 - base.I32_rotl(v27, int32(4))
		v35 = v31 ^ v22 - base.I32_rotl(v31, int32(14))
		return v5 ^ (v35 ^ v27 - base.I32_rotl(v35, int32(24)))
	}
}
func F_timetz_send(m *base.Module, l0 int32) int32 {
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
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v34 int64
	_ = v34
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
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
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
		F_enlargeStringInfo(m, v7, int32(8))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v21 = int64(56)
			v23 = int64(65280)
			v25 = int64(40)
			v28 = int64(16711680)
			v30 = int64(24)
			v32 = int64(4278190080)
			v34 = int64(8)
			*(*int64)(unsafe.Add(mBase, uint32(v18+v19))) = v14<<(uint(v21)%64) | v14&v23<<(uint(v25)%64) | (v14&v28<<(uint(v30)%64) | v14&v32<<(uint(v34)%64)) | (int64(base.Ui64(v14)>>(uint(v34)%64))&v32 | int64(base.Ui64(v14)>>(uint(v30)%64))&v28 | (int64(base.Ui64(v14)>>(uint(v25)%64))&v23 | int64(base.Ui64(v14)>>(uint(v21)%64))))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v18 + int32(8)
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
			F_enlargeStringInfo(m, v7, int32(4))
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return int32(0)
			} else {
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v67 = int32(24)
				v69 = int32(65280)
				v71 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v64+v65))) = v60<<(uint(v67)%32) | v60&v69<<(uint(v71)%32) | (int32(base.Ui32(v60)>>(uint(v71)%32))&v69 | int32(base.Ui32(v60)>>(uint(v67)%32)))
				v84 = v64 + int32(4)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v84
				v87 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				*(*int32)(unsafe.Add(mBase, uint32(v87))) = v84 << (uint(int32(2)) % 32)
				m.G0 = v7 + int32(16)
				return v87
			}
		}
	}
}
