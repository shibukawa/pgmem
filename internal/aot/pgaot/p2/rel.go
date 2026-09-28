package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RelFileLocatorSkippingWAL(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_RelFileLocatorSkippingWAL[0]))
	if v4 != 0 {
		v5 = int32(0)
		v7 = F_hash_search(m, v4, l0, v5, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v14 = base.B2i32(v7 != int32(0))
			return v14
		}
	} else {
		v14 = int32(0)
		return v14
	}
}
func F_estimate_rel_size(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int64
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 float32
	_ = v38
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 float64
	_ = v57
	var v58 float64
	_ = v58
	var v63 int32
	_ = v63
	var v68 float64
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 float32
	_ = v77
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+119)))
	switch v13 - int32(105) {
	case 0:
		v21 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v21
			if v21 == int32(0) {
				v26 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(l3))) = v26
				*(*int64)(unsafe.Add(mBase, uint32(l4))) = v26
				return
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+96))
				v34 = v21 - base.B2i32(v31 != int32(0))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+104))
				if base.Ui32(v31) < base.Ui32(int32(2)) {
					v50 = F_get_rel_data_width(m, l0, l1)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						v54 = base.I32_div_u_s(int32(_a_F_estimate_rel_size_0), v50+int32(28))
						v57 = base.F64_convert_i32_u(v54)
						v58 = base.F64_convert_i32_u(v34)
						*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_nearest(base.F64_mul(v57, v58))
						if v34 != 0 {
							v63 = v35
						} else {
							v63 = int32(0)
						}
						if v63 == int32(0) {
							*(*int64)(unsafe.Add(mBase, uint32(l4))) = int64(0)
							return
						} else {
							v68 = base.F64_convert_i32_u(v35)
							if base.F64_le(v58, v68) != 0 {
								*(*int64)(unsafe.Add(mBase, uint32(l4))) = int64(4607182418800017408)
								return
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_div(v68, v58)
								return
							}
						}
					}
				} else {
					v38 = *(*float32)(unsafe.Add(mBase, uint32(v30)+100))
					if base.F32_ge(v38, float32(0)) == int32(0) {
						v50 = F_get_rel_data_width(m, l0, l1)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							v54 = base.I32_div_u_s(int32(_a_F_estimate_rel_size_0), v50+int32(28))
							v57 = base.F64_convert_i32_u(v54)
							v58 = base.F64_convert_i32_u(v34)
							*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_nearest(base.F64_mul(v57, v58))
							if v34 != 0 {
								v63 = v35
							} else {
								v63 = int32(0)
							}
							if v63 == int32(0) {
								*(*int64)(unsafe.Add(mBase, uint32(l4))) = int64(0)
								return
							} else {
								v68 = base.F64_convert_i32_u(v35)
								if base.F64_le(v58, v68) != 0 {
									*(*int64)(unsafe.Add(mBase, uint32(l4))) = int64(4607182418800017408)
									return
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_div(v68, v58)
									return
								}
							}
						}
					} else {
						v57 = base.F64_div(base.F64_promote_f32(v38), base.F64_convert_i32_u(v31-int32(1)))
						v58 = base.F64_convert_i32_u(v34)
						*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_nearest(base.F64_mul(v57, v58))
						if v34 != 0 {
							v63 = v35
						} else {
							v63 = int32(0)
						}
						if v63 == int32(0) {
							*(*int64)(unsafe.Add(mBase, uint32(l4))) = int64(0)
							return
						} else {
							v68 = base.F64_convert_i32_u(v35)
							if base.F64_le(v58, v68) != 0 {
								*(*int64)(unsafe.Add(mBase, uint32(l4))) = int64(4607182418800017408)
								return
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_div(v68, v58)
								return
							}
						}
					}
				}
			}
		}
	default:
		v74 = *(*int32)(unsafe.Add(mBase, uint32(v12)+96))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v74
		v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v77 = *(*float32)(unsafe.Add(mBase, uint32(v76)+100))
		*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_promote_f32(v77)
		*(*int64)(unsafe.Add(mBase, uint32(l4))) = int64(0)
		return
	case 4, 9, 11:
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+164))
		m.T0[v17].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, l1, l2, l3, l4)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			return
		}
	}
}
func F_extractRelOptions(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+20)))
	if v13&int32(1) == v4 {
		v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+284)))
		if int32(0) <= v18 {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v23 = v12 + v21 + v18
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+288)))
			if v24 == int32(1) {
				v27 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+286)))
				if base.I32_popcnt(v27) != int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v27
						F_errmsg_internal(m, int32(_a_F_extractRelOptions_0), v10)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_extractRelOptions_1), int32(123), int32(_a_F_extractRelOptions_2))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					switch base.I32_ctz(v27) {
					case 0:
						v32 = int64(*(*int8)(unsafe.Add(mBase, uint32(v23))))
						v65 = v32
						v66 = int32(0)
						v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+22)))
						v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+v68)+119)))
						switch v70 - int32(73) {
						case 0, 32:
							if base.I32_wrap_i64(v65) == int32(0) {
								v110 = v66
								m.G0 = v10 + int32(16)
								return v110
							} else {
								v106 = m.T0[l2].(func(*base.Module, int64, int32) int32)(m, v65, int32(0))
								mBase = m.M
								v107 = m.ExcPending
								if v107 != 0 {
									return int32(0)
								} else {
									v110 = v106
									m.G0 = v10 + int32(16)
									return v110
								}
							}
						default:
							v110 = v66
							m.G0 = v10 + int32(16)
							return v110
						case 36, 41:
							v93 = F_build_reloptions(m, v65, int32(0), int32(1), int32(136), int32(_a_F_extractRelOptions_3), int32(26))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								v110 = v93
								m.G0 = v10 + int32(16)
								return v110
							}
						case 43:
							v78 = F_build_reloptions(m, v65, int32(0), int32(2), int32(136), int32(_a_F_extractRelOptions_3), int32(26))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int32(0)
							} else {
								if v78 == int32(0) {
									v110 = v66
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v78)+104)) = int64(-4616189618054758400)
									*(*int32)(unsafe.Add(mBase, uint32(v78)+36)) = int32(-1)
									*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = int32(100)
									v110 = v78
								}
								m.G0 = v10 + int32(16)
								return v110
							}
						case 45:
							v100 = F_build_reloptions(m, v65, int32(0), int32(512), int32(12), int32(_a_F_extractRelOptions_4), int32(3))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								v110 = v100
								m.G0 = v10 + int32(16)
								return v110
							}
						}
					case 1:
						v33 = int64(*(*int16)(unsafe.Add(mBase, uint32(v23))))
						v65 = v33
						v66 = int32(0)
						v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+22)))
						v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+v68)+119)))
						switch v70 - int32(73) {
						case 0, 32:
							if base.I32_wrap_i64(v65) == int32(0) {
								v110 = v66
								m.G0 = v10 + int32(16)
								return v110
							} else {
								v106 = m.T0[l2].(func(*base.Module, int64, int32) int32)(m, v65, int32(0))
								mBase = m.M
								v107 = m.ExcPending
								if v107 != 0 {
									return int32(0)
								} else {
									v110 = v106
									m.G0 = v10 + int32(16)
									return v110
								}
							}
						default:
							v110 = v66
							m.G0 = v10 + int32(16)
							return v110
						case 36, 41:
							v93 = F_build_reloptions(m, v65, int32(0), int32(1), int32(136), int32(_a_F_extractRelOptions_3), int32(26))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								v110 = v93
								m.G0 = v10 + int32(16)
								return v110
							}
						case 43:
							v78 = F_build_reloptions(m, v65, int32(0), int32(2), int32(136), int32(_a_F_extractRelOptions_3), int32(26))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int32(0)
							} else {
								if v78 == int32(0) {
									v110 = v66
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v78)+104)) = int64(-4616189618054758400)
									*(*int32)(unsafe.Add(mBase, uint32(v78)+36)) = int32(-1)
									*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = int32(100)
									v110 = v78
								}
								m.G0 = v10 + int32(16)
								return v110
							}
						case 45:
							v100 = F_build_reloptions(m, v65, int32(0), int32(512), int32(12), int32(_a_F_extractRelOptions_4), int32(3))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								v110 = v100
								m.G0 = v10 + int32(16)
								return v110
							}
						}
					case 2:
						v34 = int64(*(*int32)(unsafe.Add(mBase, uint32(v23))))
						v65 = v34
						v66 = int32(0)
						v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+22)))
						v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+v68)+119)))
						switch v70 - int32(73) {
						case 0, 32:
							if base.I32_wrap_i64(v65) == int32(0) {
								v110 = v66
								m.G0 = v10 + int32(16)
								return v110
							} else {
								v106 = m.T0[l2].(func(*base.Module, int64, int32) int32)(m, v65, int32(0))
								mBase = m.M
								v107 = m.ExcPending
								if v107 != 0 {
									return int32(0)
								} else {
									v110 = v106
									m.G0 = v10 + int32(16)
									return v110
								}
							}
						default:
							v110 = v66
							m.G0 = v10 + int32(16)
							return v110
						case 36, 41:
							v93 = F_build_reloptions(m, v65, int32(0), int32(1), int32(136), int32(_a_F_extractRelOptions_3), int32(26))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								v110 = v93
								m.G0 = v10 + int32(16)
								return v110
							}
						case 43:
							v78 = F_build_reloptions(m, v65, int32(0), int32(2), int32(136), int32(_a_F_extractRelOptions_3), int32(26))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int32(0)
							} else {
								if v78 == int32(0) {
									v110 = v66
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v78)+104)) = int64(-4616189618054758400)
									*(*int32)(unsafe.Add(mBase, uint32(v78)+36)) = int32(-1)
									*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = int32(100)
									v110 = v78
								}
								m.G0 = v10 + int32(16)
								return v110
							}
						case 45:
							v100 = F_build_reloptions(m, v65, int32(0), int32(512), int32(12), int32(_a_F_extractRelOptions_4), int32(3))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								v110 = v100
								m.G0 = v10 + int32(16)
								return v110
							}
						}
					case 3:
						v35 = *(*int64)(unsafe.Add(mBase, uint32(v23)))
						v65 = v35
						v66 = int32(0)
						v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+22)))
						v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+v68)+119)))
						switch v70 - int32(73) {
						case 0, 32:
							if base.I32_wrap_i64(v65) == int32(0) {
								v110 = v66
								m.G0 = v10 + int32(16)
								return v110
							} else {
								v106 = m.T0[l2].(func(*base.Module, int64, int32) int32)(m, v65, int32(0))
								mBase = m.M
								v107 = m.ExcPending
								if v107 != 0 {
									return int32(0)
								} else {
									v110 = v106
									m.G0 = v10 + int32(16)
									return v110
								}
							}
						default:
							v110 = v66
							m.G0 = v10 + int32(16)
							return v110
						case 36, 41:
							v93 = F_build_reloptions(m, v65, int32(0), int32(1), int32(136), int32(_a_F_extractRelOptions_3), int32(26))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								v110 = v93
								m.G0 = v10 + int32(16)
								return v110
							}
						case 43:
							v78 = F_build_reloptions(m, v65, int32(0), int32(2), int32(136), int32(_a_F_extractRelOptions_3), int32(26))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int32(0)
							} else {
								if v78 == int32(0) {
									v110 = v66
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v78)+104)) = int64(-4616189618054758400)
									*(*int32)(unsafe.Add(mBase, uint32(v78)+36)) = int32(-1)
									*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = int32(100)
									v110 = v78
								}
								m.G0 = v10 + int32(16)
								return v110
							}
						case 45:
							v100 = F_build_reloptions(m, v65, int32(0), int32(512), int32(12), int32(_a_F_extractRelOptions_4), int32(3))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								v110 = v100
								m.G0 = v10 + int32(16)
								return v110
							}
						}
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v27
							F_errmsg_internal(m, int32(_a_F_extractRelOptions_0), v10)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_extractRelOptions_1), int32(123), int32(_a_F_extractRelOptions_2))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
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
				v65 = base.I64_extend_i32_u(v23)
				v66 = int32(0)
				v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+22)))
				v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+v68)+119)))
				switch v70 - int32(73) {
				case 0, 32:
					if base.I32_wrap_i64(v65) == int32(0) {
						v110 = v66
						m.G0 = v10 + int32(16)
						return v110
					} else {
						v106 = m.T0[l2].(func(*base.Module, int64, int32) int32)(m, v65, int32(0))
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return int32(0)
						} else {
							v110 = v106
							m.G0 = v10 + int32(16)
							return v110
						}
					}
				default:
					v110 = v66
					m.G0 = v10 + int32(16)
					return v110
				case 36, 41:
					v93 = F_build_reloptions(m, v65, int32(0), int32(1), int32(136), int32(_a_F_extractRelOptions_3), int32(26))
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int32(0)
					} else {
						v110 = v93
						m.G0 = v10 + int32(16)
						return v110
					}
				case 43:
					v78 = F_build_reloptions(m, v65, int32(0), int32(2), int32(136), int32(_a_F_extractRelOptions_3), int32(26))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						if v78 == int32(0) {
							v110 = v66
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v78)+104)) = int64(-4616189618054758400)
							*(*int32)(unsafe.Add(mBase, uint32(v78)+36)) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = int32(100)
							v110 = v78
						}
						m.G0 = v10 + int32(16)
						return v110
					}
				case 45:
					v100 = F_build_reloptions(m, v65, int32(0), int32(512), int32(12), int32(_a_F_extractRelOptions_4), int32(3))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return int32(0)
					} else {
						v110 = v100
						m.G0 = v10 + int32(16)
						return v110
					}
				}
			}
		} else {
			v53 = F_nocachegetattr(m, l0, int32(33), l1)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				v65 = v53
				v66 = int32(0)
				v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+22)))
				v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+v68)+119)))
				switch v70 - int32(73) {
				case 0, 32:
					if base.I32_wrap_i64(v65) == int32(0) {
						v110 = v66
						m.G0 = v10 + int32(16)
						return v110
					} else {
						v106 = m.T0[l2].(func(*base.Module, int64, int32) int32)(m, v65, int32(0))
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return int32(0)
						} else {
							v110 = v106
							m.G0 = v10 + int32(16)
							return v110
						}
					}
				default:
					v110 = v66
					m.G0 = v10 + int32(16)
					return v110
				case 36, 41:
					v93 = F_build_reloptions(m, v65, int32(0), int32(1), int32(136), int32(_a_F_extractRelOptions_3), int32(26))
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int32(0)
					} else {
						v110 = v93
						m.G0 = v10 + int32(16)
						return v110
					}
				case 43:
					v78 = F_build_reloptions(m, v65, int32(0), int32(2), int32(136), int32(_a_F_extractRelOptions_3), int32(26))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						if v78 == int32(0) {
							v110 = v66
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v78)+104)) = int64(-4616189618054758400)
							*(*int32)(unsafe.Add(mBase, uint32(v78)+36)) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = int32(100)
							v110 = v78
						}
						m.G0 = v10 + int32(16)
						return v110
					}
				case 45:
					v100 = F_build_reloptions(m, v65, int32(0), int32(512), int32(12), int32(_a_F_extractRelOptions_4), int32(3))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return int32(0)
					} else {
						v110 = v100
						m.G0 = v10 + int32(16)
						return v110
					}
				}
			}
		}
	} else {
		v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+27)))
		if v55&int32(1) == int32(0) {
			v110 = v4
			m.G0 = v10 + int32(16)
			return v110
		} else {
			v61 = F_nocachegetattr(m, l0, int32(33), l1)
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return int32(0)
			} else {
				v65 = v61
				v66 = int32(0)
				v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+22)))
				v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+v68)+119)))
				switch v70 - int32(73) {
				case 0, 32:
					if base.I32_wrap_i64(v65) == int32(0) {
						v110 = v66
						m.G0 = v10 + int32(16)
						return v110
					} else {
						v106 = m.T0[l2].(func(*base.Module, int64, int32) int32)(m, v65, int32(0))
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return int32(0)
						} else {
							v110 = v106
							m.G0 = v10 + int32(16)
							return v110
						}
					}
				default:
					v110 = v66
					m.G0 = v10 + int32(16)
					return v110
				case 36, 41:
					v93 = F_build_reloptions(m, v65, int32(0), int32(1), int32(136), int32(_a_F_extractRelOptions_3), int32(26))
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int32(0)
					} else {
						v110 = v93
						m.G0 = v10 + int32(16)
						return v110
					}
				case 43:
					v78 = F_build_reloptions(m, v65, int32(0), int32(2), int32(136), int32(_a_F_extractRelOptions_3), int32(26))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						if v78 == int32(0) {
							v110 = v66
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v78)+104)) = int64(-4616189618054758400)
							*(*int32)(unsafe.Add(mBase, uint32(v78)+36)) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = int32(100)
							v110 = v78
						}
						m.G0 = v10 + int32(16)
						return v110
					}
				case 45:
					v100 = F_build_reloptions(m, v65, int32(0), int32(512), int32(12), int32(_a_F_extractRelOptions_4), int32(3))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return int32(0)
					} else {
						v110 = v100
						m.G0 = v10 + int32(16)
						return v110
					}
				}
			}
		}
	}
}
func F_get_rel_namespace(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14289(m, l0, int32(57))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_get_rel_relkind(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v5 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(l0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
			v16 = int32(*(*int8)(unsafe.Add(mBase, uint32(v13+v14)+119)))
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return base.I32_extend8_s(v16)
			}
		}
	}
}
func F_remove_rel_from_jointree(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if l0 == v5 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L19
	} else {
		goto L52
	}
L2:
	;
	m.G0 = v12 + int32(16)
	return v136
L3:
	;
	v136 = int32(0)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v17 - int32(63) {
	case 0:
		goto L8
	case 1:
		goto L6
	case 2:
		goto L7
	default:
		goto L1
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v106 = v12 + int32(12)
	v107 = F_remove_rel_from_jointree(m, v104, l1, v106, l3)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L19
	} else {
		goto L35
	}
L7:
	;
	v27 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v29 == v27 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v20 != l1 {
		v136 = l0
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v22 + int32(1)
	v136 = int32(0)
	goto L2
L10:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v87 == int32(0) {
		v99 = v92
		goto L31
	} else {
		goto L32
	}
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v87 = v32
	goto L10
L12:
	;
	goto L13
L13:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if int32(0) < v33 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v40 = v5
	v42 = v5
	goto L17
L15:
	;
	v69 = v5
	goto L16
L16:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v72 == int32(0) {
		v81 = v74
		goto L26
	} else {
		goto L27
	}
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v40<<(uint(int32(2))%32))))
	v52 = F_remove_rel_from_jointree(m, v49, l1, v12+int32(12), l3)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v69 = v58
	goto L16
L19:
	;
	return int32(0)
L20:
	;
	if v52 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v56 = F_lappend(m, v42, v52)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L19
	} else {
		goto L24
	}
L22:
	;
	v58 = v42
	goto L23
L23:
	;
	v60 = v40 + int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v60 < v61 {
		v40 = v60
		v42 = v58
		goto L17
	} else {
		goto L25
	}
L24:
	;
	v58 = v56
	goto L23
L25:
	;
	goto L18
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v81
	if v69 != 0 {
		v136 = l0
		goto L2
	} else {
		goto L30
	}
L27:
	;
	if v74 == int32(0) {
		v81 = v72
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v79 = F_list_concat(m, v72, v74)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L19
	} else {
		goto L29
	}
L29:
	;
	v81 = v79
	goto L26
L30:
	;
	v87 = v81
	goto L10
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v99
	v136 = int32(0)
	goto L2
L32:
	;
	if v92 == int32(0) {
		v99 = v87
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v97 = F_list_concat(m, v87, v92)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L19
	} else {
		goto L34
	}
L34:
	;
	v99 = v97
	goto L31
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v107
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v111 = F_remove_rel_from_jointree(m, v110, l1, v106, l3)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v111
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v111 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v116 = v114
	goto L39
L38:
	;
	v116 = int32(0)
	goto L39
L39:
	;
	if v116 != 0 {
		v136 = l0
		goto L2
	} else {
		goto L40
	}
L40:
	;
	if v114 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v117 = v114
	goto L43
L42:
	;
	v117 = v111
	goto L43
L43:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v119 == int32(0) {
		v126 = v118
		goto L44
	} else {
		goto L45
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v117
	v132 = F_list_make1_impl(m, int32(1), v12+int32(4))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L19
	} else {
		goto L50
	}
L45:
	;
	if v118 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v126 = v119
	goto L44
L47:
	;
	goto L48
L48:
	;
	v124 = F_list_concat(m, v119, v118)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L19
	} else {
		goto L49
	}
L49:
	;
	v126 = v124
	goto L44
L50:
	;
	v134 = F_makeFromExpr(m, v132, v126)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L19
	} else {
		goto L51
	}
L51:
	;
	v136 = v134
	goto L2
L52:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v153
	F_errmsg_internal(m, int32(_a_F_remove_rel_from_jointree_0), v12)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L19
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_remove_rel_from_jointree_1), int32(1392), int32(_a_F_remove_rel_from_jointree_2))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L19
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
