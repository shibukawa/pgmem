package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgss_ExecutorEnd(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
	if v7 == int64(0) {
		v53 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorEnd[0]))
		if v53 != 0 {
			m.T0[v53].(func(*base.Module, int32))(m, l0)
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return
			} else {
				return
			}
		} else {
			F_standard_ExecutorEnd(m, l0)
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		if v10 == int32(0) {
			v53 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorEnd[0]))
			if v53 != 0 {
				m.T0[v53].(func(*base.Module, int32))(m, l0)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					return
				}
			} else {
				F_standard_ExecutorEnd(m, l0)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorEnd[1]))
			if int32(0) <= v14 {
				v53 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorEnd[0]))
				if v53 != 0 {
					m.T0[v53].(func(*base.Module, int32))(m, l0)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						return
					}
				} else {
					F_standard_ExecutorEnd(m, l0)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorEnd[2]))
				if v18 != int32(2) {
					if v18 != int32(1) {
						v53 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorEnd[0]))
						if v53 != 0 {
							m.T0[v53].(func(*base.Module, int32))(m, l0)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return
							} else {
								return
							}
						} else {
							F_standard_ExecutorEnd(m, l0)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorEnd[3]))
						if v24 != 0 {
							v53 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorEnd[0]))
							if v53 != 0 {
								m.T0[v53].(func(*base.Module, int32))(m, l0)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return
								} else {
									return
								}
							} else {
								F_standard_ExecutorEnd(m, l0)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									return
								}
							}
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v6)+112))
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v6)+116))
							v29 = *(*int64)(unsafe.Add(mBase, uint32(v10)+184))
							v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							v34 = *(*int64)(unsafe.Add(mBase, uint32(v33)+120))
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)+180))
							if v39 != 0 {
								v43 = v39 + int32(8)
							} else {
								v43 = int32(0)
							}
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v33)+164))
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v33)+168))
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
							F_pgss_store(m, v25, v7, v26, v27, int32(1), base.F64_div(base.F64_convert_i64_s(v29), float64(1e+06)), v34, v10+int32(192), v10+int32(320), v43, int32(0), v45, v46, v47)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								v53 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorEnd[0]))
								if v53 != 0 {
									m.T0[v53].(func(*base.Module, int32))(m, l0)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return
									} else {
										return
									}
								} else {
									F_standard_ExecutorEnd(m, l0)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v6)+112))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v6)+116))
					v29 = *(*int64)(unsafe.Add(mBase, uint32(v10)+184))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					v34 = *(*int64)(unsafe.Add(mBase, uint32(v33)+120))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)+180))
					if v39 != 0 {
						v43 = v39 + int32(8)
					} else {
						v43 = int32(0)
					}
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v33)+164))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v33)+168))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
					F_pgss_store(m, v25, v7, v26, v27, int32(1), base.F64_div(base.F64_convert_i64_s(v29), float64(1e+06)), v34, v10+int32(192), v10+int32(320), v43, int32(0), v45, v46, v47)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorEnd[0]))
						if v53 != 0 {
							m.T0[v53].(func(*base.Module, int32))(m, l0)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return
							} else {
								return
							}
						} else {
							F_standard_ExecutorEnd(m, l0)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_pgss_ExecutorRun(m *base.Module, l0 int32, l1 int32, l2 int64) {
	var v7 int32
	_ = v7
	Fn14354(m, l0, l1, l2, int32(_a_F_pgss_ExecutorRun_0), int32(_a_F_pgss_ExecutorRun_1))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_pgss_shmem_request(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int64
	_ = v6
	var v31 int64
	_ = v31
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	v2 = m.G0
	v4 = v2 - int32(112)
	m.G0 = v4
	v6 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4)+32)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v4)+24)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v4)+44)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4)+56)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v4)+72)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v4)+80)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v4)+88)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v4)+96)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v4)+64)) = int64(1924145348632)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+104)) = int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+40)) = int32(_a_F_pgss_shmem_request_0)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+108)) = int32(_a_F_pgss_shmem_request_1)
	v31 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_request[0])))
	*(*int64)(unsafe.Add(mBase, uint32(v4)+48)) = v31
	F_ShmemRequestHashWithOpts(m, v4+int32(24))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+20)) = int32(_a_F_pgss_shmem_request_2)
		*(*int64)(unsafe.Add(mBase, uint32(v4)+12)) = int64(176)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = int32(_a_F_pgss_shmem_request_3)
		F_ShmemRequestStructWithOpts(m, v4+int32(8))
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return
		} else {
			m.G0 = v4 + int32(112)
			return
		}
	}
}
