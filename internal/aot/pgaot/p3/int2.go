package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int2_accum_inv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v43 int64
	_ = v43
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v54 int64
	_ = v54
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v76 int64
	_ = v76
	var v79 int64
	_ = v79
	var v82 int64
	_ = v82
	var v86 int64
	_ = v86
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v13 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v109 = m.ExcPending
		if v109 != 0 {
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_int2_accum_inv_0), int32(0))
			mBase = m.M
			v113 = m.ExcPending
			if v113 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_int2_accum_inv_1), int32(_a_F_int2_accum_inv_2), int32(_a_F_int2_accum_inv_3))
				mBase = m.M
				v118 = m.ExcPending
				if v118 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v15 = base.I32_wrap_i64(v14)
		if v15 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v109 = m.ExcPending
			if v109 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_int2_accum_inv_0), int32(0))
				mBase = m.M
				v113 = m.ExcPending
				if v113 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int2_accum_inv_1), int32(_a_F_int2_accum_inv_2), int32(_a_F_int2_accum_inv_3))
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
			if v18 == int32(0) {
				v21 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
				if v22 == int32(0) {
					v76 = v21 >> (uint(int64(63)) % 64)
				} else {
					v28 = v21 >> (uint(int64(63)) % 64)
					v30 = v21 * v28
					v33 = int64(32)
					v34 = int64(base.Ui64(v21) >> (uint(v33) % 64))
					v39 = int64(4294967295)
					v40 = v21 & v39
					v43 = v40 * v40
					v46 = v40 * v34
					v47 = int64(base.Ui64(v43)>>(uint(v33)%64)) + v46
					v54 = v46 + v47&v39
					*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v30 + v30 + v34*v34 + int64(base.Ui64(v47)>>(uint(v33)%64)) + int64(base.Ui64(v54)>>(uint(v33)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v11))) = v43&v39 | v54<<(uint(v33)%64)
					v65 = *(*int64)(unsafe.Add(mBase, uint32(v15)+32))
					v66 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
					*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = v65 - v66
					v69 = *(*int64)(unsafe.Add(mBase, uint32(v15)+40))
					v70 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v15)+40)) = v69 - v70 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v65) < base.Ui64(v66)))
					v76 = v28
				}
				v79 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v79 - v21
				v82 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v82 - int64(1)
				v86 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v86 - v76 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v79) < base.Ui64(v21)))
			} else {
			}
			m.G0 = v11 + int32(16)
			return v14 & int64(4294967295)
		}
	}
}
func F_int2_increment(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int64
	_ = v10
	var v16 int64
	_ = v16
	v7 = base.B2i32(l1&int64(65535) == int64(32767))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v7)
	v10 = int64(48)
	if l1&int64(65535) == int64(32767) {
		v16 = int64(0)
	} else {
		v16 = (l1<<(uint(v10)%64) - int64(-281474976710656)) >> (uint(v10) % 64)
	}
	return v16
}
func F_int2_numeric(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int64
	_ = v28
	var v35 int64
	_ = v35
	var v41 int64
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int64
	_ = v50
	var v53 int32
	_ = v53
	var v55 int64
	_ = v55
	var v58 int64
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(0)
	v17 = F_palloc(m, int32(12))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v17
		v22 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v17))) = uint16(v22)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v17 + int32(2)
		v28 = base.I64_extend16_s(v13)
		if v28 < int64(0) {
			*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(16384)
			v41 = int64(0) - v28
			v44 = v22
			v47 = v17 + int32(12)
			v50 = v41
			for {
				v53 = v47 - int32(2)
				v55 = base.I64_div_u_s(v50, int64(10000))
				v58 = v55*int64(55536) + v50
				*(*uint16)(unsafe.Add(mBase, uint32(v53))) = uint16(v58)
				v61 = v44 + int32(1)
				if base.Ui64(int64(9999)) < base.Ui64(v50) {
					v44 = v61
					v47 = v53
					v50 = v55
					continue
				} else {
					break
				}
				break
			}
			*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v53
			v65 = v61
			v69 = v44
		} else {
			v35 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v35
			if v13<<(uint(int64(48))%64) == v35 {
				v65 = v22
				v69 = int32(0)
			} else {
				v41 = v28
				v44 = v22
				v47 = v17 + int32(12)
				v50 = v41
				for {
					v53 = v47 - int32(2)
					v55 = base.I64_div_u_s(v50, int64(10000))
					v58 = v55*int64(55536) + v50
					*(*uint16)(unsafe.Add(mBase, uint32(v53))) = uint16(v58)
					v61 = v44 + int32(1)
					if base.Ui64(int64(9999)) < base.Ui64(v50) {
						v44 = v61
						v47 = v53
						v50 = v55
						continue
					} else {
						break
					}
					break
				}
				*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v53
				v65 = v61
				v69 = v44
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v69
		*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v65
		v78 = F_make_result_safe(m, v11+int32(8), int32(0))
		mBase = m.M
		v79 = m.ExcPending
		if v79 != 0 {
			return int64(0)
		} else {
			F_pfree(m, v17)
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return int64(0)
			} else {
				m.G0 = v11 + int32(32)
				return base.I64_extend_i32_u(v78)
			}
		}
	}
}
