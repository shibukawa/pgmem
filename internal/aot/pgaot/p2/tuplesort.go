package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
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
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
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
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
								*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(1)
								if v78 != 0 {
									F_s_lock(m, v74, int32(_a_F_tuplesort_begin_common_5), int32(2988), int32(_a_F_tuplesort_begin_common_6))
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
										v88 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v88 + int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v35)+232)) = v88
										v98 = int32(-1)
										*(*int32)(unsafe.Add(mBase, uint32(v35)+240)) = v98
										*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_begin_common[0])) = v31
										return v35
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
									v88 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v88 + int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v35)+232)) = v88
									v98 = int32(-1)
									*(*int32)(unsafe.Add(mBase, uint32(v35)+240)) = v98
									*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_begin_common[0])) = v31
									return v35
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v35)+232)) = int32(-1)
								v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								v98 = v96
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
			F_errmsg_internal(m, int32(_a_F_tuplesort_begin_common_7), int32(0))
			mBase = m.M
			v110 = m.ExcPending
			if v110 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_tuplesort_begin_common_5), int32(651), int32(_a_F_tuplesort_begin_common_8))
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
				*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(1846)
				v55 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v55
				*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v26
				*(*uint8)(unsafe.Add(mBase, uint32(v16)+36)) = uint8(v55)
				*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(1847)
				*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(1848)
				*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1849)
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(1850)
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
						*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(1846)
						v55 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v55
						*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v26
						*(*uint8)(unsafe.Add(mBase, uint32(v16)+36)) = uint8(v55)
						*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(1847)
						*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(1848)
						*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1849)
						*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(1850)
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
							F_errfinish(m, int32(_a_F_tuplesort_begin_datum_2), int32(686), int32(_a_F_tuplesort_begin_datum_3))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(1846)
								v55 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v55
								*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v26
								*(*uint8)(unsafe.Add(mBase, uint32(v16)+36)) = uint8(v55)
								*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(1847)
								*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(1848)
								*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1849)
								*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(1850)
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(1842)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(1843)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(1844)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(1845)
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
	v97 = v22 + v48<<(uint(int32(4))%32) + v40*int32(100)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+88))
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
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v97)+88))
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
	F_errfinish(m, int32(_a_F_tuplesort_begin_index_gin_2), int32(647), int32(_a_F_tuplesort_begin_index_gin_3))
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(_a_F_tuplesort_getbrintuple_0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getbrintuple[0]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getbrintuple[0])) = v13
	v16 = F_tuplesort_gettuple_common(m, l0, int32(1), v8)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		if v16 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getbrintuple[0])) = v11
			v34 = v3
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getbrintuple[0])) = v11
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			if v26 == int32(0) {
				v34 = v3
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v29
				v34 = v26 + int32(4)
			}
		}
		m.G0 = v8 + int32(16)
		return v34
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
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
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v164 int32
	_ = v164
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
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v230 int32
	_ = v230
	var v231 int64
	_ = v231
	var v233 int64
	_ = v233
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int64
	_ = v256
	var v258 int64
	_ = v258
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int64
	_ = v279
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int64
	_ = v291
	var v293 int64
	_ = v293
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int64
	_ = v329
	var v331 int64
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v342 int64
	_ = v342
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v424 int32
	_ = v424
	var v437 int32
	_ = v437
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v451 int64
	_ = v451
	var v453 int64
	_ = v453
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v469 int32
	_ = v469
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v499 int64
	_ = v499
	var v501 int64
	_ = v501
	var v507 int32
	_ = v507
	var v522 int32
	_ = v522
	var v523 int64
	_ = v523
	var v525 int64
	_ = v525
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int64
	_ = v541
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v611 int64
	_ = v611
	var v613 int64
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v638 int32
	_ = v638
	var v639 int64
	_ = v639
	var v641 int64
	_ = v641
	var v660 int32
	_ = v660
	var v681 int32
	_ = v681
	var v683 int64
	_ = v683
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
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
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+24))
	v52 = m.T0[v51].(func(*base.Module, int32, int32) int32)(m, v49, v50)
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
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v52
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
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v322 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v321 + v322
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v328 = v325 + v321<<(uint(int32(4))%32)
	v329 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v328)+8)) = v329
	v331 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v328))) = v331
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+68)))
	if v334 != v322 {
		goto L86
	} else {
		goto L87
	}
L14:
	;
	v313 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+144)) = uint8(v313)
	goto L13
L15:
	;
	if v84 < base.I64_extend_i32_u((v268-v76)<<(uint(int32(4))%32)) {
		goto L14
	} else {
		goto L77
	}
L16:
	;
	v265 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+144)) = uint8(v265)
	v268 = int32(134217727)
	goto L15
L17:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v248 + int32(1)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v255 = v252 + v248<<(uint(int32(4))%32)
	v256 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v255)+8)) = v256
	v258 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v255))) = v258
	F_dumptuples(m, l0, int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L7
	} else {
		goto L76
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
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
	if base.Ui32(v104) < base.Ui32(int32(134217727)) {
		v268 = v104
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
	v149 = v135
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
	v149 = v148
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
	v230 = v149 + v215<<(uint(int32(4))%32)
	v231 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v230)+8)) = v231
	v233 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v230))) = v233
	goto L12
L56:
	;
	v215 = int32(0)
	goto L55
L57:
	;
	goto L58
L58:
	;
	v164 = int32(1)
	v167 = v5
	v168 = v5
	goto L59
L59:
	;
	v177 = v167 + int32(2)
	if base.Ui32(v155) <= base.Ui32(v177) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v215 = v191
	goto L55
L61:
	;
	v191 = v164
	goto L63
L62:
	;
	v179 = int32(4)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v186 = m.T0[v185].(func(*base.Module, int32, int32, int32) int32)(m, v149+v164<<(uint(v179)%32), v149+v177<<(uint(v179)%32), l0)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L7
	} else {
		goto L64
	}
L63:
	;
	v194 = v149 + v191<<(uint(int32(4))%32)
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
	v190 = v164
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
	v215 = v168
	goto L55
L70:
	;
	goto L71
L71:
	;
	v202 = v149 + v168<<(uint(int32(4))%32)
	v203 = *(*int64)(unsafe.Add(mBase, uint32(v194)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v202)+8)) = v203
	v205 = *(*int64)(unsafe.Add(mBase, uint32(v194)))
	*(*int64)(unsafe.Add(mBase, uint32(v202))) = v205
	v207 = int32(1)
	v208 = v191 << (uint(v207) % 32)
	v210 = v208 | v207
	if base.Ui32(v210) < base.Ui32(v155) {
		v164 = v210
		v167 = v208
		v168 = v191
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
	v242 = m.ExcPending
	if v242 != 0 {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_tuplesort_puttuple_common_2), int32(1312), int32(_a_F_tuplesort_puttuple_common_3))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
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
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v276 = F_GetMemoryChunkSpace(m, v275)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L7
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v268
	v279 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v279 + base.I64_extend_i32_u(v276)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v286 = F_repalloc_huge(m, v283, v268<<(uint(int32(4))%32))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L7
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v286
	v289 = F_GetMemoryChunkSpace(m, v286)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L7
	} else {
		goto L80
	}
L80:
	;
	v291 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	v293 = v291 - base.I64_extend_i32_u(v289)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v293
	if int64(0) <= v293 {
		goto L13
	} else {
		goto L81
	}
L81:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+145)))
	if v297 != 0 {
		goto L13
	} else {
		goto L82
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L7
	} else {
		goto L83
	}
L83:
	;
	F_errmsg_internal(m, int32(_a_F_tuplesort_puttuple_common_4), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L7
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_tuplesort_puttuple_common_2), int32(1156), int32(_a_F_tuplesort_puttuple_common_5))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
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
	v681 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v333 < v681 {
		goto L159
	} else {
		goto L160
	}
L87:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v333 <= v337<<(uint(int32(1))%32) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	if v333 <= v337 {
		goto L86
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_tuplesort_puttuple_common[2])))
	if v347 != int32(1) {
		v372 = v333
		goto L94
	} else {
		goto L95
	}
L91:
	;
	v342 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	if int64(0) <= v342 {
		goto L86
	} else {
		goto L92
	}
L92:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+145)))
	if v345 != 0 {
		goto L86
	} else {
		goto L93
	}
L93:
	;
	goto L90
L94:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if int32(0) < v373 {
		goto L101
	} else {
		goto L102
	}
L95:
	;
	v352 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L7
	} else {
		goto L96
	}
L96:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v352 == int32(0) {
		v372 = v354
		goto L94
	} else {
		goto L97
	}
L97:
	;
	v359 = F_pg_rusage_show(m, l0+int32(256))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L7
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v359
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v354
	F_errmsg_internal(m, int32(_a_F_tuplesort_puttuple_common_6), v19)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L7
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_tuplesort_puttuple_common_2), int32(1248), int32(_a_F_tuplesort_puttuple_common_3))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L7
	} else {
		goto L100
	}
L100:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v372 = v371
	goto L94
L101:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v379 = int32(0)
	v381 = v376
	goto L104
L102:
	;
	goto L103
L103:
	;
	v424 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v424
	if v424 < v372 {
		goto L107
	} else {
		goto L108
	}
L104:
	;
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+8)))
	v395 = int32(1)
	v396 = v394 ^ v395
	*(*uint8)(unsafe.Add(mBase, uint32(v381)+8)) = uint8(v396)
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+9)))
	v400 = v398 ^ v395
	*(*uint8)(unsafe.Add(mBase, uint32(v381)+9)) = uint8(v400)
	v405 = v379 + v395
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v405 < v406 {
		v379 = v405
		v381 = v381 + int32(36)
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
	v437 = v5
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
	v444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v444 < v445 {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L109
L112:
	;
	v660 = v437 + int32(1)
	if v660 != v372 {
		v437 = v660
		goto L110
	} else {
		goto L158
	}
L113:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v450 = v447 + v437<<(uint(int32(4))%32)
	v451 = *(*int64)(unsafe.Add(mBase, uint32(v450)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v451
	v453 = *(*int64)(unsafe.Add(mBase, uint32(v450)))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v453
	v456 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_puttuple_common[1]))
	if v456 != 0 {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	goto L115
L115:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v529 = v437 << (uint(int32(4)) % 32)
	v531 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v532 = m.T0[v531].(func(*base.Module, int32, int32, int32) int32)(m, v527+v529, v527, l0)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L7
	} else {
		goto L127
	}
L116:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L7
	} else {
		goto L119
	}
L117:
	;
	v460 = v444
	goto L118
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v460 + int32(1)
	if v460 <= int32(0) {
		v507 = v460
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v460 = v459
	goto L118
L120:
	;
	v522 = v447 + v507<<(uint(int32(4))%32)
	v523 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v522)+8)) = v523
	v525 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v522))) = v525
	goto L112
L121:
	;
	v469 = v460
	goto L122
L122:
	;
	v484 = int32(1)
	v487 = int32(base.Ui32(v469-v484) >> (uint(v484) % 32))
	v490 = v447 + v487<<(uint(int32(4))%32)
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v492 = m.T0[v491].(func(*base.Module, int32, int32, int32) int32)(m, v19+int32(16), v490, l0)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L7
	} else {
		goto L124
	}
L123:
	;
	v507 = int32(0)
	goto L120
L124:
	;
	if int32(0) <= v492 {
		v507 = v469
		goto L120
	} else {
		goto L125
	}
L125:
	;
	v498 = v447 + v469<<(uint(int32(4))%32)
	v499 = *(*int64)(unsafe.Add(mBase, uint32(v490)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v498)+8)) = v499
	v501 = *(*int64)(unsafe.Add(mBase, uint32(v490)))
	*(*int64)(unsafe.Add(mBase, uint32(v498))) = v501
	if v487 != 0 {
		v469 = v487
		goto L122
	} else {
		goto L126
	}
L126:
	;
	goto L123
L127:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v535 = v534 + v529
	if v532 <= int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v535)))
	if v538 != 0 {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	goto L130
L130:
	;
	v558 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_puttuple_common[1]))
	if v558 != 0 {
		goto L138
	} else {
		goto L139
	}
L131:
	;
	v539 = F_GetMemoryChunkSpace(m, v538)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L7
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v552 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_puttuple_common[1]))
	if v552 == int32(0) {
		goto L112
	} else {
		goto L136
	}
L134:
	;
	v541 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v541 + base.I64_extend_i32_u(v539)
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v535)))
	F_pfree(m, v545)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L7
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v535))) = int32(0)
	goto L133
L136:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
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
	v560 = m.ExcPending
	if v560 != 0 {
		goto L7
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v561 = int32(0)
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if base.Ui32(v565) < base.Ui32(int32(2)) {
		v623 = v561
		goto L142
	} else {
		goto L143
	}
L141:
	;
	goto L140
L142:
	;
	v638 = v534 + v623<<(uint(int32(4))%32)
	v639 = *(*int64)(unsafe.Add(mBase, uint32(v535)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v638)+8)) = v639
	v641 = *(*int64)(unsafe.Add(mBase, uint32(v535)))
	*(*int64)(unsafe.Add(mBase, uint32(v638))) = v641
	goto L112
L143:
	;
	v570 = int32(1)
	v572 = v561
	v576 = v561
	goto L144
L144:
	;
	v585 = v576 + int32(2)
	if base.Ui32(v565) <= base.Ui32(v585) {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v623 = v599
	goto L142
L146:
	;
	v599 = v570
	goto L148
L147:
	;
	v587 = int32(4)
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v594 = m.T0[v593].(func(*base.Module, int32, int32, int32) int32)(m, v534+v570<<(uint(v587)%32), v534+v585<<(uint(v587)%32), l0)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L7
	} else {
		goto L149
	}
L148:
	;
	v602 = v534 + v599<<(uint(int32(4))%32)
	v603 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v604 = m.T0[v603].(func(*base.Module, int32, int32, int32) int32)(m, v535, v602, l0)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L7
	} else {
		goto L153
	}
L149:
	;
	if int32(0) < v594 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v598 = v585
	goto L152
L151:
	;
	v598 = v570
	goto L152
L152:
	;
	v599 = v598
	goto L148
L153:
	;
	if v604 <= int32(0) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v623 = v572
	goto L142
L155:
	;
	goto L156
L156:
	;
	v610 = v534 + v572<<(uint(int32(4))%32)
	v611 = *(*int64)(unsafe.Add(mBase, uint32(v602)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v610)+8)) = v611
	v613 = *(*int64)(unsafe.Add(mBase, uint32(v602)))
	*(*int64)(unsafe.Add(mBase, uint32(v610))) = v613
	v615 = int32(1)
	v616 = v599 << (uint(v615) % 32)
	v618 = v616 | v615
	if base.Ui32(v618) < base.Ui32(v565) {
		v570 = v618
		v572 = v599
		v576 = v616
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
	v683 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	if int64(0) <= v683 {
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
	v689 = m.ExcPending
	if v689 != 0 {
		goto L7
	} else {
		goto L164
	}
L162:
	;
	v686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+145)))
	if v686 != 0 {
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
	v692 = m.ExcPending
	if v692 != 0 {
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
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
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
	var v32 int32
	_ = v32
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if int32(2) <= v5 {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
		if v8 != int32(1) {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			if v27 != 0 {
				F_qsort_ssup(m, v26, v5, v27)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					return
				}
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				F_qsort_tuple(m, v26, v5, v30, l0)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			if v11 == int32(0) {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				if v27 != 0 {
					F_qsort_ssup(m, v26, v5, v27)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						return
					}
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					F_qsort_tuple(m, v26, v5, v30, l0)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
				if v14 == int32(116) {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
					F_qsort_tuple_unsigned(m, v17, v5, l0)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						return
					}
				} else {
					if v14 != int32(197) {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						if v27 != 0 {
							F_qsort_ssup(m, v26, v5, v27)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return
							} else {
								return
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							F_qsort_tuple(m, v26, v5, v30, l0)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
						F_qsort_tuple_int32(m, v22, v5, l0)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	} else {
		return
	}
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
