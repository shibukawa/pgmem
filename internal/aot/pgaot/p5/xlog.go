package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_XLOGShmemAttach(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemAttach[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+176))
	*(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemAttach[1])) = v5
	return
}
func F_XLOGShmemRequest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	v4 = m.G0
	v6 = v4 - int32(80)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemRequest[0]))
	if v9 != int32(-1) {
		v52 = F_mul_size(m, int32(128), int32(9))
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return
		} else {
			v54 = F_add_size(m, int32(448), v52)
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return
			} else {
				v58 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemRequest[0]))
				v59 = F_mul_size(m, int32(8), v58)
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return
				} else {
					v61 = F_add_size(m, v54, v59)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						v64 = F_add_size(m, v61, int32(_a_F_XLOGShmemRequest_0))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							v68 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemRequest[0]))
							v69 = F_mul_size(m, int32(_a_F_XLOGShmemRequest_0), v68)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								v71 = F_add_size(m, v64, v69)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = int32(_a_F_XLOGShmemRequest_1)
									*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v71
									*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = int32(_a_F_XLOGShmemRequest_2)
									F_ShmemRequestStructWithOpts(m, v6+int32(32))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(_a_F_XLOGShmemRequest_3)
										*(*int64)(unsafe.Add(mBase, uint32(v6)+20)) = int64(312)
										*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(_a_F_XLOGShmemRequest_4)
										F_ShmemRequestStructWithOpts(m, v6+int32(16))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											m.G0 = v6 + int32(80)
											return
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
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemRequest[1]))
		v16 = base.I32_div_s(v14, int32(32))
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemRequest[2]))
		v20 = base.I32_div_s(v18, int32(_a_F_XLOGShmemRequest_0))
		if v16 < v20 {
			v22 = v16
		} else {
			v22 = v20
		}
		if v22 <= int32(8) {
			v25 = int32(8)
		} else {
			v25 = v22
		}
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = v25
		v28 = v6 + int32(48)
		v31 = F_pg_snprintf(m, v28, int32(32), int32(_a_F_XLOGShmemRequest_5), v6)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			v34 = int32(1)
			F_SetConfigOption(m, int32(_a_F_XLOGShmemRequest_6), v28, v34, v34)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemRequest[0]))
				if v39 != int32(-1) {
					v52 = F_mul_size(m, int32(128), int32(9))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						v54 = F_add_size(m, int32(448), v52)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemRequest[0]))
							v59 = F_mul_size(m, int32(8), v58)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								v61 = F_add_size(m, v54, v59)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									v64 = F_add_size(m, v61, int32(_a_F_XLOGShmemRequest_0))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										v68 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemRequest[0]))
										v69 = F_mul_size(m, int32(_a_F_XLOGShmemRequest_0), v68)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return
										} else {
											v71 = F_add_size(m, v64, v69)
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = int32(_a_F_XLOGShmemRequest_1)
												*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v71
												*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = int32(_a_F_XLOGShmemRequest_2)
												F_ShmemRequestStructWithOpts(m, v6+int32(32))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(_a_F_XLOGShmemRequest_3)
													*(*int64)(unsafe.Add(mBase, uint32(v6)+20)) = int64(312)
													*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(_a_F_XLOGShmemRequest_4)
													F_ShmemRequestStructWithOpts(m, v6+int32(16))
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return
													} else {
														m.G0 = v6 + int32(80)
														return
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
					F_SetConfigOption(m, int32(_a_F_XLOGShmemRequest_6), v28, int32(1), int32(10))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						v52 = F_mul_size(m, int32(128), int32(9))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							v54 = F_add_size(m, int32(448), v52)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemRequest[0]))
								v59 = F_mul_size(m, int32(8), v58)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									v61 = F_add_size(m, v54, v59)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return
									} else {
										v64 = F_add_size(m, v61, int32(_a_F_XLOGShmemRequest_0))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return
										} else {
											v68 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemRequest[0]))
											v69 = F_mul_size(m, int32(_a_F_XLOGShmemRequest_0), v68)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return
											} else {
												v71 = F_add_size(m, v64, v69)
												mBase = m.M
												v72 = m.ExcPending
												if v72 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = int32(_a_F_XLOGShmemRequest_1)
													*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v71
													*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = int32(_a_F_XLOGShmemRequest_2)
													F_ShmemRequestStructWithOpts(m, v6+int32(32))
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(_a_F_XLOGShmemRequest_3)
														*(*int64)(unsafe.Add(mBase, uint32(v6)+20)) = int64(312)
														*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(_a_F_XLOGShmemRequest_4)
														F_ShmemRequestStructWithOpts(m, v6+int32(16))
														mBase = m.M
														v93 = m.ExcPending
														if v93 != 0 {
															return
														} else {
															m.G0 = v6 + int32(80)
															return
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
}
func F_xlog_identify(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	v6 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(60))+uint32(_c_F_xlog_identify[0])))
	return v6
}
