package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int4_accum_inv(m *base.Module, l0 int32) int32 {
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
	var v22 int64
	_ = v22
	var v26 int64
	_ = v26
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v39 int64
	_ = v39
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v50 int64
	_ = v50
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v81 int64
	_ = v81
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v11 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v109 = m.ExcPending
		if v109 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(336166), int32(0))
			mBase = m.M
			v113 = m.ExcPending
			if v113 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(476611), int32(6126), int32(30440))
				mBase = m.M
				v118 = m.ExcPending
				if v118 != 0 {
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
			v109 = m.ExcPending
			if v109 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(336166), int32(0))
				mBase = m.M
				v113 = m.ExcPending
				if v113 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(476611), int32(6126), int32(30440))
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
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
					v22 = base.I64_extend_i32_s(v18)
					v26 = v22 * (v22 >> (uint(int64(63)) % 64))
					v29 = int64(32)
					v30 = int64(base.Ui64(v22) >> (uint(v29) % 64))
					v35 = int64(4294967295)
					v36 = v22 & v35
					v39 = v36 * v36
					v42 = v36 * v30
					v43 = int64(base.Ui64(v39)>>(uint(v29)%64)) + v42
					v50 = v42 + v43&v35
					*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v26 + v26 + v30*v30 + int64(base.Ui64(v43)>>(uint(v29)%64)) + int64(base.Ui64(v50)>>(uint(v29)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v9))) = v39&v35 | v50<<(uint(v29)%64)
					v61 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
					v62 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
					*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v61 - v62
					v66 = v12 + int32(40)
					v67 = *(*int64)(unsafe.Add(mBase, uint32(v66)))
					v68 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v66))) = v67 - v68 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v61) < base.Ui64(v62)))
				} else {
				}
				v77 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
				v78 = base.I64_extend_i32_s(v18)
				*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v77 - v78
				v81 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v81 - int64(1)
				v86 = v12 + int32(24)
				v87 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
				*(*int64)(unsafe.Add(mBase, uint32(v86))) = v87 - v78>>(uint(int64(63))%64) - base.I64_extend_i32_u(base.B2i32(base.Ui64(v77) < base.Ui64(v78)))
			} else {
			}
			m.G0 = v9 + int32(16)
			return v12
		}
	}
}
func F_int4_cash(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v40 int64
	_ = v40
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v59 int64
	_ = v59
	var v61 int32
	_ = v61
	var v64 int64
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	v7 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
	v9 = F_PGLC_localeconv(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+41)))
		if base.Ui32(int32(10)) < base.Ui32(v13) {
			v16 = int32(2)
		} else {
			v16 = v13
		}
		v17 = base.I32_extend8_s(v16)
		if v17 <= int32(0) {
			v64 = int64(1)
		} else {
			if base.Ui32(v16) < base.Ui32(int32(8)) {
				v40 = int64(1)
			} else {
				v28 = int32(0)
				v29 = int64(1)
				for {
					v35 = v29 * int64(100000000)
					v37 = v28 + int32(8)
					if v37 != v17&int32(120) {
						v28 = v37
						v29 = v35
						continue
					} else {
						break
					}
					break
				}
				v40 = v35
			}
			if v16&int32(7) == int32(0) {
				v64 = v40
			} else {
				v52 = int32(0)
				v53 = v40
				for {
					v59 = v53 * int64(10)
					v61 = v52 + int32(1)
					if v61 != v17&int32(7) {
						v52 = v61
						v53 = v59
						continue
					} else {
						break
					}
					break
				}
				v64 = v59
			}
		}
		v71 = F_Int64GetDatum(m, v7)
		mBase = m.M
		v72 = m.ExcPending
		if v72 != 0 {
			return int32(0)
		} else {
			v73 = F_Int64GetDatum(m, v64)
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return int32(0)
			} else {
				v75 = F_DirectFunctionCall2Coll(m, int32(1278), int32(0), v71, v73)
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return int32(0)
				} else {
					v77 = *(*int64)(unsafe.Add(mBase, uint32(v75)))
					v78 = F_Int64GetDatum(m, v77)
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						return v78
					}
				}
			}
		}
	}
}
func F_int4_increment(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	v5 = base.B2i32(l1 == int32(2147483647))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v5)
	if l1 == int32(2147483647) {
		v10 = int32(0)
	} else {
		v10 = l1 + int32(1)
	}
	return v10
}
func F_int4_sum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v5 == int32(1) {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v8 == int32(1) {
			v11 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v11)
			return int32(0)
		} else {
			v15 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+28)))
			v16 = F_Int64GetDatum(m, v15)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				return v16
			}
		}
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v23 == int32(0) {
			v51 = int32(0)
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
			switch v26 - int32(429) {
			case 0:
				v51 = int32(1)
			case 1:
				v51 = int32(2)
			default:
				v51 = int32(0)
			}
		}
		v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v51 != 0 {
			v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v53 != 0 {
				v70 = v52
				return v70
			} else {
				v54 = *(*int64)(unsafe.Add(mBase, uint32(v52)))
				v55 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+28)))
				*(*int64)(unsafe.Add(mBase, uint32(v52))) = v54 + v55
				return v52
			}
		} else {
			v59 = *(*int64)(unsafe.Add(mBase, uint32(v52)))
			v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v60 == int32(1) {
				v63 = F_Int64GetDatum(m, v59)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					return v63
				}
			} else {
				v66 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+28)))
				v68 = F_Int64GetDatum(m, v59+v66)
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					v70 = v68
					return v70
				}
			}
		}
	}
}
