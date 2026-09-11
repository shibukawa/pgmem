package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int4_avg_accum_inv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v77 int64
	_ = v77
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+28)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v7 == int32(0) {
		v35 = int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		switch v10 - int32(429) {
		case 0:
			v35 = int32(1)
		case 1:
			v35 = int32(2)
		default:
			v35 = int32(0)
		}
	}
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v35 != 0 {
		v37 = F_pg_detoast_datum(m, v36)
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int32(0)
		} else {
			v43 = v37
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
			if v44 == int32(0) {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				if v47&int32(-4) == int32(160) {
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
					v72 = v43 + (v65<<(uint(int32(3))%32)+int32(23))&int32(-8)
					v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
					*(*int64)(unsafe.Add(mBase, uint32(v72))) = v73 - int64(1)
					v77 = *(*int64)(unsafe.Add(mBase, uint32(v72)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v72)+8)) = v77 - v4
					return v43
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_int4_avg_accum_inv_0), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_int4_avg_accum_inv_1), int32(_a_F_int4_avg_accum_inv_2), int32(_a_F_int4_avg_accum_inv_3))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_int4_avg_accum_inv_0), int32(0))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int4_avg_accum_inv_1), int32(_a_F_int4_avg_accum_inv_2), int32(_a_F_int4_avg_accum_inv_3))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
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
		v41 = F_pg_detoast_datum_copy(m, v36)
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return int32(0)
		} else {
			v43 = v41
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
			if v44 == int32(0) {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				if v47&int32(-4) == int32(160) {
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
					v72 = v43 + (v65<<(uint(int32(3))%32)+int32(23))&int32(-8)
					v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
					*(*int64)(unsafe.Add(mBase, uint32(v72))) = v73 - int64(1)
					v77 = *(*int64)(unsafe.Add(mBase, uint32(v72)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v72)+8)) = v77 - v4
					return v43
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_int4_avg_accum_inv_0), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_int4_avg_accum_inv_1), int32(_a_F_int4_avg_accum_inv_2), int32(_a_F_int4_avg_accum_inv_3))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_int4_avg_accum_inv_0), int32(0))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int4_avg_accum_inv_1), int32(_a_F_int4_avg_accum_inv_2), int32(_a_F_int4_avg_accum_inv_3))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func F_int4_bool(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	return base.B2i32(v2 != int32(0))
}
