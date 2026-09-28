package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int4_accum_inv(m *base.Module, l0 int32) int64 {
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
			F_errmsg_internal(m, int32(_a_F_int4_accum_inv_0), int32(0))
			mBase = m.M
			v113 = m.ExcPending
			if v113 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_int4_accum_inv_1), int32(_a_F_int4_accum_inv_2), int32(_a_F_int4_accum_inv_3))
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
				F_errmsg_internal(m, int32(_a_F_int4_accum_inv_0), int32(0))
				mBase = m.M
				v113 = m.ExcPending
				if v113 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int4_accum_inv_1), int32(_a_F_int4_accum_inv_2), int32(_a_F_int4_accum_inv_3))
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
				v21 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+40)))
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
func F_int4_cash(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int64
	_ = v27
	var v34 int64
	_ = v34
	var v36 int32
	_ = v36
	var v42 int64
	_ = v42
	var v44 int32
	_ = v44
	var v51 int64
	_ = v51
	var v62 int64
	_ = v62
	var v64 int32
	_ = v64
	var v70 int64
	_ = v70
	var v72 int32
	_ = v72
	var v75 int64
	_ = v75
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v94 int64
	_ = v94
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v105 int64
	_ = v105
	var v112 int64
	_ = v112
	var v123 int64
	_ = v123
	var v124 int64
	_ = v124
	var v128 int64
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v147 int64
	_ = v147
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_PGLC_localeconv(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+41)))
		if base.Ui32(int32(10)) < base.Ui32(v19) {
			v22 = int32(2)
		} else {
			v22 = v19
		}
		v23 = base.I32_extend8_s(v22)
		if v23 <= int32(0) {
			v75 = int64(1)
		} else {
			v27 = int64(1)
			if base.Ui32(int32(8)) <= base.Ui32(v22) {
				v34 = v27
				v36 = int32(0)
				for {
					v42 = v34 * int64(100000000)
					v44 = v36 + int32(8)
					if v44 != v23&int32(120) {
						v34 = v42
						v36 = v44
						continue
					} else {
						break
					}
					break
				}
				if v22&int32(7) == int32(0) {
					v75 = v42
				} else {
					v51 = v42
					v62 = v51
					v64 = int32(0)
					for {
						v70 = v62 * int64(10)
						v72 = v64 + int32(1)
						if v72 != v23&int32(7) {
							v62 = v70
							v64 = v72
							continue
						} else {
							break
						}
						break
					}
					v75 = v70
				}
			} else {
				v51 = v27
				v62 = v51
				v64 = int32(0)
				for {
					v70 = v62 * int64(10)
					v72 = v64 + int32(1)
					if v72 != v23&int32(7) {
						v62 = v70
						v64 = v72
						continue
					} else {
						break
					}
					break
				}
				v75 = v70
			}
		}
		v82 = base.I64_extend32_s(v13)
		v83 = int64(63)
		v91 = int64(32)
		v92 = int64(base.Ui64(v75) >> (uint(v91) % 64))
		v94 = int64(base.Ui64(v82) >> (uint(v91) % 64))
		v97 = int64(4294967295)
		v98 = v75 & v97
		v100 = v82 & v97
		v101 = v98 * v100
		v105 = int64(base.Ui64(v101)>>(uint(v91)%64)) + v98*v94
		v112 = v100*v92 + v105&v97
		*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v82*(v75>>(uint(v83)%64)) + v82>>(uint(v83)%64)*v75 + v92*v94 + int64(base.Ui64(v105)>>(uint(v91)%64)) + int64(base.Ui64(v112)>>(uint(v91)%64))
		*(*int64)(unsafe.Add(mBase, uint32(v11))) = v101&v97 | v112<<(uint(v91)%64)
		v123 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
		v124 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		if v123 == v124>>(uint(int64(63))%64) {
			v147 = v124
			m.G0 = v11 + int32(16)
			return v147
		} else {
			v128 = int64(0)
			v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v130 = F_errsave_start(m, v129)
			mBase = m.M
			v131 = m.ExcPending
			if v131 != 0 {
				return int64(0)
			} else {
				if v130 == int32(0) {
					v147 = v128
					m.G0 = v11 + int32(16)
					return v147
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v136 = m.ExcPending
					if v136 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_int4_cash_0), int32(0))
						mBase = m.M
						v140 = m.ExcPending
						if v140 != 0 {
							return int64(0)
						} else {
							F_errsave_finish(m, v129, int32(_a_F_int4_cash_1), int32(1212), int32(_a_F_int4_cash_2))
							mBase = m.M
							v145 = m.ExcPending
							if v145 != 0 {
								return int64(0)
							} else {
								v147 = v128
								m.G0 = v11 + int32(16)
								return v147
							}
						}
					}
				}
			}
		}
	}
}
func F_int4_increment(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int64
	_ = v9
	var v15 int64
	_ = v15
	v6 = base.B2i32(base.I32_wrap_i64(l1) == int32(2147483647))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v6)
	v9 = int64(32)
	if base.I32_wrap_i64(l1) == int32(2147483647) {
		v15 = int64(0)
	} else {
		v15 = (l1<<(uint(v9)%64) + int64(4294967296)) >> (uint(v9) % 64)
	}
	return v15
}
func F_int4_sum(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v3 == int32(1) {
		v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
		if v6 == int32(1) {
			v9 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v9)
			return int64(0)
		} else {
			v13 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+40)))
			return v13
		}
	} else {
		v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
		if v16 != 0 {
			v19 = v15
		} else {
			v17 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+40)))
			v19 = v17 + v15
		}
		return v19
	}
}
