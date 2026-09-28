package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_timetz_eq(m *base.Module, l0 int32) int64 {
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v12 = int64(1000000)
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	return base.I64_extend_i32_u(base.B2i32(v6 == v8) & base.B2i32(v10+base.I64_extend_i32_s(v6)*v12 == v15+base.I64_extend_i32_s(v8)*v12))
}
func F_timetz_hash(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v7 = F_DirectFunctionCall1Coll(m, int32(1397), int32(0), v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v16 = int32(711645284)
		v19 = v11 - int32(1636608428) ^ v16 - int32(1455628627)
		v24 = v19 ^ int32(-1636608428) - base.I32_rotl(v19, int32(25))
		v29 = v24 ^ v16 - base.I32_rotl(v24, int32(16))
		v33 = v29 ^ v19 - base.I32_rotl(v29, int32(4))
		v37 = v33 ^ v24 - base.I32_rotl(v33, int32(14))
		return base.I64_extend_i32_u(v37 ^ v29 - base.I32_rotl(v37, int32(24)) ^ base.I32_wrap_i64(v7))
	}
}
func F_timetz_send(m *base.Module, l0 int32) int64 {
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
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_pq_begintypsend(m, v7)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
		F_enlargeStringInfo(m, v7, int32(8))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
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
				return int64(0)
			} else {
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v69 = int32(16711935)
				*(*int32)(unsafe.Add(mBase, uint32(v64+v65))) = base.I32_rotr(v60, int32(24))&v69 | base.I32_rotr(v60&v69, int32(8))
				v78 = v64 + int32(4)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v78
				v81 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				*(*int32)(unsafe.Add(mBase, uint32(v81))) = v78 << (uint(int32(2)) % 32)
				m.G0 = v7 + int32(16)
				return base.I64_extend_i32_u(v81)
			}
		}
	}
}
