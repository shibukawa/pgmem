package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int4_accum_inv(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v38 int64
	_ = v38
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v49 int64
	_ = v49
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v77 int64
	_ = v77
	var v81 int64
	_ = v81
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v10 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v102 = m.ExcPending
		if v102 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_int4_accum_inv_0), int32(0))
			mBase = m.M
			v106 = m.ExcPending
			if v106 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_int4_accum_inv_1), int32(_a_F_int4_accum_inv_2), int32(_a_F_int4_accum_inv_3))
				mBase = m.M
				v111 = m.ExcPending
				if v111 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v11 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v102 = m.ExcPending
			if v102 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_int4_accum_inv_0), int32(0))
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int4_accum_inv_1), int32(_a_F_int4_accum_inv_2), int32(_a_F_int4_accum_inv_3))
					mBase = m.M
					v111 = m.ExcPending
					if v111 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v14 == int32(0) {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
				if v18 == int32(1) {
					v21 = base.I64_extend_i32_s(v17)
					v25 = v21 * (v21 >> (uint(int64(63)) % 64))
					v28 = int64(32)
					v29 = int64(base.Ui64(v21) >> (uint(v28) % 64))
					v34 = int64(4294967295)
					v35 = v21 & v34
					v38 = v35 * v35
					v41 = v35 * v29
					v42 = int64(base.Ui64(v38)>>(uint(v28)%64)) + v41
					v49 = v41 + v42&v34
					*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v25 + v25 + v29*v29 + int64(base.Ui64(v42)>>(uint(v28)%64)) + int64(base.Ui64(v49)>>(uint(v28)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v8))) = v38&v34 | v49<<(uint(v28)%64)
					v60 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
					v61 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
					*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v60 - v61
					v64 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
					v65 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v64 - v65 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v60) < base.Ui64(v61)))
				} else {
				}
				v73 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
				v74 = base.I64_extend_i32_s(v17)
				*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v73 - v74
				v77 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v77 - int64(1)
				v81 = *(*int64)(unsafe.Add(mBase, uint32(v11)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v81 - v74>>(uint(int64(63))%64) - base.I64_extend_i32_u(base.B2i32(base.Ui64(v73) < base.Ui64(v74)))
			} else {
			}
			m.G0 = v8 + int32(16)
			return v11
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
	var v21 int64
	_ = v21
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v34 int64
	_ = v34
	var v36 int32
	_ = v36
	var v43 int64
	_ = v43
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v58 int64
	_ = v58
	var v60 int32
	_ = v60
	var v63 int64
	_ = v63
	var v70 int32
	_ = v70
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
	var v76 int64
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
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
			v63 = int64(1)
		} else {
			v21 = int64(1)
			if base.Ui32(int32(8)) <= base.Ui32(v16) {
				v27 = int32(0)
				v28 = v21
				for {
					v34 = v28 * int64(100000000)
					v36 = v27 + int32(8)
					if v36 != v17&int32(120) {
						v27 = v36
						v28 = v34
						continue
					} else {
						break
					}
					break
				}
				if v16&int32(7) == int32(0) {
					v63 = v34
				} else {
					v43 = v34
					v51 = int32(0)
					v52 = v43
					for {
						v58 = v52 * int64(10)
						v60 = v51 + int32(1)
						if v60 != v17&int32(7) {
							v51 = v60
							v52 = v58
							continue
						} else {
							break
						}
						break
					}
					v63 = v58
				}
			} else {
				v43 = v21
				v51 = int32(0)
				v52 = v43
				for {
					v58 = v52 * int64(10)
					v60 = v51 + int32(1)
					if v60 != v17&int32(7) {
						v51 = v60
						v52 = v58
						continue
					} else {
						break
					}
					break
				}
				v63 = v58
			}
		}
		v70 = F_Int64GetDatum(m, v7)
		mBase = m.M
		v71 = m.ExcPending
		if v71 != 0 {
			return int32(0)
		} else {
			v72 = F_Int64GetDatum(m, v63)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				v74 = F_DirectFunctionCall2Coll(m, int32(1263), int32(0), v70, v72)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					v76 = *(*int64)(unsafe.Add(mBase, uint32(v74)))
					v77 = F_Int64GetDatum(m, v76)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						return v77
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
