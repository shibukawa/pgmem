package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_do_encoding_conversion(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	if l3 == int32(0) {
		v65 = l0
		m.G0 = v9 + int32(48)
		return v65
	} else {
		if l1 <= int32(0) {
			v65 = l0
			m.G0 = v9 + int32(48)
			return v65
		} else {
			if l2 == l3 {
				v65 = l0
				m.G0 = v9 + int32(48)
				return v65
			} else {
				if l2 == int32(0) {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l3*int32(28))+uint32(_consts[1226])))
					v23 = m.T0[v22].(func(*base.Module, int32, int32) int32)(m, l0, l1)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						if l1 == v23 {
							v65 = l0
							m.G0 = v9 + int32(48)
							return v65
						} else {
							F_report_invalid_encoding(m, l3, l0+v23, l1-v23)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, _consts[65]))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
					if base.B2i32(v34 == int32(2)) == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(257884), int32(0))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(494365), int32(388), int32(271158))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v39 = F_FindDefaultConversionProc(m, l2, l3)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							if v39 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(52461700))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										if base.Ui32(l2) <= base.Ui32(int32(41)) {
											v101 = *(*int32)(unsafe.Add(mBase, uint32(l2<<(uint(int32(3))%32))+uint32(_consts[363])))
											v102 = v101
										} else {
											v102 = int32(757756)
										}
										if base.Ui32(l3) <= base.Ui32(int32(41)) {
											v111 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(3))%32))+uint32(_consts[363])))
											v112 = v111
										} else {
											v112 = int32(757756)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v112
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = v102
										F_errmsg(m, int32(71206), v9)
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(494365), int32(396), int32(271158))
											mBase = m.M
											v122 = m.ExcPending
											if v122 != 0 {
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
									v126 = m.ExcPending
									if v126 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(261))
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(13904), int32(0))
											mBase = m.M
											v133 = m.ExcPending
											if v133 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
												F_errdetail(m, int32(618333), v9+int32(16))
												mBase = m.M
												v139 = m.ExcPending
												if v139 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(494365), int32(412), int32(271158))
													mBase = m.M
													v144 = m.ExcPending
													if v144 != 0 {
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
									v46 = *(*int32)(unsafe.Add(mBase, _consts[0]))
									v51 = F_MemoryContextAllocHuge(m, v46, l1<<(uint(int32(2))%32)|int32(1))
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										v54 = F_OidFunctionCall6Coll(m, v39, l2, l3, l0, v51, l1, int32(0))
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return int32(0)
										} else {
											if base.Ui32(l1) < base.Ui32(int32(1000001)) {
												v65 = v51
												m.G0 = v9 + int32(48)
												return v65
											} else {
												v58 = F_strlen(m, v51)
												mBase = m.M
												if base.Ui32(int32(1073741823)) <= base.Ui32(v58) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(261))
														mBase = m.M
														v151 = m.ExcPending
														if v151 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(13904), int32(0))
															mBase = m.M
															v155 = m.ExcPending
															if v155 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l1
																F_errdetail(m, int32(618333), v9+int32(32))
																mBase = m.M
																v161 = m.ExcPending
																if v161 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(494365), int32(440), int32(271158))
																	mBase = m.M
																	v166 = m.ExcPending
																	if v166 != 0 {
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
													v63 = F_repalloc(m, v51, v58+int32(1))
													mBase = m.M
													v64 = m.ExcPending
													if v64 != 0 {
														return int32(0)
													} else {
														v65 = v63
														m.G0 = v9 + int32(48)
														return v65
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
