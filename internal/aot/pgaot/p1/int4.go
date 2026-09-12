package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int4_accum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
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
	var v120 int32
	_ = v120
	var v123 int64
	_ = v123
	var v124 int64
	_ = v124
	var v131 int64
	_ = v131
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v140 int32
	_ = v140
	var v143 int64
	_ = v143
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v12 == int32(0) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v15 != 0 {
			v65 = v15
			v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v67 == int32(0) {
				v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
				if v71 == int32(1) {
					v75 = v10 + int32(8)
					v76 = base.I64_extend_i32_s(v70)
					v80 = v76 * (v76 >> (uint(int64(63)) % 64))
					v83 = int64(32)
					v84 = int64(base.Ui64(v76) >> (uint(v83) % 64))
					v89 = int64(4294967295)
					v90 = v76 & v89
					v93 = v90 * v90
					v96 = v90 * v84
					v97 = int64(base.Ui64(v93)>>(uint(v83)%64)) + v96
					v104 = v96 + v97&v89
					*(*int64)(unsafe.Add(mBase, uint32(v75)+8)) = v80 + v80 + v84*v84 + int64(base.Ui64(v97)>>(uint(v83)%64)) + int64(base.Ui64(v104)>>(uint(v83)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v75))) = v93&v89 | v104<<(uint(v83)%64)
					v115 = *(*int64)(unsafe.Add(mBase, uint32(v65)+32))
					v116 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
					v117 = v115 + v116
					*(*int64)(unsafe.Add(mBase, uint32(v65)+32)) = v117
					v120 = v65 + int32(40)
					v123 = *(*int64)(unsafe.Add(mBase, uint32(v120)))
					v124 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v120))) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v117) < base.Ui64(v115))) + (v123 + v124)
				} else {
				}
				v131 = *(*int64)(unsafe.Add(mBase, uint32(v65)+16))
				v132 = base.I64_extend_i32_s(v70)
				v133 = v131 + v132
				*(*int64)(unsafe.Add(mBase, uint32(v65)+16)) = v133
				v135 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v135 + int64(1)
				v140 = v65 + int32(24)
				v143 = *(*int64)(unsafe.Add(mBase, uint32(v140)))
				*(*int64)(unsafe.Add(mBase, uint32(v140))) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v133) < base.Ui64(v131))) + (v143 + v132>>(uint(int64(63))%64))
			} else {
			}
			m.G0 = v10 + int32(32)
			return v65
		} else {
			v18 = v10 + int32(28)
			v19 = int32(0)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v20 == v19 {
				v37 = int32(0)
				if v18 == v37 {
					v45 = v37
				} else {
					v40 = v37
					v41 = v19
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
					v45 = v41
				}
				v48 = v45
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				switch v23 - int32(429) {
				case 0:
					if v18 == int32(0) {
						v48 = int32(1)
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+168))
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
						v40 = v30
						v41 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
						v45 = v41
						v48 = v45
					}
				case 1:
					if v18 == int32(0) {
						v48 = int32(2)
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v20)+368))
						v40 = v35
						v41 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
						v45 = v41
						v48 = v45
					}
				default:
					v37 = int32(0)
					if v18 == v37 {
						v45 = v37
					} else {
						v40 = v37
						v41 = v19
						*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
						v45 = v41
					}
					v48 = v45
				}
			}
			if v48 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v161 = m.ExcPending
				if v161 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(60341), int32(0))
					mBase = m.M
					v165 = m.ExcPending
					if v165 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(491015), int32(5606), int32(347744))
						mBase = m.M
						v170 = m.ExcPending
						if v170 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v51 = int32(4470400)
				v52 = *(*int32)(unsafe.Add(mBase, _consts[3]))
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
				*(*int32)(unsafe.Add(mBase, _consts[3])) = v54
				v57 = F_palloc0(m, int32(48))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return int32(0)
				} else {
					v61 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v61)
					*(*int32)(unsafe.Add(mBase, _consts[3])) = v52
					v65 = v57
					v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					if v67 == int32(0) {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
						if v71 == int32(1) {
							v75 = v10 + int32(8)
							v76 = base.I64_extend_i32_s(v70)
							v80 = v76 * (v76 >> (uint(int64(63)) % 64))
							v83 = int64(32)
							v84 = int64(base.Ui64(v76) >> (uint(v83) % 64))
							v89 = int64(4294967295)
							v90 = v76 & v89
							v93 = v90 * v90
							v96 = v90 * v84
							v97 = int64(base.Ui64(v93)>>(uint(v83)%64)) + v96
							v104 = v96 + v97&v89
							*(*int64)(unsafe.Add(mBase, uint32(v75)+8)) = v80 + v80 + v84*v84 + int64(base.Ui64(v97)>>(uint(v83)%64)) + int64(base.Ui64(v104)>>(uint(v83)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v75))) = v93&v89 | v104<<(uint(v83)%64)
							v115 = *(*int64)(unsafe.Add(mBase, uint32(v65)+32))
							v116 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
							v117 = v115 + v116
							*(*int64)(unsafe.Add(mBase, uint32(v65)+32)) = v117
							v120 = v65 + int32(40)
							v123 = *(*int64)(unsafe.Add(mBase, uint32(v120)))
							v124 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v120))) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v117) < base.Ui64(v115))) + (v123 + v124)
						} else {
						}
						v131 = *(*int64)(unsafe.Add(mBase, uint32(v65)+16))
						v132 = base.I64_extend_i32_s(v70)
						v133 = v131 + v132
						*(*int64)(unsafe.Add(mBase, uint32(v65)+16)) = v133
						v135 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v135 + int64(1)
						v140 = v65 + int32(24)
						v143 = *(*int64)(unsafe.Add(mBase, uint32(v140)))
						*(*int64)(unsafe.Add(mBase, uint32(v140))) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v133) < base.Ui64(v131))) + (v143 + v132>>(uint(int64(63))%64))
					} else {
					}
					m.G0 = v10 + int32(32)
					return v65
				}
			}
		}
	} else {
		v18 = v10 + int32(28)
		v19 = int32(0)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v20 == v19 {
			v37 = int32(0)
			if v18 == v37 {
				v45 = v37
			} else {
				v40 = v37
				v41 = v19
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
				v45 = v41
			}
			v48 = v45
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
			switch v23 - int32(429) {
			case 0:
				if v18 == int32(0) {
					v48 = int32(1)
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+168))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
					v40 = v30
					v41 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
					v45 = v41
					v48 = v45
				}
			case 1:
				if v18 == int32(0) {
					v48 = int32(2)
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v20)+368))
					v40 = v35
					v41 = int32(2)
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
					v45 = v41
					v48 = v45
				}
			default:
				v37 = int32(0)
				if v18 == v37 {
					v45 = v37
				} else {
					v40 = v37
					v41 = v19
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
					v45 = v41
				}
				v48 = v45
			}
		}
		if v48 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v161 = m.ExcPending
			if v161 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(60341), int32(0))
				mBase = m.M
				v165 = m.ExcPending
				if v165 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(491015), int32(5606), int32(347744))
					mBase = m.M
					v170 = m.ExcPending
					if v170 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v51 = int32(4470400)
			v52 = *(*int32)(unsafe.Add(mBase, _consts[3]))
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
			*(*int32)(unsafe.Add(mBase, _consts[3])) = v54
			v57 = F_palloc0(m, int32(48))
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return int32(0)
			} else {
				v61 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v61)
				*(*int32)(unsafe.Add(mBase, _consts[3])) = v52
				v65 = v57
				v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
				if v67 == int32(0) {
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
					if v71 == int32(1) {
						v75 = v10 + int32(8)
						v76 = base.I64_extend_i32_s(v70)
						v80 = v76 * (v76 >> (uint(int64(63)) % 64))
						v83 = int64(32)
						v84 = int64(base.Ui64(v76) >> (uint(v83) % 64))
						v89 = int64(4294967295)
						v90 = v76 & v89
						v93 = v90 * v90
						v96 = v90 * v84
						v97 = int64(base.Ui64(v93)>>(uint(v83)%64)) + v96
						v104 = v96 + v97&v89
						*(*int64)(unsafe.Add(mBase, uint32(v75)+8)) = v80 + v80 + v84*v84 + int64(base.Ui64(v97)>>(uint(v83)%64)) + int64(base.Ui64(v104)>>(uint(v83)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v75))) = v93&v89 | v104<<(uint(v83)%64)
						v115 = *(*int64)(unsafe.Add(mBase, uint32(v65)+32))
						v116 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
						v117 = v115 + v116
						*(*int64)(unsafe.Add(mBase, uint32(v65)+32)) = v117
						v120 = v65 + int32(40)
						v123 = *(*int64)(unsafe.Add(mBase, uint32(v120)))
						v124 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v120))) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v117) < base.Ui64(v115))) + (v123 + v124)
					} else {
					}
					v131 = *(*int64)(unsafe.Add(mBase, uint32(v65)+16))
					v132 = base.I64_extend_i32_s(v70)
					v133 = v131 + v132
					*(*int64)(unsafe.Add(mBase, uint32(v65)+16)) = v133
					v135 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v135 + int64(1)
					v140 = v65 + int32(24)
					v143 = *(*int64)(unsafe.Add(mBase, uint32(v140)))
					*(*int64)(unsafe.Add(mBase, uint32(v140))) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v133) < base.Ui64(v131))) + (v143 + v132>>(uint(int64(63))%64))
				} else {
				}
				m.G0 = v10 + int32(32)
				return v65
			}
		}
	}
}
func F_int4_decrement(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	v5 = base.B2i32(l1 == int32(-2147483648))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v5)
	if l1 == int32(-2147483648) {
		v10 = int32(0)
	} else {
		v10 = l1 - int32(1)
	}
	return v10
}
