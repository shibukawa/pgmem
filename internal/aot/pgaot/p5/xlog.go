package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_XLOGShmemSize(m *base.Module) int32 {
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
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemSize[0]))
	if v9 != int32(-1) {
		v54 = F_mul_size(m, int32(128), int32(9))
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return int32(0)
		} else {
			v56 = F_add_size(m, int32(448), v54)
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				v60 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemSize[0]))
				v61 = F_mul_size(m, int32(8), v60)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					v63 = F_add_size(m, v56, v61)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						v66 = F_add_size(m, v63, int32(_a_F_XLOGShmemSize_0))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							v70 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemSize[0]))
							v71 = F_mul_size(m, int32(_a_F_XLOGShmemSize_0), v70)
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								v73 = F_add_size(m, v66, v71)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									m.G0 = v6 + int32(48)
									return v73
								}
							}
						}
					}
				}
			}
		}
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemSize[1]))
		v16 = base.I32_div_s(v14, int32(32))
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemSize[2]))
		v20 = base.I32_div_s(v18, int32(_a_F_XLOGShmemSize_0))
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
		v28 = v6 + int32(16)
		v31 = F_pg_snprintf(m, v28, int32(32), int32(_a_F_XLOGShmemSize_1), v6)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			v36 = int32(1)
			F_SetConfigOption(m, int32(_a_F_XLOGShmemSize_2), v28, v36, v36)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemSize[0]))
				if v41 != int32(-1) {
					v54 = F_mul_size(m, int32(128), int32(9))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						v56 = F_add_size(m, int32(448), v54)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemSize[0]))
							v61 = F_mul_size(m, int32(8), v60)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								v63 = F_add_size(m, v56, v61)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									v66 = F_add_size(m, v63, int32(_a_F_XLOGShmemSize_0))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										v70 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemSize[0]))
										v71 = F_mul_size(m, int32(_a_F_XLOGShmemSize_0), v70)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											v73 = F_add_size(m, v66, v71)
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int32(0)
											} else {
												m.G0 = v6 + int32(48)
												return v73
											}
										}
									}
								}
							}
						}
					}
				} else {
					F_SetConfigOption(m, int32(_a_F_XLOGShmemSize_2), v28, int32(1), int32(10))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v54 = F_mul_size(m, int32(128), int32(9))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							v56 = F_add_size(m, int32(448), v54)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemSize[0]))
								v61 = F_mul_size(m, int32(8), v60)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									v63 = F_add_size(m, v56, v61)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										v66 = F_add_size(m, v63, int32(_a_F_XLOGShmemSize_0))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											v70 = *(*int32)(unsafe.Add(mBase, _c_F_XLOGShmemSize[0]))
											v71 = F_mul_size(m, int32(_a_F_XLOGShmemSize_0), v70)
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return int32(0)
											} else {
												v73 = F_add_size(m, v66, v71)
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return int32(0)
												} else {
													m.G0 = v6 + int32(48)
													return v73
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
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v3 = int32(base.Ui32(l0) >> (uint(int32(4)) % 32))
	if v3 != int32(15) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v3<<(uint(int32(2))%32))+uint32(_c_F_xlog_identify[0])))
		v10 = v8
	} else {
		v10 = int32(0)
	}
	return v10
}
