package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tts_buffer_heap_clear(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v31 int32
	_ = v31
	v4 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v4&int32(4) != 0 {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		F_pfree(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
			v12 = v10 & int32(-5)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v12)
			v14 = v12
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			if v15 != 0 {
				F_ReleaseBuffer(m, v15)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
					v19 = v18
					v20 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v20
					*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = int64(0)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)) = uint16(v20)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(-1)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v20)
					v31 = v19 | int32(2)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v31)
					return
				}
			} else {
				v19 = v14
				v20 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v20
				*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = int64(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)) = uint16(v20)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(-1)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v20)
				v31 = v19 | int32(2)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v31)
				return
			}
		}
	} else {
		v14 = v4
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		if v15 != 0 {
			F_ReleaseBuffer(m, v15)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v19 = v18
				v20 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v20
				*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = int64(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)) = uint16(v20)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(-1)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v20)
				v31 = v19 | int32(2)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v31)
				return
			}
		} else {
			v19 = v14
			v20 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v20
			*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = int64(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)) = uint16(v20)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(-1)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v20)
			v31 = v19 | int32(2)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v31)
			return
		}
	}
}
func F_tts_buffer_heap_copyslot(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
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
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v75 int32
	_ = v75
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v5 != v6 {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
		m.T0[v13].(func(*base.Module, int32))(m, l0)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
			v18 = v16 & int32(65533)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v18)
			v20 = int32(4562080)
			v21 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v23
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+44))
			v27 = m.T0[v26].(func(*base.Module, int32) int32)(m, l1)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v27
				v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v32 = v30 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v32)
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v21
				return
			}
		}
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
		if v8&int32(4) != 0 {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
			m.T0[v13].(func(*base.Module, int32))(m, l0)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v18 = v16 & int32(65533)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v18)
				v20 = int32(4562080)
				v21 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v23
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+44))
				v27 = m.T0[v26].(func(*base.Module, int32) int32)(m, l1)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v27
					v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
					v32 = v30 | int32(4)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v32)
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v21
					return
				}
			}
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
			if v11 != 0 {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
				v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				if v37&int32(4) != 0 {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					F_pfree(m, v40)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
						v46 = v43 & int32(-5)
						v47 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v47
						*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v11
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v47)
						v53 = v46 & int32(65533)
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v53)
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v55
						v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+8)))
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)) = uint16(v57)
						v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
						if v59 == v36 {
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							v69 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
							*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v69
							v71 = *(*int64)(unsafe.Add(mBase, uint32(v68)+8))
							*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v71
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
							*(*int32)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = v75
							*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l0 + int32(48)
							return
						} else {
							if v59 != 0 {
								F_ReleaseBuffer(m, v59)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v36
									if v36 == int32(0) {
										v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										v69 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v69
										v71 = *(*int64)(unsafe.Add(mBase, uint32(v68)+8))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v71
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
										*(*int32)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = v75
										*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l0 + int32(48)
										return
									} else {
										F_IncrBufferRefCount(m, v36)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return
										} else {
											v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
											v69 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
											*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v69
											v71 = *(*int64)(unsafe.Add(mBase, uint32(v68)+8))
											*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v71
											v75 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
											*(*int32)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = v75
											*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l0 + int32(48)
											return
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v36
								if v36 == int32(0) {
									v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									v69 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v69
									v71 = *(*int64)(unsafe.Add(mBase, uint32(v68)+8))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v71
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
									*(*int32)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = v75
									*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l0 + int32(48)
									return
								} else {
									F_IncrBufferRefCount(m, v36)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										v69 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v69
										v71 = *(*int64)(unsafe.Add(mBase, uint32(v68)+8))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v71
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
										*(*int32)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = v75
										*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l0 + int32(48)
										return
									}
								}
							}
						}
					}
				} else {
					v46 = v37
					v47 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v47
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v11
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v47)
					v53 = v46 & int32(65533)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v53)
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v55
					v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+8)))
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)) = uint16(v57)
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
					if v59 == v36 {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						v69 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v69
						v71 = *(*int64)(unsafe.Add(mBase, uint32(v68)+8))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v71
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
						*(*int32)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = v75
						*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l0 + int32(48)
						return
					} else {
						if v59 != 0 {
							F_ReleaseBuffer(m, v59)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v36
								if v36 == int32(0) {
									v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									v69 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v69
									v71 = *(*int64)(unsafe.Add(mBase, uint32(v68)+8))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v71
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
									*(*int32)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = v75
									*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l0 + int32(48)
									return
								} else {
									F_IncrBufferRefCount(m, v36)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										v69 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v69
										v71 = *(*int64)(unsafe.Add(mBase, uint32(v68)+8))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v71
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
										*(*int32)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = v75
										*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l0 + int32(48)
										return
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v36
							if v36 == int32(0) {
								v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								v69 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v69
								v71 = *(*int64)(unsafe.Add(mBase, uint32(v68)+8))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v71
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
								*(*int32)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = v75
								*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l0 + int32(48)
								return
							} else {
								F_IncrBufferRefCount(m, v36)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return
								} else {
									v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									v69 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v69
									v71 = *(*int64)(unsafe.Add(mBase, uint32(v68)+8))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v71
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
									*(*int32)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = v75
									*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l0 + int32(48)
									return
								}
							}
						}
					}
				}
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
				m.T0[v13].(func(*base.Module, int32))(m, l0)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
					v18 = v16 & int32(65533)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v18)
					v20 = int32(4562080)
					v21 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v23
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+44))
					v27 = m.T0[v26].(func(*base.Module, int32) int32)(m, l1)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v27
						v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
						v32 = v30 | int32(4)
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v32)
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v21
						return
					}
				}
			}
		}
	}
}
func F_tts_buffer_heap_materialize(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v4&int32(4) == int32(0) {
		v9 = int32(4562080)
		v10 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v12
		v14 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v14)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v14
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if v18 == v14 {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v24 = F_heap_form_tuple(m, v21, v22, v23)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v24
				v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v38 = v36 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v38)
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v10
				return
			}
		} else {
			v27 = F_heap_copytuple(m, v18)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v27
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
				if v30 != 0 {
					F_ReleaseBuffer(m, v30)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = int32(0)
						v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
						v38 = v36 | int32(4)
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v38)
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v10
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = int32(0)
					v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
					v38 = v36 | int32(4)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v38)
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v10
					return
				}
			}
		}
	} else {
		return
	}
}
func F_tts_heap_get_heap_tuple(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v4 != 0 {
		v32 = v4
		return v32
	} else {
		v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
		if v6&int32(4) != 0 {
			v32 = int32(0)
			return v32
		} else {
			v9 = int32(4562080)
			v10 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v12
			v14 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v14
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v14)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v21 = F_heap_form_tuple(m, v18, v19, v20)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v21
				v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v28 = v26 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v28)
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v10
				v32 = v21
				return v32
			}
		}
	}
}
func F_tts_minimal_clear(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v3&int32(4) != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		F_pfree(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
			v12 = v9 & int32(-5)
			v13 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v13
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)) = uint16(v13)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(-1)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v13)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v13
			v24 = v12 | int32(2)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v24)
			return
		}
	} else {
		v12 = v3
		v13 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v13
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)) = uint16(v13)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(-1)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v13)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v13
		v24 = v12 | int32(2)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v24)
		return
	}
}
func F_tts_minimal_get_minimal_tuple(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v5 != 0 {
		v41 = v5
		return v41
	} else {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
		if v7&int32(4) != 0 {
			v41 = int32(0)
			return v41
		} else {
			v10 = int32(4562080)
			v11 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v13
			v15 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v15
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v15)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v23 = F_heap_form_minimal_tuple(m, v19, v20, v21, v15)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v23
				v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v30 = v28 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v30)
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				v33 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v23 - v33
				*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v32 + v33
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v11
				v41 = v23
				return v41
			}
		}
	}
}
func F_tts_minimal_materialize(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v5&int32(4) == int32(0) {
		v10 = int32(4562080)
		v11 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v13
		v15 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v15
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v15)
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		if v19 == v15 {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v26 = F_heap_form_minimal_tuple(m, v22, v23, v24, int32(0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				v31 = v26
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v31
				v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v35 = v33 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v35)
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
				v38 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v31 - v38
				*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v37 + v38
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v11
				return
			}
		} else {
			v29 = F_heap_copy_minimal_tuple(m, v19, int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				v31 = v29
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v31
				v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v35 = v33 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v35)
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
				v38 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v31 - v38
				*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v37 + v38
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v11
				return
			}
		}
	} else {
		return
	}
}
