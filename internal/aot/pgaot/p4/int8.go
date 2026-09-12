package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int8_avg(m *base.Module, l0 int32) int32 {
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v56 int64
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int64
	_ = v69
	var v72 int32
	_ = v72
	var v74 int64
	_ = v74
	var v77 int64
	_ = v77
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v112 int64
	_ = v112
	var v118 int64
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v129 int64
	_ = v129
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
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
		if v20 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v185 = m.ExcPending
			if v185 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(26116), int32(0))
				mBase = m.M
				v189 = m.ExcPending
				if v189 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(499737), int32(6928), int32(326556))
					mBase = m.M
					v194 = m.ExcPending
					if v194 != 0 {
						return int32(0)
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
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(26116), int32(0))
					mBase = m.M
					v189 = m.ExcPending
					if v189 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(499737), int32(6928), int32(326556))
						mBase = m.M
						v194 = m.ExcPending
						if v194 != 0 {
							return int32(0)
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
					v177 = int32(0)
					m.G0 = v13 + int32(32)
					return v177
				} else {
					v41 = F_palloc(m, int32(12))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v41
						v44 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v41))) = uint16(v44)
						*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v44
						*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = base.I32_wrap_i64(int64(base.Ui64(v34)>>(uint(int64(49))%64))) & int32(16384)
						v56 = v34 >> (uint(int64(63)) % 64)
						v61 = v41 + int32(12)
						v63 = v44
						v69 = v34 ^ v56 - v56
						for {
							v72 = v61 - int32(2)
							v74 = base.I64_div_u_s(v69, int64(10000))
							v77 = v74*int64(55536) + v69
							*(*uint16)(unsafe.Add(mBase, uint32(v72))) = uint16(v77)
							v80 = v63 + int32(1)
							if base.Ui64(int64(9999)) < base.Ui64(v69) {
								v61 = v72
								v63 = v80
								v69 = v74
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v63
						*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v80
						*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v72
						v86 = int32(0)
						v90 = F_make_result_opt_error(m, v13+int32(8), v86)
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v41)
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int32(0)
							} else {
								v94 = *(*int64)(unsafe.Add(mBase, uint32(v33)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = int64(0)
								v98 = F_palloc(m, int32(12))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v98
									v101 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v98))) = uint16(v101)
									*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v98 + int32(2)
									if v94 < int64(0) {
										*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = int64(16384)
										v118 = int64(0) - v94
										v121 = v98 + int32(12)
										v123 = v86
										v129 = v118
										for {
											v132 = v121 - int32(2)
											v134 = base.I64_div_u_s(v129, int64(10000))
											v137 = v134*int64(55536) + v129
											*(*uint16)(unsafe.Add(mBase, uint32(v132))) = uint16(v137)
											v140 = v123 + int32(1)
											if base.Ui64(int64(9999)) < base.Ui64(v129) {
												v121 = v132
												v123 = v140
												v129 = v134
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
										if v94 == v112 {
											v146 = v86
											v148 = int32(0)
										} else {
											v118 = v94
											v121 = v98 + int32(12)
											v123 = v86
											v129 = v118
											for {
												v132 = v121 - int32(2)
												v134 = base.I64_div_u_s(v129, int64(10000))
												v137 = v134*int64(55536) + v129
												*(*uint16)(unsafe.Add(mBase, uint32(v132))) = uint16(v137)
												v140 = v123 + int32(1)
												if base.Ui64(int64(9999)) < base.Ui64(v129) {
													v121 = v132
													v123 = v140
													v129 = v134
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
									v159 = F_make_result_opt_error(m, v13+int32(8), int32(0))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v98)
										mBase = m.M
										v162 = m.ExcPending
										if v162 != 0 {
											return int32(0)
										} else {
											v165 = F_DirectFunctionCall2Coll(m, int32(1276), int32(0), v159, v90)
											mBase = m.M
											v166 = m.ExcPending
											if v166 != 0 {
												return int32(0)
											} else {
												v177 = v165
												m.G0 = v13 + int32(32)
												return v177
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
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v77 int64
	_ = v77
	var v80 int64
	_ = v80
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v11 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v108 = m.ExcPending
		if v108 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(352996), int32(0))
			mBase = m.M
			v112 = m.ExcPending
			if v112 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(499737), int32(6172), int32(32649))
				mBase = m.M
				v117 = m.ExcPending
				if v117 != 0 {
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
			v108 = m.ExcPending
			if v108 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(352996), int32(0))
				mBase = m.M
				v112 = m.ExcPending
				if v112 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(499737), int32(6172), int32(32649))
					mBase = m.M
					v117 = m.ExcPending
					if v117 != 0 {
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
					v66 = v12 + int32(40)
					v67 = *(*int64)(unsafe.Add(mBase, uint32(v66)))
					v68 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v66))) = v67 - v68 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v61) < base.Ui64(v62)))
				} else {
				}
				v77 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v77 - v19
				v80 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v80 - int64(1)
				v85 = v12 + int32(24)
				v86 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
				*(*int64)(unsafe.Add(mBase, uint32(v85))) = v86 - v19>>(uint(int64(63))%64) - base.I64_extend_i32_u(base.B2i32(base.Ui64(v77) < base.Ui64(v19)))
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
				F_errmsg(m, int32(401737), int32(0))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(497640), int32(150), int32(558481))
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
