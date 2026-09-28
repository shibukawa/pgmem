package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_do_numeric_accum(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int64
	_ = v18
	var v22 int64
	_ = v22
	var v26 int64
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v88 int64
	_ = v88
	var v92 int32
	_ = v92
	var v95 int64
	_ = v95
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int64
	_ = v113
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if base.Ui32(int32(_a_F_do_numeric_accum_0)) <= base.Ui32(v11) {
		if v11 != int32(_a_F_do_numeric_accum_1) {
			if v11 != int32(_a_F_do_numeric_accum_2) {
				v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v26 + int64(1)
			} else {
				v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+96)) = v18 + int64(1)
			}
		} else {
			v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+104)) = v22 + int64(1)
		}
		m.G0 = v9 + int32(48)
		return
	} else {
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v35 = base.I32_extend16_s(v11)
		v37 = base.B2i32(int32(0) <= v35)
		if int32(0) <= v35 {
			v38 = int32(-8)
		} else {
			v38 = int32(-6)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(base.Ui32(int32(base.Ui32(v30)>>(uint(int32(2))%32))+v38) >> (uint(int32(1)) % 32))
		if int32(0) <= v35 {
			v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
			v53 = v43
		} else {
			v53 = v11<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v11&int32(63)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v53
		v55 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v55
		v64 = base.B2i32(v35 < v55)
		if v35 < v55 {
			v65 = int32(base.Ui32(v11)>>(uint(int32(7))%32)) & int32(63)
		} else {
			v65 = v11 & int32(_a_F_do_numeric_accum_3)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v65
		v72 = v11 & int32(_a_F_do_numeric_accum_0)
		if v72 == int32(_a_F_do_numeric_accum_4) {
			v75 = v11 << (uint(int32(1)) % 32) & int32(_a_F_do_numeric_accum_5)
		} else {
			v75 = v72
		}
		*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v75
		if v35 < v55 {
			v79 = int32(6)
		} else {
			v79 = int32(8)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = l1 + v79
		v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
		if v82 < v65 {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v65
		} else {
			if v82 != v65 {
			} else {
				v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+80))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = v88 + int64(1)
			}
		}
		v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		if v92 == int32(1) {
			v95 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v95
			*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v95
			*(*int64)(unsafe.Add(mBase, uint32(v9))) = v95
			v102 = v9 + int32(24)
			F_mul_var(m, v102, v102, v9, v65<<(uint(int32(1))%32))
			mBase = m.M
			v106 = m.ExcPending
			if v106 != 0 {
				return
			} else {
				v108 = int32(_a_F_do_numeric_accum_6)
				v109 = *(*int32)(unsafe.Add(mBase, _c_F_do_numeric_accum[0]))
				v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_accum[0])) = v111
				v113 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v113 + int64(1)
				F_accum_sum_add(m, l0+int32(16), v9+int32(24))
				mBase = m.M
				v122 = m.ExcPending
				if v122 != 0 {
					return
				} else {
					v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
					if v123 == int32(1) {
						F_accum_sum_add(m, l0+int32(44), v9)
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_accum[0])) = v109
							m.G0 = v9 + int32(48)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_accum[0])) = v109
						m.G0 = v9 + int32(48)
						return
					}
				}
			}
		} else {
			v108 = int32(_a_F_do_numeric_accum_6)
			v109 = *(*int32)(unsafe.Add(mBase, _c_F_do_numeric_accum[0]))
			v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_accum[0])) = v111
			v113 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v113 + int64(1)
			F_accum_sum_add(m, l0+int32(16), v9+int32(24))
			mBase = m.M
			v122 = m.ExcPending
			if v122 != 0 {
				return
			} else {
				v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				if v123 == int32(1) {
					F_accum_sum_add(m, l0+int32(44), v9)
					mBase = m.M
					v129 = m.ExcPending
					if v129 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_accum[0])) = v109
						m.G0 = v9 + int32(48)
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_accum[0])) = v109
					m.G0 = v9 + int32(48)
					return
				}
			}
		}
	}
}
func F_numeric_accum_inv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v4 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_numeric_accum_inv_0), int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_numeric_accum_inv_1), int32(_a_F_numeric_accum_inv_2), int32(_a_F_numeric_accum_inv_3))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v6 = base.I32_wrap_i64(v5)
		if v6 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_numeric_accum_inv_0), int32(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_numeric_accum_inv_1), int32(_a_F_numeric_accum_inv_2), int32(_a_F_numeric_accum_inv_3))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
			if v9 != 0 {
				return v5 & int64(4294967295)
			} else {
				v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v11 = F_pg_detoast_datum(m, v10)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return int64(0)
				} else {
					v15 = F_do_numeric_discard(m, v6, v11)
					mBase = m.M
					v16 = m.ExcPending
					if v16 != 0 {
						return int64(0)
					} else {
						if v15 != 0 {
							return v5 & int64(4294967295)
						} else {
							v17 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
							return int64(0)
						}
					}
				}
			}
		}
	}
}
func F_numeric_add_safe(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v152 int32
	_ = v152
	var v154 int64
	_ = v154
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v15 = base.I32_extend16_s(v14)
	if base.Ui32(v14) <= base.Ui32(int32(_a_F_numeric_add_safe_0)) {
		v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		v19 = base.I32_extend16_s(v18)
		if base.Ui32(int32(_a_F_numeric_add_safe_0)) < base.Ui32(v18) {
			v47 = v19
			if v47&int32(_a_F_numeric_add_safe_1) != int32(_a_F_numeric_add_safe_2) {
				if v14 != int32(_a_F_numeric_add_safe_3) {
					if v14 != int32(_a_F_numeric_add_safe_4) {
						if v47&int32(_a_F_numeric_add_safe_1) == int32(_a_F_numeric_add_safe_4) {
							v94 = F_make_result_safe(m, int32(_a_F_numeric_add_safe_5), int32(0))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								v202 = v94
								m.G0 = v12 + int32(80)
								return v202
							}
						} else {
							v98 = F_make_result_safe(m, int32(_a_F_numeric_add_safe_6), int32(0))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								v202 = v98
								m.G0 = v12 + int32(80)
								return v202
							}
						}
					} else {
						if v47&int32(_a_F_numeric_add_safe_1) == int32(_a_F_numeric_add_safe_3) {
							v70 = F_make_result_safe(m, int32(_a_F_numeric_add_safe_7), int32(0))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								v202 = v70
								m.G0 = v12 + int32(80)
								return v202
							}
						} else {
							v74 = F_make_result_safe(m, int32(_a_F_numeric_add_safe_5), int32(0))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								v202 = v74
								m.G0 = v12 + int32(80)
								return v202
							}
						}
					}
				} else {
					if v47&int32(_a_F_numeric_add_safe_1) == int32(_a_F_numeric_add_safe_4) {
						v82 = F_make_result_safe(m, int32(_a_F_numeric_add_safe_7), int32(0))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							v202 = v82
							m.G0 = v12 + int32(80)
							return v202
						}
					} else {
						v86 = F_make_result_safe(m, int32(_a_F_numeric_add_safe_6), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							v202 = v86
							m.G0 = v12 + int32(80)
							return v202
						}
					}
				}
			} else {
				v56 = F_make_result_safe(m, int32(_a_F_numeric_add_safe_7), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					v202 = v56
					m.G0 = v12 + int32(80)
					return v202
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v28 = base.B2i32(int32(0) <= v15)
			if int32(0) <= v15 {
				v29 = int32(-8)
			} else {
				v29 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = int32(base.Ui32(int32(base.Ui32(v22)>>(uint(int32(2))%32))+v29) >> (uint(int32(1)) % 32))
			if int32(0) <= v15 {
				v100 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
				v101 = v100
			} else {
				v101 = v14<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v14&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v101
			v103 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v103
			v112 = base.B2i32(v15 < v103)
			if v15 < v103 {
				v113 = int32(base.Ui32(v14)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v113 = v14 & int32(_a_F_numeric_add_safe_8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v113
			v120 = v14 & int32(_a_F_numeric_add_safe_2)
			if v120 == int32(_a_F_numeric_add_safe_9) {
				v123 = v14 << (uint(int32(1)) % 32) & int32(_a_F_numeric_add_safe_10)
			} else {
				v123 = v120
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v123
			if v15 < v103 {
				v127 = int32(6)
			} else {
				v127 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = l0 + v127
			v130 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v136 = base.B2i32(int32(0) <= v19)
			if int32(0) <= v19 {
				v137 = int32(-8)
			} else {
				v137 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(base.Ui32(int32(base.Ui32(v130)>>(uint(int32(2))%32))+v137) >> (uint(int32(1)) % 32))
			if int32(0) <= v19 {
				v142 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
				v152 = v142
			} else {
				v152 = v18<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v18&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v152
			v154 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v154
			*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v154
			*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v154
			v160 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v160
			v169 = base.B2i32(v19 < v160)
			if v19 < v160 {
				v170 = int32(base.Ui32(v18)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v170 = v18 & int32(_a_F_numeric_add_safe_8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v170
			v177 = v18 & int32(_a_F_numeric_add_safe_2)
			if v177 == int32(_a_F_numeric_add_safe_9) {
				v180 = v18 << (uint(int32(1)) % 32) & int32(_a_F_numeric_add_safe_10)
			} else {
				v180 = v177
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v180
			if v19 < v160 {
				v184 = int32(6)
			} else {
				v184 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = l1 + v184
			v192 = v12 + int32(8)
			F_add_var(m, v12+int32(56), v12+int32(32), v192)
			mBase = m.M
			v194 = m.ExcPending
			if v194 != 0 {
				return int32(0)
			} else {
				v195 = F_make_result_safe(m, v192, l2)
				mBase = m.M
				v196 = m.ExcPending
				if v196 != 0 {
					return int32(0)
				} else {
					v197 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
					if v197 == int32(0) {
						v202 = v195
						m.G0 = v12 + int32(80)
						return v202
					} else {
						F_pfree(m, v197)
						mBase = m.M
						v201 = m.ExcPending
						if v201 != 0 {
							return int32(0)
						} else {
							v202 = v195
							m.G0 = v12 + int32(80)
							return v202
						}
					}
				}
			}
		}
	} else {
		if v15 == int32(-16384) {
			v56 = F_make_result_safe(m, int32(_a_F_numeric_add_safe_7), int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				v202 = v56
				m.G0 = v12 + int32(80)
				return v202
			}
		} else {
			v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			v47 = v45
			if v47&int32(_a_F_numeric_add_safe_1) != int32(_a_F_numeric_add_safe_2) {
				if v14 != int32(_a_F_numeric_add_safe_3) {
					if v14 != int32(_a_F_numeric_add_safe_4) {
						if v47&int32(_a_F_numeric_add_safe_1) == int32(_a_F_numeric_add_safe_4) {
							v94 = F_make_result_safe(m, int32(_a_F_numeric_add_safe_5), int32(0))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								v202 = v94
								m.G0 = v12 + int32(80)
								return v202
							}
						} else {
							v98 = F_make_result_safe(m, int32(_a_F_numeric_add_safe_6), int32(0))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								v202 = v98
								m.G0 = v12 + int32(80)
								return v202
							}
						}
					} else {
						if v47&int32(_a_F_numeric_add_safe_1) == int32(_a_F_numeric_add_safe_3) {
							v70 = F_make_result_safe(m, int32(_a_F_numeric_add_safe_7), int32(0))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								v202 = v70
								m.G0 = v12 + int32(80)
								return v202
							}
						} else {
							v74 = F_make_result_safe(m, int32(_a_F_numeric_add_safe_5), int32(0))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								v202 = v74
								m.G0 = v12 + int32(80)
								return v202
							}
						}
					}
				} else {
					if v47&int32(_a_F_numeric_add_safe_1) == int32(_a_F_numeric_add_safe_4) {
						v82 = F_make_result_safe(m, int32(_a_F_numeric_add_safe_7), int32(0))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							v202 = v82
							m.G0 = v12 + int32(80)
							return v202
						}
					} else {
						v86 = F_make_result_safe(m, int32(_a_F_numeric_add_safe_6), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							v202 = v86
							m.G0 = v12 + int32(80)
							return v202
						}
					}
				}
			} else {
				v56 = F_make_result_safe(m, int32(_a_F_numeric_add_safe_7), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					v202 = v56
					m.G0 = v12 + int32(80)
					return v202
				}
			}
		}
	}
}
func F_numeric_avg_serialize(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v129 int64
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int64
	_ = v136
	var v138 int64
	_ = v138
	var v140 int64
	_ = v140
	var v143 int64
	_ = v143
	var v145 int64
	_ = v145
	var v147 int64
	_ = v147
	var v149 int64
	_ = v149
	var v172 int32
	_ = v172
	var v175 int64
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int64
	_ = v182
	var v184 int64
	_ = v184
	var v186 int64
	_ = v186
	var v189 int64
	_ = v189
	var v191 int64
	_ = v191
	var v193 int64
	_ = v193
	var v195 int64
	_ = v195
	var v218 int32
	_ = v218
	var v221 int64
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int64
	_ = v228
	var v230 int64
	_ = v230
	var v232 int64
	_ = v232
	var v235 int64
	_ = v235
	var v237 int64
	_ = v237
	var v239 int64
	_ = v239
	var v241 int64
	_ = v241
	var v264 int32
	_ = v264
	var v267 int64
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int64
	_ = v274
	var v276 int64
	_ = v276
	var v278 int64
	_ = v278
	var v281 int64
	_ = v281
	var v283 int64
	_ = v283
	var v285 int64
	_ = v285
	var v287 int64
	_ = v287
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 == int32(0) {
		v41 = int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		switch v16 - int32(435) {
		case 0:
			v41 = int32(1)
		case 1:
			v41 = int32(2)
		default:
			v41 = int32(0)
		}
	}
	if v41 != 0 {
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v43 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v43
		*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v43
		*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v43
		v50 = v9 + int32(32)
		F_pq_begintypsend(m, v50)
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return int64(0)
		} else {
			v55 = *(*int64)(unsafe.Add(mBase, uint32(v42)+8))
			F_enlargeStringInfo(m, v50, int32(8))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int64(0)
			} else {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
				v62 = int64(56)
				v64 = int64(65280)
				v66 = int64(40)
				v69 = int64(16711680)
				v71 = int64(24)
				v73 = int64(4278190080)
				v75 = int64(8)
				*(*int64)(unsafe.Add(mBase, uint32(v59+v60))) = v55<<(uint(v62)%64) | v55&v64<<(uint(v66)%64) | (v55&v69<<(uint(v71)%64) | v55&v73<<(uint(v75)%64)) | (int64(base.Ui64(v55)>>(uint(v75)%64))&v73 | int64(base.Ui64(v55)>>(uint(v71)%64))&v69 | (int64(base.Ui64(v55)>>(uint(v66)%64))&v64 | int64(base.Ui64(v55)>>(uint(v62)%64))))
				v98 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v59 + v98
				v104 = v9 + v98
				F_accum_sum_final(m, v42+int32(16), v104)
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return int64(0)
				} else {
					F_numericvar_serialize(m, v50, v104)
					mBase = m.M
					v108 = m.ExcPending
					if v108 != 0 {
						return int64(0)
					} else {
						v109 = *(*int32)(unsafe.Add(mBase, uint32(v42)+72))
						F_enlargeStringInfo(m, v50, int32(4))
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return int64(0)
						} else {
							v113 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
							v114 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
							v118 = int32(16711935)
							v122 = int32(8)
							*(*int32)(unsafe.Add(mBase, uint32(v113+v114))) = base.I32_rotr(v109, int32(24))&v118 | base.I32_rotr(v109&v118, v122)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v113 + int32(4)
							v129 = *(*int64)(unsafe.Add(mBase, uint32(v42)+80))
							F_enlargeStringInfo(m, v50, v122)
							mBase = m.M
							v132 = m.ExcPending
							if v132 != 0 {
								return int64(0)
							} else {
								v133 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
								v134 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
								v136 = int64(56)
								v138 = int64(65280)
								v140 = int64(40)
								v143 = int64(16711680)
								v145 = int64(24)
								v147 = int64(4278190080)
								v149 = int64(8)
								*(*int64)(unsafe.Add(mBase, uint32(v133+v134))) = v129<<(uint(v136)%64) | v129&v138<<(uint(v140)%64) | (v129&v143<<(uint(v145)%64) | v129&v147<<(uint(v149)%64)) | (int64(base.Ui64(v129)>>(uint(v149)%64))&v147 | int64(base.Ui64(v129)>>(uint(v145)%64))&v143 | (int64(base.Ui64(v129)>>(uint(v140)%64))&v138 | int64(base.Ui64(v129)>>(uint(v136)%64))))
								v172 = int32(8)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v133 + v172
								v175 = *(*int64)(unsafe.Add(mBase, uint32(v42)+88))
								F_enlargeStringInfo(m, v50, v172)
								mBase = m.M
								v178 = m.ExcPending
								if v178 != 0 {
									return int64(0)
								} else {
									v179 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
									v180 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
									v182 = int64(56)
									v184 = int64(65280)
									v186 = int64(40)
									v189 = int64(16711680)
									v191 = int64(24)
									v193 = int64(4278190080)
									v195 = int64(8)
									*(*int64)(unsafe.Add(mBase, uint32(v179+v180))) = v175<<(uint(v182)%64) | v175&v184<<(uint(v186)%64) | (v175&v189<<(uint(v191)%64) | v175&v193<<(uint(v195)%64)) | (int64(base.Ui64(v175)>>(uint(v195)%64))&v193 | int64(base.Ui64(v175)>>(uint(v191)%64))&v189 | (int64(base.Ui64(v175)>>(uint(v186)%64))&v184 | int64(base.Ui64(v175)>>(uint(v182)%64))))
									v218 = int32(8)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v179 + v218
									v221 = *(*int64)(unsafe.Add(mBase, uint32(v42)+96))
									F_enlargeStringInfo(m, v50, v218)
									mBase = m.M
									v224 = m.ExcPending
									if v224 != 0 {
										return int64(0)
									} else {
										v225 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
										v226 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
										v228 = int64(56)
										v230 = int64(65280)
										v232 = int64(40)
										v235 = int64(16711680)
										v237 = int64(24)
										v239 = int64(4278190080)
										v241 = int64(8)
										*(*int64)(unsafe.Add(mBase, uint32(v225+v226))) = v221<<(uint(v228)%64) | v221&v230<<(uint(v232)%64) | (v221&v235<<(uint(v237)%64) | v221&v239<<(uint(v241)%64)) | (int64(base.Ui64(v221)>>(uint(v241)%64))&v239 | int64(base.Ui64(v221)>>(uint(v237)%64))&v235 | (int64(base.Ui64(v221)>>(uint(v232)%64))&v230 | int64(base.Ui64(v221)>>(uint(v228)%64))))
										v264 = int32(8)
										*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v225 + v264
										v267 = *(*int64)(unsafe.Add(mBase, uint32(v42)+104))
										F_enlargeStringInfo(m, v50, v264)
										mBase = m.M
										v270 = m.ExcPending
										if v270 != 0 {
											return int64(0)
										} else {
											v271 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
											v272 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
											v274 = int64(56)
											v276 = int64(65280)
											v278 = int64(40)
											v281 = int64(16711680)
											v283 = int64(24)
											v285 = int64(4278190080)
											v287 = int64(8)
											*(*int64)(unsafe.Add(mBase, uint32(v271+v272))) = v267<<(uint(v274)%64) | v267&v276<<(uint(v278)%64) | (v267&v281<<(uint(v283)%64) | v267&v285<<(uint(v287)%64)) | (int64(base.Ui64(v267)>>(uint(v287)%64))&v285 | int64(base.Ui64(v267)>>(uint(v283)%64))&v281 | (int64(base.Ui64(v267)>>(uint(v278)%64))&v276 | int64(base.Ui64(v267)>>(uint(v274)%64))))
											*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v271 + int32(8)
											v314 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
											v315 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v314))) = v315 << (uint(int32(2)) % 32)
											v319 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
											if v319 != 0 {
												F_pfree(m, v319)
												mBase = m.M
												v321 = m.ExcPending
												if v321 != 0 {
													return int64(0)
												} else {
													m.G0 = v9 + int32(48)
													return base.I64_extend_i32_u(v314)
												}
											} else {
												m.G0 = v9 + int32(48)
												return base.I64_extend_i32_u(v314)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v330 = m.ExcPending
		if v330 != 0 {
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_numeric_avg_serialize_0), int32(0))
			mBase = m.M
			v334 = m.ExcPending
			if v334 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_numeric_avg_serialize_1), int32(_a_F_numeric_avg_serialize_2), int32(_a_F_numeric_avg_serialize_3))
				mBase = m.M
				v339 = m.ExcPending
				if v339 != 0 {
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
func F_numeric_deserialize(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
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
	var v96 int64
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int64
	_ = v117
	var v118 int32
	_ = v118
	var v120 int64
	_ = v120
	var v121 int32
	_ = v121
	var v123 int64
	_ = v123
	var v124 int32
	_ = v124
	var v126 int64
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12 == int32(0) {
		v40 = int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		switch v15 - int32(435) {
		case 0:
			v40 = int32(1)
		case 1:
			v40 = int32(2)
		default:
			v40 = int32(0)
		}
	}
	if v40 != 0 {
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v42 = F_pg_detoast_datum_packed(m, v41)
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int64(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = int64(0)
			v48 = int32(1)
			v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
			v52 = v50 & v48
			if v52 != 0 {
				v53 = v48
			} else {
				v53 = int32(4)
			}
			if v50 == int32(1) {
				v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
				if v60 == int32(18) {
					v63 = int32(16)
				} else {
					v63 = int32(0)
				}
				if base.Ui32((v60-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v70 = int32(4)
				} else {
					v70 = v63
				}
				v81 = v70
			} else {
				v71 = int32(1)
				if v52 != 0 {
					v81 = int32(base.Ui32(v50)>>(uint(v71)%32)) - v71
				} else {
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					v81 = int32(base.Ui32(v75)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v81
			*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v42 + v53
			v87 = F_palloc0(m, int32(112))
			mBase = m.M
			v88 = m.ExcPending
			if v88 != 0 {
				return int64(0)
			} else {
				v89 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v87))) = uint8(v89)
				v92 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_deserialize[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v87)+4)) = v92
				v95 = v8 + int32(32)
				v96 = F_pq_getmsgint64(m, v95)
				mBase = m.M
				v97 = m.ExcPending
				if v97 != 0 {
					return int64(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v87)+8)) = v96
					v100 = v8 + int32(8)
					F_numericvar_deserialize(m, v95, v100)
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int64(0)
					} else {
						F_accum_sum_add(m, v87+int32(16), v100)
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int64(0)
						} else {
							F_numericvar_deserialize(m, v95, v100)
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int64(0)
							} else {
								F_accum_sum_add(m, v87+int32(44), v100)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int64(0)
								} else {
									v114 = F_pq_getmsgint(m, v95, int32(4))
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v87)+72)) = v114
										v117 = F_pq_getmsgint64(m, v95)
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int64(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v87)+80)) = v117
											v120 = F_pq_getmsgint64(m, v95)
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return int64(0)
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v87)+88)) = v120
												v123 = F_pq_getmsgint64(m, v95)
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return int64(0)
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v87)+96)) = v123
													v126 = F_pq_getmsgint64(m, v95)
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return int64(0)
													} else {
														*(*int64)(unsafe.Add(mBase, uint32(v87)+104)) = v126
														F_pq_getmsgend(m, v95)
														mBase = m.M
														v130 = m.ExcPending
														if v130 != 0 {
															return int64(0)
														} else {
															v131 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
															if v131 != 0 {
																F_pfree(m, v131)
																mBase = m.M
																v133 = m.ExcPending
																if v133 != 0 {
																	return int64(0)
																} else {
																	m.G0 = v8 + int32(48)
																	return base.I64_extend_i32_u(v87)
																}
															} else {
																m.G0 = v8 + int32(48)
																return base.I64_extend_i32_u(v87)
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v142 = m.ExcPending
		if v142 != 0 {
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_numeric_deserialize_0), int32(0))
			mBase = m.M
			v146 = m.ExcPending
			if v146 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_numeric_deserialize_1), int32(_a_F_numeric_deserialize_2), int32(_a_F_numeric_deserialize_3))
				mBase = m.M
				v151 = m.ExcPending
				if v151 != 0 {
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
func F_numeric_div_safe(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int64
	_ = v207
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v273 int32
	_ = v273
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v325 int32
	_ = v325
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v19 = base.I32_extend16_s(v18)
	if base.Ui32(v18) <= base.Ui32(int32(_a_F_numeric_div_safe_0)) {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	m.G0 = v16 + int32(80)
	return v387
L2:
	;
	v367 = int32(0)
	v368 = F_errsave_start(m, l2)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L18
	} else {
		goto L110
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v151
	v153 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v153
	v162 = base.B2i32(v19 < v153)
	if v19 < v153 {
		goto L58
	} else {
		goto L59
	}
L4:
	;
	v150 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	v151 = v150
	goto L3
L5:
	;
	if v18 != int32(_a_F_numeric_div_safe_1) {
		goto L21
	} else {
		goto L22
	}
L6:
	;
	v60 = F_make_result_safe(m, int32(_a_F_numeric_div_safe_2), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L18
	} else {
		goto L19
	}
L7:
	;
	if v50&int32(_a_F_numeric_div_safe_3) != int32(_a_F_numeric_div_safe_4) {
		goto L5
	} else {
		goto L17
	}
L8:
	;
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v23 = base.I32_extend16_s(v22)
	if base.Ui32(int32(_a_F_numeric_div_safe_0)) < base.Ui32(v22) {
		v50 = v23
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	if v19 == int32(-16384) {
		goto L6
	} else {
		goto L16
	}
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = base.B2i32(int32(0) <= v19)
	if int32(0) <= v19 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v33 = int32(-8)
	goto L14
L13:
	;
	v33 = int32(-6)
	goto L14
L14:
	;
	v36 = int32(base.Ui32(int32(base.Ui32(v26)>>(uint(int32(2))%32))+v33) >> (uint(int32(1)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v36
	if int32(0) <= v19 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v151 = v18<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v18&int32(63)
	goto L3
L16:
	;
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v50 = v49
	goto L7
L17:
	;
	goto L6
L18:
	;
	return int32(0)
L19:
	;
	v387 = v60
	goto L1
L20:
	;
	v148 = F_make_result_safe(m, int32(_a_F_numeric_div_safe_5), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L18
	} else {
		goto L57
	}
L21:
	;
	if v18 != int32(_a_F_numeric_div_safe_6) {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if base.Ui32(int32(_a_F_numeric_div_safe_4)) <= base.Ui32(v50&int32(_a_F_numeric_div_safe_3)) {
		goto L41
	} else {
		goto L42
	}
L24:
	;
	if base.Ui32(int32(_a_F_numeric_div_safe_4)) <= base.Ui32(v50&int32(_a_F_numeric_div_safe_3)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v74 = F_make_result_safe(m, int32(_a_F_numeric_div_safe_2), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L18
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if int32(0) <= base.I32_extend16_s(v50) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v387 = v74
	goto L1
L29:
	;
	v84 = int32(-8)
	goto L31
L30:
	;
	v84 = int32(-6)
	goto L31
L31:
	;
	if base.Ui32(int32(base.Ui32(v76)>>(uint(int32(2))%32))+v84) < base.Ui32(int32(2)) {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v93 = v50 & int32(_a_F_numeric_div_safe_4)
	if v93 == int32(_a_F_numeric_div_safe_7) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v96 = v50 << (uint(int32(1)) % 32) & int32(_a_F_numeric_div_safe_8)
	goto L35
L34:
	;
	v96 = v93
	goto L35
L35:
	;
	if v96 != int32(_a_F_numeric_div_safe_8) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v101 = F_make_result_safe(m, int32(_a_F_numeric_div_safe_9), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L18
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v105 = F_make_result_safe(m, int32(_a_F_numeric_div_safe_10), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L18
	} else {
		goto L40
	}
L39:
	;
	v387 = v101
	goto L1
L40:
	;
	v387 = v105
	goto L1
L41:
	;
	v113 = F_make_result_safe(m, int32(_a_F_numeric_div_safe_2), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L18
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if int32(0) <= base.I32_extend16_s(v50) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v387 = v113
	goto L1
L45:
	;
	v123 = int32(-8)
	goto L47
L46:
	;
	v123 = int32(-6)
	goto L47
L47:
	;
	if base.Ui32(int32(base.Ui32(v115)>>(uint(int32(2))%32))+v123) < base.Ui32(int32(2)) {
		goto L2
	} else {
		goto L48
	}
L48:
	;
	v132 = v50 & int32(_a_F_numeric_div_safe_4)
	if v132 == int32(_a_F_numeric_div_safe_7) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v135 = v50 << (uint(int32(1)) % 32) & int32(_a_F_numeric_div_safe_8)
	goto L51
L50:
	;
	v135 = v132
	goto L51
L51:
	;
	if v135 != int32(_a_F_numeric_div_safe_8) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v140 = F_make_result_safe(m, int32(_a_F_numeric_div_safe_10), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L18
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v144 = F_make_result_safe(m, int32(_a_F_numeric_div_safe_9), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L18
	} else {
		goto L56
	}
L55:
	;
	v387 = v140
	goto L1
L56:
	;
	v387 = v144
	goto L1
L57:
	;
	v387 = v148
	goto L1
L58:
	;
	v163 = int32(base.Ui32(v18)>>(uint(int32(7))%32)) & int32(63)
	goto L60
L59:
	;
	v163 = v18 & int32(_a_F_numeric_div_safe_11)
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v163
	v170 = v18 & int32(_a_F_numeric_div_safe_4)
	if v170 == int32(_a_F_numeric_div_safe_7) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v173 = v18 << (uint(int32(1)) % 32) & int32(_a_F_numeric_div_safe_8)
	goto L63
L62:
	;
	v173 = v170
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v173
	if v19 < v153 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v177 = int32(6)
	goto L66
L65:
	;
	v177 = int32(8)
	goto L66
L66:
	;
	v178 = l0 + v177
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v178
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v186 = base.B2i32(int32(0) <= v23)
	if int32(0) <= v23 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v187 = int32(-8)
	goto L69
L68:
	;
	v187 = int32(-6)
	goto L69
L69:
	;
	v190 = int32(base.Ui32(int32(base.Ui32(v180)>>(uint(int32(2))%32))+v187) >> (uint(int32(1)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v190
	if int32(0) <= v23 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v192 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	v202 = v192
	goto L72
L71:
	;
	v202 = v22<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v22&int32(63)
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v202
	v204 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v204
	v207 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v207
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v207
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v207
	v216 = base.B2i32(v23 < v204)
	if v23 < v204 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v217 = int32(6)
	goto L75
L74:
	;
	v217 = int32(8)
	goto L75
L75:
	;
	v218 = l1 + v217
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v218
	if v23 < v204 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v226 = int32(base.Ui32(v22)>>(uint(int32(7))%32)) & int32(63)
	goto L78
L77:
	;
	v226 = v22 & int32(_a_F_numeric_div_safe_11)
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v226
	v233 = v22 & int32(_a_F_numeric_div_safe_4)
	if v233 == int32(_a_F_numeric_div_safe_7) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v236 = v22 << (uint(int32(1)) % 32) & int32(_a_F_numeric_div_safe_8)
	goto L81
L80:
	;
	v236 = v233
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v236
	if v36 != 0 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	if v190 == int32(0) {
		goto L2
	} else {
		goto L90
	}
L83:
	;
	v239 = int32(0)
	goto L86
L84:
	;
	goto L85
L85:
	;
	v273 = int32(0)
	v281 = v273
	v288 = v273
	goto L82
L86:
	;
	v256 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v178+v239<<(uint(int32(1))%32)))))
	if v256 != 0 {
		v281 = v256
		v288 = v239 - v151
		goto L82
	} else {
		goto L88
	}
L87:
	;
	goto L85
L88:
	;
	v258 = v239 + int32(1)
	if v258 != v36 {
		v239 = v258
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v294 = v204
	goto L92
L91:
	;
	v317 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v218))))
	if v317 == int32(0) {
		goto L2
	} else {
		goto L96
	}
L92:
	;
	v308 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v218+v294<<(uint(int32(1))%32)))))
	if v308 != 0 {
		v314 = v308
		v316 = v202 - v294
		goto L91
	} else {
		goto L94
	}
L93:
	;
	v312 = int32(0)
	v314 = v312
	v316 = v312
	goto L91
L94:
	;
	v310 = v294 + int32(1)
	if v310 != v190 {
		v294 = v310
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v325 = v16 + int32(8)
	v335 = (v316+v288+base.B2i32(base.I32_extend16_s(v281) <= base.I32_extend16_s(v314)))<<(uint(int32(2))%32) + int32(16)
	if v163 < v335 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v337 = v335
	goto L99
L98:
	;
	v337 = v163
	goto L99
L99:
	;
	if v226 < v337 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v339 = v337
	goto L102
L101:
	;
	v339 = v226
	goto L102
L102:
	;
	if int32(1000) <= v339 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v342 = int32(1000)
	goto L105
L104:
	;
	v342 = v339
	goto L105
L105:
	;
	v343 = int32(1)
	F_div_var(m, v16+int32(56), v16+int32(32), v325, v342, v343, v343)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L18
	} else {
		goto L106
	}
L106:
	;
	v347 = F_make_result_safe(m, v325, l2)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L18
	} else {
		goto L107
	}
L107:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	if v349 == int32(0) {
		v387 = v347
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_pfree(m, v349)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L18
	} else {
		goto L109
	}
L109:
	;
	v387 = v347
	goto L1
L110:
	;
	if v368 == int32(0) {
		v387 = v367
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errcode(m, int32(33816706))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L18
	} else {
		goto L112
	}
L112:
	;
	F_errmsg(m, int32(_a_F_numeric_div_safe_12), int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L18
	} else {
		goto L113
	}
L113:
	;
	F_errsave_finish(m, l2, int32(_a_F_numeric_div_safe_13), int32(3249), int32(_a_F_numeric_div_safe_14))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L18
	} else {
		goto L114
	}
L114:
	;
	v387 = v367
	goto L1
}
func F_numeric_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
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
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v219 int32
	_ = v219
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v275 int32
	_ = v275
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int64
	_ = v323
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int64
	_ = v344
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v374 int64
	_ = v374
	var v375 int64
	_ = v375
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v412 int64
	_ = v412
	var v416 int32
	_ = v416
	var v418 int64
	_ = v418
	var v421 int64
	_ = v421
	var v424 int32
	_ = v424
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v456 int64
	_ = v456
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v476 int64
	_ = v476
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v494 int64
	_ = v494
	var v497 int32
	_ = v497
	var v499 int64
	_ = v499
	var v502 int64
	_ = v502
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v552 int64
	_ = v552
	var v553 int64
	_ = v553
	var v555 int64
	_ = v555
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v581 int64
	_ = v581
	var v582 int64
	_ = v582
	var v586 int64
	_ = v586
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v616 int64
	_ = v616
	var v617 int64
	_ = v617
	var v620 int32
	_ = v620
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v663 int64
	_ = v663
	var v667 int32
	_ = v667
	var v669 int64
	_ = v669
	var v672 int64
	_ = v672
	var v675 int32
	_ = v675
	var v682 int32
	_ = v682
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v707 int64
	_ = v707
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v732 int64
	_ = v732
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v750 int64
	_ = v750
	var v753 int32
	_ = v753
	var v755 int64
	_ = v755
	var v758 int64
	_ = v758
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v808 int64
	_ = v808
	var v809 int64
	_ = v809
	var v812 int32
	_ = v812
	var v817 int64
	_ = v817
	var v843 int64
	_ = v843
	var v844 int64
	_ = v844
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v855 int32
	_ = v855
	var v862 int64
	_ = v862
	var v863 int64
	_ = v863
	var v867 int64
	_ = v867
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v885 int32
	_ = v885
	var v890 int32
	_ = v890
	var v897 int64
	_ = v897
	var v898 int64
	_ = v898
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v935 int64
	_ = v935
	var v939 int32
	_ = v939
	var v941 int64
	_ = v941
	var v944 int64
	_ = v944
	var v947 int32
	_ = v947
	var v954 int32
	_ = v954
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v979 int64
	_ = v979
	var v983 int32
	_ = v983
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v999 int64
	_ = v999
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1017 int64
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1022 int64
	_ = v1022
	var v1025 int64
	_ = v1025
	var v1028 int32
	_ = v1028
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1051 int32
	_ = v1051
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1063 int32
	_ = v1063
	var v1075 int64
	_ = v1075
	var v1076 int64
	_ = v1076
	var v1078 int64
	_ = v1078
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1097 int32
	_ = v1097
	var v1104 int64
	_ = v1104
	var v1105 int64
	_ = v1105
	var v1115 int32
	_ = v1115
	var v1122 int64
	_ = v1122
	var v1123 int64
	_ = v1123
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1137 int32
	_ = v1137
	var v1147 int64
	_ = v1147
	var v1153 int64
	_ = v1153
	var v1157 int32
	_ = v1157
	var v1160 int32
	_ = v1160
	var v1169 int64
	_ = v1169
	var v1173 int32
	_ = v1173
	var v1175 int64
	_ = v1175
	var v1178 int64
	_ = v1178
	var v1181 int32
	_ = v1181
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1204 int32
	_ = v1204
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1229 int64
	_ = v1229
	var v1233 int32
	_ = v1233
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1257 int32
	_ = v1257
	var v1265 int32
	_ = v1265
	var v1267 int64
	_ = v1267
	var v1272 int32
	_ = v1272
	var v1275 int32
	_ = v1275
	var v1285 int64
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1290 int64
	_ = v1290
	var v1293 int64
	_ = v1293
	var v1296 int32
	_ = v1296
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1319 int32
	_ = v1319
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1327 int32
	_ = v1327
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1334 int32
	_ = v1334
	var v1350 int32
	_ = v1350
	var v1365 int32
	_ = v1365
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1387 int32
	_ = v1387
	var v1404 int64
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1411 int32
	_ = v1411
	var v1417 int32
	_ = v1417
	var v1422 int32
	_ = v1422
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1445 int32
	_ = v1445
	var v1449 int32
	_ = v1449
	var v1467 int32
	_ = v1467
	var v1470 int32
	_ = v1470
	var v1473 int32
	_ = v1473
	var v1487 int32
	_ = v1487
	var v1493 int32
	_ = v1493
	var v1509 int32
	_ = v1509
	var v1520 int64
	_ = v1520
	v2 = int32(0)
	v14 = int64(0)
	v17 = m.G0
	v19 = v17 - int32(96)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v25 = v23
	goto L4
L1:
	;
	v65 = int32(255)
	v66 = v61 & v65
	if base.B2i32(v66 == int32(46))|base.B2i32(base.Ui32((v61-int32(48))&v65) < base.Ui32(int32(10))) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v57 = v25 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v57
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	v60 = v54
	v61 = v59
	v62 = v57
	v63 = v55
	goto L1
L3:
	;
	v54 = int32(_a_F_numeric_in_0)
	v55 = int32(_a_F_numeric_in_1)
	goto L2
L4:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if base.Ui32(v40-int32(9)) < base.Ui32(int32(5)) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v25
	v60 = v45
	v61 = v40
	v62 = v25
	v63 = v2
	goto L1
L6:
	;
	goto L5
L7:
	;
	v25 = v25 + int32(1)
	goto L4
L8:
	;
	v45 = int32(_a_F_numeric_in_2)
	switch v40 - int32(32) {
	case 0:
		goto L7
	default:
		goto L6
	case 11:
		v54 = v45
		v55 = v2
		goto L2
	case 13:
		goto L3
	}
L9:
	;
	m.G0 = v1509 + int32(96)
	return v1520
L10:
	;
	v1509 = v19
	v1520 = base.I64_extend_i32_u(v1493)
	goto L9
L11:
	;
	v1487 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v1487)
	v1509 = v1473
	v1520 = int64(0)
	goto L9
L12:
	;
	F_errsave_finish(m, v21, int32(_a_F_numeric_in_3), v1467, int32(_a_F_numeric_in_4))
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L38
	} else {
		goto L283
	}
L13:
	;
	v1439 = F_errsave_start(m, v21)
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L38
	} else {
		goto L279
	}
L14:
	;
	v1404 = int64(0)
	v1405 = F_errsave_start(m, v21)
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L38
	} else {
		goto L274
	}
L15:
	;
	v82 = v25
	v83 = int32(_a_F_numeric_in_5)
	v84 = int32(3)
	goto L20
L16:
	;
	goto L17
L17:
	;
	v323 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+56)) = v323
	*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = v323
	*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = v323
	if v66 != int32(48) {
		goto L102
	} else {
		goto L103
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v256
	v260 = v256
	goto L78
L19:
	;
	if v129 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L20:
	;
	if v84 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v129 = int32(0)
	goto L19
L22:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v87 == v88 {
		v110 = v87
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	v112 = int32(1)
	if v110 != 0 {
		v82 = v82 + v112
		v83 = v83 + v112
		v84 = v84 - v112
		goto L20
	} else {
		goto L34
	}
L26:
	;
	if base.Ui32((v87-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v98 = v87 | int32(32)
	goto L29
L28:
	;
	v98 = v87
	goto L29
L29:
	;
	if base.Ui32((v88-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v107 = v88 | int32(32)
	goto L32
L31:
	;
	v107 = v88
	goto L32
L32:
	;
	if v98 == v107 {
		v110 = v98
		goto L25
	} else {
		goto L33
	}
L33:
	;
	v129 = v98 - v107
	goto L19
L34:
	;
	goto L24
L35:
	;
	v136 = F_make_result_safe(m, int32(_a_F_numeric_in_6), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v144 = v62
	v145 = int32(_a_F_numeric_in_7)
	v146 = int32(8)
	goto L41
L38:
	;
	return int64(0)
L39:
	;
	v256 = v25 + int32(3)
	v257 = v136
	goto L18
L40:
	;
	if v191 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L41:
	;
	if v146 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v191 = int32(0)
	goto L40
L43:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	if v149 == v150 {
		v172 = v149
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	goto L42
L46:
	;
	v174 = int32(1)
	if v172 != 0 {
		v144 = v144 + v174
		v145 = v145 + v174
		v146 = v146 - v174
		goto L41
	} else {
		goto L55
	}
L47:
	;
	if base.Ui32((v149-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v160 = v149 | int32(32)
	goto L50
L49:
	;
	v160 = v149
	goto L50
L50:
	;
	if base.Ui32((v150-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v169 = v150 | int32(32)
	goto L53
L52:
	;
	v169 = v150
	goto L53
L53:
	;
	if v160 == v169 {
		v172 = v160
		goto L46
	} else {
		goto L54
	}
L54:
	;
	v191 = v160 - v169
	goto L40
L55:
	;
	goto L45
L56:
	;
	v197 = F_make_result_safe(m, v60, int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L38
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v203 = v62
	v204 = int32(_a_F_numeric_in_8)
	v205 = int32(3)
	goto L61
L59:
	;
	v256 = v62 + int32(8)
	v257 = v197
	goto L18
L60:
	;
	if v250 != 0 {
		goto L14
	} else {
		goto L76
	}
L61:
	;
	if v205 != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v250 = int32(0)
	goto L60
L63:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	if v208 == v209 {
		v231 = v208
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	goto L62
L66:
	;
	v233 = int32(1)
	if v231 != 0 {
		v203 = v203 + v233
		v204 = v204 + v233
		v205 = v205 - v233
		goto L61
	} else {
		goto L75
	}
L67:
	;
	if base.Ui32((v208-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v219 = v208 | int32(32)
	goto L70
L69:
	;
	v219 = v208
	goto L70
L70:
	;
	if base.Ui32((v209-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v228 = v209 | int32(32)
	goto L73
L72:
	;
	v228 = v209
	goto L73
L73:
	;
	if v219 == v228 {
		v231 = v219
		goto L66
	} else {
		goto L74
	}
L74:
	;
	v250 = v219 - v228
	goto L60
L75:
	;
	goto L65
L76:
	;
	v254 = F_make_result_safe(m, v60, int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L38
	} else {
		goto L77
	}
L77:
	;
	v256 = v62 + int32(3)
	v257 = v254
	goto L18
L78:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260))))
	if base.B2i32(base.Ui32(v275-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v275 == int32(32)) != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v260 = v260 + int32(1)
	goto L78
L81:
	;
	if v275 != 0 {
		goto L14
	} else {
		goto L83
	}
L83:
	;
	if v22 < int32(4) {
		v1493 = v257
		goto L10
	} else {
		goto L84
	}
L84:
	;
	v287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v257)+4)))
	if v287 == int32(_a_F_numeric_in_9) {
		v1493 = v257
		goto L10
	} else {
		goto L85
	}
L85:
	;
	v290 = F_errsave_start(m, v21)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L38
	} else {
		goto L86
	}
L86:
	;
	if v290 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L38
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v321 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v321)
	v1509 = v19
	v1520 = v14
	goto L9
L90:
	;
	F_errmsg(m, int32(_a_F_numeric_in_10), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L38
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = int32(base.Ui32(v22-int32(4)) >> (uint(int32(16)) % 32))
	v304 = int32(21)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = (v22<<(uint(v304)%32) - int32(_a_F_numeric_in_11)) >> (uint(v304) % 32)
	v314 = F_errdetail(m, int32(_a_F_numeric_in_12), v19+int32(32))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L38
	} else {
		goto L92
	}
L92:
	;
	F_errsave_finish(m, v21, int32(_a_F_numeric_in_3), int32(_a_F_numeric_in_13), int32(_a_F_numeric_in_14))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L38
	} else {
		goto L93
	}
L93:
	;
	goto L89
L94:
	;
	v1350 = v1334
	goto L263
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+76)) = v1303
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v1301
	v1319 = v19 + int32(40)
	F_add_var(m, v1319, v19+int32(72), v1319)
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L38
	} else {
		goto L257
	}
L96:
	;
	v1272 = int32(0)
	v1275 = v1219 + int32(12)
	v1285 = v1267
	goto L254
L97:
	;
	v1251 = F_errsave_start(m, v21)
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L38
	} else {
		goto L250
	}
L98:
	;
	if v1115 == int32(2) {
		goto L97
	} else {
		goto L225
	}
L99:
	;
	v867 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+88)) = v867
	*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = v867
	*(*int64)(unsafe.Add(mBase, uint32(v19)+72)) = v867
	*(*int64)(unsafe.Add(mBase, uint32(v19)+56)) = v867
	v875 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v875
	*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = v867
	v879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+2)))
	if v879 == v875 {
		goto L97
	} else {
		goto L188
	}
L100:
	;
	v586 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+88)) = v586
	*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = v586
	*(*int64)(unsafe.Add(mBase, uint32(v19)+72)) = v586
	*(*int64)(unsafe.Add(mBase, uint32(v19)+56)) = v586
	v594 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v594
	*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = v586
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+2)))
	if v598 == v594 {
		goto L97
	} else {
		goto L143
	}
L101:
	;
	v344 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+88)) = v344
	*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = v344
	*(*int64)(unsafe.Add(mBase, uint32(v19)+72)) = v344
	*(*int64)(unsafe.Add(mBase, uint32(v19)+56)) = v344
	v352 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v352
	*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = v344
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+2)))
	if v356 == v352 {
		goto L97
	} else {
		goto L106
	}
L102:
	;
	v338 = F_set_var_from_str(m, v23, v62, v19+int32(40), v19+int32(68), v21)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L38
	} else {
		goto L104
	}
L103:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	switch v331 - int32(66) {
	case 0, 32:
		goto L101
	default:
		goto L102
	case 13, 45:
		goto L99
	case 22, 54:
		goto L100
	}
L104:
	;
	if v338 == int32(0) {
		v1473 = v19
		goto L11
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v63
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v19)+68))
	v1334 = v343
	goto L94
L106:
	;
	v362 = v356
	v367 = int32(2)
	v374 = int64(1)
	v375 = v14
	goto L107
L107:
	;
	if v362&int32(254) == int32(48) {
		goto L113
	} else {
		goto L114
	}
L108:
	;
	v1115 = v574
	v1122 = v581
	v1123 = v582
	goto L98
L109:
	;
	if v569&int32(255) != 0 {
		v362 = v569
		v367 = v574
		v374 = v581
		v375 = v582
		goto L107
	} else {
		goto L142
	}
L110:
	;
	v555 = int64(1)
	v565 = v367 + int32(1)
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v565))))
	v569 = v567
	v574 = v565
	v581 = v552 << (uint(v555) % 64)
	v582 = base.I64_extend8_s(base.I64_extend_i32_u(v540)) + v553<<(uint(v555)%64) - int64(48)
	goto L109
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+76)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v510
	v528 = v19 + int32(40)
	F_add_var(m, v528, v19+int32(72), v528)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L38
	} else {
		goto L140
	}
L112:
	;
	v481 = int32(0)
	v484 = v446 + int32(12)
	v494 = v476
	goto L137
L113:
	;
	if v374 < int64(4611686018427387904) {
		v540 = v362
		v552 = v374
		v553 = v375
		goto L110
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	if v362&int32(255) != int32(95) {
		v1115 = v367
		v1122 = v374
		v1123 = v375
		goto L98
	} else {
		goto L135
	}
L116:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	if v383 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	F_pfree(m, v383)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L38
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v388 = F_palloc(m, int32(12))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L38
	} else {
		goto L121
	}
L120:
	;
	goto L119
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = v388
	v391 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v388))) = uint16(v391)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = int64(0)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	v400 = v391
	v403 = v396 + int32(12)
	v412 = v374
	goto L122
L122:
	;
	v416 = v403 - int32(2)
	v418 = base.I64_div_u_s(v412, int64(10000))
	v421 = v418*int64(55536) + v412
	*(*uint16)(unsafe.Add(mBase, uint32(v416))) = uint16(v421)
	v424 = v400 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v412) {
		v400 = v424
		v403 = v416
		v412 = v418
		goto L122
	} else {
		goto L124
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+76)) = v400
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v424
	*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v416
	v431 = v19 + int32(40)
	F_mul_var(m, v431, v19+int32(72), v431, int32(0))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L38
	} else {
		goto L125
	}
L124:
	;
	goto L123
L125:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	if v437 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	F_pfree(m, v437)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L38
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v441 = F_palloc(m, int32(12))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L38
	} else {
		goto L130
	}
L129:
	;
	goto L128
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = v441
	v444 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v441))) = uint16(v444)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v446 + int32(2)
	if v375 < int64(0) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = int64(16384)
	v476 = int64(0) - v375
	goto L112
L132:
	;
	goto L133
L133:
	;
	v456 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = v456
	if v375 != v456 {
		v476 = v375
		goto L112
	} else {
		goto L134
	}
L134:
	;
	v460 = int32(0)
	v510 = v460
	v512 = v460
	goto L111
L135:
	;
	v467 = v367 + int32(1)
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v467))))
	if base.Ui32((v469-int32(50))&int32(255)) < base.Ui32(int32(254)) {
		goto L97
	} else {
		goto L136
	}
L136:
	;
	v569 = v469
	v574 = v467
	v581 = v374
	v582 = v375
	goto L109
L137:
	;
	v497 = v484 - int32(2)
	v499 = base.I64_div_u_s(v494, int64(10000))
	v502 = v499*int64(55536) + v494
	*(*uint16)(unsafe.Add(mBase, uint32(v497))) = uint16(v502)
	v505 = v481 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v494) {
		v481 = v505
		v484 = v497
		v494 = v499
		goto L137
	} else {
		goto L139
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v497
	v510 = v505
	v512 = v481
	goto L111
L139:
	;
	goto L138
L140:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	if int32(_a_F_numeric_in_15) < v533 {
		goto L13
	} else {
		goto L141
	}
L141:
	;
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v367))))
	v540 = v536
	v552 = int64(1)
	v553 = int64(0)
	goto L110
L142:
	;
	goto L108
L143:
	;
	v604 = v598
	v609 = int32(2)
	v616 = int64(1)
	v617 = v14
	goto L144
L144:
	;
	v620 = v604 & int32(255)
	goto L150
L145:
	;
	v1115 = v855
	v1122 = v862
	v1123 = v863
	goto L98
L146:
	;
	if v850&int32(255) != 0 {
		v604 = v850
		v609 = v855
		v616 = v862
		v617 = v863
		goto L144
	} else {
		goto L187
	}
L147:
	;
	v812 = v609 + int32(1)
	v817 = base.I64_extend8_s(base.I64_extend_i32_u(v796))
	if base.Ui32((v796-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		v844 = v817 - int64(48)
		goto L181
	} else {
		goto L182
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+76)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v766
	v784 = v19 + int32(40)
	F_add_var(m, v784, v19+int32(72), v784)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L38
	} else {
		goto L179
	}
L149:
	;
	v737 = int32(0)
	v740 = v697 + int32(12)
	v750 = v732
	goto L176
L150:
	;
	if base.B2i32(base.Ui32(v620-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v620|int32(32)-int32(97)) < base.Ui32(int32(6))) != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	if v616 < int64(576460752303423488) {
		v796 = v604
		v808 = v616
		v809 = v617
		goto L147
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	if v620 != int32(95) {
		v1115 = v609
		v1122 = v616
		v1123 = v617
		goto L98
	} else {
		goto L173
	}
L154:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	if v634 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	F_pfree(m, v634)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L38
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v639 = F_palloc(m, int32(12))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L38
	} else {
		goto L159
	}
L158:
	;
	goto L157
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = v639
	v642 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v639))) = uint16(v642)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = int64(0)
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	v651 = v642
	v654 = v647 + int32(12)
	v663 = v616
	goto L160
L160:
	;
	v667 = v654 - int32(2)
	v669 = base.I64_div_u_s(v663, int64(10000))
	v672 = v669*int64(55536) + v663
	*(*uint16)(unsafe.Add(mBase, uint32(v667))) = uint16(v672)
	v675 = v651 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v663) {
		v651 = v675
		v654 = v667
		v663 = v669
		goto L160
	} else {
		goto L162
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+76)) = v651
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v675
	*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v667
	v682 = v19 + int32(40)
	F_mul_var(m, v682, v19+int32(72), v682, int32(0))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L38
	} else {
		goto L163
	}
L162:
	;
	goto L161
L163:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	if v688 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	F_pfree(m, v688)
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L38
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v692 = F_palloc(m, int32(12))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L38
	} else {
		goto L168
	}
L167:
	;
	goto L166
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = v692
	v695 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v692))) = uint16(v695)
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v697 + int32(2)
	if v617 < int64(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = int64(16384)
	v732 = int64(0) - v617
	goto L149
L170:
	;
	goto L171
L171:
	;
	v707 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = v707
	if v617 != v707 {
		v732 = v617
		goto L149
	} else {
		goto L172
	}
L172:
	;
	v711 = int32(0)
	v766 = v711
	v768 = v711
	goto L148
L173:
	;
	v716 = v609 + int32(1)
	v718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v716))))
	goto L174
L174:
	;
	if base.B2i32(base.Ui32(v718-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v718|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L97
	} else {
		goto L175
	}
L175:
	;
	v850 = v718
	v855 = v716
	v862 = v616
	v863 = v617
	goto L146
L176:
	;
	v753 = v740 - int32(2)
	v755 = base.I64_div_u_s(v750, int64(10000))
	v758 = v755*int64(55536) + v750
	*(*uint16)(unsafe.Add(mBase, uint32(v753))) = uint16(v758)
	v761 = v737 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v750) {
		v737 = v761
		v740 = v753
		v750 = v755
		goto L176
	} else {
		goto L178
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v753
	v766 = v761
	v768 = v737
	goto L148
L178:
	;
	goto L177
L179:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	if int32(_a_F_numeric_in_15) < v789 {
		goto L13
	} else {
		goto L180
	}
L180:
	;
	v792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v609))))
	v796 = v792
	v808 = int64(1)
	v809 = int64(0)
	goto L147
L181:
	;
	v848 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v812))))
	v850 = v848
	v855 = v812
	v862 = v808 << (uint(int64(4)) % 64)
	v863 = v844 + v809<<(uint(int64(4))%64)
	goto L146
L182:
	;
	if base.Ui32((v796-int32(97))&int32(255)) <= base.Ui32(int32(5)) {
		v844 = v817 - int64(87)
		goto L181
	} else {
		goto L183
	}
L183:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v796-int32(65))&int32(255)) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v843 = int64(-1)
	goto L186
L185:
	;
	v843 = v817 - int64(55)
	goto L186
L186:
	;
	v844 = v843
	goto L181
L187:
	;
	goto L145
L188:
	;
	v885 = v879
	v890 = int32(2)
	v897 = int64(1)
	v898 = v14
	goto L189
L189:
	;
	if v885&int32(248) == int32(48) {
		goto L195
	} else {
		goto L196
	}
L190:
	;
	v1115 = v1097
	v1122 = v1104
	v1123 = v1105
	goto L98
L191:
	;
	if v1092&int32(255) != 0 {
		v885 = v1092
		v890 = v1097
		v897 = v1104
		v898 = v1105
		goto L189
	} else {
		goto L224
	}
L192:
	;
	v1078 = int64(3)
	v1088 = v890 + int32(1)
	v1090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v1088))))
	v1092 = v1090
	v1097 = v1088
	v1104 = v1075 << (uint(v1078) % 64)
	v1105 = base.I64_extend8_s(base.I64_extend_i32_u(v1063)) + v1076<<(uint(v1078)%64) - int64(48)
	goto L191
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+76)) = v1035
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v1033
	v1051 = v19 + int32(40)
	F_add_var(m, v1051, v19+int32(72), v1051)
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L38
	} else {
		goto L222
	}
L194:
	;
	v1004 = int32(0)
	v1007 = v969 + int32(12)
	v1017 = v999
	goto L219
L195:
	;
	if v897 < int64(1152921504606846976) {
		v1063 = v885
		v1075 = v897
		v1076 = v898
		goto L192
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	if v885&int32(255) != int32(95) {
		v1115 = v890
		v1122 = v897
		v1123 = v898
		goto L98
	} else {
		goto L217
	}
L198:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	if v906 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	F_pfree(m, v906)
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L38
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v911 = F_palloc(m, int32(12))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L38
	} else {
		goto L203
	}
L202:
	;
	goto L201
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = v911
	v914 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v911))) = uint16(v914)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = int64(0)
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	v923 = v914
	v926 = v919 + int32(12)
	v935 = v897
	goto L204
L204:
	;
	v939 = v926 - int32(2)
	v941 = base.I64_div_u_s(v935, int64(10000))
	v944 = v941*int64(55536) + v935
	*(*uint16)(unsafe.Add(mBase, uint32(v939))) = uint16(v944)
	v947 = v923 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v935) {
		v923 = v947
		v926 = v939
		v935 = v941
		goto L204
	} else {
		goto L206
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+76)) = v923
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v947
	*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v939
	v954 = v19 + int32(40)
	F_mul_var(m, v954, v19+int32(72), v954, int32(0))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L38
	} else {
		goto L207
	}
L206:
	;
	goto L205
L207:
	;
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	if v960 != 0 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	F_pfree(m, v960)
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L38
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	v964 = F_palloc(m, int32(12))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L38
	} else {
		goto L212
	}
L211:
	;
	goto L210
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = v964
	v967 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v964))) = uint16(v967)
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v969 + int32(2)
	if v898 < int64(0) {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = int64(16384)
	v999 = int64(0) - v898
	goto L194
L214:
	;
	goto L215
L215:
	;
	v979 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = v979
	if v898 != v979 {
		v999 = v898
		goto L194
	} else {
		goto L216
	}
L216:
	;
	v983 = int32(0)
	v1033 = v983
	v1035 = v983
	goto L193
L217:
	;
	v990 = v890 + int32(1)
	v992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v990))))
	if base.Ui32((v992-int32(56))&int32(255)) < base.Ui32(int32(248)) {
		goto L97
	} else {
		goto L218
	}
L218:
	;
	v1092 = v992
	v1097 = v990
	v1104 = v897
	v1105 = v898
	goto L191
L219:
	;
	v1020 = v1007 - int32(2)
	v1022 = base.I64_div_u_s(v1017, int64(10000))
	v1025 = v1022*int64(55536) + v1017
	*(*uint16)(unsafe.Add(mBase, uint32(v1020))) = uint16(v1025)
	v1028 = v1004 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v1017) {
		v1004 = v1028
		v1007 = v1020
		v1017 = v1022
		goto L219
	} else {
		goto L221
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v1020
	v1033 = v1028
	v1035 = v1004
	goto L193
L221:
	;
	goto L220
L222:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	if int32(_a_F_numeric_in_15) < v1056 {
		goto L13
	} else {
		goto L223
	}
L223:
	;
	v1059 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v890))))
	v1063 = v1059
	v1075 = int64(1)
	v1076 = int64(0)
	goto L192
L224:
	;
	goto L190
L225:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	if v1127 != 0 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	F_pfree(m, v1127)
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L38
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	v1131 = F_palloc(m, int32(12))
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L38
	} else {
		goto L230
	}
L229:
	;
	goto L228
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = v1131
	v1134 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1131))) = uint16(v1134)
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v1137 + int32(2)
	if v1122 < int64(0) {
		goto L233
	} else {
		goto L234
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+76)) = v1188
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v1186
	v1204 = v19 + int32(40)
	F_mul_var(m, v1204, v19+int32(72), v1204, int32(0))
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L38
	} else {
		goto L240
	}
L232:
	;
	v1157 = v1134
	v1160 = v1137 + int32(12)
	v1169 = v1153
	goto L237
L233:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = int64(16384)
	v1153 = int64(0) - v1122
	goto L232
L234:
	;
	goto L235
L235:
	;
	v1147 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = v1147
	if v1122 == v1147 {
		v1186 = v1134
		v1188 = int32(0)
		goto L231
	} else {
		goto L236
	}
L236:
	;
	v1153 = v1122
	goto L232
L237:
	;
	v1173 = v1160 - int32(2)
	v1175 = base.I64_div_u_s(v1169, int64(10000))
	v1178 = v1175*int64(55536) + v1169
	*(*uint16)(unsafe.Add(mBase, uint32(v1173))) = uint16(v1178)
	v1181 = v1157 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v1169) {
		v1157 = v1181
		v1160 = v1173
		v1169 = v1175
		goto L237
	} else {
		goto L239
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v1173
	v1186 = v1181
	v1188 = v1157
	goto L231
L239:
	;
	goto L238
L240:
	;
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	if v1210 != 0 {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	F_pfree(m, v1210)
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L38
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	v1214 = F_palloc(m, int32(12))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L38
	} else {
		goto L245
	}
L244:
	;
	goto L243
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = v1214
	v1217 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1214))) = uint16(v1217)
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v1219 + int32(2)
	if v1123 < int64(0) {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = int64(16384)
	v1267 = int64(0) - v1123
	goto L96
L247:
	;
	goto L248
L248:
	;
	v1229 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = v1229
	if v1123 != v1229 {
		v1267 = v1123
		goto L96
	} else {
		goto L249
	}
L249:
	;
	v1233 = int32(0)
	v1301 = v1233
	v1303 = v1233
	goto L95
L250:
	;
	if v1251 == int32(0) {
		v1473 = v19
		goto L11
	} else {
		goto L251
	}
L251:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1257 = m.ExcPending
	if v1257 != 0 {
		goto L38
	} else {
		goto L252
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = int32(_a_F_numeric_in_16)
	F_errmsg(m, int32(_a_F_numeric_in_17), v19+int32(16))
	mBase = m.M
	v1265 = m.ExcPending
	if v1265 != 0 {
		goto L38
	} else {
		goto L253
	}
L253:
	;
	v1467 = int32(_a_F_numeric_in_18)
	goto L12
L254:
	;
	v1288 = v1275 - int32(2)
	v1290 = base.I64_div_u_s(v1285, int64(10000))
	v1293 = v1290*int64(55536) + v1285
	*(*uint16)(unsafe.Add(mBase, uint32(v1288))) = uint16(v1293)
	v1296 = v1272 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v1285) {
		v1272 = v1296
		v1275 = v1288
		v1285 = v1290
		goto L254
	} else {
		goto L256
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v1288
	v1301 = v1296
	v1303 = v1272
	goto L95
L256:
	;
	goto L255
L257:
	;
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	if int32(_a_F_numeric_in_15) < v1324 {
		goto L13
	} else {
		goto L258
	}
L258:
	;
	v1327 = v62 + v1115
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v63
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	if v1329 != 0 {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	F_pfree(m, v1329)
	mBase = m.M
	v1331 = m.ExcPending
	if v1331 != 0 {
		goto L38
	} else {
		goto L262
	}
L260:
	;
	goto L261
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v1327
	v1334 = v1327
	goto L94
L262:
	;
	goto L261
L263:
	;
	v1365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1350))))
	if base.B2i32(base.Ui32(v1365-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v1365 == int32(32)) != 0 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v1350 = v1350 + int32(1)
	goto L263
L266:
	;
	if v1365 != 0 {
		goto L14
	} else {
		goto L268
	}
L268:
	;
	v1376 = v19 + int32(40)
	v1377 = F_apply_typmod(m, v1376, v22, v21)
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L38
	} else {
		goto L269
	}
L269:
	;
	if v1377 == int32(0) {
		v1473 = v19
		goto L11
	} else {
		goto L270
	}
L270:
	;
	v1381 = F_make_result_safe(m, v1376, v21)
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L38
	} else {
		goto L271
	}
L271:
	;
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(v19)+56))
	if v1383 == int32(0) {
		v1493 = v1381
		goto L10
	} else {
		goto L272
	}
L272:
	;
	F_pfree(m, v1383)
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L38
	} else {
		goto L273
	}
L273:
	;
	v1493 = v1381
	goto L10
L274:
	;
	if v1405 == int32(0) {
		v1509 = v19
		v1520 = v1404
		goto L9
	} else {
		goto L275
	}
L275:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1411 = m.ExcPending
	if v1411 != 0 {
		goto L38
	} else {
		goto L276
	}
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(_a_F_numeric_in_16)
	F_errmsg(m, int32(_a_F_numeric_in_17), v19)
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L38
	} else {
		goto L277
	}
L277:
	;
	F_errsave_finish(m, v21, int32(_a_F_numeric_in_3), int32(789), int32(_a_F_numeric_in_19))
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L38
	} else {
		goto L278
	}
L278:
	;
	v1509 = v19
	v1520 = v1404
	goto L9
L279:
	;
	if v1439 == int32(0) {
		v1473 = v19
		goto L11
	} else {
		goto L280
	}
L280:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v1445 = m.ExcPending
	if v1445 != 0 {
		goto L38
	} else {
		goto L281
	}
L281:
	;
	F_errmsg(m, int32(_a_F_numeric_in_20), int32(0))
	mBase = m.M
	v1449 = m.ExcPending
	if v1449 != 0 {
		goto L38
	} else {
		goto L282
	}
L282:
	;
	v1467 = int32(_a_F_numeric_in_21)
	goto L12
L283:
	;
	v1473 = v19
	goto L11
}
func F_numeric_int2(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int64
	_ = v134
	var v139 int64
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v161 int64
	_ = v161
	v7 = int64(0)
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
		v18 = base.I32_extend16_s(v17)
		if base.Ui32(int32(_a_F_numeric_int2_0)) <= base.Ui32(v17) {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v22 = F_errsave_start(m, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				if v18 == int32(-16384) {
					if v22 == int32(0) {
						v161 = v7
						m.G0 = v10 - int32(-64)
						return v161
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_numeric_int2_1)
							F_errmsg(m, int32(_a_F_numeric_int2_2), v10)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int64(0)
							} else {
								F_errsave_finish(m, v21, int32(_a_F_numeric_int2_3), int32(_a_F_numeric_int2_4), int32(_a_F_numeric_int2_5))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int64(0)
								} else {
									v161 = v7
									m.G0 = v10 - int32(-64)
									return v161
								}
							}
						}
					}
				} else {
					if v22 == int32(0) {
						v161 = v7
						m.G0 = v10 - int32(-64)
						return v161
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_numeric_int2_1)
							F_errmsg(m, int32(_a_F_numeric_int2_6), v8+int32(-48))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int64(0)
							} else {
								F_errsave_finish(m, v21, int32(_a_F_numeric_int2_3), int32(_a_F_numeric_int2_7), int32(_a_F_numeric_int2_5))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int64(0)
								} else {
									v161 = v7
									m.G0 = v10 - int32(-64)
									return v161
								}
							}
						}
					}
				}
			}
		} else {
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			v64 = base.B2i32(int32(0) <= v18)
			if int32(0) <= v18 {
				v65 = int32(-8)
			} else {
				v65 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = int32(base.Ui32(int32(base.Ui32(v58)>>(uint(int32(2))%32))+v65) >> (uint(int32(1)) % 32))
			if int32(0) <= v18 {
				v70 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+6)))
				v80 = v70
			} else {
				v80 = v17<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v17&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v80
			v82 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v82
			v91 = base.B2i32(v18 < v82)
			if v18 < v82 {
				v92 = int32(base.Ui32(v17)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v92 = v17 & int32(_a_F_numeric_int2_8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v92
			v99 = v17 & int32(_a_F_numeric_int2_0)
			if v99 == int32(_a_F_numeric_int2_9) {
				v102 = v17 << (uint(int32(1)) % 32) & int32(_a_F_numeric_int2_10)
			} else {
				v102 = v99
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v102
			if v18 < v82 {
				v106 = int32(6)
			} else {
				v106 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = v13 + v106
			v113 = F_numericvar_to_int64(m, v8+int32(-24), v8+int32(-32))
			mBase = m.M
			v114 = m.ExcPending
			if v114 != 0 {
				return int64(0)
			} else {
				if v113 == int32(0) {
					v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v118 = F_errsave_start(m, v117)
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return int64(0)
					} else {
						if v118 == int32(0) {
							v161 = v7
							m.G0 = v10 - int32(-64)
							return v161
						} else {
							F_errcode(m, int32(50331778))
							mBase = m.M
							v124 = m.ExcPending
							if v124 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_numeric_int2_11), int32(0))
								mBase = m.M
								v128 = m.ExcPending
								if v128 != 0 {
									return int64(0)
								} else {
									F_errsave_finish(m, v117, int32(_a_F_numeric_int2_3), int32(_a_F_numeric_int2_12), int32(_a_F_numeric_int2_5))
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return int64(0)
									} else {
										v161 = v7
										m.G0 = v10 - int32(-64)
										return v161
									}
								}
							}
						}
					}
				} else {
					v134 = *(*int64)(unsafe.Add(mBase, uint32(v10)+32))
					if base.Ui64(int64(-65537)) < base.Ui64(v134-int64(32768)) {
						v161 = v134
						m.G0 = v10 - int32(-64)
						return v161
					} else {
						v139 = int64(0)
						v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v141 = F_errsave_start(m, v140)
						mBase = m.M
						v142 = m.ExcPending
						if v142 != 0 {
							return int64(0)
						} else {
							if v141 == int32(0) {
								v161 = v139
								m.G0 = v10 - int32(-64)
								return v161
							} else {
								F_errcode(m, int32(50331778))
								mBase = m.M
								v147 = m.ExcPending
								if v147 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_numeric_int2_11), int32(0))
									mBase = m.M
									v151 = m.ExcPending
									if v151 != 0 {
										return int64(0)
									} else {
										F_errsave_finish(m, v140, int32(_a_F_numeric_int2_3), int32(_a_F_numeric_int2_13), int32(_a_F_numeric_int2_5))
										mBase = m.M
										v156 = m.ExcPending
										if v156 != 0 {
											return int64(0)
										} else {
											v161 = v139
											m.G0 = v10 - int32(-64)
											return v161
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_numeric_int8_safe(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int64
	_ = v129
	var v134 int64
	_ = v134
	v7 = int64(0)
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v13 = base.I32_extend16_s(v12)
	if base.Ui32(int32(_a_F_numeric_int8_safe_0)) <= base.Ui32(v12) {
		v16 = F_errsave_start(m, l1)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			if v13 == int32(-16384) {
				if v16 == int32(0) {
					v134 = v7
					m.G0 = v10 - int32(-64)
					return v134
				} else {
					F_errcode(m, int32(1088))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_numeric_int8_safe_1)
						F_errmsg(m, int32(_a_F_numeric_int8_safe_2), v10)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int64(0)
						} else {
							F_errsave_finish(m, l1, int32(_a_F_numeric_int8_safe_3), int32(_a_F_numeric_int8_safe_4), int32(_a_F_numeric_int8_safe_5))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int64(0)
							} else {
								v134 = v7
								m.G0 = v10 - int32(-64)
								return v134
							}
						}
					}
				}
			} else {
				if v16 == int32(0) {
					v134 = v7
					m.G0 = v10 - int32(-64)
					return v134
				} else {
					F_errcode(m, int32(1088))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_numeric_int8_safe_1)
						F_errmsg(m, int32(_a_F_numeric_int8_safe_6), v8+int32(-48))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int64(0)
						} else {
							F_errsave_finish(m, l1, int32(_a_F_numeric_int8_safe_3), int32(_a_F_numeric_int8_safe_7), int32(_a_F_numeric_int8_safe_5))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int64(0)
							} else {
								v134 = v7
								m.G0 = v10 - int32(-64)
								return v134
							}
						}
					}
				}
			}
		}
	} else {
		v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v60 = base.B2i32(int32(0) <= v13)
		if int32(0) <= v13 {
			v61 = int32(-8)
		} else {
			v61 = int32(-6)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = int32(base.Ui32(int32(base.Ui32(v54)>>(uint(int32(2))%32))+v61) >> (uint(int32(1)) % 32))
		if int32(0) <= v13 {
			v66 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
			v76 = v66
		} else {
			v76 = v12<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v12&int32(63)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v76
		v78 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v78
		v87 = base.B2i32(v13 < v78)
		if v13 < v78 {
			v88 = int32(base.Ui32(v12)>>(uint(int32(7))%32)) & int32(63)
		} else {
			v88 = v12 & int32(_a_F_numeric_int8_safe_8)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v88
		v95 = v12 & int32(_a_F_numeric_int8_safe_0)
		if v95 == int32(_a_F_numeric_int8_safe_9) {
			v98 = v12 << (uint(int32(1)) % 32) & int32(_a_F_numeric_int8_safe_10)
		} else {
			v98 = v95
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v98
		if v13 < v78 {
			v102 = int32(6)
		} else {
			v102 = int32(8)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = l0 + v102
		v109 = F_numericvar_to_int64(m, v8+int32(-24), v8+int32(-32))
		mBase = m.M
		v110 = m.ExcPending
		if v110 != 0 {
			return int64(0)
		} else {
			if v109 == int32(0) {
				v113 = F_errsave_start(m, l1)
				mBase = m.M
				v114 = m.ExcPending
				if v114 != 0 {
					return int64(0)
				} else {
					if v113 == int32(0) {
						v134 = v7
						m.G0 = v10 - int32(-64)
						return v134
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_numeric_int8_safe_11), int32(0))
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return int64(0)
							} else {
								F_errsave_finish(m, l1, int32(_a_F_numeric_int8_safe_3), int32(_a_F_numeric_int8_safe_12), int32(_a_F_numeric_int8_safe_5))
								mBase = m.M
								v128 = m.ExcPending
								if v128 != 0 {
									return int64(0)
								} else {
									v134 = v7
									m.G0 = v10 - int32(-64)
									return v134
								}
							}
						}
					}
				}
			} else {
				v129 = *(*int64)(unsafe.Add(mBase, uint32(v10)+32))
				v134 = v129
				m.G0 = v10 - int32(-64)
				return v134
			}
		}
	}
}
func F_numeric_larger(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v9 = F_pg_detoast_datum(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+4)))
			v23 = base.I32_extend16_s(v22)
			v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+4)))
			if base.Ui32(int32(_a_F_numeric_larger_0)) <= base.Ui32(v24) {
				if v24 != int32(_a_F_numeric_larger_1) {
					if v24 != int32(_a_F_numeric_larger_0) {
						if v23 != int32(-4096) {
							v43 = int32(-1)
						} else {
							v43 = int32(0)
						}
						v164 = v43
					} else {
						v164 = base.B2i32(v23 != int32(-16384))
					}
				} else {
					if v23 == int32(-16384) {
						v38 = int32(-1)
					} else {
						v38 = base.B2i32(v23 != int32(-12288))
					}
					v164 = v38
				}
			} else {
				if base.Ui32(int32(-16384)) <= base.Ui32(v23) {
					if v23 == int32(-4096) {
						v50 = int32(1)
					} else {
						v50 = int32(-1)
					}
					v164 = v50
				} else {
					v52 = v4 + int32(6)
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
					v58 = base.I32_extend16_s(v24)
					v60 = base.B2i32(int32(0) <= v58)
					if int32(0) <= v58 {
						v61 = int32(-8)
					} else {
						v61 = int32(-6)
					}
					if int32(0) <= v58 {
						v63 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52))))
						v73 = v63
					} else {
						v73 = v24<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v24&int32(63)
					}
					v75 = int32(base.Ui32(int32(base.Ui32(v53)>>(uint(int32(2))%32))+v61) >> (uint(int32(1)) % 32))
					v77 = v9 + int32(6)
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v84 = base.B2i32(int32(0) <= v23)
					if int32(0) <= v23 {
						v85 = int32(-8)
					} else {
						v85 = int32(-6)
					}
					if int32(0) <= v23 {
						v87 = int32(*(*int16)(unsafe.Add(mBase, uint32(v77))))
						v97 = v87
					} else {
						v97 = v22<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v22&int32(63)
					}
					v98 = int32(1)
					v99 = int32(base.Ui32(int32(base.Ui32(v78)>>(uint(int32(2))%32))+v85) >> (uint(v98) % 32))
					v105 = v22 & int32(_a_F_numeric_larger_0)
					if v105 == int32(_a_F_numeric_larger_2) {
						v108 = v22 << (uint(v98) % 32) & int32(_a_F_numeric_larger_3)
					} else {
						v108 = v105
					}
					if v75 == int32(0) {
						if v99 == int32(0) {
							v164 = int32(0)
						} else {
							if v108 == int32(_a_F_numeric_larger_3) {
								v118 = int32(1)
							} else {
								v118 = int32(-1)
							}
							v164 = v118
						}
					} else {
						v124 = v24 & int32(_a_F_numeric_larger_0)
						if v124 == int32(_a_F_numeric_larger_2) {
							v127 = v24 << (uint(int32(1)) % 32) & int32(_a_F_numeric_larger_3)
						} else {
							v127 = v124
						}
						if v99 == int32(0) {
							if v127 != 0 {
								v132 = int32(-1)
							} else {
								v132 = int32(1)
							}
							v164 = v132
						} else {
							if v58 < int32(0) {
								v137 = v52
							} else {
								v137 = v4 + int32(8)
							}
							if v23 < int32(0) {
								v142 = v77
							} else {
								v142 = v9 + int32(8)
							}
							if v127 == int32(0) {
								if v108 == int32(_a_F_numeric_larger_3) {
									v164 = int32(1)
								} else {
									v148 = F_cmp_abs_common(m, v137, v75, v73, v142, v99, v97)
									mBase = m.M
									v164 = v148
								}
							} else {
								if v108 == int32(0) {
									v164 = int32(-1)
								} else {
									v152 = F_cmp_abs_common(m, v142, v99, v97, v137, v75, v73)
									mBase = m.M
									v164 = v152
								}
							}
						}
					}
				}
			}
			if int32(0) < v164 {
				v167 = v4
			} else {
				v167 = v9
			}
			return base.I64_extend_i32_u(v167)
		}
	}
}
func F_numeric_min_scale(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+4)))
	if base.Ui32(int32(_a_F_numeric_min_scale_0)) <= base.Ui32(v11) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v14 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
	return int64(0)
L4:
	;
	goto L5
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v23 = base.I32_extend16_s(v11)
	v25 = base.B2i32(int32(0) <= v23)
	if int32(0) <= v23 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v26 = int32(-8)
	goto L8
L7:
	;
	v26 = int32(-6)
	goto L8
L8:
	;
	if int32(0) <= v23 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(v7)+6)))
	v38 = v28
	goto L11
L10:
	;
	v38 = v11<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v11&int32(63)
	goto L11
L11:
	;
	v41 = int32(0)
	if v23 < v41 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v46 = int32(6)
	goto L14
L13:
	;
	v46 = int32(8)
	goto L14
L14:
	;
	v48 = int32(base.Ui32(int32(base.Ui32(v18)>>(uint(int32(2))%32))+v26) >> (uint(int32(1)) % 32))
	goto L16
L15:
	;
	return base.I64_extend_i32_s(v90)
L16:
	;
	if v48 <= int32(0) {
		v90 = v41
		goto L15
	} else {
		goto L18
	}
L17:
	;
	v65 = (v56 - v38) << (uint(int32(2)) % 32)
	if v65 <= int32(0) {
		v90 = v41
		goto L15
	} else {
		goto L20
	}
L18:
	;
	v55 = int32(1)
	v56 = v48 - v55
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7+v46+v56<<(uint(v55)%32)))))
	if v60 == int32(0) {
		v48 = v56
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v70 = base.I32_rem_s(base.I32_extend16_s(v60), int32(10))
	if v70 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return base.I64_extend_i32_s(v65)
L22:
	;
	goto L23
L23:
	;
	v74 = v60
	v75 = v65
	goto L24
L24:
	;
	v79 = v75 - int32(1)
	v81 = int32(10)
	v82 = base.I32_div_s(base.I32_extend16_s(v74), v81)
	v85 = base.I32_rem_s(base.I32_extend16_s(v82), v81)
	if v85 == int32(0) {
		v74 = v82
		v75 = v79
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v90 = v79
	goto L15
L26:
	;
	goto L25
}
func F_numeric_poly_avg(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int64
	_ = v22
	var v26 int32
	_ = v26
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int64
	_ = v44
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int64
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v92 int64
	_ = v92
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v123 int64
	_ = v123
	var v124 int64
	_ = v124
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v131 int64
	_ = v131
	var v134 int64
	_ = v134
	var v137 int64
	_ = v137
	var v140 int64
	_ = v140
	var v141 int64
	_ = v141
	var v145 int64
	_ = v145
	var v152 int64
	_ = v152
	var v164 int32
	_ = v164
	var v165 int64
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int64
	_ = v174
	var v177 int64
	_ = v177
	var v179 int32
	_ = v179
	var v182 int64
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v203 int64
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v221 int64
	_ = v221
	var v225 int32
	_ = v225
	var v227 int64
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int64
	_ = v238
	var v245 int32
	_ = v245
	var v247 int64
	_ = v247
	var v250 int64
	_ = v250
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v290 int64
	_ = v290
	var v291 int32
	_ = v291
	var v305 int64
	_ = v305
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v18 != 0 {
		v26 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
		v305 = int64(0)
		m.G0 = v16 + int32(80)
		return v305
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v19 == int32(0) {
			v26 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
			v305 = int64(0)
			m.G0 = v16 + int32(80)
			return v305
		} else {
			v22 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
			if v22 != int64(0) {
				v29 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
				v30 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
				v32 = F_palloc(m, int32(22))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v32
					v37 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v32))) = uint16(v37)
					*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v37
					*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v32 + int32(2)
					v44 = int64(0)
					if v30 == v44 {
						v50 = base.B2i32(v29 != v44)
					} else {
						v50 = base.B2i32(v44 < v30)
					}
					v51 = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = (v50 - base.B2i32(v30 < v51)) & int32(_a_F_numeric_poly_avg_0)
					v57 = int32(0)
					if v29|v30 != v51 {
						v63 = v57
						v66 = v32 + int32(22)
						v70 = v29
						v71 = v30
						for {
							v76 = int32(16)
							v77 = v16 + v76
							v79 = int64(0)
							v83 = m.G0
							v85 = v83 - v76
							m.G0 = v85
							v87 = int64(63)
							v88 = v71 >> (uint(v87) % 64)
							v89 = v70 ^ v88
							v92 = v89 + int64(base.Ui64(v71)>>(uint(v87)%64))
							F___udivmodti4(m, v85, v92, base.I64_extend_i32_u(base.B2i32(base.Ui64(v92) < base.Ui64(v89)))+(v88^v71), int64(10000), base.I64_extend_i32_u(int32(0))+v79)
							mBase = m.M
							v108 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
							v109 = v88 ^ v79
							v110 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
							v111 = v109 ^ v110
							*(*int64)(unsafe.Add(mBase, uint32(v77))) = v111 - v109
							*(*int64)(unsafe.Add(mBase, uint32(v77)+8)) = v109 ^ v108 - v109 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v111) < base.Ui64(v109)))
							m.G0 = v85 + v76
							v123 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
							v124 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
							v125 = int64(4294957296)
							v126 = int64(0)
							v131 = int64(32)
							v134 = int64(base.Ui64(v123) >> (uint(v131) % 64))
							v137 = int64(4294967295)
							v140 = v123 & v137
							v141 = v125 * v140
							v145 = int64(base.Ui64(v141)>>(uint(v131)%64)) + v125*v134
							v152 = v140*v126 + v145&v137
							*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v123*v126 + v124*v125 + v126*v134 + int64(base.Ui64(v145)>>(uint(v131)%64)) + int64(base.Ui64(v152)>>(uint(v131)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v16))) = v141&v137 | v152<<(uint(v131)%64)
							v164 = v66 - int32(2)
							v165 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
							v167 = base.I32_wrap_i64(v165 + v70)
							v169 = v167 >> (uint(int32(31)) % 32)
							v171 = v167 ^ v169 - v169
							*(*uint16)(unsafe.Add(mBase, uint32(v164))) = uint16(v171)
							v174 = v70 + int64(9999)
							v177 = v71 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v174) < base.Ui64(v70)))
							v179 = v63 + int32(1)
							v182 = int64(0)
							if v177 == v182 {
								v186 = base.B2i32(base.Ui64(int64(19998)) < base.Ui64(v174))
							} else {
								v186 = base.B2i32(v177 != v182)
							}
							if v186 != 0 {
								v63 = v179
								v66 = v164
								v70 = v123
								v71 = v124
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v164
						v188 = v179
						v190 = v63
					} else {
						v188 = v57
						v190 = int32(0)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v190
					*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v188
					v203 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v16)+64)) = int64(0)
					v207 = F_palloc(m, int32(12))
					mBase = m.M
					v208 = m.ExcPending
					if v208 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v207
						v210 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v207))) = uint16(v210)
						*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v207 + int32(2)
						if v203 < int64(0) {
							*(*int64)(unsafe.Add(mBase, uint32(v16)+64)) = int64(16384)
							v227 = int64(0) - v203
							v231 = int32(0)
							v234 = v207 + int32(12)
							v238 = v227
							for {
								v245 = v234 - int32(2)
								v247 = base.I64_div_u_s(v238, int64(10000))
								v250 = v247*int64(55536) + v238
								*(*uint16)(unsafe.Add(mBase, uint32(v245))) = uint16(v250)
								v253 = v231 + int32(1)
								if base.Ui64(int64(9999)) < base.Ui64(v238) {
									v231 = v253
									v234 = v245
									v238 = v247
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v245
							v257 = v253
							v259 = v231
						} else {
							v221 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v16)+64)) = v221
							if v203 != v221 {
								v227 = v203
								v231 = int32(0)
								v234 = v207 + int32(12)
								v238 = v227
								for {
									v245 = v234 - int32(2)
									v247 = base.I64_div_u_s(v238, int64(10000))
									v250 = v247*int64(55536) + v238
									*(*uint16)(unsafe.Add(mBase, uint32(v245))) = uint16(v250)
									v253 = v231 + int32(1)
									if base.Ui64(int64(9999)) < base.Ui64(v238) {
										v231 = v253
										v234 = v245
										v238 = v247
										continue
									} else {
										break
									}
									break
								}
								*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v245
								v257 = v253
								v259 = v231
							} else {
								v225 = int32(0)
								v257 = v225
								v259 = v225
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v259
						*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v257
						v275 = F_make_result_safe(m, v16+int32(56), int32(0))
						mBase = m.M
						v276 = m.ExcPending
						if v276 != 0 {
							return int64(0)
						} else {
							F_pfree(m, v207)
							mBase = m.M
							v278 = m.ExcPending
							if v278 != 0 {
								return int64(0)
							} else {
								v282 = F_make_result_safe(m, v16+int32(32), int32(0))
								mBase = m.M
								v283 = m.ExcPending
								if v283 != 0 {
									return int64(0)
								} else {
									F_pfree(m, v32)
									mBase = m.M
									v285 = m.ExcPending
									if v285 != 0 {
										return int64(0)
									} else {
										v290 = F_DirectFunctionCall2Coll(m, int32(1391), int32(0), base.I64_extend_i32_u(v282), base.I64_extend_i32_u(v275))
										mBase = m.M
										v291 = m.ExcPending
										if v291 != 0 {
											return int64(0)
										} else {
											v305 = v290
											m.G0 = v16 + int32(80)
											return v305
										}
									}
								}
							}
						}
					}
				}
			} else {
				v26 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
				v305 = int64(0)
				m.G0 = v16 + int32(80)
				return v305
			}
		}
	}
}
func F_numeric_poly_stddev_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v42 int64
	_ = v42
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v91 int64
	_ = v91
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v123 int32
	_ = v123
	var v124 int64
	_ = v124
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v132 int64
	_ = v132
	var v135 int64
	_ = v135
	var v138 int64
	_ = v138
	var v141 int64
	_ = v141
	var v142 int64
	_ = v142
	var v146 int64
	_ = v146
	var v153 int64
	_ = v153
	var v165 int32
	_ = v165
	var v166 int64
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int64
	_ = v175
	var v178 int64
	_ = v178
	var v180 int32
	_ = v180
	var v183 int64
	_ = v183
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v211 int32
	_ = v211
	var v212 int64
	_ = v212
	var v213 int64
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v225 int64
	_ = v225
	var v231 int32
	_ = v231
	var v232 int64
	_ = v232
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int64
	_ = v256
	var v257 int64
	_ = v257
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int64
	_ = v265
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int64
	_ = v273
	var v274 int64
	_ = v274
	var v275 int64
	_ = v275
	var v278 int64
	_ = v278
	var v294 int64
	_ = v294
	var v295 int64
	_ = v295
	var v296 int64
	_ = v296
	var v297 int64
	_ = v297
	var v309 int64
	_ = v309
	var v310 int64
	_ = v310
	var v311 int64
	_ = v311
	var v312 int64
	_ = v312
	var v317 int64
	_ = v317
	var v320 int64
	_ = v320
	var v323 int64
	_ = v323
	var v326 int64
	_ = v326
	var v327 int64
	_ = v327
	var v331 int64
	_ = v331
	var v338 int64
	_ = v338
	var v350 int32
	_ = v350
	var v351 int64
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v360 int64
	_ = v360
	var v363 int64
	_ = v363
	var v365 int32
	_ = v365
	var v368 int64
	_ = v368
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	v5 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(208)
	m.G0 = v18
	base.MemoryFill(m, v18+int32(96), v5, int32(112))
	if l0 != 0 {
		v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v18)+104)) = v25
		v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
		v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v30 = F_palloc(m, int32(22))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = v30
			v35 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v30))) = uint16(v35)
			*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v35
			*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = v30 + int32(2)
			v42 = int64(0)
			if v28 == v42 {
				v48 = base.B2i32(v27 != v42)
			} else {
				v48 = base.B2i32(v42 < v28)
			}
			v49 = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = (v48 - base.B2i32(v28 < v49)) & int32(_a_F_numeric_poly_stddev_internal_0)
			if v27|v28 != v49 {
				v66 = v30 + int32(22)
				v67 = v5
				v69 = v27
				v70 = v28
				for {
					v76 = v18 + int32(48)
					v78 = int64(0)
					v82 = m.G0
					v83 = int32(16)
					v84 = v82 - v83
					m.G0 = v84
					v86 = int64(63)
					v87 = v70 >> (uint(v86) % 64)
					v88 = v69 ^ v87
					v91 = v88 + int64(base.Ui64(v70)>>(uint(v86)%64))
					F___udivmodti4(m, v84, v91, base.I64_extend_i32_u(base.B2i32(base.Ui64(v91) < base.Ui64(v88)))+(v87^v70), int64(10000), base.I64_extend_i32_u(int32(0))+v78)
					mBase = m.M
					v107 = *(*int64)(unsafe.Add(mBase, uint32(v84)+8))
					v108 = v87 ^ v78
					v109 = *(*int64)(unsafe.Add(mBase, uint32(v84)))
					v110 = v108 ^ v109
					*(*int64)(unsafe.Add(mBase, uint32(v76))) = v110 - v108
					*(*int64)(unsafe.Add(mBase, uint32(v76)+8)) = v108 ^ v107 - v108 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v110) < base.Ui64(v108)))
					m.G0 = v84 + v83
					v123 = v18 + int32(32)
					v124 = *(*int64)(unsafe.Add(mBase, uint32(v18)+48))
					v125 = *(*int64)(unsafe.Add(mBase, uint32(v18)+56))
					v126 = int64(4294957296)
					v127 = int64(0)
					v132 = int64(32)
					v135 = int64(base.Ui64(v124) >> (uint(v132) % 64))
					v138 = int64(4294967295)
					v141 = v124 & v138
					v142 = v126 * v141
					v146 = int64(base.Ui64(v142)>>(uint(v132)%64)) + v126*v135
					v153 = v141*v127 + v146&v138
					*(*int64)(unsafe.Add(mBase, uint32(v123)+8)) = v124*v127 + v125*v126 + v127*v135 + int64(base.Ui64(v146)>>(uint(v132)%64)) + int64(base.Ui64(v153)>>(uint(v132)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v123))) = v142&v138 | v153<<(uint(v132)%64)
					v165 = v66 - int32(2)
					v166 = *(*int64)(unsafe.Add(mBase, uint32(v18)+32))
					v168 = base.I32_wrap_i64(v166 + v69)
					v170 = v168 >> (uint(int32(31)) % 32)
					v172 = v168 ^ v170 - v170
					*(*uint16)(unsafe.Add(mBase, uint32(v165))) = uint16(v172)
					v175 = v69 + int64(9999)
					v178 = v70 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v175) < base.Ui64(v69)))
					v180 = v67 + int32(1)
					v183 = int64(0)
					if v178 == v183 {
						v187 = base.B2i32(base.Ui64(int64(19998)) < base.Ui64(v175))
					} else {
						v187 = base.B2i32(v178 != v183)
					}
					if v187 != 0 {
						v66 = v165
						v67 = v180
						v69 = v124
						v70 = v125
						continue
					} else {
						break
					}
					break
				}
				*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = v165
				v194 = v67
				v196 = v180
			} else {
				v194 = v5
				v196 = v5
			}
			*(*int32)(unsafe.Add(mBase, uint32(v18)+76)) = v194
			*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v196
			F_accum_sum_add(m, v18+int32(112), v18+int32(72))
			mBase = m.M
			v211 = m.ExcPending
			if v211 != 0 {
				return int32(0)
			} else {
				v212 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
				v213 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
				F_pfree(m, v30)
				mBase = m.M
				v215 = m.ExcPending
				if v215 != 0 {
					return int32(0)
				} else {
					v217 = F_palloc(m, int32(22))
					mBase = m.M
					v218 = m.ExcPending
					if v218 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = v217
						v220 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v217))) = uint16(v220)
						*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v220
						v225 = int64(0)
						if v213 == v225 {
							v231 = base.B2i32(v212 != v225)
						} else {
							v231 = base.B2i32(v225 < v213)
						}
						v232 = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = (v231 - base.B2i32(v213 < v232)) & int32(_a_F_numeric_poly_stddev_internal_0)
						*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = v217 + int32(2)
						if v212|v213 != v232 {
							v253 = v217 + int32(22)
							v254 = v220
							v256 = v212
							v257 = v213
							for {
								v262 = int32(16)
								v263 = v18 + v262
								v265 = int64(0)
								v269 = m.G0
								v271 = v269 - v262
								m.G0 = v271
								v273 = int64(63)
								v274 = v257 >> (uint(v273) % 64)
								v275 = v256 ^ v274
								v278 = v275 + int64(base.Ui64(v257)>>(uint(v273)%64))
								F___udivmodti4(m, v271, v278, base.I64_extend_i32_u(base.B2i32(base.Ui64(v278) < base.Ui64(v275)))+(v274^v257), int64(10000), base.I64_extend_i32_u(int32(0))+v265)
								mBase = m.M
								v294 = *(*int64)(unsafe.Add(mBase, uint32(v271)+8))
								v295 = v274 ^ v265
								v296 = *(*int64)(unsafe.Add(mBase, uint32(v271)))
								v297 = v295 ^ v296
								*(*int64)(unsafe.Add(mBase, uint32(v263))) = v297 - v295
								*(*int64)(unsafe.Add(mBase, uint32(v263)+8)) = v295 ^ v294 - v295 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v297) < base.Ui64(v295)))
								m.G0 = v271 + v262
								v309 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
								v310 = *(*int64)(unsafe.Add(mBase, uint32(v18)+24))
								v311 = int64(4294957296)
								v312 = int64(0)
								v317 = int64(32)
								v320 = int64(base.Ui64(v309) >> (uint(v317) % 64))
								v323 = int64(4294967295)
								v326 = v309 & v323
								v327 = v311 * v326
								v331 = int64(base.Ui64(v327)>>(uint(v317)%64)) + v311*v320
								v338 = v326*v312 + v331&v323
								*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v309*v312 + v310*v311 + v312*v320 + int64(base.Ui64(v331)>>(uint(v317)%64)) + int64(base.Ui64(v338)>>(uint(v317)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v18))) = v327&v323 | v338<<(uint(v317)%64)
								v350 = v253 - int32(2)
								v351 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
								v353 = base.I32_wrap_i64(v351 + v256)
								v355 = v353 >> (uint(int32(31)) % 32)
								v357 = v353 ^ v355 - v355
								*(*uint16)(unsafe.Add(mBase, uint32(v350))) = uint16(v357)
								v360 = v256 + int64(9999)
								v363 = v257 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v360) < base.Ui64(v256)))
								v365 = v254 + int32(1)
								v368 = int64(0)
								if v363 == v368 {
									v372 = base.B2i32(base.Ui64(int64(19998)) < base.Ui64(v360))
								} else {
									v372 = base.B2i32(v363 != v368)
								}
								if v372 != 0 {
									v253 = v350
									v254 = v365
									v256 = v309
									v257 = v310
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = v350
							v379 = v254
							v381 = v365
						} else {
							v379 = int32(0)
							v381 = v220
						}
						*(*int32)(unsafe.Add(mBase, uint32(v18)+76)) = v379
						*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v381
						F_accum_sum_add(m, v18+int32(140), v18+int32(72))
						mBase = m.M
						v396 = m.ExcPending
						if v396 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v217)
							mBase = m.M
							v398 = m.ExcPending
							if v398 != 0 {
								return int32(0)
							} else {
								v416 = F_numeric_stddev_internal(m, v18+int32(96), l1, l2, l3)
								mBase = m.M
								v417 = m.ExcPending
								if v417 != 0 {
									return int32(0)
								} else {
									v418 = *(*int32)(unsafe.Add(mBase, uint32(v18)+112))
									if int32(0) < v418 {
										v421 = *(*int32)(unsafe.Add(mBase, uint32(v18)+132))
										F_pfree(m, v421)
										mBase = m.M
										v423 = m.ExcPending
										if v423 != 0 {
											return int32(0)
										} else {
											v424 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
											F_pfree(m, v424)
											mBase = m.M
											v426 = m.ExcPending
											if v426 != 0 {
												return int32(0)
											} else {
												v427 = *(*int32)(unsafe.Add(mBase, uint32(v18)+140))
												if int32(0) < v427 {
													v430 = *(*int32)(unsafe.Add(mBase, uint32(v18)+160))
													F_pfree(m, v430)
													mBase = m.M
													v432 = m.ExcPending
													if v432 != 0 {
														return int32(0)
													} else {
														v433 = *(*int32)(unsafe.Add(mBase, uint32(v18)+164))
														F_pfree(m, v433)
														mBase = m.M
														v435 = m.ExcPending
														if v435 != 0 {
															return int32(0)
														} else {
															m.G0 = v18 + int32(208)
															return v416
														}
													}
												} else {
													m.G0 = v18 + int32(208)
													return v416
												}
											}
										}
									} else {
										v427 = *(*int32)(unsafe.Add(mBase, uint32(v18)+140))
										if int32(0) < v427 {
											v430 = *(*int32)(unsafe.Add(mBase, uint32(v18)+160))
											F_pfree(m, v430)
											mBase = m.M
											v432 = m.ExcPending
											if v432 != 0 {
												return int32(0)
											} else {
												v433 = *(*int32)(unsafe.Add(mBase, uint32(v18)+164))
												F_pfree(m, v433)
												mBase = m.M
												v435 = m.ExcPending
												if v435 != 0 {
													return int32(0)
												} else {
													m.G0 = v18 + int32(208)
													return v416
												}
											}
										} else {
											m.G0 = v18 + int32(208)
											return v416
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v416 = F_numeric_stddev_internal(m, v18+int32(96), l1, l2, l3)
		mBase = m.M
		v417 = m.ExcPending
		if v417 != 0 {
			return int32(0)
		} else {
			v418 = *(*int32)(unsafe.Add(mBase, uint32(v18)+112))
			if int32(0) < v418 {
				v421 = *(*int32)(unsafe.Add(mBase, uint32(v18)+132))
				F_pfree(m, v421)
				mBase = m.M
				v423 = m.ExcPending
				if v423 != 0 {
					return int32(0)
				} else {
					v424 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
					F_pfree(m, v424)
					mBase = m.M
					v426 = m.ExcPending
					if v426 != 0 {
						return int32(0)
					} else {
						v427 = *(*int32)(unsafe.Add(mBase, uint32(v18)+140))
						if int32(0) < v427 {
							v430 = *(*int32)(unsafe.Add(mBase, uint32(v18)+160))
							F_pfree(m, v430)
							mBase = m.M
							v432 = m.ExcPending
							if v432 != 0 {
								return int32(0)
							} else {
								v433 = *(*int32)(unsafe.Add(mBase, uint32(v18)+164))
								F_pfree(m, v433)
								mBase = m.M
								v435 = m.ExcPending
								if v435 != 0 {
									return int32(0)
								} else {
									m.G0 = v18 + int32(208)
									return v416
								}
							}
						} else {
							m.G0 = v18 + int32(208)
							return v416
						}
					}
				}
			} else {
				v427 = *(*int32)(unsafe.Add(mBase, uint32(v18)+140))
				if int32(0) < v427 {
					v430 = *(*int32)(unsafe.Add(mBase, uint32(v18)+160))
					F_pfree(m, v430)
					mBase = m.M
					v432 = m.ExcPending
					if v432 != 0 {
						return int32(0)
					} else {
						v433 = *(*int32)(unsafe.Add(mBase, uint32(v18)+164))
						F_pfree(m, v433)
						mBase = m.M
						v435 = m.ExcPending
						if v435 != 0 {
							return int32(0)
						} else {
							m.G0 = v18 + int32(208)
							return v416
						}
					}
				} else {
					m.G0 = v18 + int32(208)
					return v416
				}
			}
		}
	}
}
func F_numeric_poly_stddev_pop(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(0)
	v4 = Fn14335(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_numeric_poly_stddev_samp(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14335(m, l0, int32(1), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_numeric_sum(m *base.Module, l0 int32) int64 {
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
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v21 int64
	_ = v21
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v11 != 0 {
		v28 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
		v82 = int32(0)
		m.G0 = v9 + int32(32)
		return base.I64_extend_i32_u(v82)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v12 == int32(0) {
			v28 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
			v82 = int32(0)
			m.G0 = v9 + int32(32)
			return base.I64_extend_i32_u(v82)
		} else {
			v15 = *(*int64)(unsafe.Add(mBase, uint32(v12)+96))
			v16 = *(*int64)(unsafe.Add(mBase, uint32(v12)+88))
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
			v21 = *(*int64)(unsafe.Add(mBase, uint32(v12)+104))
			if v15+(v16+v17) != int64(0)-v21 {
				if int64(0) < v16 {
					v35 = F_make_result_safe(m, int32(_a_F_numeric_sum_0), int32(0))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int64(0)
					} else {
						v82 = v35
						m.G0 = v9 + int32(32)
						return base.I64_extend_i32_u(v82)
					}
				} else {
					v39 = int64(0)
					if base.B2i32(v15 <= v39)|base.B2i32(v21 <= v39) == int32(0) {
						v48 = F_make_result_safe(m, int32(_a_F_numeric_sum_0), int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int64(0)
						} else {
							v82 = v48
							m.G0 = v9 + int32(32)
							return base.I64_extend_i32_u(v82)
						}
					} else {
						if int64(0) < v15 {
							v54 = F_make_result_safe(m, int32(_a_F_numeric_sum_1), int32(0))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int64(0)
							} else {
								v82 = v54
								m.G0 = v9 + int32(32)
								return base.I64_extend_i32_u(v82)
							}
						} else {
							if int64(0) < v21 {
								v60 = F_make_result_safe(m, int32(_a_F_numeric_sum_2), int32(0))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int64(0)
								} else {
									v82 = v60
									m.G0 = v9 + int32(32)
									return base.I64_extend_i32_u(v82)
								}
							} else {
								v62 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v62
								*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v62
								*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v62
								v71 = v9 + int32(8)
								F_accum_sum_final(m, v12+int32(16), v71)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int64(0)
								} else {
									v75 = F_make_result_safe(m, v71, int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int64(0)
									} else {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
										if v77 == int32(0) {
											v82 = v75
											m.G0 = v9 + int32(32)
											return base.I64_extend_i32_u(v82)
										} else {
											F_pfree(m, v77)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int64(0)
											} else {
												v82 = v75
												m.G0 = v9 + int32(32)
												return base.I64_extend_i32_u(v82)
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v28 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
				v82 = int32(0)
				m.G0 = v9 + int32(32)
				return base.I64_extend_i32_u(v82)
			}
		}
	}
}
func F_numeric_uplus(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		v12 = F_palloc(m, int32(base.Ui32(v9)>>(uint(int32(2))%32)))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			v16 = int32(base.Ui32(v14) >> (uint(int32(2)) % 32))
			if v16 != 0 {
				base.MemoryCopy(m, v12, v5, v16)
			} else {
			}
			return base.I64_extend_i32_u(v12)
		}
	}
}
func F_numeric_var_pop(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14336(m, l0, int32(0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_numeric_var_samp(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(1)
	v4 = Fn14336(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
