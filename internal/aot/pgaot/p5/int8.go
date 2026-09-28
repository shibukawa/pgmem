package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int8_accum_inv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v44 int64
	_ = v44
	var v48 int64
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int64
	_ = v58
	var v62 int32
	_ = v62
	var v64 int64
	_ = v64
	var v67 int64
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v15 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v118 = m.ExcPending
		if v118 != 0 {
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_int8_accum_inv_0), int32(0))
			mBase = m.M
			v122 = m.ExcPending
			if v122 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_int8_accum_inv_1), int32(_a_F_int8_accum_inv_2), int32(_a_F_int8_accum_inv_3))
				mBase = m.M
				v127 = m.ExcPending
				if v127 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v17 = base.I32_wrap_i64(v16)
		if v17 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v118 = m.ExcPending
			if v118 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_int8_accum_inv_0), int32(0))
				mBase = m.M
				v122 = m.ExcPending
				if v122 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int8_accum_inv_1), int32(_a_F_int8_accum_inv_2), int32(_a_F_int8_accum_inv_3))
					mBase = m.M
					v127 = m.ExcPending
					if v127 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
			if v20 == int32(0) {
				v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
				*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = int64(0)
				v27 = F_palloc(m, int32(12))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v27
					v32 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v27))) = uint16(v32)
					*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v27 + int32(2)
					if v23 < int64(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = int64(16384)
						v48 = int64(0) - v23
						v51 = v32
						v54 = v27 + int32(12)
						v58 = v48
						for {
							v62 = v54 - int32(2)
							v64 = base.I64_div_u_s(v58, int64(10000))
							v67 = v64*int64(55536) + v58
							*(*uint16)(unsafe.Add(mBase, uint32(v62))) = uint16(v67)
							v70 = v51 + int32(1)
							if base.Ui64(int64(9999)) < base.Ui64(v58) {
								v51 = v70
								v54 = v62
								v58 = v64
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v62
						v74 = v70
						v79 = v51
					} else {
						v44 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v44
						if v23 == v44 {
							v74 = v32
							v79 = int32(0)
						} else {
							v48 = v23
							v51 = v32
							v54 = v27 + int32(12)
							v58 = v48
							for {
								v62 = v54 - int32(2)
								v64 = base.I64_div_u_s(v58, int64(10000))
								v67 = v64*int64(55536) + v58
								*(*uint16)(unsafe.Add(mBase, uint32(v62))) = uint16(v67)
								v70 = v51 + int32(1)
								if base.Ui64(int64(9999)) < base.Ui64(v58) {
									v51 = v70
									v54 = v62
									v58 = v64
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v62
							v74 = v70
							v79 = v51
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v79
					*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v74
					v89 = F_make_result_safe(m, v13+int32(8), int32(0))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int64(0)
					} else {
						F_pfree(m, v27)
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int64(0)
						} else {
							v93 = F_do_numeric_discard(m, v17, v89)
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int64(0)
							} else {
								if v93 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v131 = m.ExcPending
									if v131 != 0 {
										return int64(0)
									} else {
										F_errmsg_internal(m, int32(_a_F_int8_accum_inv_4), int32(0))
										mBase = m.M
										v135 = m.ExcPending
										if v135 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_int8_accum_inv_1), int32(_a_F_int8_accum_inv_5), int32(_a_F_int8_accum_inv_3))
											mBase = m.M
											v140 = m.ExcPending
											if v140 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									m.G0 = v13 + int32(32)
									return v16 & int64(4294967295)
								}
							}
						}
					}
				}
			} else {
				m.G0 = v13 + int32(32)
				return v16 & int64(4294967295)
			}
		}
	}
}
func F_int8_numeric(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v34 int64
	_ = v34
	var v38 int64
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int64
	_ = v47
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v55 int64
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(0)
	v17 = F_palloc(m, int32(12))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v17
		v22 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v17))) = uint16(v22)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v17 + int32(2)
		if v13 < int64(0) {
			*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(16384)
			v38 = int64(0) - v13
			v41 = v22
			v44 = v17 + int32(12)
			v47 = v38
			for {
				v50 = v44 - int32(2)
				v52 = base.I64_div_u_s(v47, int64(10000))
				v55 = v52*int64(55536) + v47
				*(*uint16)(unsafe.Add(mBase, uint32(v50))) = uint16(v55)
				v58 = v41 + int32(1)
				if base.Ui64(int64(9999)) < base.Ui64(v47) {
					v41 = v58
					v44 = v50
					v47 = v52
					continue
				} else {
					break
				}
				break
			}
			*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v50
			v62 = v58
			v66 = v41
		} else {
			v34 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v34
			if v13 == v34 {
				v62 = v22
				v66 = int32(0)
			} else {
				v38 = v13
				v41 = v22
				v44 = v17 + int32(12)
				v47 = v38
				for {
					v50 = v44 - int32(2)
					v52 = base.I64_div_u_s(v47, int64(10000))
					v55 = v52*int64(55536) + v47
					*(*uint16)(unsafe.Add(mBase, uint32(v50))) = uint16(v55)
					v58 = v41 + int32(1)
					if base.Ui64(int64(9999)) < base.Ui64(v47) {
						v41 = v58
						v44 = v50
						v47 = v52
						continue
					} else {
						break
					}
					break
				}
				*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v50
				v62 = v58
				v66 = v41
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v66
		*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v62
		v75 = F_make_result_safe(m, v11+int32(8), int32(0))
		mBase = m.M
		v76 = m.ExcPending
		if v76 != 0 {
			return int64(0)
		} else {
			F_pfree(m, v17)
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int64(0)
			} else {
				m.G0 = v11 + int32(32)
				return base.I64_extend_i32_u(v75)
			}
		}
	}
}
func F_int8_sum(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v44 int64
	_ = v44
	var v48 int64
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int64
	_ = v57
	var v61 int32
	_ = v61
	var v63 int64
	_ = v63
	var v66 int64
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int64
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v119 int64
	_ = v119
	var v123 int64
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int64
	_ = v132
	var v136 int32
	_ = v136
	var v138 int64
	_ = v138
	var v141 int64
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int64
	_ = v169
	var v170 int32
	_ = v170
	var v180 int64
	_ = v180
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v14 == int32(1) {
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
		if v17 == int32(1) {
			v20 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v20)
			v180 = int64(0)
			m.G0 = v12 + int32(32)
			return v180
		} else {
			v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
			*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(0)
			v27 = F_palloc(m, int32(12))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v27
				v32 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v27))) = uint16(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v27 + int32(2)
				if v23 < int64(0) {
					*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(16384)
					v48 = int64(0) - v23
					v51 = v32
					v54 = v27 + int32(12)
					v57 = v48
					for {
						v61 = v54 - int32(2)
						v63 = base.I64_div_u_s(v57, int64(10000))
						v66 = v63*int64(55536) + v57
						*(*uint16)(unsafe.Add(mBase, uint32(v61))) = uint16(v66)
						v69 = v51 + int32(1)
						if base.Ui64(int64(9999)) < base.Ui64(v57) {
							v51 = v69
							v54 = v61
							v57 = v63
							continue
						} else {
							break
						}
						break
					}
					*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v61
					v73 = v69
					v77 = v51
				} else {
					v44 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v44
					if v23 == v44 {
						v73 = v32
						v77 = v2
					} else {
						v48 = v23
						v51 = v32
						v54 = v27 + int32(12)
						v57 = v48
						for {
							v61 = v54 - int32(2)
							v63 = base.I64_div_u_s(v57, int64(10000))
							v66 = v63*int64(55536) + v57
							*(*uint16)(unsafe.Add(mBase, uint32(v61))) = uint16(v66)
							v69 = v51 + int32(1)
							if base.Ui64(int64(9999)) < base.Ui64(v57) {
								v51 = v69
								v54 = v61
								v57 = v63
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v61
						v73 = v69
						v77 = v51
					}
				}
				*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v77
				*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v73
				v87 = F_make_result_safe(m, v12+int32(8), int32(0))
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return int64(0)
				} else {
					F_pfree(m, v27)
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int64(0)
					} else {
						v180 = base.I64_extend_i32_u(v87)
						m.G0 = v12 + int32(32)
						return v180
					}
				}
			}
		}
	} else {
		v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v93 = F_pg_detoast_datum(m, v92)
		mBase = m.M
		v94 = m.ExcPending
		if v94 != 0 {
			return int64(0)
		} else {
			v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
			if v96 == int32(1) {
				v180 = base.I64_extend_i32_u(v93)
				m.G0 = v12 + int32(32)
				return v180
			} else {
				v99 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(0)
				v103 = F_palloc(m, int32(12))
				mBase = m.M
				v104 = m.ExcPending
				if v104 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v103
					v106 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v103))) = uint16(v106)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v103 + int32(2)
					if v99 < int64(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(16384)
						v123 = int64(0) - v99
						v126 = v106
						v129 = v103 + int32(12)
						v132 = v123
						for {
							v136 = v129 - int32(2)
							v138 = base.I64_div_u_s(v132, int64(10000))
							v141 = v138*int64(55536) + v132
							*(*uint16)(unsafe.Add(mBase, uint32(v136))) = uint16(v141)
							v144 = v126 + int32(1)
							if base.Ui64(int64(9999)) < base.Ui64(v132) {
								v126 = v144
								v129 = v136
								v132 = v138
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v136
						v148 = v144
						v152 = v126
					} else {
						v119 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v119
						if v99 == v119 {
							v148 = v106
							v152 = v2
						} else {
							v123 = v99
							v126 = v106
							v129 = v103 + int32(12)
							v132 = v123
							for {
								v136 = v129 - int32(2)
								v138 = base.I64_div_u_s(v132, int64(10000))
								v141 = v138*int64(55536) + v132
								*(*uint16)(unsafe.Add(mBase, uint32(v136))) = uint16(v141)
								v144 = v126 + int32(1)
								if base.Ui64(int64(9999)) < base.Ui64(v132) {
									v126 = v144
									v129 = v136
									v132 = v138
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v136
							v148 = v144
							v152 = v126
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v152
					*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v148
					v162 = F_make_result_safe(m, v12+int32(8), int32(0))
					mBase = m.M
					v163 = m.ExcPending
					if v163 != 0 {
						return int64(0)
					} else {
						F_pfree(m, v103)
						mBase = m.M
						v165 = m.ExcPending
						if v165 != 0 {
							return int64(0)
						} else {
							v169 = F_DirectFunctionCall2Coll(m, int32(1406), int32(0), base.I64_extend_i32_u(v93), base.I64_extend_i32_u(v162))
							mBase = m.M
							v170 = m.ExcPending
							if v170 != 0 {
								return int64(0)
							} else {
								v180 = v169
								m.G0 = v12 + int32(32)
								return v180
							}
						}
					}
				}
			}
		}
	}
}
