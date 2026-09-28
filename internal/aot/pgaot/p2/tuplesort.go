package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_tuplesort_begin_common(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
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
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	if l2&int32(1) != 0 {
		v11 = l1
	} else {
		v11 = int32(0)
	}
	if v11 == int32(0) {
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_begin_common[0]))
		v20 = F_AllocSetContextCreateInternal(m, v15, int32(_a_F_tuplesort_begin_common_0), int32(0), int32(_a_F_tuplesort_begin_common_1), int32(_a_F_tuplesort_begin_common_2))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v28 = F_AllocSetContextCreateInternal(m, v20, int32(_a_F_tuplesort_begin_common_3), int32(0), int32(_a_F_tuplesort_begin_common_1), int32(_a_F_tuplesort_begin_common_2))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = int32(_a_F_tuplesort_begin_common_4)
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_begin_common[0]))
				*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_begin_common[0])) = v20
				v35 = F_palloc0(m, int32(424))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_tuplesort_begin_common[1])))
					if v38 == int32(1) {
						F_getrusage(m, v35+int32(272))
						mBase = m.M
						F_gettimeofday(m, v35+int32(256))
						mBase = m.M
					} else {
					}
					*(*int64)(unsafe.Add(mBase, uint32(v35)+248)) = int64(10)
					v49 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v35)+56)) = uint8(v49)
					*(*int32)(unsafe.Add(mBase, uint32(v35)+52)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v28
					*(*int32)(unsafe.Add(mBase, uint32(v35)+140)) = int32(1024)
					*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v20
					*(*int32)(unsafe.Add(mBase, uint32(v35)+132)) = int32(0)
					v58 = int32(64)
					if l0 <= v58 {
						v61 = v58
					} else {
						v61 = l0
					}
					*(*int64)(unsafe.Add(mBase, uint32(v35)+96)) = base.I64_extend_i32_u(v61) << (uint(int64(10)) % 64)
					F_tuplesort_begin_batch(m, v35)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						if l1 == int32(0) {
							*(*int64)(unsafe.Add(mBase, uint32(v35)+232)) = int64(4294967295)
							v98 = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v35)+240)) = v98
							*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_begin_common[0])) = v31
							return v35
						} else {
							v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
							v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v35)+236)) = v74
							if v73 == int32(1) {
								v80 = base.AtomicRmwXchg32(m, v74, int32(0), int32(1))
								if v80 != 0 {
									F_s_lock(m, v74, int32(_a_F_tuplesort_begin_common_5))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return int32(0)
									} else {
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v84 + int32(1)
										v88 = int32(0)
										atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v74))), uint32(v88))
										*(*int32)(unsafe.Add(mBase, uint32(v35)+232)) = v84
										v98 = int32(-1)
										*(*int32)(unsafe.Add(mBase, uint32(v35)+240)) = v98
										*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_begin_common[0])) = v31
										return v35
									}
								} else {
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v84 + int32(1)
									v88 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v74))), uint32(v88))
									*(*int32)(unsafe.Add(mBase, uint32(v35)+232)) = v84
									v98 = int32(-1)
									*(*int32)(unsafe.Add(mBase, uint32(v35)+240)) = v98
									*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_begin_common[0])) = v31
									return v35
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v35)+232)) = int32(-1)
								v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								v98 = v95
								*(*int32)(unsafe.Add(mBase, uint32(v35)+240)) = v98
								*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_begin_common[0])) = v31
								return v35
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v106 = m.ExcPending
		if v106 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_tuplesort_begin_common_6), int32(0))
			mBase = m.M
			v110 = m.ExcPending
			if v110 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_tuplesort_begin_common_7), int32(556), int32(_a_F_tuplesort_begin_common_8))
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_tuplesort_begin_datum(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	v4 = l3
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = F_tuplesort_begin_common(m, l4, int32(0), l5)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = int32(_a_F_tuplesort_begin_datum_0)
		v21 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_begin_datum[0]))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
		*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_begin_datum[0])) = v23
		v26 = F_palloc(m, int32(8))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v29 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_tuplesort_begin_datum[1])))
			if v29 != int32(1) {
				*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(2069)
				v55 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v55
				*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v26
				*(*uint8)(unsafe.Add(mBase, uint32(v16)+36)) = uint8(v55)
				*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(2070)
				*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(2071)
				*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(2072)
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(2073)
				*(*int32)(unsafe.Add(mBase, uint32(v26))) = l0
				F_get_typlenbyval(m, l0, v13+int32(14), v13+int32(13))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+14)))
					*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v75
					v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)))
					v79 = v77 ^ int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v16)+56)) = uint8(v79)
					v82 = F_palloc0(m, int32(36))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v82
						v86 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_begin_datum[0]))
						*(*int32)(unsafe.Add(mBase, uint32(v82))) = v86
						v88 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
						*(*int32)(unsafe.Add(mBase, uint32(v88)+4)) = l2
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
						*(*uint8)(unsafe.Add(mBase, uint32(v90)+9)) = uint8(v4)
						v92 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
						v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)))
						v95 = v93 ^ int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v92)+20)) = uint8(v95)
						v97 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
						F_PrepareSortSupportFromOrderingOp(m, l1, v97)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
							v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+24))
							if v101 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v100
							} else {
							}
							*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_begin_datum[0])) = v21
							m.G0 = v13 + int32(16)
							return v16
						}
					}
				}
			} else {
				v34 = F_errstart(m, int32(15), int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					if v34 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(2069)
						v55 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v55
						*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v26
						*(*uint8)(unsafe.Add(mBase, uint32(v16)+36)) = uint8(v55)
						*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(2070)
						*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(2071)
						*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(2072)
						*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(2073)
						*(*int32)(unsafe.Add(mBase, uint32(v26))) = l0
						F_get_typlenbyval(m, l0, v13+int32(14), v13+int32(13))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+14)))
							*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v75
							v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)))
							v79 = v77 ^ int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v16)+56)) = uint8(v79)
							v82 = F_palloc0(m, int32(36))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v82
								v86 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_begin_datum[0]))
								*(*int32)(unsafe.Add(mBase, uint32(v82))) = v86
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
								*(*int32)(unsafe.Add(mBase, uint32(v88)+4)) = l2
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
								*(*uint8)(unsafe.Add(mBase, uint32(v90)+9)) = uint8(v4)
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
								v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)))
								v95 = v93 ^ int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v92)+20)) = uint8(v95)
								v97 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
								F_PrepareSortSupportFromOrderingOp(m, l1, v97)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
									v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+24))
									if v101 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v100
									} else {
									}
									*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_begin_datum[0])) = v21
									m.G0 = v13 + int32(16)
									return v16
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13))) = l4
						if l5&int32(1) != 0 {
							v43 = int32(116)
						} else {
							v43 = int32(102)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v43
						F_errmsg_internal(m, int32(_a_F_tuplesort_begin_datum_1), v13)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_tuplesort_begin_datum_2), int32(687), int32(_a_F_tuplesort_begin_datum_3))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(2069)
								v55 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v55
								*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v26
								*(*uint8)(unsafe.Add(mBase, uint32(v16)+36)) = uint8(v55)
								*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(2070)
								*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(2071)
								*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(2072)
								*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(2073)
								*(*int32)(unsafe.Add(mBase, uint32(v26))) = l0
								F_get_typlenbyval(m, l0, v13+int32(14), v13+int32(13))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+14)))
									*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v75
									v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)))
									v79 = v77 ^ int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v16)+56)) = uint8(v79)
									v82 = F_palloc0(m, int32(36))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v82
										v86 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_begin_datum[0]))
										*(*int32)(unsafe.Add(mBase, uint32(v82))) = v86
										v88 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
										*(*int32)(unsafe.Add(mBase, uint32(v88)+4)) = l2
										v90 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
										*(*uint8)(unsafe.Add(mBase, uint32(v90)+9)) = uint8(v4)
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
										v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)))
										v95 = v93 ^ int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v92)+20)) = uint8(v95)
										v97 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
										F_PrepareSortSupportFromOrderingOp(m, l1, v97)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return int32(0)
										} else {
											v100 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
											v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+24))
											if v101 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v100
											} else {
											}
											*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_begin_datum[0])) = v21
											m.G0 = v13 + int32(16)
											return v16
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
func F_tuplesort_begin_index_gin(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v121 int32
	_ = v121
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = F_tuplesort_begin_common(m, l1, l2, int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = int32(_a_F_tuplesort_begin_index_gin_0)
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_begin_index_gin[0]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_begin_index_gin[0])) = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v27 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+10)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v27
	v31 = F_palloc0(m, v27*int32(36))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v31
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
	if int32(0) < v34 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L21
	}
L5:
	;
	v40 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	v121 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v121
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+36)) = uint8(v121)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(2065)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(2066)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(2067)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(2068)
	*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_begin_index_gin[0])) = v21
	m.G0 = v13 + int32(16)
	return v16
L8:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v52 = v49 + v40*int32(36)
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_begin_index_gin[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56+v40<<(uint(int32(2))%32))))
	v61 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+20)) = uint8(v61)
	v64 = v40 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v52)+10)) = uint16(v64)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+9)) = uint8(v61)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v60
	if v60 == v61 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = int32(100)
	goto L12
L11:
	;
	goto L12
L12:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76)+6)))
	v85 = int32(4)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v75+v77*(base.I32_extend16_s(v64)-int32(1))<<(uint(int32(2))%32)+v85-v85)))
	goto L13
L13:
	;
	if v89 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v97 = v22 + v48<<(uint(int32(3))%32) + v40*int32(100)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+96))
	v100 = F_lookup_type_cache(m, v98, int32(64))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	v106 = v89
	goto L16
L16:
	;
	F_PrepareSortSupportComparisonShim(m, v106, v52)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v100)+108))
	if v102 == int32(0) {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v106 = v102
	goto L16
L19:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
	if v64 < v109 {
		v40 = v64
		goto L8
	} else {
		goto L20
	}
L20:
	;
	goto L9
L21:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v97)+96))
	v147 = F_format_type_be(m, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v147
	F_errmsg(m, int32(_a_F_tuplesort_begin_index_gin_1), v13)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_tuplesort_begin_index_gin_2), int32(648), int32(_a_F_tuplesort_begin_index_gin_3))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tuplesort_getbrintuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(_a_F_tuplesort_getbrintuple_0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getbrintuple[0]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getbrintuple[0])) = v13
	v18 = F_tuplesort_gettuple_common(m, l0, int32(1), v8+int32(8))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		if v18 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getbrintuple[0])) = v11
			v36 = v3
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getbrintuple[0])) = v11
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
			if v28 == int32(0) {
				v36 = v3
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v31
				v36 = v28 + int32(4)
			}
		}
		m.G0 = v8 + int32(32)
		return v36
	}
}
func F_tuplesort_method_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	if base.Ui32(l0) <= base.Ui32(int32(8)) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_tuplesort_method_name[0])))
		v8 = v6
	} else {
		v8 = int32(_a_F_tuplesort_method_name_0)
	}
	return v8
}
func F_tuplesort_puttuple_common(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v30 int64
	_ = v30
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v91 int32
	_ = v91
	var v97 float64
	_ = v97
	var v98 float64
	_ = v98
	var v101 float64
	_ = v101
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int64
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int64
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int64
	_ = v203
	var v205 int64
	_ = v205
	var v207 int64
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v232 int32
	_ = v232
	var v233 int64
	_ = v233
	var v235 int64
	_ = v235
	var v237 int64
	_ = v237
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int64
	_ = v260
	var v262 int64
	_ = v262
	var v264 int64
	_ = v264
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int64
	_ = v285
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int64
	_ = v297
	var v299 int64
	_ = v299
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int64
	_ = v335
	var v337 int64
	_ = v337
	var v339 int64
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v350 int64
	_ = v350
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v432 int32
	_ = v432
	var v445 int32
	_ = v445
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int64
	_ = v459
	var v461 int64
	_ = v461
	var v463 int64
	_ = v463
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v477 int32
	_ = v477
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v509 int64
	_ = v509
	var v511 int64
	_ = v511
	var v513 int64
	_ = v513
	var v517 int32
	_ = v517
	var v534 int32
	_ = v534
	var v535 int64
	_ = v535
	var v537 int64
	_ = v537
	var v539 int64
	_ = v539
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int64
	_ = v555
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v625 int64
	_ = v625
	var v627 int64
	_ = v627
	var v629 int64
	_ = v629
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v639 int32
	_ = v639
	var v654 int32
	_ = v654
	var v655 int64
	_ = v655
	var v657 int64
	_ = v657
	var v659 int64
	_ = v659
	var v678 int32
	_ = v678
	var v699 int32
	_ = v699
	var v701 int64
	_ = v701
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	v5 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v21 = int32(_a_F_tuplesort_puttuple_common_0)
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_puttuple_common[0]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_puttuple_common[0])) = v24
	v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	v27 = base.I64_extend_i32_u(l3)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v26 - v27
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)+80))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = v30 + v27
	if l2 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	switch v74 {
	case 0:
		goto L20
	case 1:
		goto L19
	case 2:
		goto L17
	default:
		goto L18
	}
L2:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v35 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+16)) = v56
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v59 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+28)) = v59
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+32)) = v59
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	m.T0[v69].(func(*base.Module, int32, int32, int32))(m, l0, v67, v68)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L7
	} else {
		goto L11
	}
L4:
	;
	v49 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+24))
	v52 = m.T0[v51].(func(*base.Module, int64, int32) int64)(m, v49, v50)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L7
	} else {
		goto L10
	}
L5:
	;
	v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+248))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if base.I64_extend_i32_s(v37) < v36 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v36 << (uint(int64(1)) % 64)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+28))
	v45 = m.T0[v44].(func(*base.Module, int32, int32) int32)(m, v37, v43)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	if v45 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	goto L4
L10:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v52
	goto L1
L11:
	;
	goto L1
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_puttuple_common[0])) = v22
	m.G0 = v19 + int32(32)
	return
L13:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v328 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v327 + v328
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v334 = v331 + v327*int32(24)
	v335 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v334)+16)) = v335
	v337 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v334)+8)) = v337
	v339 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v334))) = v339
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+68)))
	if v342 != v328 {
		goto L86
	} else {
		goto L87
	}
L14:
	;
	v319 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+144)) = uint8(v319)
	goto L13
L15:
	;
	if v84 < base.I64_extend_i32_u((v274-v76)*int32(24)) {
		goto L14
	} else {
		goto L77
	}
L16:
	;
	v271 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+144)) = uint8(v271)
	v274 = int32(89478485)
	goto L15
L17:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v252 + int32(1)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v259 = v256 + v252*int32(24)
	v260 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v259)+16)) = v260
	v262 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v259)+8)) = v262
	v264 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v259))) = v264
	F_dumptuples(m, l0, int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L7
	} else {
		goto L76
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L7
	} else {
		goto L73
	}
L19:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v112 = m.T0[v111].(func(*base.Module, int32, int32, int32) int32)(m, l1, v110, l0)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L7
	} else {
		goto L35
	}
L20:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v75 < v76-int32(1) {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+144)))
	if v80 != int32(1) {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	v83 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
	v84 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	v85 = v83 - v84
	if v85 <= v84 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	if v76 != int32(2147483647) {
		goto L16
	} else {
		goto L34
	}
L24:
	;
	if v104 <= v76 {
		goto L14
	} else {
		goto L32
	}
L25:
	;
	if int32(1073741822) < v76 {
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v91 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+144)) = uint8(v91)
	v97 = base.F64_mul(base.F64_div(base.F64_convert_i64_s(v83), base.F64_convert_i64_s(v85)), base.F64_convert_i32_s(v76))
	v98 = float64(2.147483647e+09)
	if base.F64_lt(v97, v98) != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v104 = v76 << (uint(int32(1)) % 32)
	goto L24
L29:
	;
	v101 = v97
	goto L31
L30:
	;
	v101 = v98
	goto L31
L31:
	;
	v104 = base.I32_trunc_sat_f64_s(v101)
	goto L24
L32:
	;
	if base.Ui32(v104) < base.Ui32(int32(89478485)) {
		v274 = v104
		goto L15
	} else {
		goto L33
	}
L33:
	;
	goto L16
L34:
	;
	goto L14
L35:
	;
	if v112 <= int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v116 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	if v136 != 0 {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	v117 = F_GetMemoryChunkSpace(m, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L7
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_puttuple_common[1]))
	if v130 == int32(0) {
		goto L12
	} else {
		goto L44
	}
L42:
	;
	v119 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v119 + base.I64_extend_i32_u(v117)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_pfree(m, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	goto L41
L44:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L7
	} else {
		goto L45
	}
L45:
	;
	goto L12
L46:
	;
	v137 = F_GetMemoryChunkSpace(m, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L7
	} else {
		goto L49
	}
L47:
	;
	v150 = v135
	goto L48
L48:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_puttuple_common[1]))
	if v152 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v139 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v139 + base.I64_extend_i32_u(v137)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	F_pfree(m, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L7
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = int32(0)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v150 = v148
	goto L48
L51:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L7
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if base.Ui32(v155) < base.Ui32(int32(2)) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L53
L55:
	;
	v232 = v150 + v217*int32(24)
	v233 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v232)+16)) = v233
	v235 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v232)+8)) = v235
	v237 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v232))) = v237
	goto L12
L56:
	;
	v217 = int32(0)
	goto L55
L57:
	;
	goto L58
L58:
	;
	v162 = int32(1)
	v167 = v5
	v168 = v5
	goto L59
L59:
	;
	v177 = v168 + int32(2)
	if base.Ui32(v155) <= base.Ui32(v177) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v217 = v191
	goto L55
L61:
	;
	v191 = v162
	goto L63
L62:
	;
	v179 = int32(24)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v186 = m.T0[v185].(func(*base.Module, int32, int32, int32) int32)(m, v150+v162*v179, v150+v177*v179, l0)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L7
	} else {
		goto L64
	}
L63:
	;
	v194 = v150 + v191*int32(24)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v196 = m.T0[v195].(func(*base.Module, int32, int32, int32) int32)(m, l1, v194, l0)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L7
	} else {
		goto L68
	}
L64:
	;
	if int32(0) < v186 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v190 = v177
	goto L67
L66:
	;
	v190 = v162
	goto L67
L67:
	;
	v191 = v190
	goto L63
L68:
	;
	if v196 <= int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v217 = v167
	goto L55
L70:
	;
	goto L71
L71:
	;
	v202 = v150 + v167*int32(24)
	v203 = *(*int64)(unsafe.Add(mBase, uint32(v194)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v202)+16)) = v203
	v205 = *(*int64)(unsafe.Add(mBase, uint32(v194)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v202)+8)) = v205
	v207 = *(*int64)(unsafe.Add(mBase, uint32(v194)))
	*(*int64)(unsafe.Add(mBase, uint32(v202))) = v207
	v209 = int32(1)
	v210 = v191 << (uint(v209) % 32)
	v212 = v210 | v209
	if base.Ui32(v212) < base.Ui32(v155) {
		v162 = v212
		v167 = v191
		v168 = v210
		goto L59
	} else {
		goto L72
	}
L72:
	;
	goto L60
L73:
	;
	F_errmsg_internal(m, int32(_a_F_tuplesort_puttuple_common_1), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_tuplesort_puttuple_common_2), int32(1209), int32(_a_F_tuplesort_puttuple_common_3))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L7
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	goto L12
L77:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v282 = F_GetMemoryChunkSpace(m, v281)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L7
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v274
	v285 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v285 + base.I64_extend_i32_u(v282)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v292 = F_repalloc_huge(m, v289, v274*int32(24))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L7
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v292
	v295 = F_GetMemoryChunkSpace(m, v292)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L7
	} else {
		goto L80
	}
L80:
	;
	v297 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	v299 = v297 - base.I64_extend_i32_u(v295)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v299
	if int64(0) <= v299 {
		goto L13
	} else {
		goto L81
	}
L81:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+145)))
	if v303 != 0 {
		goto L13
	} else {
		goto L82
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L7
	} else {
		goto L83
	}
L83:
	;
	F_errmsg_internal(m, int32(_a_F_tuplesort_puttuple_common_4), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L7
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_tuplesort_puttuple_common_2), int32(1053), int32(_a_F_tuplesort_puttuple_common_5))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L7
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v341 < v699 {
		goto L159
	} else {
		goto L160
	}
L87:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v341 <= v345<<(uint(int32(1))%32) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	if v341 <= v345 {
		goto L86
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v355 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_tuplesort_puttuple_common[2])))
	if v355 != int32(1) {
		v380 = v341
		goto L94
	} else {
		goto L95
	}
L91:
	;
	v350 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	if int64(0) <= v350 {
		goto L86
	} else {
		goto L92
	}
L92:
	;
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+145)))
	if v353 != 0 {
		goto L86
	} else {
		goto L93
	}
L93:
	;
	goto L90
L94:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if int32(0) < v381 {
		goto L101
	} else {
		goto L102
	}
L95:
	;
	v360 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L7
	} else {
		goto L96
	}
L96:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v360 == int32(0) {
		v380 = v362
		goto L94
	} else {
		goto L97
	}
L97:
	;
	v367 = F_pg_rusage_show(m, l0+int32(256))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L7
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v367
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v362
	F_errmsg_internal(m, int32(_a_F_tuplesort_puttuple_common_6), v19)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L7
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_tuplesort_puttuple_common_2), int32(1145), int32(_a_F_tuplesort_puttuple_common_3))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L7
	} else {
		goto L100
	}
L100:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v380 = v379
	goto L94
L101:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v387 = int32(0)
	v389 = v384
	goto L104
L102:
	;
	goto L103
L103:
	;
	v432 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v432
	if v432 < v380 {
		goto L107
	} else {
		goto L108
	}
L104:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+8)))
	v403 = int32(1)
	v404 = v402 ^ v403
	*(*uint8)(unsafe.Add(mBase, uint32(v389)+8)) = uint8(v404)
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+9)))
	v408 = v406 ^ v403
	*(*uint8)(unsafe.Add(mBase, uint32(v389)+9)) = uint8(v408)
	v413 = v387 + v403
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v413 < v414 {
		v387 = v413
		v389 = v389 + int32(36)
		goto L104
	} else {
		goto L106
	}
L105:
	;
	goto L103
L106:
	;
	goto L105
L107:
	;
	v445 = v5
	goto L110
L108:
	;
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(1)
	goto L12
L110:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v452 < v453 {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L109
L112:
	;
	v678 = v445 + int32(1)
	if v678 != v380 {
		v445 = v678
		goto L110
	} else {
		goto L158
	}
L113:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v458 = v455 + v445*int32(24)
	v459 = *(*int64)(unsafe.Add(mBase, uint32(v458)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v459
	v461 = *(*int64)(unsafe.Add(mBase, uint32(v458)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v461
	v463 = *(*int64)(unsafe.Add(mBase, uint32(v458)))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v463
	v466 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_puttuple_common[1]))
	if v466 != 0 {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	goto L115
L115:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v543 = v445 * int32(24)
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v546 = m.T0[v545].(func(*base.Module, int32, int32, int32) int32)(m, v541+v543, v541, l0)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L7
	} else {
		goto L127
	}
L116:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L7
	} else {
		goto L119
	}
L117:
	;
	v470 = v452
	goto L118
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v470 + int32(1)
	if v470 <= int32(0) {
		v517 = v470
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v470 = v469
	goto L118
L120:
	;
	v534 = v455 + v517*int32(24)
	v535 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v534)+16)) = v535
	v537 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v534)+8)) = v537
	v539 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v534))) = v539
	goto L112
L121:
	;
	v477 = v470
	goto L122
L122:
	;
	v494 = int32(1)
	v497 = int32(base.Ui32(v477-v494) >> (uint(v494) % 32))
	v500 = v455 + v497*int32(24)
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v502 = m.T0[v501].(func(*base.Module, int32, int32, int32) int32)(m, v19+int32(8), v500, l0)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L7
	} else {
		goto L124
	}
L123:
	;
	v517 = int32(0)
	goto L120
L124:
	;
	if int32(0) <= v502 {
		v517 = v477
		goto L120
	} else {
		goto L125
	}
L125:
	;
	v508 = v455 + v477*int32(24)
	v509 = *(*int64)(unsafe.Add(mBase, uint32(v500)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v508)+16)) = v509
	v511 = *(*int64)(unsafe.Add(mBase, uint32(v500)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v508)+8)) = v511
	v513 = *(*int64)(unsafe.Add(mBase, uint32(v500)))
	*(*int64)(unsafe.Add(mBase, uint32(v508))) = v513
	if v497 != 0 {
		v477 = v497
		goto L122
	} else {
		goto L126
	}
L126:
	;
	goto L123
L127:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v549 = v548 + v543
	if v546 <= int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v549)))
	if v552 != 0 {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	goto L130
L130:
	;
	v572 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_puttuple_common[1]))
	if v572 != 0 {
		goto L138
	} else {
		goto L139
	}
L131:
	;
	v553 = F_GetMemoryChunkSpace(m, v552)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L7
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v566 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_puttuple_common[1]))
	if v566 == int32(0) {
		goto L112
	} else {
		goto L136
	}
L134:
	;
	v555 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v555 + base.I64_extend_i32_u(v553)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v549)))
	F_pfree(m, v559)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L7
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v549))) = int32(0)
	goto L133
L136:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L7
	} else {
		goto L137
	}
L137:
	;
	goto L112
L138:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L7
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v575 = int32(0)
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if base.Ui32(v579) < base.Ui32(int32(2)) {
		v639 = v575
		goto L142
	} else {
		goto L143
	}
L141:
	;
	goto L140
L142:
	;
	v654 = v548 + v639*int32(24)
	v655 = *(*int64)(unsafe.Add(mBase, uint32(v549)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v654)+16)) = v655
	v657 = *(*int64)(unsafe.Add(mBase, uint32(v549)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v654)+8)) = v657
	v659 = *(*int64)(unsafe.Add(mBase, uint32(v549)))
	*(*int64)(unsafe.Add(mBase, uint32(v654))) = v659
	goto L112
L143:
	;
	v583 = int32(1)
	v586 = v575
	v589 = v575
	goto L144
L144:
	;
	v599 = v589 + int32(2)
	if base.Ui32(v579) <= base.Ui32(v599) {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v639 = v613
	goto L142
L146:
	;
	v613 = v583
	goto L148
L147:
	;
	v601 = int32(24)
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v608 = m.T0[v607].(func(*base.Module, int32, int32, int32) int32)(m, v548+v583*v601, v548+v599*v601, l0)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L7
	} else {
		goto L149
	}
L148:
	;
	v616 = v548 + v613*int32(24)
	v617 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v618 = m.T0[v617].(func(*base.Module, int32, int32, int32) int32)(m, v549, v616, l0)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L7
	} else {
		goto L153
	}
L149:
	;
	if int32(0) < v608 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v612 = v599
	goto L152
L151:
	;
	v612 = v583
	goto L152
L152:
	;
	v613 = v612
	goto L148
L153:
	;
	if v618 <= int32(0) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v639 = v586
	goto L142
L155:
	;
	goto L156
L156:
	;
	v624 = v548 + v586*int32(24)
	v625 = *(*int64)(unsafe.Add(mBase, uint32(v616)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v624)+16)) = v625
	v627 = *(*int64)(unsafe.Add(mBase, uint32(v616)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v624)+8)) = v627
	v629 = *(*int64)(unsafe.Add(mBase, uint32(v616)))
	*(*int64)(unsafe.Add(mBase, uint32(v624))) = v629
	v631 = int32(1)
	v632 = v613 << (uint(v631) % 32)
	v634 = v632 | v631
	if base.Ui32(v634) < base.Ui32(v579) {
		v583 = v634
		v586 = v613
		v589 = v632
		goto L144
	} else {
		goto L157
	}
L157:
	;
	goto L145
L158:
	;
	goto L111
L159:
	;
	v701 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	if int64(0) <= v701 {
		goto L12
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	F_inittapes(m, l0, int32(1))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L7
	} else {
		goto L164
	}
L162:
	;
	v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+145)))
	if v704 != 0 {
		goto L12
	} else {
		goto L163
	}
L163:
	;
	goto L161
L164:
	;
	F_dumptuples(m, l0, int32(0))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L7
	} else {
		goto L165
	}
L165:
	;
	goto L12
}
func F_tuplesort_readtup_alloc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	if base.Ui32(l1) <= base.Ui32(int32(1024)) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
		if v6 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v14
			return v6
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v9 = F_MemoryContextAlloc(m, v8, l1)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v9
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v9 = F_MemoryContextAlloc(m, v8, l1)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v9
		}
	}
}
func F_tuplesort_sort_memtuples(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v99 int64
	_ = v99
	var v101 int64
	_ = v101
	var v103 int64
	_ = v103
	var v105 int64
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v127 int64
	_ = v127
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v143 int32
	_ = v143
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var __phi178 int32
	_ = __phi178
	var v180 int32
	_ = v180
	var __phi180 int32
	_ = __phi180
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v16 < int32(2) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(32)
	return
L2:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v19 != int32(1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_qsort_tuple(m, v169, v159, v212, l0)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L15
	} else {
		goto L62
	}
L4:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v206 != 0 {
		goto L57
	} else {
		goto L58
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if base.B2i32(v22 == int32(0))|base.B2i32(base.Ui32(v16) < base.Ui32(int32(40))) != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	if base.B2i32(base.B2i32(v28 == int32(199))|base.B2i32(v28 == int32(118)) == int32(0))&base.B2i32(v28 != int32(202)) != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+9)))
	v43 = int32(0)
	goto L9
L8:
	;
	v67 = v16 - int32(1)
	if base.Ui32(v65) < base.Ui32(v67) {
		goto L18
	} else {
		goto L19
	}
L9:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v43*int32(24))+16)))
	if v56 != v40 {
		v65 = v43
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v65 = v16
	goto L8
L11:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_sort_memtuples[0]))
	if v59 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v63 = v43 + int32(1)
	if v63 != v16 {
		v43 = v63
		goto L9
	} else {
		goto L17
	}
L15:
	;
	return
L16:
	;
	goto L14
L17:
	;
	goto L10
L18:
	;
	v71 = v39 + v65*int32(24)
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v71)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v72
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v71)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v74
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v71)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v76
	v79 = v65
	v80 = v65
	goto L21
L19:
	;
	v143 = v65
	goto L20
L20:
	;
	v155 = v39 + v143*int32(24)
	v156 = v16 - v143
	v158 = v40 & int32(1)
	if v158 != 0 {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	v89 = int32(24)
	v91 = v39 + v79*v89
	v94 = v39 + v80*v89
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v94)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v91)+16)) = v95
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v94)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v91)+8)) = v97
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v94)))
	*(*int64)(unsafe.Add(mBase, uint32(v91))) = v99
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v91)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v94))) = v101
	v103 = *(*int64)(unsafe.Add(mBase, uint32(v91)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v94)+8)) = v103
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v91)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v94)+16)) = v105
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_sort_memtuples[0]))
	if v108 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v119 = int32(24)
	v121 = v39 + v67*v119
	v124 = v39 + v115*v119
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v124)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v121)+16)) = v125
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v124)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v121)+8)) = v127
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v124)))
	*(*int64)(unsafe.Add(mBase, uint32(v121))) = v129
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v124))) = v131
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v124)+8)) = v133
	v135 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v124)+16)) = v135
	v143 = v115 + base.B2i32(base.I32_wrap_i64(v135)&int32(255) == v40)
	goto L20
L23:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L15
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v115 = base.B2i32(base.I32_wrap_i64(v105)&int32(255) == v40) + v80
	v117 = v79 + int32(1)
	if v117 != v67 {
		v79 = v117
		v80 = v115
		goto L21
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	goto L22
L28:
	;
	v159 = v156
	goto L30
L29:
	;
	v159 = v143
	goto L30
L30:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v160 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v158 != 0 {
		goto L41
	} else {
		goto L42
	}
L32:
	;
	if v158 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v161 = v143
	goto L35
L34:
	;
	v161 = v156
	goto L35
L35:
	;
	if base.Ui32(v161) < base.Ui32(int32(2)) {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	if v158 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v164 = v39
	goto L39
L38:
	;
	v164 = v155
	goto L39
L39:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_qsort_tuple(m, v164, v161, v165, l0)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L15
	} else {
		goto L40
	}
L40:
	;
	goto L31
L41:
	;
	v169 = v155
	goto L43
L42:
	;
	v169 = v39
	goto L43
L43:
	;
	if base.Ui32(v159) < base.Ui32(int32(40)) {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	v172 = int32(24)
	__phi178 = v169 + v172
	__phi180 = v169
	v178 = __phi178
	v180 = __phi180
	goto L45
L45:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v189 = m.T0[v188].(func(*base.Module, int32, int32, int32) int32)(m, v180, v178, l0)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L15
	} else {
		goto L47
	}
L46:
	;
	F_radix_sort_recursive(m, v169, v159, int32(0), l0)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L15
	} else {
		goto L56
	}
L47:
	;
	if v189 <= int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_sort_memtuples[0]))
	if v194 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	goto L46
L51:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L15
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v198 = v178 + int32(24)
	if base.Ui32(v198) < base.Ui32(v169+v159*v172) {
		__phi178 = v198
		__phi180 = v178
		v178 = __phi178
		v180 = __phi180
		goto L45
	} else {
		goto L55
	}
L54:
	;
	goto L53
L55:
	;
	goto L1
L56:
	;
	goto L1
L57:
	;
	F_qsort_ssup(m, v205, v16, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L15
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_qsort_tuple(m, v205, v16, v209, l0)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L15
	} else {
		goto L61
	}
L60:
	;
	goto L1
L61:
	;
	goto L1
L62:
	;
	goto L1
}
func F_tuplesort_space_type_name(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	if l0 != 0 {
		v4 = int32(_a_F_tuplesort_space_type_name_0)
	} else {
		v4 = int32(_a_F_tuplesort_space_type_name_1)
	}
	return v4
}
