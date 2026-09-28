package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CacheInvalidateSmgr(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	v4 = m.G0
	v5 = int32(16)
	v6 = v4 - v5
	m.G0 = v6
	v8 = int32(253)
	*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v8)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+2)) = uint16(v10)
	v13 = int32(base.Ui32(v10) >> (uint(v5) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)) = uint8(v13)
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v6)+4)) = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v17
	F_SIInsertDataEntries(m, v6, int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		m.G0 = v6 + int32(16)
		return
	}
}
func F_PlanCacheRoleCallback(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	v4 = int32(0)
	if base.B2i32(l2 == v4)|base.B2i32(l1 != int32(21)) == v4 {
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_PlanCacheRoleCallback[0]))
		if l2 != v13 {
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_PlanCacheRoleCallback[1]))
			if base.B2i32(v16 == int32(0))|base.B2i32(v16 == int32(_a_F_PlanCacheRoleCallback_0)) != 0 {
			} else {
				v23 = v16
				for {
					v27 = v23 - int32(5)
					v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
					if v28 != int32(1) {
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v23-int32(96))))
						if v33 != 0 {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
							switch v37 - int32(137) {
							case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
								v41 = int32(1)
							default:
								v41 = int32(0)
							}
							if v41 != 0 {
								v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23-int32(15)))))
								if v71 == int32(1) {
									v74 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v74)
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v23-int32(12))))
									if v78 != 0 {
										v87 = v78
										v88 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v87)+10)) = uint8(v88)
									} else {
									}
								} else {
									v81 = *(*int32)(unsafe.Add(mBase, uint32(v23-int32(12))))
									if v81 == int32(0) {
									} else {
										v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+16)))
										if v84 != int32(1) {
										} else {
											v87 = v81
											v88 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v87)+10)) = uint8(v88)
										}
									}
								}
							} else {
							}
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v23-int32(92))))
							if v44 == int32(0) {
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
								if v48 != int32(6) {
									v63 = int32(1)
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v44)+28))
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
									v55 = v53 - int32(201)
									if base.Ui32(int32(41)) < base.Ui32(v55) {
										v63 = int32(0)
									} else {
										v63 = base.I32_wrap_i64(int64(base.Ui64(int64(3298534887425)) >> (uint(base.I64_extend_i32_u(v55)) % 64)))
									}
								}
								if v63&int32(1) == int32(0) {
								} else {
									v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23-int32(15)))))
									if v71 == int32(1) {
										v74 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v74)
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v23-int32(12))))
										if v78 != 0 {
											v87 = v78
											v88 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v87)+10)) = uint8(v88)
										} else {
										}
									} else {
										v81 = *(*int32)(unsafe.Add(mBase, uint32(v23-int32(12))))
										if v81 == int32(0) {
										} else {
											v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+16)))
											if v84 != int32(1) {
											} else {
												v87 = v81
												v88 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v87)+10)) = uint8(v88)
											}
										}
									}
								}
							}
						}
					}
					v92 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
					if v92 != int32(_a_F_PlanCacheRoleCallback_0) {
						v23 = v92
						continue
					} else {
						break
					}
					break
				}
			}
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_PlanCacheRoleCallback[1]))
		if base.B2i32(v16 == int32(0))|base.B2i32(v16 == int32(_a_F_PlanCacheRoleCallback_0)) != 0 {
		} else {
			v23 = v16
			for {
				v27 = v23 - int32(5)
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
				if v28 != int32(1) {
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v23-int32(96))))
					if v33 != 0 {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
						switch v37 - int32(137) {
						case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
							v41 = int32(1)
						default:
							v41 = int32(0)
						}
						if v41 != 0 {
							v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23-int32(15)))))
							if v71 == int32(1) {
								v74 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v74)
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v23-int32(12))))
								if v78 != 0 {
									v87 = v78
									v88 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v87)+10)) = uint8(v88)
								} else {
								}
							} else {
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v23-int32(12))))
								if v81 == int32(0) {
								} else {
									v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+16)))
									if v84 != int32(1) {
									} else {
										v87 = v81
										v88 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v87)+10)) = uint8(v88)
									}
								}
							}
						} else {
						}
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v23-int32(92))))
						if v44 == int32(0) {
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
							if v48 != int32(6) {
								v63 = int32(1)
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v44)+28))
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
								v55 = v53 - int32(201)
								if base.Ui32(int32(41)) < base.Ui32(v55) {
									v63 = int32(0)
								} else {
									v63 = base.I32_wrap_i64(int64(base.Ui64(int64(3298534887425)) >> (uint(base.I64_extend_i32_u(v55)) % 64)))
								}
							}
							if v63&int32(1) == int32(0) {
							} else {
								v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23-int32(15)))))
								if v71 == int32(1) {
									v74 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v74)
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v23-int32(12))))
									if v78 != 0 {
										v87 = v78
										v88 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v87)+10)) = uint8(v88)
									} else {
									}
								} else {
									v81 = *(*int32)(unsafe.Add(mBase, uint32(v23-int32(12))))
									if v81 == int32(0) {
									} else {
										v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+16)))
										if v84 != int32(1) {
										} else {
											v87 = v81
											v88 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v87)+10)) = uint8(v88)
										}
									}
								}
							}
						}
					}
				}
				v92 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
				if v92 != int32(_a_F_PlanCacheRoleCallback_0) {
					v23 = v92
					continue
				} else {
					break
				}
				break
			}
		}
	}
	return
}
