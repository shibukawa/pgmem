package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int2_accum(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v93 int64
	_ = v93
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v104 int64
	_ = v104
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v121 int64
	_ = v121
	var v122 int64
	_ = v122
	var v126 int64
	_ = v126
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v132 int64
	_ = v132
	var v138 int64
	_ = v138
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v13 == int32(0) {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v16 != 0 {
			v66 = v16
			v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
			if v68 == int32(0) {
				v71 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
				v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
				if v72 == int32(0) {
					v126 = v71 >> (uint(int64(63)) % 64)
				} else {
					v78 = v71 >> (uint(int64(63)) % 64)
					v80 = v71 * v78
					v83 = int64(32)
					v84 = int64(base.Ui64(v71) >> (uint(v83) % 64))
					v89 = int64(4294967295)
					v90 = v71 & v89
					v93 = v90 * v90
					v96 = v90 * v84
					v97 = int64(base.Ui64(v93)>>(uint(v83)%64)) + v96
					v104 = v96 + v97&v89
					*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v80 + v80 + v84*v84 + int64(base.Ui64(v97)>>(uint(v83)%64)) + int64(base.Ui64(v104)>>(uint(v83)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v11))) = v93&v89 | v104<<(uint(v83)%64)
					v115 = *(*int64)(unsafe.Add(mBase, uint32(v66)+32))
					v116 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
					v117 = v115 + v116
					*(*int64)(unsafe.Add(mBase, uint32(v66)+32)) = v117
					v121 = *(*int64)(unsafe.Add(mBase, uint32(v66)+40))
					v122 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v66)+40)) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v117) < base.Ui64(v115))) + (v121 + v122)
					v126 = v78
				}
				v129 = *(*int64)(unsafe.Add(mBase, uint32(v66)+16))
				v130 = v71 + v129
				*(*int64)(unsafe.Add(mBase, uint32(v66)+16)) = v130
				v132 = *(*int64)(unsafe.Add(mBase, uint32(v66)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v66)+8)) = v132 + int64(1)
				v138 = *(*int64)(unsafe.Add(mBase, uint32(v66)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v66)+24)) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v130) < base.Ui64(v129))) + (v138 + v126)
			} else {
			}
			m.G0 = v11 + int32(32)
			return base.I64_extend_i32_u(v66)
		} else {
			v19 = v11 + int32(28)
			v20 = int32(0)
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v21 == v20 {
				v38 = int32(0)
				if v19 == v38 {
					v46 = v38
				} else {
					v41 = v38
					v42 = v20
					*(*int32)(unsafe.Add(mBase, uint32(v19))) = v41
					v46 = v42
				}
				v49 = v46
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				switch v24 - int32(435) {
				case 0:
					if v19 == int32(0) {
						v49 = int32(1)
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+168))
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
						v41 = v31
						v42 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v41
						v46 = v42
						v49 = v46
					}
				case 1:
					if v19 == int32(0) {
						v49 = int32(2)
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v21)+376))
						v41 = v36
						v42 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v41
						v46 = v42
						v49 = v46
					}
				default:
					v38 = int32(0)
					if v19 == v38 {
						v46 = v38
					} else {
						v41 = v38
						v42 = v20
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v41
						v46 = v42
					}
					v49 = v46
				}
			}
			if v49 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v154 = m.ExcPending
				if v154 != 0 {
					return int64(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_int2_accum_0), int32(0))
					mBase = m.M
					v158 = m.ExcPending
					if v158 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_int2_accum_1), int32(_a_F_int2_accum_2), int32(_a_F_int2_accum_3))
						mBase = m.M
						v163 = m.ExcPending
						if v163 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v52 = int32(_a_F_int2_accum_4)
				v53 = *(*int32)(unsafe.Add(mBase, _c_F_int2_accum[0]))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
				*(*int32)(unsafe.Add(mBase, _c_F_int2_accum[0])) = v55
				v58 = F_palloc0(m, int32(48))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int64(0)
				} else {
					v62 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v62)
					*(*int32)(unsafe.Add(mBase, _c_F_int2_accum[0])) = v53
					v66 = v58
					v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
					if v68 == int32(0) {
						v71 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
						if v72 == int32(0) {
							v126 = v71 >> (uint(int64(63)) % 64)
						} else {
							v78 = v71 >> (uint(int64(63)) % 64)
							v80 = v71 * v78
							v83 = int64(32)
							v84 = int64(base.Ui64(v71) >> (uint(v83) % 64))
							v89 = int64(4294967295)
							v90 = v71 & v89
							v93 = v90 * v90
							v96 = v90 * v84
							v97 = int64(base.Ui64(v93)>>(uint(v83)%64)) + v96
							v104 = v96 + v97&v89
							*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v80 + v80 + v84*v84 + int64(base.Ui64(v97)>>(uint(v83)%64)) + int64(base.Ui64(v104)>>(uint(v83)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v11))) = v93&v89 | v104<<(uint(v83)%64)
							v115 = *(*int64)(unsafe.Add(mBase, uint32(v66)+32))
							v116 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
							v117 = v115 + v116
							*(*int64)(unsafe.Add(mBase, uint32(v66)+32)) = v117
							v121 = *(*int64)(unsafe.Add(mBase, uint32(v66)+40))
							v122 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v66)+40)) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v117) < base.Ui64(v115))) + (v121 + v122)
							v126 = v78
						}
						v129 = *(*int64)(unsafe.Add(mBase, uint32(v66)+16))
						v130 = v71 + v129
						*(*int64)(unsafe.Add(mBase, uint32(v66)+16)) = v130
						v132 = *(*int64)(unsafe.Add(mBase, uint32(v66)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v66)+8)) = v132 + int64(1)
						v138 = *(*int64)(unsafe.Add(mBase, uint32(v66)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v66)+24)) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v130) < base.Ui64(v129))) + (v138 + v126)
					} else {
					}
					m.G0 = v11 + int32(32)
					return base.I64_extend_i32_u(v66)
				}
			}
		}
	} else {
		v19 = v11 + int32(28)
		v20 = int32(0)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v21 == v20 {
			v38 = int32(0)
			if v19 == v38 {
				v46 = v38
			} else {
				v41 = v38
				v42 = v20
				*(*int32)(unsafe.Add(mBase, uint32(v19))) = v41
				v46 = v42
			}
			v49 = v46
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			switch v24 - int32(435) {
			case 0:
				if v19 == int32(0) {
					v49 = int32(1)
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+168))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
					v41 = v31
					v42 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v19))) = v41
					v46 = v42
					v49 = v46
				}
			case 1:
				if v19 == int32(0) {
					v49 = int32(2)
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v21)+376))
					v41 = v36
					v42 = int32(2)
					*(*int32)(unsafe.Add(mBase, uint32(v19))) = v41
					v46 = v42
					v49 = v46
				}
			default:
				v38 = int32(0)
				if v19 == v38 {
					v46 = v38
				} else {
					v41 = v38
					v42 = v20
					*(*int32)(unsafe.Add(mBase, uint32(v19))) = v41
					v46 = v42
				}
				v49 = v46
			}
		}
		if v49 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v154 = m.ExcPending
			if v154 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_int2_accum_0), int32(0))
				mBase = m.M
				v158 = m.ExcPending
				if v158 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int2_accum_1), int32(_a_F_int2_accum_2), int32(_a_F_int2_accum_3))
					mBase = m.M
					v163 = m.ExcPending
					if v163 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v52 = int32(_a_F_int2_accum_4)
			v53 = *(*int32)(unsafe.Add(mBase, _c_F_int2_accum[0]))
			v55 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
			*(*int32)(unsafe.Add(mBase, _c_F_int2_accum[0])) = v55
			v58 = F_palloc0(m, int32(48))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int64(0)
			} else {
				v62 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v62)
				*(*int32)(unsafe.Add(mBase, _c_F_int2_accum[0])) = v53
				v66 = v58
				v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
				if v68 == int32(0) {
					v71 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
					v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
					if v72 == int32(0) {
						v126 = v71 >> (uint(int64(63)) % 64)
					} else {
						v78 = v71 >> (uint(int64(63)) % 64)
						v80 = v71 * v78
						v83 = int64(32)
						v84 = int64(base.Ui64(v71) >> (uint(v83) % 64))
						v89 = int64(4294967295)
						v90 = v71 & v89
						v93 = v90 * v90
						v96 = v90 * v84
						v97 = int64(base.Ui64(v93)>>(uint(v83)%64)) + v96
						v104 = v96 + v97&v89
						*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v80 + v80 + v84*v84 + int64(base.Ui64(v97)>>(uint(v83)%64)) + int64(base.Ui64(v104)>>(uint(v83)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v11))) = v93&v89 | v104<<(uint(v83)%64)
						v115 = *(*int64)(unsafe.Add(mBase, uint32(v66)+32))
						v116 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
						v117 = v115 + v116
						*(*int64)(unsafe.Add(mBase, uint32(v66)+32)) = v117
						v121 = *(*int64)(unsafe.Add(mBase, uint32(v66)+40))
						v122 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v66)+40)) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v117) < base.Ui64(v115))) + (v121 + v122)
						v126 = v78
					}
					v129 = *(*int64)(unsafe.Add(mBase, uint32(v66)+16))
					v130 = v71 + v129
					*(*int64)(unsafe.Add(mBase, uint32(v66)+16)) = v130
					v132 = *(*int64)(unsafe.Add(mBase, uint32(v66)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v66)+8)) = v132 + int64(1)
					v138 = *(*int64)(unsafe.Add(mBase, uint32(v66)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v66)+24)) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v130) < base.Ui64(v129))) + (v138 + v126)
				} else {
				}
				m.G0 = v11 + int32(32)
				return base.I64_extend_i32_u(v66)
			}
		}
	}
}
func F_int2_bytea(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_int2send(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_int2_decrement(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int64
	_ = v10
	var v16 int64
	_ = v16
	v7 = base.B2i32(l1&int64(65535) == int64(32768))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v7)
	v10 = int64(48)
	if l1&int64(65535) == int64(32768) {
		v16 = int64(0)
	} else {
		v16 = (l1<<(uint(v10)%64) + int64(-281474976710656)) >> (uint(v10) % 64)
	}
	return v16
}
