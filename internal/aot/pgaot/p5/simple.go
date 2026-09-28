package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SimpleLruDoesPhysicalPageExist(m *base.Module, l0 int32, l1 int64) int32 {
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
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v110 int32
	_ = v110
	v9 = m.G0
	v11 = v9 - int32(1056)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	v16 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[0])) = uint8(v16)
	*(*uint8)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[1])) = uint8(v16)
	v22 = v14 << (uint(int32(6)) % 32)
	v25 = *(*int64)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_SimpleLruDoesPhysicalPageExist[2])))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_SimpleLruDoesPhysicalPageExist[2]))) = v25 + int64(1)
	v30 = base.I64_div_s(l1, int64(32))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v32 == v16 {
		*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v30
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = v31
		v41 = F_pg_snprintf(m, v11+int32(32), int32(1024), int32(_a_F_SimpleLruDoesPhysicalPageExist_0), v11)
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int32(0)
		} else {
			v55 = int32(0)
			v59 = F_OpenTransientFile(m, v11+int32(32), v55)
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return int32(0)
			} else {
				if v59 < int32(0) {
					v64 = *(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[3]))
					if v64 == int32(44) {
						v110 = v55
						m.G0 = v11 + int32(1056)
						return v110
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[4])) = v64
						v70 = int32(0)
						*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[5])) = v70
						F_SlruReportIOError(m, l0, l1, v70)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							v76 = int64(0)
							v78 = F___lseek(m, v59, v76, int32(2))
							mBase = m.M
							if v78 < v76 {
								*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[5])) = int32(1)
								v86 = *(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[3]))
								*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[4])) = v86
								F_SlruReportIOError(m, l0, l1, int32(0))
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return int32(0)
								} else {
									v91 = F_CloseTransientFile(m, v59)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										if v91 != 0 {
											*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[5])) = int32(5)
											v98 = *(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[3]))
											*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[4])) = v98
											v110 = v55
										} else {
											v110 = base.B2i32(base.I64_extend_i32_s(base.I32_wrap_i64(l1-v30<<(uint(int64(5))%64))<<(uint(int32(13))%32)-int32(-8192)) <= v78)
										}
										m.G0 = v11 + int32(1056)
										return v110
									}
								}
							} else {
								v91 = F_CloseTransientFile(m, v59)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									if v91 != 0 {
										*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[5])) = int32(5)
										v98 = *(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[3]))
										*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[4])) = v98
										v110 = v55
									} else {
										v110 = base.B2i32(base.I64_extend_i32_s(base.I32_wrap_i64(l1-v30<<(uint(int64(5))%64))<<(uint(int32(13))%32)-int32(-8192)) <= v78)
									}
									m.G0 = v11 + int32(1056)
									return v110
								}
							}
						}
					}
				} else {
					v76 = int64(0)
					v78 = F___lseek(m, v59, v76, int32(2))
					mBase = m.M
					if v78 < v76 {
						*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[5])) = int32(1)
						v86 = *(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[3]))
						*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[4])) = v86
						F_SlruReportIOError(m, l0, l1, int32(0))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							v91 = F_CloseTransientFile(m, v59)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								if v91 != 0 {
									*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[5])) = int32(5)
									v98 = *(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[3]))
									*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[4])) = v98
									v110 = v55
								} else {
									v110 = base.B2i32(base.I64_extend_i32_s(base.I32_wrap_i64(l1-v30<<(uint(int64(5))%64))<<(uint(int32(13))%32)-int32(-8192)) <= v78)
								}
								m.G0 = v11 + int32(1056)
								return v110
							}
						}
					} else {
						v91 = F_CloseTransientFile(m, v59)
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							if v91 != 0 {
								*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[5])) = int32(5)
								v98 = *(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[3]))
								*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[4])) = v98
								v110 = v55
							} else {
								v110 = base.B2i32(base.I64_extend_i32_s(base.I32_wrap_i64(l1-v30<<(uint(int64(5))%64))<<(uint(int32(13))%32)-int32(-8192)) <= v78)
							}
							m.G0 = v11 + int32(1056)
							return v110
						}
					}
				}
			}
		}
	} else {
		*(*uint32)(unsafe.Add(mBase, uint32(v11)+20)) = uint32(v30)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v31
		v53 = F_pg_snprintf(m, v11+int32(32), int32(1024), int32(_a_F_SimpleLruDoesPhysicalPageExist_1), v11+int32(16))
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return int32(0)
		} else {
			v55 = int32(0)
			v59 = F_OpenTransientFile(m, v11+int32(32), v55)
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return int32(0)
			} else {
				if v59 < int32(0) {
					v64 = *(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[3]))
					if v64 == int32(44) {
						v110 = v55
						m.G0 = v11 + int32(1056)
						return v110
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[4])) = v64
						v70 = int32(0)
						*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[5])) = v70
						F_SlruReportIOError(m, l0, l1, v70)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							v76 = int64(0)
							v78 = F___lseek(m, v59, v76, int32(2))
							mBase = m.M
							if v78 < v76 {
								*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[5])) = int32(1)
								v86 = *(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[3]))
								*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[4])) = v86
								F_SlruReportIOError(m, l0, l1, int32(0))
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return int32(0)
								} else {
									v91 = F_CloseTransientFile(m, v59)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										if v91 != 0 {
											*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[5])) = int32(5)
											v98 = *(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[3]))
											*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[4])) = v98
											v110 = v55
										} else {
											v110 = base.B2i32(base.I64_extend_i32_s(base.I32_wrap_i64(l1-v30<<(uint(int64(5))%64))<<(uint(int32(13))%32)-int32(-8192)) <= v78)
										}
										m.G0 = v11 + int32(1056)
										return v110
									}
								}
							} else {
								v91 = F_CloseTransientFile(m, v59)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									if v91 != 0 {
										*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[5])) = int32(5)
										v98 = *(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[3]))
										*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[4])) = v98
										v110 = v55
									} else {
										v110 = base.B2i32(base.I64_extend_i32_s(base.I32_wrap_i64(l1-v30<<(uint(int64(5))%64))<<(uint(int32(13))%32)-int32(-8192)) <= v78)
									}
									m.G0 = v11 + int32(1056)
									return v110
								}
							}
						}
					}
				} else {
					v76 = int64(0)
					v78 = F___lseek(m, v59, v76, int32(2))
					mBase = m.M
					if v78 < v76 {
						*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[5])) = int32(1)
						v86 = *(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[3]))
						*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[4])) = v86
						F_SlruReportIOError(m, l0, l1, int32(0))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							v91 = F_CloseTransientFile(m, v59)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								if v91 != 0 {
									*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[5])) = int32(5)
									v98 = *(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[3]))
									*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[4])) = v98
									v110 = v55
								} else {
									v110 = base.B2i32(base.I64_extend_i32_s(base.I32_wrap_i64(l1-v30<<(uint(int64(5))%64))<<(uint(int32(13))%32)-int32(-8192)) <= v78)
								}
								m.G0 = v11 + int32(1056)
								return v110
							}
						}
					} else {
						v91 = F_CloseTransientFile(m, v59)
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							if v91 != 0 {
								*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[5])) = int32(5)
								v98 = *(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[3]))
								*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruDoesPhysicalPageExist[4])) = v98
								v110 = v55
							} else {
								v110 = base.B2i32(base.I64_extend_i32_s(base.I32_wrap_i64(l1-v30<<(uint(int64(5))%64))<<(uint(int32(13))%32)-int32(-8192)) <= v78)
							}
							m.G0 = v11 + int32(1056)
							return v110
						}
					}
				}
			}
		}
	}
}
