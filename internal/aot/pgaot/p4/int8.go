package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int8_avg(m *base.Module, l0 int32) int32 {
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
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v36 int32
	_ = v36
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
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	var v75 int64
	_ = v75
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v110 int64
	_ = v110
	var v116 int64
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int64
	_ = v126
	var v129 int32
	_ = v129
	var v131 int64
	_ = v131
	var v134 int64
	_ = v134
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
		if v19 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v180 = m.ExcPending
			if v180 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_int8_avg_0), int32(0))
				mBase = m.M
				v184 = m.ExcPending
				if v184 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int8_avg_1), int32(_a_F_int8_avg_2), int32(_a_F_int8_avg_3))
					mBase = m.M
					v189 = m.ExcPending
					if v189 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			if v20&int32(-4) != int32(160) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v180 = m.ExcPending
				if v180 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_int8_avg_0), int32(0))
					mBase = m.M
					v184 = m.ExcPending
					if v184 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int8_avg_1), int32(_a_F_int8_avg_2), int32(_a_F_int8_avg_3))
						mBase = m.M
						v189 = m.ExcPending
						if v189 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
				v32 = v15 + (v25<<(uint(int32(3))%32)+int32(23))&int32(-8)
				v33 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
				if v33 == int64(0) {
					v36 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v36)
					v172 = int32(0)
					m.G0 = v12 + int32(32)
					return v172
				} else {
					v40 = F_palloc(m, int32(12))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v40
						v43 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v40))) = uint16(v43)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v43
						*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = base.I32_wrap_i64(int64(base.Ui64(v33)>>(uint(int64(49))%64))) & int32(_a_F_int8_avg_4)
						v55 = v33 >> (uint(int64(63)) % 64)
						v60 = v40 + int32(12)
						v62 = v43
						v67 = v33 ^ v55 - v55
						for {
							v70 = v60 - int32(2)
							v72 = base.I64_div_u_s(v67, int64(10000))
							v75 = v72*int64(55536) + v67
							*(*uint16)(unsafe.Add(mBase, uint32(v70))) = uint16(v75)
							v78 = v62 + int32(1)
							if base.Ui64(int64(9999)) < base.Ui64(v67) {
								v60 = v70
								v62 = v78
								v67 = v72
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v62
						*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v78
						*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v70
						v84 = int32(0)
						v88 = F_make_result_opt_error(m, v12+int32(8), v84)
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v40)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								v92 = *(*int64)(unsafe.Add(mBase, uint32(v32)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(0)
								v96 = F_palloc(m, int32(12))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v96
									v99 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v96))) = uint16(v99)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v96 + int32(2)
									if v92 < int64(0) {
										*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(16384)
										v116 = int64(0) - v92
										v119 = v96 + int32(12)
										v121 = v84
										v126 = v116
										for {
											v129 = v119 - int32(2)
											v131 = base.I64_div_u_s(v126, int64(10000))
											v134 = v131*int64(55536) + v126
											*(*uint16)(unsafe.Add(mBase, uint32(v129))) = uint16(v134)
											v137 = v121 + int32(1)
											if base.Ui64(int64(9999)) < base.Ui64(v126) {
												v119 = v129
												v121 = v137
												v126 = v131
												continue
											} else {
												break
											}
											break
										}
										*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v129
										v143 = v137
										v145 = v121
									} else {
										v110 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v110
										if v92 == v110 {
											v143 = v84
											v145 = int32(0)
										} else {
											v116 = v92
											v119 = v96 + int32(12)
											v121 = v84
											v126 = v116
											for {
												v129 = v119 - int32(2)
												v131 = base.I64_div_u_s(v126, int64(10000))
												v134 = v131*int64(55536) + v126
												*(*uint16)(unsafe.Add(mBase, uint32(v129))) = uint16(v134)
												v137 = v121 + int32(1)
												if base.Ui64(int64(9999)) < base.Ui64(v126) {
													v119 = v129
													v121 = v137
													v126 = v131
													continue
												} else {
													break
												}
												break
											}
											*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v129
											v143 = v137
											v145 = v121
										}
									}
									*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v145
									*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v143
									v155 = F_make_result_opt_error(m, v12+int32(8), int32(0))
									mBase = m.M
									v156 = m.ExcPending
									if v156 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v96)
										mBase = m.M
										v158 = m.ExcPending
										if v158 != 0 {
											return int32(0)
										} else {
											v161 = F_DirectFunctionCall2Coll(m, int32(1260), int32(0), v155, v88)
											mBase = m.M
											v162 = m.ExcPending
											if v162 != 0 {
												return int32(0)
											} else {
												v172 = v161
												m.G0 = v12 + int32(32)
												return v172
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
func F_int8_avg_accum_inv(m *base.Module, l0 int32) int32 {
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
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
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
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
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
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v11 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v102 = m.ExcPending
		if v102 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_int8_avg_accum_inv_0), int32(0))
			mBase = m.M
			v106 = m.ExcPending
			if v106 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_int8_avg_accum_inv_1), int32(_a_F_int8_avg_accum_inv_2), int32(_a_F_int8_avg_accum_inv_3))
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
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v12 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v102 = m.ExcPending
			if v102 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_int8_avg_accum_inv_0), int32(0))
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int8_avg_accum_inv_1), int32(_a_F_int8_avg_accum_inv_2), int32(_a_F_int8_avg_accum_inv_3))
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
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v15 == int32(0) {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v19 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
				if v20 == int32(1) {
					v26 = v19 * (v19 >> (uint(int64(63)) % 64))
					v29 = int64(32)
					v30 = int64(base.Ui64(v19) >> (uint(v29) % 64))
					v35 = int64(4294967295)
					v36 = v19 & v35
					v39 = v36 * v36
					v42 = v36 * v30
					v43 = int64(base.Ui64(v39)>>(uint(v29)%64)) + v42
					v50 = v42 + v43&v35
					*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v26 + v26 + v30*v30 + int64(base.Ui64(v43)>>(uint(v29)%64)) + int64(base.Ui64(v50)>>(uint(v29)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v9))) = v39&v35 | v50<<(uint(v29)%64)
					v61 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
					v62 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
					*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v61 - v62
					v65 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
					v66 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v65 - v66 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v61) < base.Ui64(v62)))
				} else {
				}
				v74 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v74 - v19
				v77 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v77 - int64(1)
				v81 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v81 - v19>>(uint(int64(63))%64) - base.I64_extend_i32_u(base.B2i32(base.Ui64(v74) < base.Ui64(v19)))
			} else {
			}
			m.G0 = v9 + int32(16)
			return v12
		}
	}
}
func F_int8_mul_cash(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v34 int64
	_ = v34
	var v41 int64
	_ = v41
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v10 = int64(63)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	v20 = int64(32)
	v21 = int64(base.Ui64(v13) >> (uint(v20) % 64))
	v23 = int64(base.Ui64(v9) >> (uint(v20) % 64))
	v26 = int64(4294967295)
	v27 = v13 & v26
	v29 = v9 & v26
	v30 = v27 * v29
	v34 = int64(base.Ui64(v30)>>(uint(v20)%64)) + v27*v23
	v41 = v29*v21 + v34&v26
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v9*(v13>>(uint(v10)%64)) + v9>>(uint(v10)%64)*v13 + v21*v23 + int64(base.Ui64(v34)>>(uint(v20)%64)) + int64(base.Ui64(v41)>>(uint(v20)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v30&v26 | v41<<(uint(v20)%64)
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	if v52 != v53>>(uint(int64(63))%64) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v62 = m.ExcPending
		if v62 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int8_mul_cash_0), int32(0))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int8_mul_cash_1), int32(150), int32(_a_F_int8_mul_cash_2))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v75 = F_Int64GetDatum(m, v53)
		mBase = m.M
		v76 = m.ExcPending
		if v76 != 0 {
			return int32(0)
		} else {
			m.G0 = v6 + int32(16)
			return v75
		}
	}
}
