package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SlruInternalDeleteSegment(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	v2 = l1
	v5 = m.G0
	v7 = v5 - int32(1072)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v9 != int32(5) {
		v12 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v7)+48)) = v12
		*(*int64)(unsafe.Add(mBase, uint32(v7)+56)) = v12
		*(*int64)(unsafe.Add(mBase, uint32(v7)+64)) = v2
		*(*uint16)(unsafe.Add(mBase, uint32(v7)+48)) = uint16(v9)
		v22 = F_RegisterSyncRequest(m, v7+int32(48), int32(2), int32(1))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
			if v25 == int32(1) {
				*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v2
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v24
				v36 = F_pg_snprintf(m, v7+int32(48), int32(1024), int32(_a_F_SlruInternalDeleteSegment_0), v7+int32(16))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					v50 = F_errstart(m, int32(13), int32(0))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						if v50 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 + int32(48)
							F_errmsg_internal(m, int32(_a_F_SlruInternalDeleteSegment_1), v7)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_SlruInternalDeleteSegment_2), int32(1568), int32(_a_F_SlruInternalDeleteSegment_3))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									v65 = F_unlink(m, v7+int32(48))
									mBase = m.M
									m.G0 = v7 + int32(1072)
									return
								}
							}
						} else {
							v65 = F_unlink(m, v7+int32(48))
							mBase = m.M
							m.G0 = v7 + int32(1072)
							return
						}
					}
				}
			} else {
				*(*uint32)(unsafe.Add(mBase, uint32(v7)+36)) = uint32(v2)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v24
				v46 = F_pg_snprintf(m, v7+int32(48), int32(1024), int32(_a_F_SlruInternalDeleteSegment_4), v7+int32(32))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					v50 = F_errstart(m, int32(13), int32(0))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						if v50 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 + int32(48)
							F_errmsg_internal(m, int32(_a_F_SlruInternalDeleteSegment_1), v7)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_SlruInternalDeleteSegment_2), int32(1568), int32(_a_F_SlruInternalDeleteSegment_3))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									v65 = F_unlink(m, v7+int32(48))
									mBase = m.M
									m.G0 = v7 + int32(1072)
									return
								}
							}
						} else {
							v65 = F_unlink(m, v7+int32(48))
							mBase = m.M
							m.G0 = v7 + int32(1072)
							return
						}
					}
				}
			}
		}
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
		if v25 == int32(1) {
			*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v2
			*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v24
			v36 = F_pg_snprintf(m, v7+int32(48), int32(1024), int32(_a_F_SlruInternalDeleteSegment_0), v7+int32(16))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				v50 = F_errstart(m, int32(13), int32(0))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					if v50 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 + int32(48)
						F_errmsg_internal(m, int32(_a_F_SlruInternalDeleteSegment_1), v7)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_SlruInternalDeleteSegment_2), int32(1568), int32(_a_F_SlruInternalDeleteSegment_3))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								v65 = F_unlink(m, v7+int32(48))
								mBase = m.M
								m.G0 = v7 + int32(1072)
								return
							}
						}
					} else {
						v65 = F_unlink(m, v7+int32(48))
						mBase = m.M
						m.G0 = v7 + int32(1072)
						return
					}
				}
			}
		} else {
			*(*uint32)(unsafe.Add(mBase, uint32(v7)+36)) = uint32(v2)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v24
			v46 = F_pg_snprintf(m, v7+int32(48), int32(1024), int32(_a_F_SlruInternalDeleteSegment_4), v7+int32(32))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return
			} else {
				v50 = F_errstart(m, int32(13), int32(0))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					if v50 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 + int32(48)
						F_errmsg_internal(m, int32(_a_F_SlruInternalDeleteSegment_1), v7)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_SlruInternalDeleteSegment_2), int32(1568), int32(_a_F_SlruInternalDeleteSegment_3))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								v65 = F_unlink(m, v7+int32(48))
								mBase = m.M
								m.G0 = v7 + int32(1072)
								return
							}
						}
					} else {
						v65 = F_unlink(m, v7+int32(48))
						mBase = m.M
						m.G0 = v7 + int32(1072)
						return
					}
				}
			}
		}
	}
}
func F_SlruScanDirCbDeleteCutoff(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int64
	_ = v22
	var v24 int32
	_ = v24
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v8 = m.T0[v7].(func(*base.Module, int64, int64) int32)(m, l2, v6)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 == int32(0) {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v17 = m.T0[v16].(func(*base.Module, int64, int64) int32)(m, l2+int64(31), v6)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				if v17 == int32(0) {
					return int32(0)
				} else {
					v22 = base.I64_div_s(l2, int64(32))
					F_SlruInternalDeleteSegment(m, l0, v22)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				}
			}
		}
	}
}
