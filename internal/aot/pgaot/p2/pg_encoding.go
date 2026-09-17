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
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	v5 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	if base.B2i32(l3 == v5)|base.B2i32(l1 <= v5)|base.B2i32(l2 == l3) != 0 {
		v67 = l0
		m.G0 = v9 + int32(48)
		return v67
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
					v67 = l0
					m.G0 = v9 + int32(48)
					return v67
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
				v77 = m.ExcPending
				if v77 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_pg_do_encoding_conversion_0), int32(0))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_do_encoding_conversion_1), int32(388), int32(_a_F_pg_do_encoding_conversion_2))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
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
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(52461700))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int32(0)
							} else {
								if base.Ui32(l2) <= base.Ui32(int32(41)) {
									v98 = *(*int32)(unsafe.Add(mBase, uint32(l2<<(uint(int32(3))%32))+uint32(_c_F_pg_do_encoding_conversion[2])))
									v100 = v98
								} else {
									v100 = int32(_a_F_pg_do_encoding_conversion_3)
								}
								if base.Ui32(l3) <= base.Ui32(int32(41)) {
									v105 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(3))%32))+uint32(_c_F_pg_do_encoding_conversion[2])))
									v107 = v105
								} else {
									v107 = int32(_a_F_pg_do_encoding_conversion_3)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v107
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v100
								F_errmsg(m, int32(_a_F_pg_do_encoding_conversion_4), v9)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_pg_do_encoding_conversion_1), int32(396), int32(_a_F_pg_do_encoding_conversion_2))
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
						}
					} else {
						if base.Ui32(int32(536870911)) <= base.Ui32(l1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(261))
								mBase = m.M
								v124 = m.ExcPending
								if v124 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_pg_do_encoding_conversion_5), int32(0))
									mBase = m.M
									v128 = m.ExcPending
									if v128 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
										F_errdetail(m, int32(_a_F_pg_do_encoding_conversion_6), v9+int32(16))
										mBase = m.M
										v134 = m.ExcPending
										if v134 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_pg_do_encoding_conversion_1), int32(412), int32(_a_F_pg_do_encoding_conversion_2))
											mBase = m.M
											v139 = m.ExcPending
											if v139 != 0 {
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
							v48 = *(*int32)(unsafe.Add(mBase, _c_F_pg_do_encoding_conversion[3]))
							v53 = F_MemoryContextAllocHuge(m, v48, l1<<(uint(int32(2))%32)|int32(1))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								v56 = F_OidFunctionCall6Coll(m, v41, l2, l3, l0, v53, l1, int32(0))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									if base.Ui32(l1) < base.Ui32(int32(_a_F_pg_do_encoding_conversion_7)) {
										v67 = v53
										m.G0 = v9 + int32(48)
										return v67
									} else {
										v60 = F_strlen(m, v53)
										mBase = m.M
										if base.Ui32(int32(1073741823)) <= base.Ui32(v60) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v143 = m.ExcPending
											if v143 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(261))
												mBase = m.M
												v146 = m.ExcPending
												if v146 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_pg_do_encoding_conversion_5), int32(0))
													mBase = m.M
													v150 = m.ExcPending
													if v150 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l1
														F_errdetail(m, int32(_a_F_pg_do_encoding_conversion_6), v9+int32(32))
														mBase = m.M
														v156 = m.ExcPending
														if v156 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_pg_do_encoding_conversion_1), int32(440), int32(_a_F_pg_do_encoding_conversion_2))
															mBase = m.M
															v161 = m.ExcPending
															if v161 != 0 {
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
											v65 = F_repalloc(m, v53, v60+int32(1))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												v67 = v65
												m.G0 = v9 + int32(48)
												return v67
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
