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
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	if l2&int32(1) != 0 {
		v11 = l1
	} else {
		v11 = int32(0)
	}
	if v11 == int32(0) {
		v15 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		v20 = F_AllocSetContextCreateInternal(m, v15, int32(261504), int32(0), int32(8192), int32(8388608))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v28 = F_AllocSetContextCreateInternal(m, v20, int32(74128), int32(0), int32(8192), int32(8388608))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = int32(4425280)
				v31 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v20
				v35 = F_palloc0(m, int32(424))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v38 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1095])))
					if v38 == int32(1) {
						F_getrusage(m, v35+int32(272))
						mBase = m.M
						F___gettimeofday(m, v35+int32(256))
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
							v100 = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v35)+240)) = v100
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v31
							return v35
						} else {
							v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
							v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v35)+236)) = v74
							if v73 == int32(1) {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
								*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(1)
								if v78 != 0 {
									F_s_lock(m, v74, int32(462590), int32(2988), int32(209176))
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
										v88 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v88 + int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v35)+232)) = v88
										v100 = int32(-1)
										*(*int32)(unsafe.Add(mBase, uint32(v35)+240)) = v100
										*(*int32)(unsafe.Add(mBase, _consts[0])) = v31
										return v35
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
									v88 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v88 + int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v35)+232)) = v88
									v100 = int32(-1)
									*(*int32)(unsafe.Add(mBase, uint32(v35)+240)) = v100
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v31
									return v35
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v35)+232)) = int32(-1)
								v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								v100 = v96
								*(*int32)(unsafe.Add(mBase, uint32(v35)+240)) = v100
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v31
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
		v108 = m.ExcPending
		if v108 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(74216), int32(0))
			mBase = m.M
			v112 = m.ExcPending
			if v112 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(462590), int32(651), int32(230233))
				mBase = m.M
				v117 = m.ExcPending
				if v117 != 0 {
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
		v20 = int32(4425280)
		v21 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v23
		v26 = F_palloc(m, int32(8))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v29 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1095])))
			if v29 != int32(1) {
				*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(1861)
				v55 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v55
				*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v26
				*(*uint8)(unsafe.Add(mBase, uint32(v16)+36)) = uint8(v55)
				*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(1862)
				*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(1863)
				*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1864)
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(1865)
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
						v86 = *(*int32)(unsafe.Add(mBase, _consts[0]))
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
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v21
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
						*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(1861)
						v55 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v55
						*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v26
						*(*uint8)(unsafe.Add(mBase, uint32(v16)+36)) = uint8(v55)
						*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(1862)
						*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(1863)
						*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1864)
						*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(1865)
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
								v86 = *(*int32)(unsafe.Add(mBase, _consts[0]))
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
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v21
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
						F_errmsg_internal(m, int32(470991), v13)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(463301), int32(686), int32(269033))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(1861)
								v55 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v55
								*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v26
								*(*uint8)(unsafe.Add(mBase, uint32(v16)+36)) = uint8(v55)
								*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(1862)
								*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(1863)
								*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1864)
								*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(1865)
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
										v86 = *(*int32)(unsafe.Add(mBase, _consts[0]))
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
											*(*int32)(unsafe.Add(mBase, _consts[0])) = v21
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v126 int32
	_ = v126
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v17 = F_tuplesort_begin_common(m, l1, l2, int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = int32(4425280)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+10)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v28
	v32 = F_palloc0(m, v28*int32(36))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v32
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	if int32(0) < v35 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L21
	}
L5:
	;
	v43 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	v126 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v126
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+36)) = uint8(v126)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = int32(1857)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(1858)
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(1859)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = int32(1860)
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v22
	m.G0 = v14 + int32(16)
	return v17
L8:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v17)+44))
	v56 = v53 + v43*int32(36)
	v58 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60+v43<<(uint(int32(2))%32))))
	v65 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+20)) = uint8(v65)
	v68 = v43 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+10)) = uint16(v68)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+9)) = uint8(v65)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v64
	if v64 == v65 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = int32(100)
	goto L12
L11:
	;
	goto L12
L12:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+6)))
	v89 = int32(4)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v79+v81*(base.I32_extend16_s(v68)-int32(1))<<(uint(int32(2))%32)+v89-v89)))
	goto L13
L13:
	;
	if v93 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v101 = v23 + int32(88) + v52<<(uint(int32(4))%32) + v43*int32(100)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v104 = F_lookup_type_cache(m, v102, int32(64))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	v110 = v93
	goto L16
L16:
	;
	F_PrepareSortSupportComparisonShim(m, v110, v56)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v104)+108))
	if v106 == int32(0) {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v110 = v106
	goto L16
L19:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	if v68 < v113 {
		v43 = v68
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
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v152 = F_format_type_be(m, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v152
	F_errmsg(m, int32(177662), v14)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(463301), int32(647), int32(260333))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
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
	v10 = int32(4425280)
	v11 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v13
	v16 = F_tuplesort_gettuple_common(m, l0, int32(1), v8)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		if v16 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v11
			v34 = v3
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v11
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	if base.Ui32(l0) <= base.Ui32(int32(8)) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[1242])))
		v11 = v10
	} else {
		v11 = int32(228441)
	}
	return v11
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v93 int32
	_ = v93
	var v99 float64
	_ = v99
	var v100 float64
	_ = v100
	var v103 float64
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int64
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int64
	_ = v209
	var v211 int64
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v236 int32
	_ = v236
	var v237 int64
	_ = v237
	var v239 int64
	_ = v239
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
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
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v348 int64
	_ = v348
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v431 int32
	_ = v431
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int64
	_ = v460
	var v462 int64
	_ = v462
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v478 int32
	_ = v478
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v508 int64
	_ = v508
	var v510 int64
	_ = v510
	var v516 int32
	_ = v516
	var v532 int32
	_ = v532
	var v533 int64
	_ = v533
	var v535 int64
	_ = v535
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int64
	_ = v551
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v621 int64
	_ = v621
	var v623 int64
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v648 int32
	_ = v648
	var v649 int64
	_ = v649
	var v651 int64
	_ = v651
	var v670 int32
	_ = v670
	var v691 int32
	_ = v691
	var v693 int64
	_ = v693
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	v5 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v21 = int32(4425280)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v24
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
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	switch v76 {
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
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+16)) = v57
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v60 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v59)+24)) = v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+28)) = v60
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v65)+32)) = v60
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	m.T0[v70].(func(*base.Module, int32, int32, int32))(m, l0, v68, v69)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L7
	} else {
		goto L11
	}
L4:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+24))
	v53 = m.T0[v52].(func(*base.Module, int32, int32) int32)(m, v50, v51)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v53
	goto L1
L11:
	;
	goto L1
L12:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v22
	m.G0 = v19 + int32(32)
	return
L13:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v328 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v327 + v328
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v334 = v331 + v327<<(uint(int32(4))%32)
	v335 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v334)+8)) = v335
	v337 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v334))) = v337
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+68)))
	if v340 != v328 {
		goto L89
	} else {
		goto L90
	}
L14:
	;
	v319 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+144)) = uint8(v319)
	goto L13
L15:
	;
	if v86 < base.I64_extend_i32_u((v274-v78)<<(uint(int32(4))%32)) {
		goto L14
	} else {
		goto L80
	}
L16:
	;
	v271 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+144)) = uint8(v271)
	v274 = int32(134217727)
	goto L15
L17:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v254 + int32(1)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v261 = v258 + v254<<(uint(int32(4))%32)
	v262 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v261)+8)) = v262
	v264 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v261))) = v264
	F_dumptuples(m, l0, int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L7
	} else {
		goto L79
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L7
	} else {
		goto L76
	}
L19:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v118 = m.T0[v117].(func(*base.Module, int32, int32, int32) int32)(m, l1, v116, l0)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L7
	} else {
		goto L38
	}
L20:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v77 < v78-int32(1) {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+144)))
	if v82 != int32(1) {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	v85 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
	v86 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	v87 = v85 - v86
	if v87 <= v86 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	if v78 != int32(2147483647) {
		goto L16
	} else {
		goto L37
	}
L24:
	;
	if v110 <= v78 {
		goto L14
	} else {
		goto L35
	}
L25:
	;
	if int32(1073741822) < v78 {
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v93 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+144)) = uint8(v93)
	v99 = base.F64_mul(base.F64_div(base.F64_convert_i64_s(v85), base.F64_convert_i64_s(v87)), base.F64_convert_i32_s(v78))
	v100 = float64(2.147483647e+09)
	if base.F64_lt(v99, v100) != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v110 = v78 << (uint(int32(1)) % 32)
	goto L24
L29:
	;
	v103 = v99
	goto L31
L30:
	;
	v103 = v100
	goto L31
L31:
	;
	if base.F64_lt(base.F64_abs(v103), float64(2.147483648e+09)) != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v107 = base.I32_trunc_f64_s(v103)
	v110 = v107
	goto L24
L33:
	;
	goto L34
L34:
	;
	v110 = int32(-2147483648)
	goto L24
L35:
	;
	if base.Ui32(v110) < base.Ui32(int32(134217727)) {
		v274 = v110
		goto L15
	} else {
		goto L36
	}
L36:
	;
	goto L16
L37:
	;
	goto L14
L38:
	;
	if v118 <= int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v122 != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	if v142 != 0 {
		goto L49
	} else {
		goto L50
	}
L42:
	;
	v123 = F_GetMemoryChunkSpace(m, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L7
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v136 == int32(0) {
		goto L12
	} else {
		goto L47
	}
L45:
	;
	v125 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v125 + base.I64_extend_i32_u(v123)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_pfree(m, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	goto L44
L47:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	goto L12
L49:
	;
	v143 = F_GetMemoryChunkSpace(m, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L7
	} else {
		goto L52
	}
L50:
	;
	v155 = v141
	goto L51
L51:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v158 != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v145 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v145 + base.I64_extend_i32_u(v143)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	F_pfree(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L7
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v141))) = int32(0)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v155 = v154
	goto L51
L54:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L7
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if base.Ui32(v161) < base.Ui32(int32(2)) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L56
L58:
	;
	v236 = v155 + v221<<(uint(int32(4))%32)
	v237 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v236))) = v237
	v239 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v236)+8)) = v239
	goto L12
L59:
	;
	v221 = int32(0)
	goto L58
L60:
	;
	goto L61
L61:
	;
	v170 = int32(1)
	v171 = v5
	v172 = v5
	goto L62
L62:
	;
	v183 = v172 + int32(2)
	if base.Ui32(v161) <= base.Ui32(v183) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v221 = v197
	goto L58
L64:
	;
	v197 = v170
	goto L66
L65:
	;
	v185 = int32(4)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v192 = m.T0[v191].(func(*base.Module, int32, int32, int32) int32)(m, v155+v170<<(uint(v185)%32), v155+v183<<(uint(v185)%32), l0)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L7
	} else {
		goto L67
	}
L66:
	;
	v200 = v155 + v197<<(uint(int32(4))%32)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v202 = m.T0[v201].(func(*base.Module, int32, int32, int32) int32)(m, l1, v200, l0)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L7
	} else {
		goto L71
	}
L67:
	;
	if int32(0) < v192 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v196 = v183
	goto L70
L69:
	;
	v196 = v170
	goto L70
L70:
	;
	v197 = v196
	goto L66
L71:
	;
	if v202 <= int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v221 = v171
	goto L58
L73:
	;
	goto L74
L74:
	;
	v208 = v155 + v171<<(uint(int32(4))%32)
	v209 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
	*(*int64)(unsafe.Add(mBase, uint32(v208))) = v209
	v211 = *(*int64)(unsafe.Add(mBase, uint32(v200)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v208)+8)) = v211
	v213 = int32(1)
	v214 = v197 << (uint(v213) % 32)
	v216 = v214 | v213
	if base.Ui32(v216) < base.Ui32(v161) {
		v170 = v216
		v171 = v197
		v172 = v214
		goto L62
	} else {
		goto L75
	}
L75:
	;
	goto L63
L76:
	;
	F_errmsg_internal(m, int32(330319), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L7
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(462590), int32(1312), int32(230406))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L7
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	goto L12
L80:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v282 = F_GetMemoryChunkSpace(m, v281)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L7
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v274
	v285 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v285 + base.I64_extend_i32_u(v282)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v292 = F_repalloc_huge(m, v289, v274<<(uint(int32(4))%32))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L7
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v292
	v295 = F_GetMemoryChunkSpace(m, v292)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L7
	} else {
		goto L83
	}
L83:
	;
	v297 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	v299 = v297 - base.I64_extend_i32_u(v295)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v299
	if int64(0) <= v299 {
		goto L13
	} else {
		goto L84
	}
L84:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+145)))
	if v303 != 0 {
		goto L13
	} else {
		goto L85
	}
L85:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L7
	} else {
		goto L86
	}
L86:
	;
	F_errmsg_internal(m, int32(73985), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L7
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(462590), int32(1156), int32(152549))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L7
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v339 < v691 {
		goto L164
	} else {
		goto L165
	}
L90:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v339 <= v343<<(uint(int32(1))%32) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	if v339 <= v343 {
		goto L89
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v353 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1095])))
	if v353 != int32(1) {
		v379 = v339
		goto L97
	} else {
		goto L98
	}
L94:
	;
	v348 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	if int64(0) <= v348 {
		goto L89
	} else {
		goto L95
	}
L95:
	;
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+145)))
	if v351 != 0 {
		goto L89
	} else {
		goto L96
	}
L96:
	;
	goto L93
L97:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if int32(0) < v380 {
		goto L104
	} else {
		goto L105
	}
L98:
	;
	v358 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L7
	} else {
		goto L99
	}
L99:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v358 == int32(0) {
		v379 = v360
		goto L97
	} else {
		goto L100
	}
L100:
	;
	v365 = F_pg_rusage_show(m, l0+int32(256))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L7
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v365
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v360
	F_errmsg_internal(m, int32(188410), v19)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L7
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(462590), int32(1248), int32(230406))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L7
	} else {
		goto L103
	}
L103:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v379 = v377
	goto L97
L104:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v386 = int32(0)
	v388 = v383
	goto L107
L105:
	;
	goto L106
L106:
	;
	v431 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v431
	if v431 < v379 {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388)+8)))
	v402 = int32(1)
	v403 = v401 ^ v402
	*(*uint8)(unsafe.Add(mBase, uint32(v388)+8)) = uint8(v403)
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388)+9)))
	v407 = v405 ^ v402
	*(*uint8)(unsafe.Add(mBase, uint32(v388)+9)) = uint8(v407)
	v412 = v386 + v402
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v412 < v413 {
		v386 = v412
		v388 = v388 + int32(36)
		goto L107
	} else {
		goto L109
	}
L108:
	;
	goto L106
L109:
	;
	goto L108
L110:
	;
	v445 = v5
	goto L113
L111:
	;
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(1)
	goto L12
L113:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v451 < v452 {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	goto L112
L115:
	;
	v670 = v445 + int32(1)
	if v670 != v379 {
		v445 = v670
		goto L113
	} else {
		goto L163
	}
L116:
	;
	v455 = v19 + int32(24)
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v459 = v456 + v445<<(uint(int32(4))%32)
	v460 = *(*int64)(unsafe.Add(mBase, uint32(v459)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v455))) = v460
	v462 = *(*int64)(unsafe.Add(mBase, uint32(v459)))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v462
	v465 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v465 != 0 {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	goto L118
L118:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v539 = v445 << (uint(int32(4)) % 32)
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v542 = m.T0[v541].(func(*base.Module, int32, int32, int32) int32)(m, v537+v539, v537, l0)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L7
	} else {
		goto L132
	}
L119:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L7
	} else {
		goto L122
	}
L120:
	;
	v469 = v451
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v469 + int32(1)
	if v469 <= int32(0) {
		v516 = v469
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v469 = v468
	goto L121
L123:
	;
	v532 = v456 + v516<<(uint(int32(4))%32)
	v533 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v532))) = v533
	v535 = *(*int64)(unsafe.Add(mBase, uint32(v455)))
	*(*int64)(unsafe.Add(mBase, uint32(v532)+8)) = v535
	goto L115
L124:
	;
	v478 = v469
	goto L125
L125:
	;
	v493 = int32(1)
	v494 = v478 - v493
	v496 = int32(base.Ui32(v494) >> (uint(v493) % 32))
	v499 = v456 + v496<<(uint(int32(4))%32)
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v501 = m.T0[v500].(func(*base.Module, int32, int32, int32) int32)(m, v19+int32(16), v499, l0)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L7
	} else {
		goto L127
	}
L126:
	;
	v516 = v496
	goto L123
L127:
	;
	if int32(0) <= v501 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v516 = v478
	goto L123
L129:
	;
	goto L130
L130:
	;
	v507 = v456 + v478<<(uint(int32(4))%32)
	v508 = *(*int64)(unsafe.Add(mBase, uint32(v499)))
	*(*int64)(unsafe.Add(mBase, uint32(v507))) = v508
	v510 = *(*int64)(unsafe.Add(mBase, uint32(v499)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v507)+8)) = v510
	if base.Ui32(int32(1)) < base.Ui32(v494) {
		v478 = v496
		goto L125
	} else {
		goto L131
	}
L131:
	;
	goto L126
L132:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v545 = v544 + v539
	if v542 <= int32(0) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v545)))
	if v548 != 0 {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	goto L135
L135:
	;
	v568 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v568 != 0 {
		goto L143
	} else {
		goto L144
	}
L136:
	;
	v549 = F_GetMemoryChunkSpace(m, v548)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L7
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v562 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v562 == int32(0) {
		goto L115
	} else {
		goto L141
	}
L139:
	;
	v551 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v551 + base.I64_extend_i32_u(v549)
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v545)))
	F_pfree(m, v555)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L7
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v545))) = int32(0)
	goto L138
L141:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L7
	} else {
		goto L142
	}
L142:
	;
	goto L115
L143:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L7
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v571 = int32(0)
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if base.Ui32(v575) < base.Ui32(int32(2)) {
		v633 = v571
		goto L147
	} else {
		goto L148
	}
L146:
	;
	goto L145
L147:
	;
	v648 = v544 + v633<<(uint(int32(4))%32)
	v649 = *(*int64)(unsafe.Add(mBase, uint32(v545)))
	*(*int64)(unsafe.Add(mBase, uint32(v648))) = v649
	v651 = *(*int64)(unsafe.Add(mBase, uint32(v545)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v648)+8)) = v651
	goto L115
L148:
	;
	v580 = int32(1)
	v582 = v571
	v584 = v571
	goto L149
L149:
	;
	v595 = v584 + int32(2)
	if base.Ui32(v575) <= base.Ui32(v595) {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v633 = v609
	goto L147
L151:
	;
	v609 = v580
	goto L153
L152:
	;
	v597 = int32(4)
	v603 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v604 = m.T0[v603].(func(*base.Module, int32, int32, int32) int32)(m, v544+v580<<(uint(v597)%32), v544+v595<<(uint(v597)%32), l0)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L7
	} else {
		goto L154
	}
L153:
	;
	v612 = v544 + v609<<(uint(int32(4))%32)
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v614 = m.T0[v613].(func(*base.Module, int32, int32, int32) int32)(m, v545, v612, l0)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L7
	} else {
		goto L158
	}
L154:
	;
	if int32(0) < v604 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v608 = v595
	goto L157
L156:
	;
	v608 = v580
	goto L157
L157:
	;
	v609 = v608
	goto L153
L158:
	;
	if v614 <= int32(0) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v633 = v582
	goto L147
L160:
	;
	goto L161
L161:
	;
	v620 = v544 + v582<<(uint(int32(4))%32)
	v621 = *(*int64)(unsafe.Add(mBase, uint32(v612)))
	*(*int64)(unsafe.Add(mBase, uint32(v620))) = v621
	v623 = *(*int64)(unsafe.Add(mBase, uint32(v612)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v620)+8)) = v623
	v625 = int32(1)
	v626 = v609 << (uint(v625) % 32)
	v628 = v626 | v625
	if base.Ui32(v628) < base.Ui32(v575) {
		v580 = v628
		v582 = v609
		v584 = v626
		goto L149
	} else {
		goto L162
	}
L162:
	;
	goto L150
L163:
	;
	goto L114
L164:
	;
	v693 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	if int64(0) <= v693 {
		goto L12
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	F_inittapes(m, l0, int32(1))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L7
	} else {
		goto L169
	}
L167:
	;
	v696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+145)))
	if v696 != 0 {
		goto L12
	} else {
		goto L168
	}
L168:
	;
	goto L166
L169:
	;
	F_dumptuples(m, l0, int32(0))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L7
	} else {
		goto L170
	}
L170:
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
		v4 = int32(13029)
	} else {
		v4 = int32(295605)
	}
	return v4
}
