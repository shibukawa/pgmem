package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_executeItemUnwrapTargetArray(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v11 != int32(18) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v20
			F_errmsg_internal(m, int32(485246), v9)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(499757), int32(1680), int32(26220))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
		v31 = int32(1)
		v35 = F_executeAnyItem(m, l0, l1, v30, l3, v31, v31, v31, int32(0), l4)
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			m.G0 = v9 + int32(16)
			return v35
		}
	}
}
func F_executeNextItem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	if l2 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		if int32(0) < v12 {
			v25 = l2
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			v27 = F_executeItemOptUnwrapTarget(m, l0, v25, l3, l4, v26)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v67 = v27
				m.G0 = v10 + int32(48)
				return v67
			}
		} else {
			v30 = int32(0)
			if l4 == v30 {
				v67 = v30
				m.G0 = v10 + int32(48)
				return v67
			} else {
				if l5 == int32(0) {
					v44 = l3
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
					if v45 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v44
						*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v45
						*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v45
						*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v44
						v54 = F_list_make2_impl(m, v10+int32(8), v10+int32(4))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v54
							v67 = v30
							m.G0 = v10 + int32(48)
							return v67
						}
					} else {
						v59 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
						if v59 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v44
							v67 = v30
							m.G0 = v10 + int32(48)
							return v67
						} else {
							v63 = F_lappend(m, v59, v44)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v63
								v67 = v30
								m.G0 = v10 + int32(48)
								return v67
							}
						}
					}
				} else {
					v36 = F_palloc(m, int32(20))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v38
						v40 = *(*int64)(unsafe.Add(mBase, uint32(l3)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v40
						v42 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
						*(*int64)(unsafe.Add(mBase, uint32(v36))) = v42
						v44 = v36
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
						if v45 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v44
							*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v45
							*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v45
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v44
							v54 = F_list_make2_impl(m, v10+int32(8), v10+int32(4))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v54
								v67 = v30
								m.G0 = v10 + int32(48)
								return v67
							}
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
							if v59 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v44
								v67 = v30
								m.G0 = v10 + int32(48)
								return v67
							} else {
								v63 = F_lappend(m, v59, v44)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v63
									v67 = v30
									m.G0 = v10 + int32(48)
									return v67
								}
							}
						}
					}
				}
			}
		}
	} else {
		v16 = v10 + int32(12)
		v19 = F_jspGetNext(m, l1, v16)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			if v19 == int32(0) {
				v30 = int32(0)
				if l4 == v30 {
					v67 = v30
					m.G0 = v10 + int32(48)
					return v67
				} else {
					if l5 == int32(0) {
						v44 = l3
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
						if v45 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v44
							*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v45
							*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v45
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v44
							v54 = F_list_make2_impl(m, v10+int32(8), v10+int32(4))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v54
								v67 = v30
								m.G0 = v10 + int32(48)
								return v67
							}
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
							if v59 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v44
								v67 = v30
								m.G0 = v10 + int32(48)
								return v67
							} else {
								v63 = F_lappend(m, v59, v44)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v63
									v67 = v30
									m.G0 = v10 + int32(48)
									return v67
								}
							}
						}
					} else {
						v36 = F_palloc(m, int32(20))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v38
							v40 = *(*int64)(unsafe.Add(mBase, uint32(l3)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v40
							v42 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
							*(*int64)(unsafe.Add(mBase, uint32(v36))) = v42
							v44 = v36
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							if v45 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v44
								*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v45
								*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v45
								*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v44
								v54 = F_list_make2_impl(m, v10+int32(8), v10+int32(4))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v54
									v67 = v30
									m.G0 = v10 + int32(48)
									return v67
								}
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
								if v59 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(l4))) = v44
									v67 = v30
									m.G0 = v10 + int32(48)
									return v67
								} else {
									v63 = F_lappend(m, v59, v44)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v63
										v67 = v30
										m.G0 = v10 + int32(48)
										return v67
									}
								}
							}
						}
					}
				}
			} else {
				v25 = v16
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
				v27 = F_executeItemOptUnwrapTarget(m, l0, v25, l3, l4, v26)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v67 = v27
					m.G0 = v10 + int32(48)
					return v67
				}
			}
		}
	}
}
