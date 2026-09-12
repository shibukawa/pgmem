package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pq_endmsgread(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[475])) = uint8(v2)
	return
}
func F_pq_getmsgint(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v66 int32
	_ = v66
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	switch l1 - int32(1) {
	case 0:
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v11-v12 <= int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16908800))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(422245), int32(0))
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(515266), int32(533), int32(167949))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
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
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v12))))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v12 + int32(1)
			v66 = v18
			m.G0 = v7 + int32(16)
			return v66
		}
	case 1:
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v22-v23 <= int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v93 = m.ExcPending
			if v93 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16908800))
				mBase = m.M
				v96 = m.ExcPending
				if v96 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(422245), int32(0))
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(515266), int32(533), int32(167949))
						mBase = m.M
						v105 = m.ExcPending
						if v105 != 0 {
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
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27+v23))))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v23 + int32(2)
			v33 = int32(8)
			v66 = (v29<<(uint(v33)%32) | int32(base.Ui32(v29)>>(uint(v33)%32))) & int32(65535)
			m.G0 = v7 + int32(16)
			return v66
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v125 = m.ExcPending
		if v125 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
			F_errmsg_internal(m, int32(495542), v7)
			mBase = m.M
			v129 = m.ExcPending
			if v129 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(515266), int32(437), int32(95324))
				mBase = m.M
				v134 = m.ExcPending
				if v134 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 3:
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v40-v41 <= int32(3) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v109 = m.ExcPending
			if v109 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16908800))
				mBase = m.M
				v112 = m.ExcPending
				if v112 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(422245), int32(0))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(515266), int32(533), int32(167949))
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
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
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v45+v41)))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v41 + int32(4)
			v51 = int32(24)
			v53 = int32(65280)
			v55 = int32(8)
			v66 = v47<<(uint(v51)%32) | v47&v53<<(uint(v55)%32) | (int32(base.Ui32(v47)>>(uint(v55)%32))&v53 | int32(base.Ui32(v47)>>(uint(v51)%32)))
			m.G0 = v7 + int32(16)
			return v66
		}
	}
}
func F_pq_getmsgint64(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v37 int64
	_ = v37
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v4-v5 <= int32(7) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(16908800))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(422245), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(515266), int32(533), int32(167949))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v29 = *(*int64)(unsafe.Add(mBase, uint32(v27+v5)))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v5 + int32(8)
		v33 = int64(56)
		v35 = int64(65280)
		v37 = int64(40)
		v40 = int64(16711680)
		v42 = int64(24)
		v44 = int64(4278190080)
		v46 = int64(8)
		return v29<<(uint(v33)%64) | v29&v35<<(uint(v37)%64) | (v29&v40<<(uint(v42)%64) | v29&v44<<(uint(v46)%64)) | (int64(base.Ui64(v29)>>(uint(v46)%64))&v44 | int64(base.Ui64(v29)>>(uint(v42)%64))&v40 | (int64(base.Ui64(v29)>>(uint(v37)%64))&v35 | int64(base.Ui64(v29)>>(uint(v33)%64))))
	}
}
func F_pq_setkeepalivescount(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
	if l1 == int32(0) {
	} else {
		v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
		if v11 == int32(1) {
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+408))
			if l0 == v14 {
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+392))
				if int32(0) < v16 {
					if l0 == int32(0) {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+392))
						*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v47
					} else {
					}
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+408)) = v50
				} else {
					v19 = int32(0)
					v21 = m.G0
					v23 = v21 - int32(16)
					m.G0 = v23
					if l1 == v19 {
						v39 = v19
					} else {
						v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
						if v27 == int32(1) {
							v39 = v19
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+408))
							if v30 != 0 {
								v39 = v30
							} else {
								v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+392))
								if v31 != 0 {
									v39 = v31
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = int32(4)
									v37 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(392))))
									v39 = v37
								}
							}
						}
					}
					m.G0 = v23 + int32(16)
					if int32(0) <= v39 {
						if l0 == int32(0) {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+392))
							*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v47
						} else {
						}
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+408)) = v50
					} else {
					}
				}
			}
		}
	}
	m.G0 = v6 + int32(16)
	return
}
