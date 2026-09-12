package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int2_accum_inv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int64
	_ = v23
	var v27 int64
	_ = v27
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v40 int64
	_ = v40
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v51 int64
	_ = v51
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v78 int64
	_ = v78
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v11 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v111 = m.ExcPending
		if v111 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(340827), int32(0))
			mBase = m.M
			v115 = m.ExcPending
			if v115 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(483087), int32(6101), int32(31264))
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v12 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v111 = m.ExcPending
			if v111 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(340827), int32(0))
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(483087), int32(6101), int32(31264))
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v15 == int32(0) {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
				if v19 == int32(1) {
					v23 = base.I64_extend16_s(base.I64_extend_i32_u(v18))
					v27 = v23 * (v23 >> (uint(int64(63)) % 64))
					v30 = int64(32)
					v31 = int64(base.Ui64(v23) >> (uint(v30) % 64))
					v36 = int64(4294967295)
					v37 = v23 & v36
					v40 = v37 * v37
					v43 = v37 * v31
					v44 = int64(base.Ui64(v40)>>(uint(v30)%64)) + v43
					v51 = v43 + v44&v36
					*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v27 + v27 + v31*v31 + int64(base.Ui64(v44)>>(uint(v30)%64)) + int64(base.Ui64(v51)>>(uint(v30)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v9))) = v40&v36 | v51<<(uint(v30)%64)
					v62 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
					v63 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
					*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v62 - v63
					v67 = v12 + int32(40)
					v68 = *(*int64)(unsafe.Add(mBase, uint32(v67)))
					v69 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v67))) = v68 - v69 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v62) < base.Ui64(v63)))
				} else {
				}
				v78 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v78 - int64(1)
				v82 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
				v84 = base.I64_extend16_s(base.I64_extend_i32_u(v18))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v82 - v84
				v88 = v12 + int32(24)
				v89 = *(*int64)(unsafe.Add(mBase, uint32(v88)))
				*(*int64)(unsafe.Add(mBase, uint32(v88))) = v89 - v84>>(uint(int64(63))%64) - base.I64_extend_i32_u(base.B2i32(base.Ui64(v82) < base.Ui64(v84)))
			} else {
			}
			m.G0 = v9 + int32(16)
			return v12
		}
	}
}
func F_int2_increment(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	v7 = base.B2i32(l1&int32(65535) == int32(32767))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v7)
	v10 = int32(16)
	if l1&int32(65535) == int32(32767) {
		v16 = int32(0)
	} else {
		v16 = (l1<<(uint(v10)%32) + int32(65536)) >> (uint(v10) % 32)
	}
	return v16
}
func F_int2_numeric(m *base.Module, l0 int32) int32 {
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
	v13 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+20)))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(0)
	v17 = F_palloc(m, int32(12))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
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
		v78 = F_make_result_opt_error(m, v11+int32(8), int32(0))
		mBase = m.M
		v79 = m.ExcPending
		if v79 != 0 {
			return int32(0)
		} else {
			F_pfree(m, v17)
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return int32(0)
			} else {
				m.G0 = v11 + int32(32)
				return v78
			}
		}
	}
}
