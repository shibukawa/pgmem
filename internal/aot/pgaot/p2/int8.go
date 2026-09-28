package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int8_avg_accum(m *base.Module, l0 int32) int64 {
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
				v71 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
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
					F_errmsg_internal(m, int32(_a_F_int8_avg_accum_0), int32(0))
					mBase = m.M
					v158 = m.ExcPending
					if v158 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_int8_avg_accum_1), int32(_a_F_int8_avg_accum_2), int32(_a_F_int8_avg_accum_3))
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
				v52 = int32(_a_F_int8_avg_accum_4)
				v53 = *(*int32)(unsafe.Add(mBase, _c_F_int8_avg_accum[0]))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
				*(*int32)(unsafe.Add(mBase, _c_F_int8_avg_accum[0])) = v55
				v58 = F_palloc0(m, int32(48))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int64(0)
				} else {
					v62 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v62)
					*(*int32)(unsafe.Add(mBase, _c_F_int8_avg_accum[0])) = v53
					v66 = v58
					v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
					if v68 == int32(0) {
						v71 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
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
				F_errmsg_internal(m, int32(_a_F_int8_avg_accum_0), int32(0))
				mBase = m.M
				v158 = m.ExcPending
				if v158 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int8_avg_accum_1), int32(_a_F_int8_avg_accum_2), int32(_a_F_int8_avg_accum_3))
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
			v52 = int32(_a_F_int8_avg_accum_4)
			v53 = *(*int32)(unsafe.Add(mBase, _c_F_int8_avg_accum[0]))
			v55 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
			*(*int32)(unsafe.Add(mBase, _c_F_int8_avg_accum[0])) = v55
			v58 = F_palloc0(m, int32(48))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int64(0)
			} else {
				v62 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v62)
				*(*int32)(unsafe.Add(mBase, _c_F_int8_avg_accum[0])) = v53
				v66 = v58
				v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
				if v68 == int32(0) {
					v71 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
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
func F_int8_avg_combine(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v117 int64
	_ = v117
	var v120 int64
	_ = v120
	var v123 int64
	_ = v123
	var v124 int64
	_ = v124
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v130 int64
	_ = v130
	var v134 int32
	_ = v134
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = v11 + int32(8)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v16 == v2 {
		v33 = int32(0)
		if v14 == v33 {
			v41 = v33
		} else {
			v36 = v33
			v37 = v2
			*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
			v41 = v37
		}
		v44 = v41
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		switch v19 - int32(435) {
		case 0:
			if v14 == int32(0) {
				v44 = int32(1)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+168))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
				v36 = v26
				v37 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
				v41 = v37
				v44 = v41
			}
		case 1:
			if v14 == int32(0) {
				v44 = int32(2)
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v16)+376))
				v36 = v31
				v37 = int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
				v41 = v37
				v44 = v41
			}
		default:
			v33 = int32(0)
			if v14 == v33 {
				v41 = v33
			} else {
				v36 = v33
				v37 = v2
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
				v41 = v37
			}
			v44 = v41
		}
	}
	if v44 != 0 {
		v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v45 == int32(0) {
			v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v49 = v48
		} else {
			v49 = v2
		}
		v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
		if v50 == int32(0) {
			v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			if v53 != 0 {
				if v49 == int32(0) {
					v60 = int32(_a_F_int8_avg_combine_0)
					v61 = *(*int32)(unsafe.Add(mBase, _c_F_int8_avg_combine[0]))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
					*(*int32)(unsafe.Add(mBase, _c_F_int8_avg_combine[0])) = v63
					v66 = v11 + int32(12)
					v67 = int32(0)
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v68 == v67 {
						v85 = int32(0)
						if v66 == v85 {
							v93 = v85
						} else {
							v88 = v85
							v89 = v67
							*(*int32)(unsafe.Add(mBase, uint32(v66))) = v88
							v93 = v89
						}
						v96 = v93
					} else {
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
						switch v71 - int32(435) {
						case 0:
							if v66 == int32(0) {
								v96 = int32(1)
							} else {
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+168))
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
								v88 = v78
								v89 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v66))) = v88
								v93 = v89
								v96 = v93
							}
						case 1:
							if v66 == int32(0) {
								v96 = int32(2)
							} else {
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v68)+376))
								v88 = v83
								v89 = int32(2)
								*(*int32)(unsafe.Add(mBase, uint32(v66))) = v88
								v93 = v89
								v96 = v93
							}
						default:
							v85 = int32(0)
							if v66 == v85 {
								v93 = v85
							} else {
								v88 = v85
								v89 = v67
								*(*int32)(unsafe.Add(mBase, uint32(v66))) = v88
								v93 = v89
							}
							v96 = v93
						}
					}
					if v96 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v161 = m.ExcPending
						if v161 != 0 {
							return int64(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_int8_avg_combine_1), int32(0))
							mBase = m.M
							v165 = m.ExcPending
							if v165 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_int8_avg_combine_2), int32(_a_F_int8_avg_combine_3), int32(_a_F_int8_avg_combine_4))
								mBase = m.M
								v170 = m.ExcPending
								if v170 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v100 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
						*(*int32)(unsafe.Add(mBase, _c_F_int8_avg_combine[0])) = v100
						v103 = F_palloc0(m, int32(48))
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int64(0)
						} else {
							v107 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v103))) = uint8(v107)
							v109 = *(*int64)(unsafe.Add(mBase, uint32(v53)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v103)+8)) = v109
							v111 = *(*int64)(unsafe.Add(mBase, uint32(v53)+16))
							v112 = *(*int64)(unsafe.Add(mBase, uint32(v53)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v103)+24)) = v112
							*(*int64)(unsafe.Add(mBase, uint32(v103)+16)) = v111
							*(*int32)(unsafe.Add(mBase, _c_F_int8_avg_combine[0])) = v61
							v134 = v103
							m.G0 = v11 + int32(16)
							return base.I64_extend_i32_u(v134)
						}
					}
				} else {
					v117 = *(*int64)(unsafe.Add(mBase, uint32(v53)+8))
					if v117 <= int64(0) {
						v134 = v49
					} else {
						v120 = *(*int64)(unsafe.Add(mBase, uint32(v49)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v120 + v117
						v123 = *(*int64)(unsafe.Add(mBase, uint32(v53)+24))
						v124 = *(*int64)(unsafe.Add(mBase, uint32(v49)+16))
						v125 = *(*int64)(unsafe.Add(mBase, uint32(v53)+16))
						v126 = v124 + v125
						*(*int64)(unsafe.Add(mBase, uint32(v49)+16)) = v126
						v130 = *(*int64)(unsafe.Add(mBase, uint32(v49)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v49)+24)) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v126) < base.Ui64(v124))) + (v123 + v130)
						v134 = v49
					}
					m.G0 = v11 + int32(16)
					return base.I64_extend_i32_u(v134)
				}
			} else {
				if v49 != 0 {
					v134 = v49
				} else {
					v55 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v55)
					v134 = int32(0)
				}
				m.G0 = v11 + int32(16)
				return base.I64_extend_i32_u(v134)
			}
		} else {
			if v49 != 0 {
				v134 = v49
			} else {
				v55 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v55)
				v134 = int32(0)
			}
			m.G0 = v11 + int32(16)
			return base.I64_extend_i32_u(v134)
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v148 = m.ExcPending
		if v148 != 0 {
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_int8_avg_combine_1), int32(0))
			mBase = m.M
			v152 = m.ExcPending
			if v152 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_int8_avg_combine_2), int32(_a_F_int8_avg_combine_5), int32(_a_F_int8_avg_combine_6))
				mBase = m.M
				v157 = m.ExcPending
				if v157 != 0 {
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
func F_int8_dist(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14242(m, l0, int32(_a_F_int8_dist_0), int32(_a_F_int8_dist_1), int32(_a_F_int8_dist_2))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_int8_to_char(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v184 float64
	_ = v184
	var v186 int64
	_ = v186
	var v187 int32
	_ = v187
	var v188 int64
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int64
	_ = v194
	var v197 int64
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v351 int32
	_ = v351
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v17 = F_pg_detoast_datum_packed(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v13 + int32(48)
	return base.I64_extend_i32_u(v351)
L2:
	;
	if base.Ui32(int32(268435454)) <= base.Ui32(v50-int32(1)) {
		goto L15
	} else {
		goto L16
	}
L3:
	;
	return int64(0)
L4:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v21 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v27 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v38 = int32(1)
	if v21&v38 != 0 {
		v50 = int32(base.Ui32(v21)>>(uint(v38)%32)) - v38
		goto L2
	} else {
		goto L14
	}
L8:
	;
	v30 = int32(16)
	goto L10
L9:
	;
	v30 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v27-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v37 = int32(4)
	goto L13
L12:
	;
	v37 = v30
	goto L13
L13:
	;
	v50 = v37
	goto L2
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
	goto L2
L15:
	;
	v56 = F_cstring_to_text(m, int32(_a_F_int8_to_char_0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L3
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v62 = F_palloc0(m, v50<<(uint(int32(3))%32)|int32(5))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L3
	} else {
		goto L19
	}
L18:
	;
	v351 = v56
	goto L1
L19:
	;
	v68 = F_NUM_cache(m, v50, v13+int32(12), v17, v13+int32(11))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	if v70&int32(1024) != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v330 = v62 + int32(4)
	F_NUM_processor(m, v68, v13+int32(12), v330, v320, int32(0), v323, v325, int32(1))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L3
	} else {
		goto L104
	}
L22:
	;
	if base.Ui64(int64(4294967296)) <= base.Ui64(v15+int64(2147483648)) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	if v70&int32(_a_F_int8_to_char_1) != 0 {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v79 = int32(2147483647)
	goto L27
L26:
	;
	v79 = base.I32_wrap_i64(v15)
	goto L27
L27:
	;
	v80 = F_int_to_roman(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	v320 = v80
	v323 = v2
	v325 = v2
	goto L21
L29:
	;
	v84 = F_int64_to_numeric(m, v15)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L3
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if v70&int32(2048) != 0 {
		goto L59
	} else {
		goto L60
	}
L32:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v87 = F_numeric_out_sci(m, v84, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v89 == int32(45) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v320 = v87
	v323 = v2
	v325 = v2
	goto L21
L35:
	;
	goto L36
L36:
	;
	v92 = F_strlen(m, v87)
	mBase = m.M
	v95 = F_palloc(m, v92+int32(2))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v97 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v95))) = uint8(v97)
	v100 = v95 + int32(1)
	if (v87^v100)&int32(3) != 0 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v320 = v95
	v323 = v2
	v325 = v2
	goto L21
L39:
	;
	goto L38
L40:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v155))) = uint8(v154)
	if v154&int32(255) == int32(0) {
		goto L39
	} else {
		goto L55
	}
L41:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v153 = v87
	v154 = v106
	v155 = v100
	goto L40
L42:
	;
	goto L43
L43:
	;
	if v87&int32(3) != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v110 = v87
	v112 = v100
	goto L47
L45:
	;
	v124 = v87
	v126 = v100
	goto L46
L46:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v131 = int32(-2139062144)
	if (int32(16843008)-v128|v128)&v131 != v131 {
		v153 = v124
		v154 = v128
		v155 = v126
		goto L40
	} else {
		goto L51
	}
L47:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	*(*uint8)(unsafe.Add(mBase, uint32(v112))) = uint8(v113)
	if v113 == int32(0) {
		goto L39
	} else {
		goto L49
	}
L48:
	;
	v124 = v120
	v126 = v118
	goto L46
L49:
	;
	v117 = int32(1)
	v118 = v112 + v117
	v120 = v110 + v117
	if v120&int32(3) != 0 {
		v110 = v120
		v112 = v118
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v136 = v124
	v137 = v128
	v138 = v126
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v137
	v140 = int32(4)
	v141 = v138 + v140
	v143 = v136 + v140
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	v148 = int32(-2139062144)
	if (int32(16843008)-v145|v145)&v148 == v148 {
		v136 = v143
		v137 = v145
		v138 = v141
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v153 = v143
	v154 = v145
	v155 = v141
	goto L40
L54:
	;
	goto L53
L55:
	;
	v162 = v153
	v164 = v155
	goto L56
L56:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v164)+1)) = uint8(v165)
	v167 = int32(1)
	if v165 != 0 {
		v162 = v162 + v167
		v164 = v164 + v167
		goto L56
	} else {
		goto L58
	}
L57:
	;
	goto L39
L58:
	;
	goto L57
L59:
	;
	v178 = int32(0)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
	v184 = F_pow(m, float64(10), base.F64_convert_i32_s(v182))
	mBase = m.M
	v186 = F_DirectFunctionCall1Coll(m, int32(1440), v178, base.I64_reinterpret_f64(v184))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L3
	} else {
		goto L62
	}
L60:
	;
	v194 = v15
	goto L61
L61:
	;
	v197 = F_DirectFunctionCall1Coll(m, int32(1441), int32(0), v194)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L3
	} else {
		goto L64
	}
L62:
	;
	v188 = F_DirectFunctionCall2Coll(m, int32(1439), v178, v15, v186)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L3
	} else {
		goto L63
	}
L63:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v182 + v190
	v194 = v188
	goto L61
L64:
	;
	v199 = base.I32_wrap_i64(v197)
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
	v202 = base.B2i32(v200 == int32(45))
	v203 = v199 + v202
	v204 = F_strlen(m, v203)
	mBase = m.M
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v205 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v209 = F_palloc(m, v204+v205+int32(2))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L3
	} else {
		goto L68
	}
L66:
	;
	v295 = v203
	goto L67
L67:
	;
	if v200 == int32(45) {
		goto L93
	} else {
		goto L94
	}
L68:
	;
	if (v203^v209)&int32(3) != 0 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v285 = v209 + v204
	v286 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v285))) = uint8(v286)
	if v205 != 0 {
		goto L90
	} else {
		goto L91
	}
L70:
	;
	goto L69
L71:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v265))) = uint8(v264)
	if v264&int32(255) == int32(0) {
		goto L70
	} else {
		goto L86
	}
L72:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	v263 = v203
	v264 = v216
	v265 = v209
	goto L71
L73:
	;
	goto L74
L74:
	;
	if v203&int32(3) != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v220 = v203
	v222 = v209
	goto L78
L76:
	;
	v234 = v203
	v236 = v209
	goto L77
L77:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	v241 = int32(-2139062144)
	if (int32(16843008)-v238|v238)&v241 != v241 {
		v263 = v234
		v264 = v238
		v265 = v236
		goto L71
	} else {
		goto L82
	}
L78:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220))))
	*(*uint8)(unsafe.Add(mBase, uint32(v222))) = uint8(v223)
	if v223 == int32(0) {
		goto L70
	} else {
		goto L80
	}
L79:
	;
	v234 = v230
	v236 = v228
	goto L77
L80:
	;
	v227 = int32(1)
	v228 = v222 + v227
	v230 = v220 + v227
	if v230&int32(3) != 0 {
		v220 = v230
		v222 = v228
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v246 = v234
	v247 = v238
	v248 = v236
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v248))) = v247
	v250 = int32(4)
	v251 = v248 + v250
	v253 = v246 + v250
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v246)+4))
	v258 = int32(-2139062144)
	if (int32(16843008)-v255|v255)&v258 == v258 {
		v246 = v253
		v247 = v255
		v248 = v251
		goto L83
	} else {
		goto L85
	}
L84:
	;
	v263 = v253
	v264 = v255
	v265 = v251
	goto L71
L85:
	;
	goto L84
L86:
	;
	v272 = v263
	v274 = v265
	goto L87
L87:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v274)+1)) = uint8(v275)
	v277 = int32(1)
	if v275 != 0 {
		v272 = v272 + v277
		v274 = v274 + v277
		goto L87
	} else {
		goto L89
	}
L88:
	;
	goto L70
L89:
	;
	goto L88
L90:
	;
	base.MemoryFill(m, v285+int32(1), int32(48), v205)
	goto L92
L91:
	;
	goto L92
L92:
	;
	v293 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v285+v205)+1)) = uint8(v293)
	v295 = v209
	goto L67
L93:
	;
	v299 = int32(45)
	goto L95
L94:
	;
	v299 = int32(43)
	goto L95
L95:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if base.Ui32(v204) < base.Ui32(v300) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v320 = v295
	v323 = v300 - v204
	v325 = v299
	goto L21
L97:
	;
	goto L98
L98:
	;
	v303 = int32(0)
	if base.Ui32(v204) <= base.Ui32(v300) {
		v320 = v295
		v323 = v303
		v325 = v299
		goto L21
	} else {
		goto L99
	}
L99:
	;
	v305 = v300 + v205
	v308 = F_palloc(m, v305+int32(2))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L3
	} else {
		goto L100
	}
L100:
	;
	v311 = v305 + int32(1)
	if v311 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	base.MemoryFill(m, v308, int32(35), v311)
	goto L103
L102:
	;
	goto L103
L103:
	;
	v315 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v308+v311))) = uint8(v315)
	v318 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v308+v300))) = uint8(v318)
	v320 = v308
	v323 = v303
	v325 = v299
	goto L21
L104:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)))
	if v335 == int32(1) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	F_pfree(m, v68)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L3
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v340 = F_strlen(m, v330)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v340<<(uint(int32(2))%32) + int32(16)
	v351 = v62
	goto L1
L108:
	;
	goto L107
}
