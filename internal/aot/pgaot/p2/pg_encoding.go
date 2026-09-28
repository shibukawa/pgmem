package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_do_encoding_conversion(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	v5 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	if base.B2i32(l3 == v5)|base.B2i32(l1 <= v5)|base.B2i32(l2 == l3) != 0 {
		v72 = l0
		m.G0 = v9 + int32(48)
		return v72
	} else {
		if l2 == int32(0) {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l3*int32(28))+uint32(_c_F_pg_do_encoding_conversion[0])))
			v25 = m.T0[v24].(func(*base.Module, int32, int32) int32)(m, l0, l1)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				if l1 == v25 {
					v72 = l0
					m.G0 = v9 + int32(48)
					return v72
				} else {
					F_report_invalid_encoding(m, l3, l0+v25, l1-v25)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, _c_F_pg_do_encoding_conversion[1]))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
			if base.B2i32(v36 == int32(2)) == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_pg_do_encoding_conversion_0), int32(0))
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_do_encoding_conversion_1), int32(388), int32(_a_F_pg_do_encoding_conversion_2))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v41 = F_FindDefaultConversionProc(m, l2, l3)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					if v41 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(52461700))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return int32(0)
							} else {
								if base.B2i32(l2 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(l2)) != 0 {
									v110 = int32(_a_F_pg_do_encoding_conversion_3)
								} else {
									v109 = *(*int32)(unsafe.Add(mBase, uint32(l2<<(uint(int32(3))%32))+uint32(_c_F_pg_do_encoding_conversion[2])))
									v110 = v109
								}
								if base.B2i32(l3 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(l3)) != 0 {
									v122 = int32(_a_F_pg_do_encoding_conversion_3)
								} else {
									v121 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(3))%32))+uint32(_c_F_pg_do_encoding_conversion[2])))
									v122 = v121
								}
								*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v122
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v110
								F_errmsg(m, int32(_a_F_pg_do_encoding_conversion_4), v9)
								mBase = m.M
								v127 = m.ExcPending
								if v127 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_pg_do_encoding_conversion_1), int32(396), int32(_a_F_pg_do_encoding_conversion_2))
									mBase = m.M
									v132 = m.ExcPending
									if v132 != 0 {
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
						if base.Ui32(int32(536870911)) <= base.Ui32(l1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v136 = m.ExcPending
							if v136 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(261))
								mBase = m.M
								v139 = m.ExcPending
								if v139 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_pg_do_encoding_conversion_5), int32(0))
									mBase = m.M
									v143 = m.ExcPending
									if v143 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
										v148 = F_errdetail(m, int32(_a_F_pg_do_encoding_conversion_6), v9+int32(16))
										mBase = m.M
										v149 = m.ExcPending
										if v149 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_pg_do_encoding_conversion_1), int32(412), int32(_a_F_pg_do_encoding_conversion_2))
											mBase = m.M
											v154 = m.ExcPending
											if v154 != 0 {
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
							v51 = *(*int32)(unsafe.Add(mBase, _c_F_pg_do_encoding_conversion[3]))
							v56 = F_MemoryContextAllocHuge(m, v51, l1<<(uint(int32(2))%32)|int32(1))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								v61 = F_OidFunctionCall6Coll(m, v41, base.I64_extend_i32_s(l2), base.I64_extend_i32_s(l3), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(v56), base.I64_extend_i32_u(l1), int64(0))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									if base.Ui32(l1) < base.Ui32(int32(_a_F_pg_do_encoding_conversion_7)) {
										v72 = v56
										m.G0 = v9 + int32(48)
										return v72
									} else {
										v65 = F_strlen(m, v56)
										mBase = m.M
										if base.Ui32(int32(1073741823)) <= base.Ui32(v65) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v158 = m.ExcPending
											if v158 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(261))
												mBase = m.M
												v161 = m.ExcPending
												if v161 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_pg_do_encoding_conversion_5), int32(0))
													mBase = m.M
													v165 = m.ExcPending
													if v165 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l1
														v170 = F_errdetail(m, int32(_a_F_pg_do_encoding_conversion_6), v9+int32(32))
														mBase = m.M
														v171 = m.ExcPending
														if v171 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_pg_do_encoding_conversion_1), int32(440), int32(_a_F_pg_do_encoding_conversion_2))
															mBase = m.M
															v176 = m.ExcPending
															if v176 != 0 {
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
											v70 = F_repalloc(m, v56, v65+int32(1))
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												v72 = v70
												m.G0 = v9 + int32(48)
												return v72
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
