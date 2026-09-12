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
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v111 int32
	_ = v111
	v9 = m.G0
	v11 = v9 - int32(1056)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	v16 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[89])) = uint8(v16)
	*(*uint8)(unsafe.Add(mBase, _consts[90])) = uint8(v16)
	v22 = v14 << (uint(int32(6)) % 32)
	v25 = *(*int64)(unsafe.Add(mBase, uint32(v22)+uint32(_consts[92])))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+uint32(_consts[92]))) = v25 + int64(1)
	v30 = l0 + int32(16)
	v32 = base.I64_div_s(l1, int64(32))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v33 == v16 {
		*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v32
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = v30
		v42 = F_pg_snprintf(m, v11+int32(32), int32(1024), int32(506646), v11)
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			v56 = int32(0)
			v60 = F_OpenTransientFile(m, v11+int32(32), v56)
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				if v60 < int32(0) {
					v65 = *(*int32)(unsafe.Add(mBase, _consts[86]))
					if v65 == int32(44) {
						v111 = v56
						m.G0 = v11 + int32(1056)
						return v111
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[93])) = v65
						v71 = int32(0)
						*(*int32)(unsafe.Add(mBase, _consts[94])) = v71
						F_SlruReportIOError(m, l0, l1, v71)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							v77 = int64(0)
							v79 = F___lseek(m, v60, v77, int32(2))
							mBase = m.M
							if v79 < v77 {
								*(*int32)(unsafe.Add(mBase, _consts[94])) = int32(1)
								v87 = *(*int32)(unsafe.Add(mBase, _consts[86]))
								*(*int32)(unsafe.Add(mBase, _consts[93])) = v87
								F_SlruReportIOError(m, l0, l1, int32(0))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int32(0)
								} else {
									v92 = F_CloseTransientFile(m, v60)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return int32(0)
									} else {
										if v92 != 0 {
											*(*int32)(unsafe.Add(mBase, _consts[94])) = int32(5)
											v99 = *(*int32)(unsafe.Add(mBase, _consts[86]))
											*(*int32)(unsafe.Add(mBase, _consts[93])) = v99
											v111 = v56
										} else {
											v111 = base.B2i32(base.I64_extend_i32_s(base.I32_wrap_i64(l1-v32<<(uint(int64(5))%64))<<(uint(int32(13))%32)-int32(-8192)) <= v79)
										}
										m.G0 = v11 + int32(1056)
										return v111
									}
								}
							} else {
								v92 = F_CloseTransientFile(m, v60)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int32(0)
								} else {
									if v92 != 0 {
										*(*int32)(unsafe.Add(mBase, _consts[94])) = int32(5)
										v99 = *(*int32)(unsafe.Add(mBase, _consts[86]))
										*(*int32)(unsafe.Add(mBase, _consts[93])) = v99
										v111 = v56
									} else {
										v111 = base.B2i32(base.I64_extend_i32_s(base.I32_wrap_i64(l1-v32<<(uint(int64(5))%64))<<(uint(int32(13))%32)-int32(-8192)) <= v79)
									}
									m.G0 = v11 + int32(1056)
									return v111
								}
							}
						}
					}
				} else {
					v77 = int64(0)
					v79 = F___lseek(m, v60, v77, int32(2))
					mBase = m.M
					if v79 < v77 {
						*(*int32)(unsafe.Add(mBase, _consts[94])) = int32(1)
						v87 = *(*int32)(unsafe.Add(mBase, _consts[86]))
						*(*int32)(unsafe.Add(mBase, _consts[93])) = v87
						F_SlruReportIOError(m, l0, l1, int32(0))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							v92 = F_CloseTransientFile(m, v60)
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int32(0)
							} else {
								if v92 != 0 {
									*(*int32)(unsafe.Add(mBase, _consts[94])) = int32(5)
									v99 = *(*int32)(unsafe.Add(mBase, _consts[86]))
									*(*int32)(unsafe.Add(mBase, _consts[93])) = v99
									v111 = v56
								} else {
									v111 = base.B2i32(base.I64_extend_i32_s(base.I32_wrap_i64(l1-v32<<(uint(int64(5))%64))<<(uint(int32(13))%32)-int32(-8192)) <= v79)
								}
								m.G0 = v11 + int32(1056)
								return v111
							}
						}
					} else {
						v92 = F_CloseTransientFile(m, v60)
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int32(0)
						} else {
							if v92 != 0 {
								*(*int32)(unsafe.Add(mBase, _consts[94])) = int32(5)
								v99 = *(*int32)(unsafe.Add(mBase, _consts[86]))
								*(*int32)(unsafe.Add(mBase, _consts[93])) = v99
								v111 = v56
							} else {
								v111 = base.B2i32(base.I64_extend_i32_s(base.I32_wrap_i64(l1-v32<<(uint(int64(5))%64))<<(uint(int32(13))%32)-int32(-8192)) <= v79)
							}
							m.G0 = v11 + int32(1056)
							return v111
						}
					}
				}
			}
		}
	} else {
		*(*uint32)(unsafe.Add(mBase, uint32(v11)+20)) = uint32(v32)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v30
		v54 = F_pg_snprintf(m, v11+int32(32), int32(1024), int32(507033), v11+int32(16))
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return int32(0)
		} else {
			v56 = int32(0)
			v60 = F_OpenTransientFile(m, v11+int32(32), v56)
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				if v60 < int32(0) {
					v65 = *(*int32)(unsafe.Add(mBase, _consts[86]))
					if v65 == int32(44) {
						v111 = v56
						m.G0 = v11 + int32(1056)
						return v111
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[93])) = v65
						v71 = int32(0)
						*(*int32)(unsafe.Add(mBase, _consts[94])) = v71
						F_SlruReportIOError(m, l0, l1, v71)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							v77 = int64(0)
							v79 = F___lseek(m, v60, v77, int32(2))
							mBase = m.M
							if v79 < v77 {
								*(*int32)(unsafe.Add(mBase, _consts[94])) = int32(1)
								v87 = *(*int32)(unsafe.Add(mBase, _consts[86]))
								*(*int32)(unsafe.Add(mBase, _consts[93])) = v87
								F_SlruReportIOError(m, l0, l1, int32(0))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int32(0)
								} else {
									v92 = F_CloseTransientFile(m, v60)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return int32(0)
									} else {
										if v92 != 0 {
											*(*int32)(unsafe.Add(mBase, _consts[94])) = int32(5)
											v99 = *(*int32)(unsafe.Add(mBase, _consts[86]))
											*(*int32)(unsafe.Add(mBase, _consts[93])) = v99
											v111 = v56
										} else {
											v111 = base.B2i32(base.I64_extend_i32_s(base.I32_wrap_i64(l1-v32<<(uint(int64(5))%64))<<(uint(int32(13))%32)-int32(-8192)) <= v79)
										}
										m.G0 = v11 + int32(1056)
										return v111
									}
								}
							} else {
								v92 = F_CloseTransientFile(m, v60)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int32(0)
								} else {
									if v92 != 0 {
										*(*int32)(unsafe.Add(mBase, _consts[94])) = int32(5)
										v99 = *(*int32)(unsafe.Add(mBase, _consts[86]))
										*(*int32)(unsafe.Add(mBase, _consts[93])) = v99
										v111 = v56
									} else {
										v111 = base.B2i32(base.I64_extend_i32_s(base.I32_wrap_i64(l1-v32<<(uint(int64(5))%64))<<(uint(int32(13))%32)-int32(-8192)) <= v79)
									}
									m.G0 = v11 + int32(1056)
									return v111
								}
							}
						}
					}
				} else {
					v77 = int64(0)
					v79 = F___lseek(m, v60, v77, int32(2))
					mBase = m.M
					if v79 < v77 {
						*(*int32)(unsafe.Add(mBase, _consts[94])) = int32(1)
						v87 = *(*int32)(unsafe.Add(mBase, _consts[86]))
						*(*int32)(unsafe.Add(mBase, _consts[93])) = v87
						F_SlruReportIOError(m, l0, l1, int32(0))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							v92 = F_CloseTransientFile(m, v60)
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int32(0)
							} else {
								if v92 != 0 {
									*(*int32)(unsafe.Add(mBase, _consts[94])) = int32(5)
									v99 = *(*int32)(unsafe.Add(mBase, _consts[86]))
									*(*int32)(unsafe.Add(mBase, _consts[93])) = v99
									v111 = v56
								} else {
									v111 = base.B2i32(base.I64_extend_i32_s(base.I32_wrap_i64(l1-v32<<(uint(int64(5))%64))<<(uint(int32(13))%32)-int32(-8192)) <= v79)
								}
								m.G0 = v11 + int32(1056)
								return v111
							}
						}
					} else {
						v92 = F_CloseTransientFile(m, v60)
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int32(0)
						} else {
							if v92 != 0 {
								*(*int32)(unsafe.Add(mBase, _consts[94])) = int32(5)
								v99 = *(*int32)(unsafe.Add(mBase, _consts[86]))
								*(*int32)(unsafe.Add(mBase, _consts[93])) = v99
								v111 = v56
							} else {
								v111 = base.B2i32(base.I64_extend_i32_s(base.I32_wrap_i64(l1-v32<<(uint(int64(5))%64))<<(uint(int32(13))%32)-int32(-8192)) <= v79)
							}
							m.G0 = v11 + int32(1056)
							return v111
						}
					}
				}
			}
		}
	}
}
func F_SimpleLruWritePage(m *base.Module, l0 int32, l1 int32) {
	var v5 int32
	_ = v5
	F_SlruInternalWritePage(m, l0, l1, int32(0))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
