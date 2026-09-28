package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_relation_is_publishable(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	v4 = int64(0)
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_SearchSysCache1(m, int32(57), v7&int64(4294967295))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		if v10 == int32(0) {
			v16 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v16)
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+22)))
			v22 = v20 + v21
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+119)))
			switch v23 - int32(112) {
			case 0, 2:
				v28 = base.I32_wrap_i64(v7)
				if base.Ui32(v28) < base.Ui32(int32(_a_F_pg_relation_is_publishable_0)) {
					v39 = v4
				} else {
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+118)))
					v39 = base.I64_extend_i32_u(base.B2i32(v31 == int32(112)) & base.B2i32(base.Ui32(int32(_a_F_pg_relation_is_publishable_1)) < base.Ui32(v28)))
				}
			case 1:
				v39 = v4
			default:
				if v23 != int32(83) {
					v39 = v4
				} else {
					v28 = base.I32_wrap_i64(v7)
					if base.Ui32(v28) < base.Ui32(int32(_a_F_pg_relation_is_publishable_0)) {
						v39 = v4
					} else {
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+118)))
						v39 = base.I64_extend_i32_u(base.B2i32(v31 == int32(112)) & base.B2i32(base.Ui32(int32(_a_F_pg_relation_is_publishable_1)) < base.Ui32(v28)))
					}
				}
			}
			F_ReleaseCatCache(m, v10)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int64(0)
			} else {
				return v39
			}
		}
	}
}
