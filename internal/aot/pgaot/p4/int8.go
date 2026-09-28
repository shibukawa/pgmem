package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int8_avg(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v55 int64
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int64
	_ = v67
	var v71 int32
	_ = v71
	var v73 int64
	_ = v73
	var v76 int64
	_ = v76
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int64
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v112 int64
	_ = v112
	var v118 int64
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v128 int64
	_ = v128
	var v132 int32
	_ = v132
	var v134 int64
	_ = v134
	var v137 int64
	_ = v137
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int64
	_ = v166
	var v167 int32
	_ = v167
	var v176 int64
	_ = v176
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int64(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
		if v20 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v185 = m.ExcPending
			if v185 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_int8_avg_0), int32(0))
				mBase = m.M
				v189 = m.ExcPending
				if v189 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int8_avg_1), int32(_a_F_int8_avg_2), int32(_a_F_int8_avg_3))
					mBase = m.M
					v194 = m.ExcPending
					if v194 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			if v21&int32(-4) != int32(160) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v185 = m.ExcPending
				if v185 != 0 {
					return int64(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_int8_avg_0), int32(0))
					mBase = m.M
					v189 = m.ExcPending
					if v189 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_int8_avg_1), int32(_a_F_int8_avg_2), int32(_a_F_int8_avg_3))
						mBase = m.M
						v194 = m.ExcPending
						if v194 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
				v33 = v16 + (v26<<(uint(int32(3))%32)+int32(23))&int32(-8)
				v34 = *(*int64)(unsafe.Add(mBase, uint32(v33)))
				if v34 == int64(0) {
					v37 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v37)
					v176 = int64(0)
					m.G0 = v13 + int32(32)
					return v176
				} else {
					v40 = F_palloc(m, int32(12))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v40
						v43 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v40))) = uint16(v43)
						*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v43
						*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = base.I32_wrap_i64(int64(base.Ui64(v34)>>(uint(int64(49))%64))) & int32(_a_F_int8_avg_4)
						v55 = v34 >> (uint(int64(63)) % 64)
						v60 = v40 + int32(12)
						v62 = v43
						v67 = v34 ^ v55 - v55
						for {
							v71 = v60 - int32(2)
							v73 = base.I64_div_u_s(v67, int64(10000))
							v76 = v73*int64(55536) + v67
							*(*uint16)(unsafe.Add(mBase, uint32(v71))) = uint16(v76)
							v79 = v62 + int32(1)
							if base.Ui64(int64(9999)) < base.Ui64(v67) {
								v60 = v71
								v62 = v79
								v67 = v73
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v62
						*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v79
						*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v71
						v85 = int32(0)
						v89 = F_make_result_safe(m, v13+int32(8), v85)
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int64(0)
						} else {
							F_pfree(m, v40)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int64(0)
							} else {
								v93 = *(*int64)(unsafe.Add(mBase, uint32(v33)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = int64(0)
								v97 = F_palloc(m, int32(12))
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v97
									v100 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v97))) = uint16(v100)
									*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v97 + int32(2)
									if v93 < int64(0) {
										*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = int64(16384)
										v118 = int64(0) - v93
										v121 = v97 + int32(12)
										v123 = v85
										v128 = v118
										for {
											v132 = v121 - int32(2)
											v134 = base.I64_div_u_s(v128, int64(10000))
											v137 = v134*int64(55536) + v128
											*(*uint16)(unsafe.Add(mBase, uint32(v132))) = uint16(v137)
											v140 = v123 + int32(1)
											if base.Ui64(int64(9999)) < base.Ui64(v128) {
												v121 = v132
												v123 = v140
												v128 = v134
												continue
											} else {
												break
											}
											break
										}
										*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v132
										v146 = v140
										v148 = v123
									} else {
										v112 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v112
										if v93 == v112 {
											v146 = v85
											v148 = int32(0)
										} else {
											v118 = v93
											v121 = v97 + int32(12)
											v123 = v85
											v128 = v118
											for {
												v132 = v121 - int32(2)
												v134 = base.I64_div_u_s(v128, int64(10000))
												v137 = v134*int64(55536) + v128
												*(*uint16)(unsafe.Add(mBase, uint32(v132))) = uint16(v137)
												v140 = v123 + int32(1)
												if base.Ui64(int64(9999)) < base.Ui64(v128) {
													v121 = v132
													v123 = v140
													v128 = v134
													continue
												} else {
													break
												}
												break
											}
											*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v132
											v146 = v140
											v148 = v123
										}
									}
									*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v148
									*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v146
									v159 = F_make_result_safe(m, v13+int32(8), int32(0))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return int64(0)
									} else {
										F_pfree(m, v97)
										mBase = m.M
										v162 = m.ExcPending
										if v162 != 0 {
											return int64(0)
										} else {
											v166 = F_DirectFunctionCall2Coll(m, int32(1391), int32(0), base.I64_extend_i32_u(v159), base.I64_extend_i32_u(v89))
											mBase = m.M
											v167 = m.ExcPending
											if v167 != 0 {
												return int64(0)
											} else {
												v176 = v166
												m.G0 = v13 + int32(32)
												return v176
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
func F_int8_avg_accum_inv(m *base.Module, l0 int32) int64 {
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
			F_errmsg_internal(m, int32(_a_F_int8_avg_accum_inv_0), int32(0))
			mBase = m.M
			v113 = m.ExcPending
			if v113 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_int8_avg_accum_inv_1), int32(_a_F_int8_avg_accum_inv_2), int32(_a_F_int8_avg_accum_inv_3))
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
				F_errmsg_internal(m, int32(_a_F_int8_avg_accum_inv_0), int32(0))
				mBase = m.M
				v113 = m.ExcPending
				if v113 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int8_avg_accum_inv_1), int32(_a_F_int8_avg_accum_inv_2), int32(_a_F_int8_avg_accum_inv_3))
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
				v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
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
func F_int8_mul_cash(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v32 int64
	_ = v32
	var v39 int64
	_ = v39
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v9 = int64(63)
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = int64(32)
	v19 = int64(base.Ui64(v11) >> (uint(v18) % 64))
	v21 = int64(base.Ui64(v8) >> (uint(v18) % 64))
	v24 = int64(4294967295)
	v25 = v11 & v24
	v27 = v8 & v24
	v28 = v25 * v27
	v32 = int64(base.Ui64(v28)>>(uint(v18)%64)) + v25*v21
	v39 = v27*v19 + v32&v24
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v8*(v11>>(uint(v9)%64)) + v8>>(uint(v9)%64)*v11 + v19*v21 + int64(base.Ui64(v32)>>(uint(v18)%64)) + int64(base.Ui64(v39)>>(uint(v18)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v28&v24 | v39<<(uint(v18)%64)
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	if v50 != v51>>(uint(int64(63))%64) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v60 = m.ExcPending
		if v60 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int8_mul_cash_0), int32(0))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int8_mul_cash_1), int32(151), int32(_a_F_int8_mul_cash_2))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		m.G0 = v6 + int32(16)
		return v51
	}
}
