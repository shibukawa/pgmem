package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__ltxtq_exec(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14226(m, l0, int32(_a_F__ltxtq_exec_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_ltxtq_out(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		if v15 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(16801924))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_ltxtq_out_0), int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int64(0)
					} else {
						v31 = F_errdetail(m, int32(_a_F_ltxtq_out_1), int32(0))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_ltxtq_out_2), int32(617), int32(_a_F_ltxtq_out_3))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		} else {
			v38 = int32(32)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v38
			v41 = v11 + int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v41
			v45 = F_palloc_mul(m, int32(1), v38)
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v45
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v45
				v49 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v49)
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
				v52 = int32(12)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v41 + v51*v52
				F_infix_2(m, v8+v52, int32(1))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return int64(0)
				} else {
					v61 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v8)+16)))
					m.G0 = v8 + int32(32)
					return v61
				}
			}
		}
	}
}
