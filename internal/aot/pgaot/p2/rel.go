package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RelFileLocatorSkippingWAL(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v5 = *(*int32)(unsafe.Add(mBase, _consts[295]))
	if v5 != 0 {
		v6 = int32(0)
		v8 = F_hash_search(m, v5, l0, v6, v6)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v14 = base.B2i32(v8 != int32(0))
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
						v54 = base.I32_div_u_s(int32(8168), v50+int32(28))
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
							v54 = base.I32_div_u_s(int32(8168), v50+int32(28))
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)))
	if v12&int32(1) == v4 {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+532))
		if int32(0) <= v17 {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)))
			v22 = v11 + v20 + v17
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+538)))
			if v23 != int32(1) {
				v62 = v22
				v63 = int32(0)
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+22)))
				v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+v65)+119)))
				switch v67 - int32(73) {
				case 0, 32:
					if v62 == int32(0) {
						v106 = v63
						m.G0 = v9 + int32(16)
						return v106
					} else {
						v102 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, v62, int32(0))
						mBase = m.M
						v103 = m.ExcPending
						if v103 != 0 {
							return int32(0)
						} else {
							v106 = v102
							m.G0 = v9 + int32(16)
							return v106
						}
					}
				default:
					v106 = v63
					m.G0 = v9 + int32(16)
					return v106
				case 36, 41:
					v90 = F_build_reloptions(m, v62, int32(0), int32(1), int32(128), int32(733312), int32(24))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						v106 = v90
						m.G0 = v9 + int32(16)
						return v106
					}
				case 43:
					v75 = F_build_reloptions(m, v62, int32(0), int32(2), int32(128), int32(733312), int32(24))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						if v75 == int32(0) {
							v106 = v63
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v75)+96)) = int64(-4616189618054758400)
							*(*int32)(unsafe.Add(mBase, uint32(v75)+32)) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = int32(100)
							v106 = v75
						}
						m.G0 = v9 + int32(16)
						return v106
					}
				case 45:
					v97 = F_build_reloptions(m, v62, int32(0), int32(512), int32(12), int32(733696), int32(3))
					mBase = m.M
					v98 = m.ExcPending
					if v98 != 0 {
						return int32(0)
					} else {
						v106 = v97
						m.G0 = v9 + int32(16)
						return v106
					}
				}
			} else {
				v26 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+536)))
				switch v26&int32(65535) - int32(1) {
				case 0:
					v31 = int32(*(*int8)(unsafe.Add(mBase, uint32(v22))))
					v62 = v31
					v63 = int32(0)
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+22)))
					v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+v65)+119)))
					switch v67 - int32(73) {
					case 0, 32:
						if v62 == int32(0) {
							v106 = v63
							m.G0 = v9 + int32(16)
							return v106
						} else {
							v102 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, v62, int32(0))
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
								return int32(0)
							} else {
								v106 = v102
								m.G0 = v9 + int32(16)
								return v106
							}
						}
					default:
						v106 = v63
						m.G0 = v9 + int32(16)
						return v106
					case 36, 41:
						v90 = F_build_reloptions(m, v62, int32(0), int32(1), int32(128), int32(733312), int32(24))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							v106 = v90
							m.G0 = v9 + int32(16)
							return v106
						}
					case 43:
						v75 = F_build_reloptions(m, v62, int32(0), int32(2), int32(128), int32(733312), int32(24))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							if v75 == int32(0) {
								v106 = v63
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v75)+96)) = int64(-4616189618054758400)
								*(*int32)(unsafe.Add(mBase, uint32(v75)+32)) = int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = int32(100)
								v106 = v75
							}
							m.G0 = v9 + int32(16)
							return v106
						}
					case 45:
						v97 = F_build_reloptions(m, v62, int32(0), int32(512), int32(12), int32(733696), int32(3))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return int32(0)
						} else {
							v106 = v97
							m.G0 = v9 + int32(16)
							return v106
						}
					}
				case 1:
					v32 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22))))
					v62 = v32
					v63 = int32(0)
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+22)))
					v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+v65)+119)))
					switch v67 - int32(73) {
					case 0, 32:
						if v62 == int32(0) {
							v106 = v63
							m.G0 = v9 + int32(16)
							return v106
						} else {
							v102 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, v62, int32(0))
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
								return int32(0)
							} else {
								v106 = v102
								m.G0 = v9 + int32(16)
								return v106
							}
						}
					default:
						v106 = v63
						m.G0 = v9 + int32(16)
						return v106
					case 36, 41:
						v90 = F_build_reloptions(m, v62, int32(0), int32(1), int32(128), int32(733312), int32(24))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							v106 = v90
							m.G0 = v9 + int32(16)
							return v106
						}
					case 43:
						v75 = F_build_reloptions(m, v62, int32(0), int32(2), int32(128), int32(733312), int32(24))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							if v75 == int32(0) {
								v106 = v63
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v75)+96)) = int64(-4616189618054758400)
								*(*int32)(unsafe.Add(mBase, uint32(v75)+32)) = int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = int32(100)
								v106 = v75
							}
							m.G0 = v9 + int32(16)
							return v106
						}
					case 45:
						v97 = F_build_reloptions(m, v62, int32(0), int32(512), int32(12), int32(733696), int32(3))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return int32(0)
						} else {
							v106 = v97
							m.G0 = v9 + int32(16)
							return v106
						}
					}
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v26
						F_errmsg_internal(m, int32(475879), v9)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(321683), int32(70), int32(67010))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 3:
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
					v62 = v33
					v63 = int32(0)
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+22)))
					v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+v65)+119)))
					switch v67 - int32(73) {
					case 0, 32:
						if v62 == int32(0) {
							v106 = v63
							m.G0 = v9 + int32(16)
							return v106
						} else {
							v102 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, v62, int32(0))
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
								return int32(0)
							} else {
								v106 = v102
								m.G0 = v9 + int32(16)
								return v106
							}
						}
					default:
						v106 = v63
						m.G0 = v9 + int32(16)
						return v106
					case 36, 41:
						v90 = F_build_reloptions(m, v62, int32(0), int32(1), int32(128), int32(733312), int32(24))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							v106 = v90
							m.G0 = v9 + int32(16)
							return v106
						}
					case 43:
						v75 = F_build_reloptions(m, v62, int32(0), int32(2), int32(128), int32(733312), int32(24))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							if v75 == int32(0) {
								v106 = v63
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v75)+96)) = int64(-4616189618054758400)
								*(*int32)(unsafe.Add(mBase, uint32(v75)+32)) = int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = int32(100)
								v106 = v75
							}
							m.G0 = v9 + int32(16)
							return v106
						}
					case 45:
						v97 = F_build_reloptions(m, v62, int32(0), int32(512), int32(12), int32(733696), int32(3))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return int32(0)
						} else {
							v106 = v97
							m.G0 = v9 + int32(16)
							return v106
						}
					}
				}
			}
		} else {
			v50 = F_nocachegetattr(m, l0, int32(33), l1)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				v62 = v50
				v63 = int32(0)
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+22)))
				v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+v65)+119)))
				switch v67 - int32(73) {
				case 0, 32:
					if v62 == int32(0) {
						v106 = v63
						m.G0 = v9 + int32(16)
						return v106
					} else {
						v102 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, v62, int32(0))
						mBase = m.M
						v103 = m.ExcPending
						if v103 != 0 {
							return int32(0)
						} else {
							v106 = v102
							m.G0 = v9 + int32(16)
							return v106
						}
					}
				default:
					v106 = v63
					m.G0 = v9 + int32(16)
					return v106
				case 36, 41:
					v90 = F_build_reloptions(m, v62, int32(0), int32(1), int32(128), int32(733312), int32(24))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						v106 = v90
						m.G0 = v9 + int32(16)
						return v106
					}
				case 43:
					v75 = F_build_reloptions(m, v62, int32(0), int32(2), int32(128), int32(733312), int32(24))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						if v75 == int32(0) {
							v106 = v63
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v75)+96)) = int64(-4616189618054758400)
							*(*int32)(unsafe.Add(mBase, uint32(v75)+32)) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = int32(100)
							v106 = v75
						}
						m.G0 = v9 + int32(16)
						return v106
					}
				case 45:
					v97 = F_build_reloptions(m, v62, int32(0), int32(512), int32(12), int32(733696), int32(3))
					mBase = m.M
					v98 = m.ExcPending
					if v98 != 0 {
						return int32(0)
					} else {
						v106 = v97
						m.G0 = v9 + int32(16)
						return v106
					}
				}
			}
		}
	} else {
		v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+27)))
		if v52&int32(1) == int32(0) {
			v106 = v4
			m.G0 = v9 + int32(16)
			return v106
		} else {
			v58 = F_nocachegetattr(m, l0, int32(33), l1)
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				v62 = v58
				v63 = int32(0)
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+22)))
				v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+v65)+119)))
				switch v67 - int32(73) {
				case 0, 32:
					if v62 == int32(0) {
						v106 = v63
						m.G0 = v9 + int32(16)
						return v106
					} else {
						v102 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, v62, int32(0))
						mBase = m.M
						v103 = m.ExcPending
						if v103 != 0 {
							return int32(0)
						} else {
							v106 = v102
							m.G0 = v9 + int32(16)
							return v106
						}
					}
				default:
					v106 = v63
					m.G0 = v9 + int32(16)
					return v106
				case 36, 41:
					v90 = F_build_reloptions(m, v62, int32(0), int32(1), int32(128), int32(733312), int32(24))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						v106 = v90
						m.G0 = v9 + int32(16)
						return v106
					}
				case 43:
					v75 = F_build_reloptions(m, v62, int32(0), int32(2), int32(128), int32(733312), int32(24))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						if v75 == int32(0) {
							v106 = v63
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v75)+96)) = int64(-4616189618054758400)
							*(*int32)(unsafe.Add(mBase, uint32(v75)+32)) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = int32(100)
							v106 = v75
						}
						m.G0 = v9 + int32(16)
						return v106
					}
				case 45:
					v97 = F_build_reloptions(m, v62, int32(0), int32(512), int32(12), int32(733696), int32(3))
					mBase = m.M
					v98 = m.ExcPending
					if v98 != 0 {
						return int32(0)
					} else {
						v106 = v97
						m.G0 = v9 + int32(16)
						return v106
					}
				}
			}
		}
	}
}
func F_get_rel_namespace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, int32(57), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+68))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_get_rel_relkind(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, int32(57), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v13)+119)))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return base.I32_extend8_s(v15)
			}
		}
	}
}
func F_remove_rel_from_query(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
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
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v144 int32
	_ = v144
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
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
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v467 int32
	_ = v467
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v494 int32
	_ = v494
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v524 int32
	_ = v524
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v14 = F_adjust_relid_set(m, v12, v13, l2)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v14
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v18 = F_adjust_relid_set(m, v17, v13, l2)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v18
	if l3 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v23 = F_bms_del_member(m, v21, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v31 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v23
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v28 = F_bms_del_member(m, v26, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v28
	goto L6
L9:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v144 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v34 <= int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v42 = int32(0)
	goto L12
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v42<<(uint(int32(2))%32))))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v54 = F_bms_copy(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L9
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v54
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v58 = F_bms_copy(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v58
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v62 = F_bms_copy(m, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v62
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	v66 = F_bms_copy(m, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = v66
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v70 = F_adjust_relid_set(m, v69, v13, l2)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v70
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v74 = F_adjust_relid_set(m, v73, v13, l2)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v74
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v78 = F_adjust_relid_set(m, v77, v13, l2)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v78
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	v82 = F_adjust_relid_set(m, v81, v13, l2)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = v82
	if l3 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v130 = v42 + int32(1)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v130 < v131 {
		v42 = v130
		goto L12
	} else {
		goto L35
	}
L23:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v87 = F_bms_del_member(m, v85, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v52)+52))
	F_ChangeVarNodesExtended(m, v125, v13, l2, int32(827))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L34
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v87
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v92 = F_bms_del_member(m, v90, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v92
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v97 = F_bms_del_member(m, v95, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v97
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v102 = F_bms_del_member(m, v100, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = v102
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v52)+28))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v107 = F_bms_del_member(m, v105, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+28)) = v107
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v52)+32))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v112 = F_bms_del_member(m, v110, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+32)) = v112
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v52)+36))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v117 = F_bms_del_member(m, v115, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+36)) = v117
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v52)+40))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v122 = F_bms_del_member(m, v120, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+40)) = v122
	goto L22
L34:
	;
	goto L22
L35:
	;
	goto L13
L36:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v297 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L37:
	;
	v154 = int32(0)
	v157 = v144
	goto L38
L38:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if v159 <= v154 {
		goto L36
	} else {
		goto L40
	}
L39:
	;
	goto L36
L40:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v161+v154<<(uint(int32(2))%32))))
	if l3 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	v247 = F_adjust_relid_set(m, v246, v13, l2)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L64
	}
L42:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v165)+20))
	v169 = int32(0)
	if v168 == v169 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v222 == int32(0) {
		goto L41
	} else {
		goto L57
	}
L44:
	;
	v222 = int32(1)
	goto L43
L45:
	;
	goto L46
L46:
	;
	if l4 == int32(0) {
		v213 = v169
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v222 = v213
	goto L43
L48:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v179 < v178 {
		v213 = v169
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v181 = int32(1)
	if v178 <= v181 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v184 = v181
	goto L52
L51:
	;
	v184 = v178
	goto L52
L52:
	;
	v185 = int32(8)
	v190 = int32(0)
	goto L53
L53:
	;
	v197 = v190 << (uint(int32(2)) % 32)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v168+v185+v197)))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v197+(l4+v185))))
	v204 = v199 & (v201 ^ int32(-1))
	v206 = base.B2i32(v204 == int32(0))
	if v204 != 0 {
		v213 = v206
		goto L47
	} else {
		goto L55
	}
L54:
	;
	v213 = v206
	goto L47
L55:
	;
	v208 = v190 + int32(1)
	if v208 != v184 {
		v190 = v208
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	v226 = F_bms_is_member(m, v13, v225)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	if v226 == int32(0) {
		goto L41
	} else {
		goto L59
	}
L59:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	v232 = F_bms_is_member(m, v230, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	if v232 != 0 {
		goto L41
	} else {
		goto L61
	}
L61:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v235 = F_list_delete_nth_cell(m, v234, v154)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v235
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v238+v239<<(uint(int32(2))%32)))) = int32(0)
	if v235 != 0 {
		v157 = v235
		goto L38
	} else {
		goto L63
	}
L63:
	;
	goto L36
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+12)) = v247
	if l3 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v251 = F_adjust_relid_set(m, v247, v250, l2)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v254 = int32(0)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v165)+20))
	v257 = F_bms_is_member(m, v254, v256)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L69
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+12)) = v251
	goto L67
L69:
	;
	if v257 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v260 = F_bms_make_singleton(m, int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	v262 = v254
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+20)) = v262
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v265 = F_adjust_relid_set(m, v264, v13, l2)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L74
	}
L73:
	;
	v262 = v260
	goto L72
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+16)) = v265
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	v269 = F_bms_difference(m, v265, v268)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+16)) = v269
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v245)+8))
	v273 = F_adjust_relid_set(m, v272, v13, l2)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v245)+8)) = v273
	if l3 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v277 = F_adjust_relid_set(m, v273, v276, l2)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	F_ChangeVarNodesExtended(m, v280, v13, l2, int32(827))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L81
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v245)+8)) = v277
	goto L79
L81:
	;
	if v157 != 0 {
		v154 = v154 + int32(1)
		goto L38
	} else {
		goto L82
	}
L82:
	;
	goto L39
L83:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.Ui32(int32(2)) <= base.Ui32(v494) {
		goto L137
	} else {
		goto L138
	}
L84:
	;
	v300 = int32(0)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	if v301 <= v300 {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v308 = v300
	goto L86
L86:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v297)+12))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v315+v308<<(uint(int32(2))%32))))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)+36))
	v321 = F_bms_is_member(m, v13, v320)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L88
	}
L87:
	;
	goto L83
L88:
	;
	if l3 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	v480 = v308 + int32(1)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	if v480 < v481 {
		v308 = v480
		goto L86
	} else {
		goto L136
	}
L90:
	;
	v347 = int32(0)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v319)+16))
	if v348 == v347 {
		goto L101
	} else {
		goto L102
	}
L91:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v343 = F_adjust_relid_set(m, v341, v342, l2)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L100
	}
L92:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v319)+36))
	v336 = F_adjust_relid_set(m, v335, v13, l2)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L98
	}
L93:
	;
	if v321 != 0 {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v319)+36))
	v327 = F_bms_is_member(m, v325, v326)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	if v327 == int32(0) {
		goto L89
	} else {
		goto L96
	}
L96:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v319)+36))
	v332 = F_adjust_relid_set(m, v331, v13, l2)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v319)+36)) = v332
	v341 = v332
	goto L91
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v319)+36)) = v336
	if l3 == int32(0) {
		goto L90
	} else {
		goto L99
	}
L99:
	;
	v341 = v336
	goto L91
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v319)+36)) = v343
	goto L90
L101:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v319)+24))
	if v420 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L102:
	;
	v352 = v347
	v356 = v348
	goto L103
L103:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v356)+4))
	if v362 <= v352 {
		goto L101
	} else {
		goto L105
	}
L104:
	;
	goto L101
L105:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v356)+12))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v364+v352<<(uint(int32(2))%32))))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)+8))
	v370 = F_bms_is_member(m, v13, v369)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L109
	}
L106:
	;
	if v405 != 0 {
		v352 = v404 + int32(1)
		v356 = v405
		goto L103
	} else {
		goto L122
	}
L107:
	;
	if v397 != 0 {
		v404 = v352
		v405 = v356
		goto L106
	} else {
		goto L120
	}
L108:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v394 = F_adjust_relid_set(m, v392, v393, l2)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L119
	}
L109:
	;
	if v370 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	if l3 == int32(0) {
		v404 = v352
		v405 = v356
		goto L106
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v368)+8))
	v387 = F_adjust_relid_set(m, v386, v13, l2)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L117
	}
L113:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v368)+8))
	v378 = F_bms_is_member(m, v376, v377)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	if v378 == int32(0) {
		v404 = v352
		v405 = v356
		goto L106
	} else {
		goto L115
	}
L115:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v368)+8))
	v383 = F_adjust_relid_set(m, v382, v13, l2)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v368)+8)) = v383
	v392 = v383
	goto L108
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v368)+8)) = v387
	if l3 == int32(0) {
		v397 = v387
		goto L107
	} else {
		goto L118
	}
L118:
	;
	v392 = v387
	goto L108
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v368)+8)) = v394
	v397 = v394
	goto L107
L120:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v319)+16))
	v399 = F_list_delete_nth_cell(m, v398, v352)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v319)+16)) = v399
	v404 = v352 - int32(1)
	v405 = v399
	goto L106
L122:
	;
	goto L104
L123:
	;
	F_ec_clear_derived_clauses(m, v319)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L135
	}
L124:
	;
	v423 = int32(0)
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
	if v424 <= v423 {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v428 = v423
	goto L126
L126:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v420)+12))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v438+v428<<(uint(int32(2))%32))))
	if l3 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	goto L123
L128:
	;
	v452 = v428 + int32(1)
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
	if v452 < v453 {
		v428 = v452
		goto L126
	} else {
		goto L134
	}
L129:
	;
	F_ChangeVarNodesExtended(m, v442, v13, l2, int32(827))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	F_remove_rel_from_restrictinfo(m, v442, v13, v448)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L133
	}
L132:
	;
	goto L128
L133:
	;
	goto L128
L134:
	;
	goto L127
L135:
	;
	goto L89
L136:
	;
	goto L87
L137:
	;
	v507 = int32(1)
	goto L140
L138:
	;
	goto L139
L139:
	;
	return
L140:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v509+v507<<(uint(int32(2))%32))))
	if v513 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	goto L139
L142:
	;
	v581 = v507 + int32(1)
	v582 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.Ui32(v581) < base.Ui32(v582) {
		v507 = v581
		goto L140
	} else {
		goto L157
	}
L143:
	;
	v516 = int32(*(*int16)(unsafe.Add(mBase, uint32(v513)+82)))
	v517 = int32(*(*int16)(unsafe.Add(mBase, uint32(v513)+80)))
	v518 = v516 - v517
	if int32(0) <= v518 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v524 = v518
	goto L147
L145:
	;
	goto L146
L146:
	;
	if l2 <= int32(0) {
		goto L142
	} else {
		goto L155
	}
L147:
	;
	v532 = int32(0)
	v535 = v524 << (uint(int32(2)) % 32)
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v513)+84))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v535+v536)))
	v539 = F_bms_is_member(m, v532, v538)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L1
	} else {
		goto L149
	}
L148:
	;
	goto L146
L149:
	;
	if v539 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v542 = F_bms_make_singleton(m, int32(0))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L153
	}
L151:
	;
	v544 = v532
	goto L152
L152:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v513)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v545+v535))) = v544
	if int32(0) < v524 {
		v524 = v524 - int32(1)
		goto L147
	} else {
		goto L154
	}
L153:
	;
	v544 = v542
	goto L152
L154:
	;
	goto L148
L155:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v513)+100))
	F_ChangeVarNodesExtended(m, v565, v13, l2, int32(827))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	goto L142
L157:
	;
	goto L141
}
func F_remove_rel_from_restrictinfo(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_bms_copy(m, v7)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v8
	v11 = F_bms_del_member(m, v8, l1)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v11
	v14 = F_bms_del_member(m, v11, l2)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v14
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v18 = F_bms_copy(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v18
	v21 = F_bms_del_member(m, v18, l1)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v21
	v24 = F_bms_del_member(m, v21, l2)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v24
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	goto L9
L8:
	;
	return
L9:
	;
	if base.B2i32(v27 != int32(0)) == int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	if v33 == int32(0) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v36 <= int32(0) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v44 = int32(0)
	goto L13
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v44<<(uint(int32(2))%32))))
	if v49 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L8
L15:
	;
	v89 = v44 + int32(1)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v89 < v90 {
		v44 = v89
		goto L13
	} else {
		goto L27
	}
L16:
	;
	F_remove_rel_from_restrictinfo(m, v49, l1, l2)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L26
	}
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v52 != int32(21) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v55 != 0 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	if v56 == int32(0) {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v59 = int32(0)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v60 <= v59 {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v63 = v59
	goto L22
L22:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69+v63<<(uint(int32(2))%32))))
	F_remove_rel_from_restrictinfo(m, v73, l1, l2)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	goto L15
L24:
	;
	v77 = v63 + int32(1)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v77 < v78 {
		v63 = v77
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	goto L15
L27:
	;
	goto L14
}
