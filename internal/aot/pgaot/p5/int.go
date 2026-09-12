package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__int_contains_joinsel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v9 = F_DirectFunctionCall5Coll(m, int32(4013), int32(0), v4, int32(2751), v6, v7, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F__int_contains_sel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v8 = F_DirectFunctionCall4Coll(m, int32(4012), int32(0), v4, int32(2751), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_call_int_check_hook(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = int32(1)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v14 == int32(0) {
		v81 = v13
		m.G0 = v11 - int32(-64)
		return v81
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[332])) = int32(50856066)
		v21 = int32(0)
		*(*int32)(unsafe.Add(mBase, _consts[330])) = v21
		*(*int32)(unsafe.Add(mBase, _consts[87])) = v21
		*(*int32)(unsafe.Add(mBase, _consts[1282])) = v21
		v29 = m.T0[v14].(func(*base.Module, int32, int32, int32) int32)(m, l1, l2, l3)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			if v29 != 0 {
				v81 = v13
				m.G0 = v11 - int32(-64)
				return v81
			} else {
				v34 = F_errstart(m, l4, int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					if v34 != 0 {
						v37 = *(*int32)(unsafe.Add(mBase, _consts[332]))
						F_errcode(m, v37)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, _consts[330]))
							if v41 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v41
								F_errmsg_internal(m, int32(216894), v9+int32(-16))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, _consts[87]))
									if v59 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v59
										F_errdetail_internal(m, int32(216894), v9+int32(-48))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, _consts[1282]))
											if v67 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
												F_errhint(m, int32(216894), v11)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(525050), int32(6870), int32(330754))
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
														return int32(0)
													} else {
														F_FlushErrorState(m)
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return int32(0)
														} else {
															v81 = int32(0)
															m.G0 = v11 - int32(-64)
															return v81
														}
													}
												}
											} else {
												F_errfinish(m, int32(525050), int32(6870), int32(330754))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													F_FlushErrorState(m)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														v81 = int32(0)
														m.G0 = v11 - int32(-64)
														return v81
													}
												}
											}
										}
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, _consts[1282]))
										if v67 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
											F_errhint(m, int32(216894), v11)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(525050), int32(6870), int32(330754))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													F_FlushErrorState(m)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														v81 = int32(0)
														m.G0 = v11 - int32(-64)
														return v81
													}
												}
											}
										} else {
											F_errfinish(m, int32(525050), int32(6870), int32(330754))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												F_FlushErrorState(m)
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return int32(0)
												} else {
													v81 = int32(0)
													m.G0 = v11 - int32(-64)
													return v81
												}
											}
										}
									}
								}
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v49
								*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v48
								F_errmsg(m, int32(511026), v9+int32(-32))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, _consts[87]))
									if v59 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v59
										F_errdetail_internal(m, int32(216894), v9+int32(-48))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, _consts[1282]))
											if v67 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
												F_errhint(m, int32(216894), v11)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(525050), int32(6870), int32(330754))
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
														return int32(0)
													} else {
														F_FlushErrorState(m)
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return int32(0)
														} else {
															v81 = int32(0)
															m.G0 = v11 - int32(-64)
															return v81
														}
													}
												}
											} else {
												F_errfinish(m, int32(525050), int32(6870), int32(330754))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													F_FlushErrorState(m)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														v81 = int32(0)
														m.G0 = v11 - int32(-64)
														return v81
													}
												}
											}
										}
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, _consts[1282]))
										if v67 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v67
											F_errhint(m, int32(216894), v11)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(525050), int32(6870), int32(330754))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													F_FlushErrorState(m)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														v81 = int32(0)
														m.G0 = v11 - int32(-64)
														return v81
													}
												}
											}
										} else {
											F_errfinish(m, int32(525050), int32(6870), int32(330754))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												F_FlushErrorState(m)
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return int32(0)
												} else {
													v81 = int32(0)
													m.G0 = v11 - int32(-64)
													return v81
												}
											}
										}
									}
								}
							}
						}
					} else {
						F_FlushErrorState(m)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							v81 = int32(0)
							m.G0 = v11 - int32(-64)
							return v81
						}
					}
				}
			}
		}
	}
}
