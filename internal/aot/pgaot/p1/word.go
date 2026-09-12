package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_compareWORD(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v5 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	if v5 == int32(0) {
		v13 = int32(0)
		if v13 < v7 {
			v16 = int32(-1)
		} else {
			v16 = v13
		}
		v33 = v16
	} else {
		if v7 == int32(0) {
			v33 = base.B2i32(int32(0) < v5)
		} else {
			if base.Ui32(v5) < base.Ui32(v7) {
				v22 = v5
			} else {
				v22 = v7
			}
			v23 = F_memcmp(m, v4, v6, v22)
			mBase = m.M
			if v23 != 0 {
				v31 = v23
				v33 = v31
			} else {
				if v5 == v7 {
					v33 = int32(0)
				} else {
					if v5 < v7 {
						v30 = int32(-1)
					} else {
						v30 = int32(1)
					}
					v31 = v30
					v33 = v31
				}
			}
		}
	}
	if v33 != 0 {
		v44 = v33
	} else {
		v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
		v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
		if v35 == v36 {
			v44 = int32(0)
		} else {
			if base.Ui32(v36) < base.Ui32(v35) {
				v41 = int32(1)
			} else {
				v41 = int32(-1)
			}
			v44 = v41
		}
	}
	return v44
}
func F_compareWordEntryPos(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	v4 = int32(16383)
	v5 = v3 & v4
	v6 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v8 = v6 & v4
	return base.B2i32(base.Ui32(v8) < base.Ui32(v5)) - base.B2i32(base.Ui32(v5) < base.Ui32(v8))
}
func F_word_similarity_dist_op(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 float32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v17 = v10 + int32(1)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v19 = F_pg_detoast_datum_packed(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
			v22 = int32(1)
			v23 = v21 & v22
			if v21 == v22 {
				v26 = int32(4)
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				if v28&int32(254) == int32(2) {
					v37 = v26
				} else {
					v37 = base.B2i32(v28 == int32(18)) << (uint(v26) % 32)
				}
				if v28 == int32(1) {
					v40 = v26
				} else {
					v40 = v37
				}
				v51 = v40
			} else {
				v41 = int32(1)
				if v23 != 0 {
					v51 = int32(base.Ui32(v21)>>(uint(v41)%32)) - v41
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			if v23 != 0 {
				v52 = v17
			} else {
				v52 = v10 + int32(4)
			}
			v53 = int32(1)
			v54 = v19 + v53
			v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
			v59 = v57 & v53
			if v59 != 0 {
				v60 = v54
			} else {
				v60 = v19 + int32(4)
			}
			if v57 == int32(1) {
				v63 = int32(4)
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
				if v65&int32(254) == int32(2) {
					v74 = v63
				} else {
					v74 = base.B2i32(v65 == int32(18)) << (uint(v63) % 32)
				}
				if v65 == int32(1) {
					v77 = v63
				} else {
					v77 = v74
				}
				v88 = v77
			} else {
				v78 = int32(1)
				if v59 != 0 {
					v88 = int32(base.Ui32(v57)>>(uint(v78)%32)) - v78
				} else {
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					v88 = int32(base.Ui32(v82)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v90 = F_calc_word_similarity(m, v52, v51, v60, v88, int32(0))
			mBase = m.M
			v91 = m.ExcPending
			if v91 != 0 {
				return int32(0)
			} else {
				v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v92 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return int32(0)
					} else {
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v96 != v19 {
							F_pfree(m, v19)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								return base.I32_reinterpret_f32(base.F32_sub(float32(1), v90))
							}
						} else {
							return base.I32_reinterpret_f32(base.F32_sub(float32(1), v90))
						}
					}
				} else {
					v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v96 != v19 {
						F_pfree(m, v19)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							return base.I32_reinterpret_f32(base.F32_sub(float32(1), v90))
						}
					} else {
						return base.I32_reinterpret_f32(base.F32_sub(float32(1), v90))
					}
				}
			}
		}
	}
}
